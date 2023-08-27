package service

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"cookbook/devto-grpc/repo"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewKeyboard return a keyboard instance
func NewKeyboard() *repo.Keyboard {
	return &repo.Keyboard{
		Layout:  randLayout(),
		Backlit: randBool(),
	}
}

func randBool() bool {
	return rand.Intn(2) == 1
}

func randLayout() repo.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return repo.Keyboard_QWERTY
	case 2:
		return repo.Keyboard_QWERTZ
	case 3:
		return repo.Keyboard_AZERTY
	default:
		return repo.Keyboard_UNKNOWN
	}
}

func randEnumSet(e ...string) string {
	n := len(e)
	if n == 0 {
		return ""
	}
	return e[rand.Intn(n)]
}

func randCPUBrand() string {
	return randEnumSet("Intel", "AMD")
}

func randCPUName(brand string) string {
	if brand == "Intel" {
		return randEnumSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randEnumSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// NewCPU return a new CPU instance
func NewCPU() *repo.CPU {
	brand := randCPUBrand()
	name := randCPUName(brand)
	minGhz := randFloat(2, 3)
	maxGhz := randFloat(minGhz, 6)
	cores := randInt(4, 10)
	threads := randInt(cores, 12)

	return &repo.CPU{
		Brand:   brand,
		Name:    name,
		MinGhz:  minGhz,
		MaxGhz:  maxGhz,
		Cores:   uint32(cores),
		Threads: uint32(threads),
	}
}

func randGPUBrand() string {
	return randEnumSet("Nvidia", "AMD")
}

func randGPUName(brand string) string {
	if brand == "Nvidia" {
		return randEnumSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randEnumSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

// NewGPU return a new GPU instance
func NewGPU() *repo.GPU {
	brand := randGPUBrand()
	name := randGPUName(brand)
	minGhz := randFloat(2, 5)
	maxGhz := randFloat(minGhz, 10)
	memGB := randInt(8, 256)

	return &repo.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &repo.Memory{
			Unit:  repo.Memory_GIGABYTE,
			Value: uint64(memGB),
		},
	}
}

// NewRAM return a new RAM instance
func NewRAM() *repo.Memory {
	return &repo.Memory{
		Unit:  repo.Memory_GIGABYTE,
		Value: uint64(randInt(8, 128)),
	}
}

// NewSSD return a new SSD storage
func NewSSD() *repo.Storage {
	return &repo.Storage{
		Driver: repo.Storage_HDD,
		Memory: &repo.Memory{
			Unit:  repo.Memory_GIGABYTE,
			Value: uint64(randInt(512, 1024)),
		},
	}
}

// NewHHD return a new HDD storage
func NewHHD() *repo.Storage {
	return &repo.Storage{
		Driver: repo.Storage_HDD,
		Memory: &repo.Memory{
			Unit:  repo.Memory_TERABYTE,
			Value: uint64(randInt(1, 6)),
		},
	}
}

func randScreenResolution() *repo.Screen_Resolution {
	height := randInt(1080, 4320)
	width := height * 16 / 9
	return &repo.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
}

func randScreenPanel() repo.Screen_Panel {
	if rand.Intn(2) == 1 {
		return repo.Screen_IPS
	}
	return repo.Screen_OLED
}

// NewScreen return a new Screen instance
func NewScreen() *repo.Screen {
	return &repo.Screen{
		Panel:      randScreenPanel(),
		SizeInch:   float32(randFloat(13, 17)),
		Resolution: randScreenResolution(),
		Multitouch: randBool(),
	}
}

func randUUID() string {
	return uuid.New().String()
}

func randLaptopBrand() string {
	return randEnumSet("Apple", "Dell", "Lenovo")
}

func randLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randEnumSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randEnumSet("XPS", "Vostro", "Latitude")
	default:
		return randEnumSet("Thinkpad X1", "Thinkpad T405")
	}
}

// NewLaptop return a new Laptop instance
func NewLaptop() *repo.Laptop {
	brand := randLaptopBrand()
	name := randLaptopName(brand)

	return &repo.Laptop{
		Id:          randUUID(),
		Brand:       brand,
		Name:        name,
		Ram:         NewRAM(),
		Cpu:         NewCPU(),
		Gpus:        []*repo.GPU{NewGPU()},
		Storages:    []*repo.Storage{NewSSD(), NewHHD()},
		Screen:      NewScreen(),
		Keyboard:    NewKeyboard(),
		Weight:      &repo.Laptop_WeightKg{WeightKg: randFloat(1, 2)},
		PriceUsd:    randFloat(1500, 4500),
		ReleaseYear: uint32(randInt(2020, 2023)),
		UpdateAt:    timestamppb.Now(),
	}
}
