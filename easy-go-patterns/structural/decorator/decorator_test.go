package decorator

import (
	"log"
	"strings"
	"testing"
)

func TestDecorator(t *testing.T) {
	type args struct {
		phone Phone
	}
	tests := []struct {
		name string
		args args
	}{
		{"A", args{new(Huawei)}},
		{"B", args{new(Xiaomi)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.phone.Show()

			log.Print(strings.Repeat("-", 20))
			filmedPhone := NewFilmDecorator(tt.args.phone)
			filmedPhone.Show()

			log.Print(strings.Repeat("-", 20))
			shelledPhone := NewShellDecorator(filmedPhone)
			shelledPhone.Show()
		})
	}
}
