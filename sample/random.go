package sample

import (
	"math/rand"
	"time"

	"github.com/albugowy15/pcbook/pb"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomID() string {
	return uuid.New().String()
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
		case 1:
			return pb.Keyboard_QWERTY	
		case 2:
			return pb.Keyboard_AZERTY
		default:
			return pb.Keyboard_QWERTZ
	}

}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Width: uint32(width),
		Height: uint32(height),
	}

	return resolution
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo", "Asus")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Pentium Gold",
			"Core 19-1245P",
			"Core 15-1240H",
			"Core 13-1156G7",
			"Core 17-12450K",
		)
	}

	return randomStringFromSet(
		"Athlon Gold",
		"Ryzen 5 3600",
		"Ryzen 3 3200G",
		"Ryzen 7 3700X",
		"Ryzen 9 4900H",
	)
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 3060",
			"GTX 1660",
			"GTX 1650",
		)
	}

	return randomStringFromSet(
		"RX 5500",
		"RX 5600",
		"RX 5700",
		"RX 5800",
	)
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air M1", "Macbook Pro M2", "Macbook M2 Max")
		case "Dell": 
			return randomStringFromSet("XPS 13", "XPS 15", "XPS 17")
		case "Lenovo":
			return randomStringFromSet("Thinkpad X1", "Thinkpad X1 Extreme", "Thinkpad X1 Yoga")
		case "Asus":
			return randomStringFromSet("Zenbook 13", "Zenbook 14", "Zenbook 15")
		default:
			return randomStringFromSet("Laptop 1", "Laptop 2", "Laptop 3")
	}
} 

func randomStringFromSet(set ...string) string {
	n := len(set)
	if n == 0 {
		return ""
	}
	return set[rand.Intn(2)]
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}