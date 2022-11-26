package factory

import "log"

// Fruit of abstract
type Fruit interface {
	Show()
}

// Apple of concrete
type Apple struct {
	Fruit
}

// NewApple return a new Apple instance
func NewApple() *Apple { return new(Apple) }

// Show method of Apple
func (a *Apple) Show() {
	log.Printf("This is an Apple")
}

// Pear of concrete
type Pear struct {
	Fruit
}

// NewPear return a new Pear instance
func NewPear() *Pear { return new(Pear) }

// Show method of Pear
func (p *Pear) Show() {
	log.Printf("This is a Pear")
}

// Factory of concrete
type Factory struct{}

// CreateFruit product all kinds of fruit
func (f *Factory) CreateFruit(kind string) Fruit {
	if kind == "apple" {
		return NewApple()
	}

	if kind == "pear" {
		return NewPear()
	}

	return nil
}
