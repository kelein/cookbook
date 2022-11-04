package abstractfactory

import "log"

// Computer of abstract
type Computer struct {
	cpu CPU
	gpu GPU
	ram RAM
}

// Work for computer
func (c *Computer) Work() {
	c.cpu.Calculate()
	c.gpu.Display()
	c.ram.Storage()
}

/* ------------------ Abstract Products ----------------- */

// CPU for abstract
type CPU interface {
	Calculate()
}

// RAM for abstract
type RAM interface {
	Storage()
}

// GPU for abstract
type GPU interface {
	Display()
}

/* ------------------ Abstract Factory ------------------ */

// AbstractFactory create all kinds of products
type AbstractFactory interface {
	CreateCPU() CPU
	CreateGPU() GPU
	CreateRAM() RAM
}

/* ------------------ Concrete Products ----------------- */

// IntelCPU CPU created by Intel
type IntelCPU struct{}

// Calculate by Intel CPU
func (i *IntelCPU) Calculate() {
	log.Print("IntelCPU Calculate ...")
}

// IntelGPU GPU created by Intel
type IntelGPU struct{}

// Display by Intel GPU
func (i *IntelGPU) Display() {
	log.Print("IntelGPU Display ...")
}

// IntelMemory Memory created by Intel
type IntelMemory struct{}

// Storage by intel memory
func (i *IntelMemory) Storage() {
	log.Print("IntelMemory Storage ...")
}

/* ------------------ Concrete Factory ------------------ */

// IntelFactory create products of Intel
type IntelFactory struct{}

// CreateCPU create CPU by Intel
func (f *IntelFactory) CreateCPU() CPU {
	return new(IntelCPU)
}

// CreateGPU create GPU by Intel
func (f *IntelFactory) CreateGPU() GPU {
	return new(IntelGPU)
}

// CreateRAM create memory by Intel
func (f *IntelFactory) CreateRAM() RAM {
	return new(IntelMemory)
}

// NvidiaCPU CPU create by Nvidia
type NvidiaCPU struct{}

// Calculate by Nvidia CPU
func (n *NvidiaCPU) Calculate() {
	log.Print("NvidiaCPU Calculate ...")
}

// NvidiaGPU GPU create by Nvidia
type NvidiaGPU struct{}

// Display by Nvidia GPU
func (n *NvidiaGPU) Display() {
	log.Print("NvidiaGPU Display ...")
}

// NvidiaMemory memory by Nvidia
type NvidiaMemory struct{}

// Storage by Nvidia Memory
func (n *NvidiaMemory) Storage() {
	log.Print("NvidiaMemory Storage ...")
}

// NvidiaFactory create products of Nvidia
type NvidiaFactory struct{}

// CreateCPU create CPU by Nvidia
func (f *NvidiaFactory) CreateCPU() CPU {
	return new(NvidiaCPU)
}

// CreateGPU create GPU by Nvidia
func (f *NvidiaFactory) CreateGPU() GPU {
	return new(NvidiaGPU)
}

// CreateRAM create memory by Nvidia
func (f *NvidiaFactory) CreateRAM() RAM {
	return new(NvidiaMemory)
}

// KindstonCPU CPU created by Kindston
type KindstonCPU struct{}

// Calculate by Kindston CPU
func (k *KindstonCPU) Calculate() {
	log.Print("KindstonCPU Calculate ...")
}

// KindstonGPU GPU created by Kindston
type KindstonGPU struct{}

// Display by Kindston GPU
func (k *KindstonGPU) Display() {
	log.Print("KindstonGPU Display ...")
}

// KindstonRAM RAM created by Kindston
type KindstonRAM struct{}

// Storage by Kindston Memory
func (k *KindstonRAM) Storage() {
	log.Print("KindstonRAM Storage ...")
}

// KingstonFactory create products of Kindston
type KingstonFactory struct{}

// CreateCPU create CPU by Kindston
func (f *KingstonFactory) CreateCPU() CPU {
	return new(KindstonCPU)
}

// CreateGPU create GPU by Kindston
func (f *KingstonFactory) CreateGPU() GPU {
	return new(KindstonGPU)
}

// CreateRAM create RAM by Kindston
func (f *KingstonFactory) CreateRAM() RAM {
	return new(KindstonRAM)
}
