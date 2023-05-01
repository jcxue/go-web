package sample

import (
	"github.com/golang/protobuf/ptypes"
	api "github.com/jcxue/go-web/proto"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *api.Keyboard {
	keyboard := &api.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

// NewCPU returns a new sample CPU
func NewCPU() *api.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &api.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *api.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &api.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &api.Memory{
			Value: uint64(memGB),
			Unit:  api.Memory_GIGABYTE,
		},
	}

	return gpu
}

// NewRAM returns a new sample RAM
func NewRAM() *api.Memory {
	memGB := randomInt(4, 64)

	ram := &api.Memory{
		Value: uint64(memGB),
		Unit:  api.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new sample SSD
func NewSSD() *api.Storage {
	memGB := randomInt(128, 1024)

	ssd := &api.Storage{
		Driver: api.Storage_SSD,
		Memory: &api.Memory{
			Value: uint64(memGB),
			Unit:  api.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *api.Storage {
	memTB := randomInt(1, 6)

	hdd := &api.Storage{
		Driver: api.Storage_HDD,
		Memory: &api.Memory{
			Value: uint64(memTB),
			Unit:  api.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new sample Screen
func NewScreen() *api.Screen {
	screen := &api.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop returns a new sample Laptop
func NewLaptop() *api.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &api.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*api.GPU{NewGPU()},
		Storages: []*api.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &api.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}

	return laptop
}
