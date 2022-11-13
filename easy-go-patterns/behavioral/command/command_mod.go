package command

import "log"

// Cooker stands for command Receiver
type Cooker struct{}

// MakeChicken action of Receiver
func (c *Cooker) MakeChicken() {
	log.Print("Cooker make chicken")
}

// MakeSkewers action of Receiver
func (c *Cooker) MakeSkewers() {
	log.Print("Cooker make skewers")
}

// CookerCommand for abstract
type CookerCommand interface {
	Make()
}

// CookChickenCommand concrete command
type CookChickenCommand struct {
	cooker *Cooker
}

// Make of CookChickenCommand
func (cmd *CookChickenCommand) Make() {
	cmd.cooker.MakeChicken()
}

// CookSkewersCommand concrete command
type CookSkewersCommand struct {
	cooker *Cooker
}

// Make of CookSkewersCommand
func (cmd *CookSkewersCommand) Make() {
	cmd.cooker.MakeSkewers()
}

// Waiter stands for command invoker
type Waiter struct {
	commands []CookerCommand
}

// Notify make command to be executed
func (w *Waiter) Notify() {
	log.Print("Waiter notify to cooking ...")
	for _, cmd := range w.commands {
		cmd.Make()
	}
}
