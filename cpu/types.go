package cpu

import "github.com/racerxdl/goboy/gameboy"

type GBInstruction func(*Core)

type MemoryInterface interface {
	Write(addr uint16, val uint8)
	Read(addr uint16) uint8
}

type Catridge interface {
	LoadRom([]byte)
	RomName() string
	CatridgeRamSize() gameboy.RamSize
	RomSize() gameboy.RomSize
	MBCType() gameboy.MBCType

	// Memory Interface
	Write(addr uint16, val uint8)
	Read(addr uint16) uint8

	// Tools
	Reset()
	Randomize()
	DumpRam() []byte
	LoadRam([]byte)
}
