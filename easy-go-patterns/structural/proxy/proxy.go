package proxy

import "log"

// Goods stands for products
type Goods struct {
	Kind string
	Fact bool
}

// Shopping of abstract
type Shopping interface {
	Buy(goods *Goods)
}

// KoreaShopping of concrete
type KoreaShopping struct{}

// Buy of concrete shopping
func (ks *KoreaShopping) Buy(goods *Goods) {
	log.Printf("buy %s from Korea", goods.Kind)
}

// AmericanShopping of concrete
type AmericanShopping struct{}

// Buy of concrete shopping
func (as *AmericanShopping) Buy(goods *Goods) {
	log.Printf("buy %s from American", goods.Kind)
}

// AfricaShopping of concrete
type AfricaShopping struct{}

// Buy of concrete shopping
func (af *AfricaShopping) Buy(goods *Goods) {
	log.Printf("buy %s from Africa", goods.Kind)
}

// OverseaProxy proxy for shopping
type OverseaProxy struct {
	shopping Shopping
}

// NewOverseaProxy create a new OverseaProxy instance
func NewOverseaProxy(shopping Shopping) Shopping {
	return &OverseaProxy{shopping}
}

// Buy of oversea proxy
func (op *OverseaProxy) Buy(goods *Goods) {
	if op.distinguish(goods) {
		op.shopping.Buy(goods)
		op.check(goods)
	}
}

func (op *OverseaProxy) distinguish(goods *Goods) bool {
	log.Printf("proxy start check [%s]", goods.Kind)
	if !goods.Fact {
		log.Printf("proxy found fake [%s],", goods.Kind)
		return false
	}
	return true
}

func (op *OverseaProxy) check(goods *Goods) {
	log.Printf("proxy checked passport [%s]", goods.Kind)
}
