package strategy

import (
	"testing"
)

func TestPayContext_Pay(t *testing.T) {
	type args struct {
		behavior PayBehavior
	}
	tests := []struct {
		name string
		args args
	}{
		{"WechatPay", args{&WechatPay{}}},
		{"ThirdPartyPay", args{&ThirdPartyPay{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewPayContext(tt.args.behavior)
			ctx.Pay()
		})
	}
}

func TestPayContext_setBehavior(t *testing.T) {
	type args struct {
		first  PayBehavior
		second PayBehavior
	}
	tests := []struct {
		name string
		args args
	}{
		{"SwitchPayMethod", args{&WechatPay{}, &ThirdPartyPay{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewPayContext(tt.args.first)
			ctx.Pay()

			ctx.setBehavior(tt.args.second)
			ctx.Pay()
		})
	}
}
