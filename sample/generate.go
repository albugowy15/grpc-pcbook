package sample

import (
	"github.com/albugowy15/pcbook/pb"
	"github.com/golang/protobuf/ptypes"
)

// NewKeyboard creates a new keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard {
		Layout: randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

// NewCPU creates a new CPU
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGHz := randomFloat64(2.0, 3.5)
	maxGHz := randomFloat64(minGHz, 5.0)

	cpu := &pb.CPU{
		Brand: brand,
		Name: name,
		NumberCores: uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz: minGHz,
		MaxGhz: maxGHz,
	}

	return cpu
}

// NewGPU creates a new GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit: pb.Memory_GIGABYTE,
	}
	gpu := &pb.GPU{
		Brand: brand,
		Name: name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

// NewRAM creates a new RAM
func NewRAM() *pb.Memory {
	memory := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit: pb.Memory_GIGABYTE,
	}

	return memory
}

// NewSSD creates a new SSD
func NewSSD() *pb.Storage {
	storage := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit: pb.Memory_GIGABYTE,
		},
	}

	return storage
}

// NewHDD creates a new HDD
func NewHDD() *pb.Storage {
	storage := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit: pb.Memory_GIGABYTE,
		},
	}

	return storage
}

// NewScreen creates a new screen
func NewScreen() *pb.Screen {
	
	screen := &pb.Screen{
		SizeInch: randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel: randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id: randomID(),
		Brand: brand,
		Name: name,
		Cpu: NewCPU(),
		Memory: NewRAM(),
		Gpus: []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen: NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd: randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2020)),
		UpdatedAt: ptypes.TimestampNow(),
	}

	return laptop
}
