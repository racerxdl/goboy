package cpu

import (
	"golang.org/x/image/colornames"
	"image/color"
)

// Timings
const (
	// Divide by 4 since we use Processor Cycles instead Clock Cycles

	horizontalBlankCycles = 207 / 4
	verticalBlankCycles   = 4560 / 4
	oamCycles             = 83 / 4
	vRamCycles            = 175 / 4
	CpuClock              = 4194304
	CpuPeriodMs           = 1000 / CpuClock
)

// Addresses
const (
	VRamBase       = 0x8000
	AddrIntVblank  = 0x40
	AddrIntLcdstat = 0x48
	AddrIntTimer   = 0x50
	AddrIntSerial  = 0x58
	AddrIntJoypad  = 0x60
)

// Default Paletes
var defaultBgPallete = []color.RGBA{
	colornames.Black,
	colornames.Darkgray,
	colornames.Gray,
	colornames.White,
}

var defaultObj0Pallete = []color.RGBA{
	colornames.Black,
	colornames.Darkgray,
	colornames.Gray,
	colornames.White,
}

var defaultObj1Pallete = []color.RGBA{
	colornames.Black,
	colornames.Darkgray,
	colornames.Gray,
	colornames.White,
}
