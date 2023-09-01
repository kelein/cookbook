package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cookbook/devto-grpc/repo"
	"cookbook/devto-grpc/service"
	"cookbook/devto-grpc/store"
)

func TestLaptopServer_CreateLaptop(t *testing.T) {
	t.Parallel()

	/* --------- Test Data Begin ------------- */
	laptopNoID := service.NewLaptop()
	laptopNoID.Id = ""

	laptopBadID := service.NewLaptop()
	laptopBadID.Id = "x-x-x-x-x-x"

	laptopDupID := service.NewLaptop()
	laptopDupStore := store.NewMemoryLaptopStore()
	err := laptopDupStore.Save(laptopDupID)
	require.Nil(t, err)
	/* --------- Test Data Done -------------- */

	tests := []struct {
		name   string
		code   codes.Code
		laptop *repo.Laptop
		store  store.LaptopStore
	}{
		{
			name:   "success_with_id",
			laptop: service.NewLaptop(),
			code:   codes.OK,
			store:  store.NewMemoryLaptopStore(),
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			code:   codes.OK,
			store:  store.NewMemoryLaptopStore(),
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopBadID,
			code:   codes.InvalidArgument,
			store:  store.NewMemoryLaptopStore(),
		},
		{
			name:   "failure_dup_id",
			laptop: laptopDupID,
			code:   codes.AlreadyExists,
			store:  laptopDupStore,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := &repo.CreateLaptopRequest{Laptop: tt.laptop}
			imgStore := store.NewDiskImageStore("../tests/imgs")
			server := service.NewLaptopServer(tt.store, imgStore)
			res, err := server.CreateLaptop(context.Background(), req)
			t.Logf("CreateLaptop() code=[%v], res=[%v], err=[%v]", tt.code, res, err)
			if tt.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				require.Equal(t, tt.laptop.Id, res.Id)
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				state, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.code, state.Code())
			}
		})
	}
}
