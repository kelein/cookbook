package strategy

import "log"

// PromotionStrategy of abstract
type PromotionStrategy interface {
	GetPrice(price float64) float64
}

// CouponStrategy of concrete
type CouponStrategy struct{}

// GetPrice of CouponStrategy
func (c *CouponStrategy) GetPrice(price float64) float64 {
	log.Print("use CouponStrategy: 100 Coupon when price beyond 200")
	if price >= 200 {
		price -= 100
	}
	return price
}

// ScoreStrategy of concrete
type ScoreStrategy struct{}

// GetPrice of ScoreStrategy
func (s *ScoreStrategy) GetPrice(price float64) float64 {
	log.Print("use ScoreStrategy: All goods with eight snap")
	return price * 0.8
}

// Goods for promotion
type Goods struct {
	Strategy PromotionStrategy
	Price    float64
}

// SetStrategy of Goods
func (g *Goods) SetStrategy(p PromotionStrategy) {
	g.Strategy = p
}

// SalePrice of Goods
func (g *Goods) SalePrice() float64 {
	return g.Strategy.GetPrice(g.Price)
}
