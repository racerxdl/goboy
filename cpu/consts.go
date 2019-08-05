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
	colornames.White,
	colornames.Gray,
	colornames.Darkgray,
	colornames.Black,
}

// For CGB Only
var defaultObjPallete = []color.RGBA{
	colornames.White,
	colornames.White,
	colornames.White,
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

// Color map for 5 bit color
var colorMap = []uint8{
	0x0,
	0x8,
	0x10,
	0x18,
	0x20,
	0x29,
	0x31,
	0x39,
	0x41,
	0x4a,
	0x52,
	0x5a,
	0x62,
	0x6a,
	0x73,
	0x7b,
	0x83,
	0x8b,
	0x94,
	0x9c,
	0xa4,
	0xac,
	0xb4,
	0xbd,
	0xc5,
	0xcd,
	0xd5,
	0xde,
	0xe6,
	0xee,
	0xf6,
	0xff,
}
