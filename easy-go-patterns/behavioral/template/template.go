package template

import "log"

// Beverage for abstract
type Beverage interface {
	BoilWater()
	Brew()
	PourCup()
	Relish()
	WantRelished() bool
}

// BeverageMaker for abstract
type BeverageMaker interface {
	MakeBeverage()
}

type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourCup()

	if t.b.WantRelished() {
		t.b.Relish()
	}
}

// MakeCaffee make caffee with template
type MakeCaffee struct {
	template
	wantRelished bool
}

// NewMakeCaffee create a MakeCaffee instance
func NewMakeCaffee(wantRelished bool) *MakeCaffee {
	mc := new(MakeCaffee)
	mc.b = mc
	mc.wantRelished = wantRelished
	return mc
}

// BoilWater of MakeCaffee
func (mc *MakeCaffee) BoilWater() {
	log.Print("[MakeCaffee] boil water to 100 °C")
}

// Brew of MakeCaffee
func (mc *MakeCaffee) Brew() {
	log.Print("[MakeCaffee] brew with water and caffee beans")
}

// PourCup of MakeCaffee
func (mc *MakeCaffee) PourCup() {
	log.Print("[MakeCaffee] pour into crystal cup")
}

// Relish of MakeCaffee
func (mc *MakeCaffee) Relish() {
	log.Print("[MakeCaffee] relish with milk and sugar")
}

// WantRelished of MakeCaffee
func (mc *MakeCaffee) WantRelished() bool {
	return true
}

// MakeTea make tea with template
type MakeTea struct {
	template
	wantRelished bool
}

// NewMakeTea create a MakeTea instance
func NewMakeTea(wantRelished bool) *MakeTea {
	mt := new(MakeTea)
	mt.b = mt
	mt.wantRelished = wantRelished
	return mt
}

// BoilWater of MakeTea
func (mt *MakeTea) BoilWater() {
	log.Print("[MakeTea] boil water to 80 °C")
}

// Brew of MakeTea
func (mt *MakeTea) Brew() {
	log.Print("[MakeTea] brew with water and tea leaf")
}

// PourCup of MakeTea
func (mt *MakeTea) PourCup() {
	log.Print("[MakeTea] pour into tea cup")
}

// Relish of MakeTea
func (mt *MakeTea) Relish() {
	log.Print("[MakeTea] relish with lemon")
}

// WantRelished of MakeTea
func (mt *MakeTea) WantRelished() bool {
	return mt.wantRelished
}
