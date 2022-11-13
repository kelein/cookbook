package command

import "log"

// Doctor .
type Doctor struct{}

func (d *Doctor) treatEye() {
	log.Print("Doctor treat Eyes")
}

func (d *Doctor) treatNose() {
	log.Print("Doctor treat Nose")
}

// Command of abstract
type Command interface {
	Treat()
}

// TreatEyeCommand treat Eyes with command
type TreatEyeCommand struct {
	doctor *Doctor
}

// Treat of TreatEyeCommand
func (te *TreatEyeCommand) Treat() {
	te.doctor.treatEye()
}

// TreatNoseCommand treat Nose with command
type TreatNoseCommand struct {
	doctor *Doctor
}

// Treat of TreatNoseCommand
func (tn *TreatNoseCommand) Treat() {
	tn.doctor.treatNose()
}

// Nurse stands for command invoker
type Nurse struct {
	commands []Command
}

// Notify command to execute
func (n *Nurse) Notify() {
	log.Print("Nurse notify to treating ...")
	for _, cmd := range n.commands {
		cmd.Treat()
	}
}
