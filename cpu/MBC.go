package cpu

import (
	"bytes"
	"math/rand"
)

// No MBC
type MBC0 struct {
	romData []byte
	ramData []byte
}

func MakeMBC0() *MBC0 {
	return &MBC0{
		romData: make([]byte, 0x8000),
		ramData: make([]byte, 0x2000),
	}
}

func (m *MBC0) Reset() {
	for i := 0; i < 0x8000; i++ {
		m.romData[i] = 0x00
	}
	for i := 0; i < 0x2000; i++ {
		m.ramData[i] = 0x00
	}
}

func (m *MBC0) Randomize() {
	for i := 0; i < 0x7FFF; i++ {
		m.romData[i] = byte(rand.Int31n(255))
	}

	for i := 0; i < 0x2000; i++ {
		m.ramData[i] = byte(rand.Int31n(255))
	}
}

func (m *MBC0) LoadRom(data []byte) {
	copy(m.romData, data)
}

func (m *MBC0) RomName() string {
	o := m.romData[0x134 : 0x134+0xE]
	n := bytes.Index(o, []byte{0x00})
	if n != -1 {
		o = o[:n]
	}
	return string(o)
}

func (m *MBC0) CatridgeRamSize() RamSize {
	return RamSize(m.romData[0x149])
}

func (m *MBC0) RomSize() RomSize {
	return RomSize(m.romData[0x148])
}

func (m *MBC0) MBCType() MBCType {
	return MBCNone
}

func (m *MBC0) Read(addr uint16) uint8 {
	switch {
	case addr <= 0x3FFF:
		return m.romData[addr]
	case addr >= 0x4000 && addr <= 0x7FFF:
		return m.romData[addr]
	case addr >= 0xA000 && addr <= 0xBFFF:
		return m.ramData[addr-0xA000]
	}

	return 0x00
}

func (m *MBC0) Write(addr uint16, val uint8) {
	switch {
	case addr < 0x3FFF: // Catridge ROM
		// Do nothing, not allowed to write catridge
	case addr >= 0x4000 && addr <= 0x7FFF: // Catridge Bank N
		// Do nothing, not allowed to write catrige bank N
	case addr >= 0xA000 && addr <= 0xBFFF: // Catridge RAM
		m.ramData[addr-0xA000] = val
	}
}
