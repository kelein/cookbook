package template

import "testing"

func Test_template_MakeBeverage(t *testing.T) {
	type args struct {
		maker BeverageMaker
	}
	tests := []struct {
		name string
		args args
	}{
		{"CaffeeNoRelish", args{NewMakeCaffee(false)}},
		{"TeaNoRelish", args{NewMakeTea(false)}},
		{"CaffeeRelished", args{NewMakeCaffee(true)}},
		{"TeaNoRelished", args{NewMakeTea(true)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.maker.MakeBeverage()
		})
	}
}
