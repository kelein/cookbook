package tests

import (
	"log"
	"testing"
)

func TestWallet_Deposit(t *testing.T) {
	wallet := Wallet{100}

	type args struct {
		amount Bitcoin
	}
	tests := []struct {
		name string
		args args
		want Bitcoin
	}{
		{"A", args{10}, 110},
		{"B", args{20}, 130},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallet.Deposit(tt.args.amount)
			log.Printf("address of balance in Wallets: %p", &wallet.balance)
			if got := wallet.Ballence(); got != tt.want {
				t.Errorf("%+v Deposit(%v) then got %q, want %q", tt, tt.args.amount, got, tt.want)
			}
		})
	}
}

func TestWallet_Withdraw(t *testing.T) {
	wallet := Wallet{100}

	type args struct {
		amount Bitcoin
	}
	tests := []struct {
		name    string
		args    args
		want    Bitcoin
		wantErr bool
	}{
		{"A", args{10}, 90, false},
		{"B", args{20}, 70, false},
		{"C", args{110}, 70, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Printf("address of balance in Wallets: %p", &wallet.balance)
			err := wallet.Withdraw(tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("wallet.Withdraw(%v) err = %v", tt.args.amount, err)
			}
			if got := wallet.Ballence(); got != tt.want {
				t.Errorf("%+v Withdraw(%v) then got %q, want %q", tt, tt.args.amount, got, tt.want)
			}
		})
	}
}
