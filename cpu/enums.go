package cpu

type RamSize byte
type RomSize byte
type GPUMode uint

const (
	RamNone    RamSize = 0
	Ram16kbit          = 1
	Ram64kbit          = 2
	Ram256kbit         = 3
	Ram1mbit           = 4
)

const (
	Rom256kbit RomSize = 0
	Rom512kbit         = 1
	Rom1mbit           = 2
	Rom2mbit           = 3
	Rom4mbit           = 4
	Rom8mbit           = 5
	Rom16mbit          = 6
	Rom9mbit           = 0x52
	Rom10mbit          = 0x53
	Rom12mbit          = 0x54
)

const (
	HBlank   GPUMode = 0
	VBlank           = 1
	OamRead          = 2
	VramRead         = 3
)
