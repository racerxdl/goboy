package cpu

type GBInstruction func(*Core)

type MemoryInterface interface {
	Write(addr uint16, val uint8)
	Read(addr uint16) uint8
}

type Catridge interface {
	LoadRom([]byte)
	RomName() string
	CatridgeRamSize() RamSize
	RomSize() RomSize
	MBCType() MBCType

	// Memory Interface
	Write(addr uint16, val uint8)
	Read(addr uint16) uint8

	// Tools
	Reset()
	Randomize()
}
