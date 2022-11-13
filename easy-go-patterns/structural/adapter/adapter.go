package adapter

import "log"

// V5 of abstract
type V5 interface {
	Use5V()
}

// Phone with V5
type Phone struct {
	v V5
}

// NewPhone return a Phone instance
func NewPhone(v V5) *Phone {
	return &Phone{v}
}

// Charge of Phone
func (p *Phone) Charge() {
	log.Print("Phone start charging ...")
	p.v.Use5V()
}

// V220 Adaptee
type V220 struct{}

// Use220V of V220
func (v *V220) Use220V() {
	log.Print("charged by Use220V")
}

// Adapter of charge
type Adapter struct {
	v220 *V220
}

// NewAdapter create a adapter instance
func NewAdapter(v *V220) *Adapter {
	return &Adapter{v}
}

// Use5V of adapter
func (a *Adapter) Use5V() {
	log.Print("charged with Adapter")
	a.v220.Use220V()
}
