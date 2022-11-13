package facade

import "testing"

func Test_facade(t *testing.T) {
	t.Run("WithoutFacade", func(t *testing.T) {
		a := new(subsystemA)
		a.MethodA()

		b := new(subsystemB)
		b.MethodB()
	})

	t.Run("WithFacade", func(t *testing.T) {
		f := facade{
			a: new(subsystemA),
			b: new(subsystemB),
			c: new(subsystemC),
			d: new(subsystemD),
		}

		f.MethodOne()
		f.MethodTwo()
	})
}
