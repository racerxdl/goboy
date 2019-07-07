package gameboy

import "github.com/quan-to/slog"

type RamSize byte
type RomSize byte
type MBCType byte
type GPUMode uint

const (
	RamNone    RamSize = 0
	Ram16kbit          = 1
	Ram64kbit          = 2
	Ram256kbit         = 3
	Ram1mbit           = 4
)

func (s RamSize) String() string {
	return [...]string{"None", "16 kbits", "64 kbits", "256 kbits", "1 mbit"}[s]
}

var NumRamBanks = map[RamSize]int{
	RamNone:    0,
	Ram16kbit:  1,
	Ram64kbit:  1,
	Ram256kbit: 4,
	Ram1mbit:   16,
}

const (
	Rom256kbit RomSize = 0
	Rom512kbit         = 1
	Rom1mbit           = 2
	Rom2mbit           = 3
	Rom4mbit           = 4
	Rom8mbit           = 5
	Rom9mbit           = 0x52
	Rom10mbit          = 0x53
	Rom12mbit          = 0x54
	Rom16mbit          = 6
	Rom32mbit          = 7
)

func (s RomSize) String() string {
	return [...]string{
		"256 kbits",
		"512 kbits",
		"1 mbit",
		"2 mbits",
		"4 mbits",
		"8 mbits",
		"9 mbits",
		"10 mbits",
		"12 mbits",
		"16 mbits",
		"32 mbits",
	}[s]
}

var NumRomBanks = map[RomSize]int{
	Rom256kbit: 2,
	Rom512kbit: 4,
	Rom1mbit:   8,
	Rom2mbit:   16,
	Rom4mbit:   32,
	Rom8mbit:   64,
	Rom9mbit:   72,
	Rom10mbit:  80,
	Rom12mbit:  96,
	Rom16mbit:  128,
	Rom32mbit:  256,
}

const (
	MBCNone MBCType = 0
	MBC1            = 1
	MBC2            = 2
	MBC3            = 3
	MBC4            = 4
	MBC5            = 5
)

func CatridgeTypeToMBC(ct uint8) MBCType {
	switch {
	case ct == 0:
		return MBCNone
	case ct > 0 && ct <= 0x3:
		return MBC1
	case ct == 0x5 || ct == 0x6:
		return MBC2
	case ct >= 0xF && ct <= 0x13:
		return MBC3
	case ct >= 0x15 && ct <= 0x17:
		return MBC4
	case ct >= 0x19 && ct <= 0x1E:
		return MBC5
	}

	slog.Error("Unknown MBC Type: %02X", ct)

	return MBCNone
}

func (s MBCType) String() string {
	return [...]string{
		"None",
		"MBC1",
		"MBC2",
		"MBC3",
		"MBC4",
		"MBC5",
	}[s]
}

const (
	HBlank   GPUMode = 0
	VBlank           = 1
	OamRead          = 2
	VramRead         = 3
)
