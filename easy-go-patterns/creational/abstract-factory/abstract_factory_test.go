package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	t.Run("IntelFactory", func(t *testing.T) {
		fact := new(IntelFactory)
		com := new(Computer)
		com.cpu = fact.CreateCPU()
		com.gpu = fact.CreateGPU()
		com.ram = fact.CreateRAM()
		com.Work()
	})

	t.Run("NvidiaFactory", func(t *testing.T) {
		fact := new(NvidiaFactory)
		com := new(Computer)
		com.cpu = fact.CreateCPU()
		com.gpu = fact.CreateGPU()
		com.ram = fact.CreateRAM()
		com.Work()
	})

	t.Run("KingstonFactory", func(t *testing.T) {
		fact := new(KingstonFactory)
		com := new(Computer)
		com.cpu = fact.CreateCPU()
		com.gpu = fact.CreateGPU()
		com.ram = fact.CreateRAM()
		com.Work()
	})

	t.Run("MixedFactory", func(t *testing.T) {
		com := new(Computer)
		com.cpu = new(IntelFactory).CreateCPU()
		com.gpu = new(NvidiaFactory).CreateGPU()
		com.ram = new(KingstonFactory).CreateRAM()
		com.Work()
	})
}
