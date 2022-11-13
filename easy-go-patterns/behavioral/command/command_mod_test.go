package command

import "testing"

func TestWaiter_Notify(t *testing.T) {
	t.Run("WaiterNotify", func(t *testing.T) {
		cooker := new(Cooker)
		cmdChicken := &CookChickenCommand{cooker}
		cmdSkewers := &CookSkewersCommand{cooker}

		w := &Waiter{}
		w.commands = []CookerCommand{
			cmdChicken,
			cmdSkewers,
		}

		w.Notify()
	})

}
