package strategy

import (
	"log"
	"testing"
)

func TestGoods_SalePrice(t *testing.T) {
	type args struct {
		Price    float64
		Strategy PromotionStrategy
	}
	tests := []struct {
		name string
		args args
	}{
		{"A", args{200, new(CouponStrategy)}},
		{"B", args{100, new(ScoreStrategy)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Goods{Price: tt.args.Price}
			g.SetStrategy(tt.args.Strategy)
			price := g.SalePrice()
			log.Printf("orig price: ￥%.2f", g.Price)
			log.Printf("final price: ￥%.2f", price)
		})
	}
}
