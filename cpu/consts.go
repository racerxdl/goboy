package cpu

import (
	"golang.org/x/image/colornames"
	"image/color"
	"time"
)

// Timings
const (
	// Divide by 4 since we use Processor Cycles instead Clock Cycles

	horizontalBlankCycles = 207 / 4
	verticalBlankCycles   = 4560 / 4
	oamCycles             = 83 / 4
	vRamCycles            = 175 / 4
	Clock                 = 4194304
	ColorModeClock        = Clock * 2
	Period                = time.Second / Clock
	ColorModePeriod       = time.Second / ColorModeClock
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
