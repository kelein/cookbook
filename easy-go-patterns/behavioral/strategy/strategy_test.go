package strategy

import "testing"

func TestHero_Attack(t *testing.T) {
	type args struct {
		w WeaponStrategy
	}
	tests := []struct {
		name string
		args args
	}{
		{"A", args{new(Rifle)}},
		{"B", args{new(Knife)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hero{}
			h.SetWeaponStrategy(tt.args.w)
			h.Attack()
		})
	}
}
