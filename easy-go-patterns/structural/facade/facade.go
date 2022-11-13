package facade

import "log"

type subsystemA struct{}

func (a *subsystemA) MethodA() {
	log.Print("Sub system method A")
}

type subsystemB struct{}

func (b *subsystemB) MethodB() {
	log.Print("Sub system method B")
}

type subsystemC struct{}

func (c *subsystemC) MethodC() {
	log.Print("Sub system method C")
}

type subsystemD struct{}

func (d *subsystemD) MethodD() {
	log.Print("Sub system method D")
}

type facade struct {
	a *subsystemA
	b *subsystemB
	c *subsystemC
	d *subsystemD
}

func (f *facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f *facade) MethodTwo() {
	f.c.MethodC()
	f.d.MethodD()
}
