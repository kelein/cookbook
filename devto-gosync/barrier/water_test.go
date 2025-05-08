package barrier

import (
	"strconv"
	"testing"
)

func TestWaterFactory(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		args    args
		wantErr bool
	}{
		{args{-200}, true},
		{args{-10}, true},
		{args{0}, true},
		{args{1}, false},
		{args{5}, false},
		{args{10}, false},
		{args{100}, false},
		{args{200}, false},
		{args{300}, false},
		{args{500}, false},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.args.count), func(t *testing.T) {
			if err := WaterFactory(tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("WaterFactory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
