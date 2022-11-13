package decorator

import "log"

// Phone of abstract
type Phone interface {
	Show()
}

// Decorator for Phone
type Decorator struct {
	phone Phone
}

// Show stands for abstract method
func (d *Decorator) Show() {}

// Huawei Phone
type Huawei struct{}

// Show for Huawei Phone
func (h *Huawei) Show() {
	log.Print("Huawei Phone")
}

// Xiaomi Phone
type Xiaomi struct{}

// Show for Xiaomi Phone
func (x *Xiaomi) Show() {
	log.Printf("Xiaomi Phone")
}

// FilmDecorator for Phone Film
type FilmDecorator struct {
	Decorator
}

// NewFilmDecorator create a FilmDecorator instance
func NewFilmDecorator(phone Phone) Phone {
	return &FilmDecorator{Decorator{phone}}
}

// Show of FilmDecorator
func (fd *FilmDecorator) Show() {
	fd.phone.Show()
	log.Printf("decorator with film")
}

// ShellDecorator for Phone Shell
type ShellDecorator struct {
	Decorator
}

// NewShellDecorator create a ShellDecorator instance
func NewShellDecorator(phone Phone) Phone {
	return &ShellDecorator{Decorator{phone}}
}

// Show of ShellDecorator
func (sd *ShellDecorator) Show() {
	sd.phone.Show()
	log.Printf("decorator with shell")
}
