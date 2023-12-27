package errcode

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRPCCode(t *testing.T) {
	tests := []struct {
		name string
		code int
		want codes.Code
	}{
		{
			name: "Success",
			code: 0,
			want: codes.OK,
		},
		{
			name: "Fail",
			code: 10000000,
			want: codes.Internal,
		},
		{
			name: "InvalidParams",
			code: 10000001,
			want: codes.InvalidArgument,
		},
		{
			name: "Unauthorized",
			code: 10000002,
			want: codes.Unauthenticated,
		},
		{
			name: "NotFound",
			code: 10000003,
			want: codes.NotFound,
		},
		{
			name: "Unknown",
			code: 10000004,
			want: codes.Unknown,
		},
		{
			name: "DeadlineExceeded",
			code: 10000005,
			want: codes.DeadlineExceeded,
		},
		{
			name: "AccessDenied",
			code: 10000006,
			want: codes.PermissionDenied,
		},
		{
			name: "LimitExceed",
			code: 10000007,
			want: codes.ResourceExhausted,
		},
		{
			name: "MethodNotAllowed",
			code: 10000008,
			want: codes.Unimplemented,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RPCCode(tt.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RPCCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromError2(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *CommonStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromError(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		want    *status.Status
		wantErr bool
	}{
		{
			name: "success",
			err:  errors.New("test error"),
			want: status.New(codes.Unknown, "test error"),
		},
		{
			name: "nil error",
			err:  nil,
			want: status.New(codes.OK, ""),
		},
		{
			name: "unknown error",
			err:  fmt.Errorf("unknown error"),
			want: status.New(codes.Unknown, "unknown error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromError(tt.err)
			if got.String() != tt.want.String() {
				t.Errorf("FromError() = %q, want %q", got, tt.want)
			}
		})
	}
}
