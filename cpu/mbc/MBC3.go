package mbc

import (
	"bytes"
	"github.com/quan-to/slog"
	"github.com/racerxdl/goboy/gameboy"
	"math/rand"
)

var mbc3log = slog.Scope("MBC3")

// No MBC
type MBC3 struct {
	romBanks      [][0x4000]byte
	ramBanks      [][0x2000]byte
	activeRomBank int
	activeRamBank int
	ramEnabled    bool
	ramRomMode    bool
}

func MakeMBC3() *MBC3 {
	return &MBC3{
		romBanks:      make([][0x4000]byte, 128),
		ramBanks:      make([][0x2000]byte, 4),
		activeRomBank: 1,
		activeRamBank: 0,
		ramEnabled:    false,
		ramRomMode:    false,
	}
}

func (m *MBC3) Reset() {
	m.activeRomBank = 1
	m.activeRamBank = 0
}

func (m *MBC3) Randomize() {
	for i := 0; i < len(m.romBanks[0]); i++ {
		m.romBanks[0][i] = byte(rand.Int31n(255))
	}

	for i := 0; i < len(m.ramBanks); i++ {
		for z := 0; z < len(m.ramBanks[i]); z++ {
			m.ramBanks[i][z] = byte(rand.Int31n(255))
		}
	}
}

func (m *MBC3) LoadRom(data []byte) {
	copy(m.romBanks[0][:], data) // Copy first rom bank
	data = data[0x4000:]
	n := 1

	for len(data) > 0 {
		copy(m.romBanks[n][:], data)
		data = data[0x4000:]
		n++
	}

	mbc3log.Debug("Loaded %d banks", n)
}

func (m *MBC3) RomName() string {
	o := m.romBanks[0][0x134 : 0x134+0xE]
	n := bytes.Index(o, []byte{0x00})
	if n != -1 {
		o = o[:n]
	}
	return string(o)
}

func (m *MBC3) CatridgeRamSize() gameboy.RamSize {
	return gameboy.RamSize(m.romBanks[0][0x149])
}

func (m *MBC3) RomSize() gameboy.RomSize {
	return gameboy.RomSize(m.romBanks[0][0x148])
}

func (m *MBC3) MBCType() gameboy.MBCType {
	return gameboy.MBC3
}

func (m *MBC3) Read(addr uint16) uint8 {
	switch {
	case addr <= 0x3FFF:
		return m.romBanks[0][addr]
	case addr >= 0x4000 && addr <= 0x7FFF:
		return m.romBanks[m.activeRomBank][addr&0x3FFF]
	case addr >= 0xA000 && addr <= 0xBFFF:
		if m.ramEnabled {
			return m.ramBanks[m.activeRamBank][addr-0xA000]
		} else {
			return 0xFF
		}
	}

	return 0x00
}

func (m *MBC3) Write(addr uint16, val uint8) {
	switch {
	case addr < 0x2000: // Enable RAM
		m.ramEnabled = val&0xF == 0xA
		//mbc3log.Debug("Changed Ram Enable to %v", m.ramEnabled)
	case addr >= 0x2000 && addr < 0x4000: // Select ROM Bank
		if val == 0 {
			val = 1
		}
		m.activeRomBank = int(val & 0x7F)
		//mbc3log.Debug("Changed Rom Bank to %d", m.activeRomBank)
	case addr >= 0x4000 && addr < 0x5FFF:
		m.activeRamBank = int(val & 0x3)
		//mbc3log.Debug("Changed Ram Bank to %d", m.activeRamBank)
	case addr >= 0x6000 && addr <= 0x7FFF:
		m.ramRomMode = val > 0
		//mbc3log.Debug("Changed Rom/Ram Mode %v", m.ramRomMode)
	case addr >= 0xA000 && addr <= 0xBFFF: // Catridge RAM
		if m.ramEnabled {
			m.ramBanks[m.activeRamBank][addr-0xA000] = val
		}
	}
}

func (m *MBC3) LoadRam(data []byte) {
	copy(m.ramBanks[0][:], data) // Copy first rom bank
	if len(data) < 0x2000 {
		return
	}
	data = data[0x2000:]
	n := 1

	for len(data) > 0 {
		copy(m.ramBanks[n][:], data)
		data = data[0x2000:]
		n++
	}

	mbc3log.Debug("Loaded %d ram banks", n)
}

func (m *MBC3) DumpRam() []byte {
	c := make([]byte, 0x2000*len(m.ramBanks))

	for i, v := range m.ramBanks {
		copy(c[0x2000*i:], v[:])
	}

	return c
}

func (m *MBC3) GBC() bool {
	return m.romBanks[0][0x143] == 0x80 || m.romBanks[0][0x143] == 0xC0
}
