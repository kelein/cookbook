package factory

import (
	"testing"
)

func TestFactory_CreateFruit(t *testing.T) {
	var fruit Fruit
	type args struct {
		kind string
	}
	tests := []struct {
		name string
		args args
		want Fruit
	}{
		{"A", args{"apple"}, fruit},
		{"B", args{"pear"}, fruit},
		{"C", args{"banana"}, fruit},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Factory{}
			got := f.CreateFruit(tt.args.kind)
			if got != nil {
				got.Show()
			}
		})
	}
}
