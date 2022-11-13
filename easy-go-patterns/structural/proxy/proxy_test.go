package proxy

import (
	"testing"
)

func TestOverseaProxy(t *testing.T) {
	type args struct {
		gs       []Goods
		shopping Shopping
	}
	tests := []struct {
		name string
		args args
	}{
		{"A", args{[]Goods{{"Phone", true}}, new(KoreaShopping)}},
		{"B", args{[]Goods{{"Car", true}}, new(AmericanShopping)}},
		{"C", args{[]Goods{{"Agate", false}, {"Iron", true}}, new(AfricaShopping)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proxy := NewOverseaProxy(tt.args.shopping)
			for _, g := range tt.args.gs {
				proxy.Buy(&g)
			}
		})
	}
}
