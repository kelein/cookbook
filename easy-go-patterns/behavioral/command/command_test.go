package command

import "testing"

func TestNurse_Notify(t *testing.T) {
	t.Run("Notify", func(t *testing.T) {
		doctor := new(Doctor)
		cmdEye := TreatEyeCommand{doctor}
		cmdNose := TreatNoseCommand{doctor}

		n := &Nurse{}
		n.commands = []Command{&cmdEye, &cmdNose}
		n.Notify()
	})
}
