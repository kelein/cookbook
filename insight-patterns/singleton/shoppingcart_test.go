package singleton

import (
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/sync/errgroup"
)

func BenchmarkGetShoppingCartManager(b *testing.B) {
	innerCart := GetShoppingCartManager()
	for i := 0; i < b.N; i++ {
		got := GetShoppingCartManager()
		if !reflect.DeepEqual(got, innerCart) {
			b.Errorf("got unequal instance, got: %v, want: %v", got, innerCart)
		}
	}
}

func TestGetShoppingCartManager(t *testing.T) {
	innerCart := GetShoppingCartManager()

	tests := []struct {
		name  string
		count int
		want  *ShoppingCartManager
	}{
		{"A", 100, innerCart},
		{"B", 1000, innerCart},
		{"C", 99999, innerCart},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eg := errgroup.Group{}
			for i := 0; i < tt.count; i++ {
				eg.Go(func() error {
					got := GetShoppingCartManager()
					if !reflect.DeepEqual(got, tt.want) {
						return fmt.Errorf("got=%v, want %v", got, tt.want)
					}
					return nil
				})
			}
			if err := eg.Wait(); err != nil {
				t.Errorf("GetShoppingCartManager() occurs error: %v", err)
			}
		})
	}
}

func TestShoppingCartManager_AddProduct(t *testing.T) {
	cart := GetShoppingCartManager()

	type args struct {
		name     string
		quantity int
	}
	tests := []struct {
		name string
		args args
	}{
		{"A", args{"book", 10}},
		{"B", args{"pen", 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cart.AddProduct(tt.args.name, tt.args.quantity)
			cart.ShowProduct()
		})
	}
}
