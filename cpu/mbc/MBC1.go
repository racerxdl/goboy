package mbc

import (
	"bytes"
	"github.com/quan-to/slog"
	"github.com/racerxdl/goboy/gameboy"
	"math/rand"
)

var mbc1log = slog.Scope("MBC1")

// No MBC
type MBC1 struct {
	romBanks      [][0x8000]byte
	ramBanks      [][0x2000]byte
	activeRomBank int
	activeRamBank int
	ramEnabled    bool
	ramRomMode    bool
}

func MakeMBC1() *MBC1 {
	return &MBC1{
		romBanks:      make([][0x8000]byte, 1),
		ramBanks:      make([][0x2000]byte, 4),
		activeRomBank: 0,
		activeRamBank: 0,
		ramEnabled:    false,
		ramRomMode:    false,
	}
}

func (m *MBC1) Reset() {
	for i := 0; i < len(m.romBanks[0]); i++ {
		m.romBanks[0][i] = 0x00
	}

	for i := 0; i < len(m.ramBanks); i++ {
		for z := 0; z < len(m.ramBanks[i]); z++ {
			m.ramBanks[i][z] = 0x00
		}
	}
}

func (m *MBC1) Randomize() {
	for i := 0; i < len(m.romBanks[0]); i++ {
		m.romBanks[0][i] = byte(rand.Int31n(255))
	}

	for i := 0; i < len(m.ramBanks); i++ {
		for z := 0; z < len(m.ramBanks[i]); z++ {
			m.ramBanks[i][z] = byte(rand.Int31n(255))
		}
	}
}

func (m *MBC1) LoadRom(data []byte) {
	mbc1log.Debug("Loading Bank 0")
	copy(m.romBanks[0][:], data) // Copy first rom bank
	data = data[0x8000:]
	for len(data) > 0 {
		mbc1log.Debug("Loading Bank %d", len(m.romBanks))
		b := [0x8000]byte{}
		copy(b[:], data)
		m.romBanks = append(m.romBanks, b)
		data = data[0x8000:]
	}
}

func (m *MBC1) RomName() string {
	o := m.romBanks[0][0x134 : 0x134+0xE]
	n := bytes.Index(o, []byte{0x00})
	if n != -1 {
		o = o[:n]
	}
	return string(o)
}

func (m *MBC1) CatridgeRamSize() gameboy.RamSize {
	return gameboy.RamSize(m.romBanks[0][0x149])
}

func (m *MBC1) RomSize() gameboy.RomSize {
	return gameboy.RomSize(m.romBanks[0][0x148])
}

func (m *MBC1) MBCType() gameboy.MBCType {
	return gameboy.MBC1
}

func (m *MBC1) Read(addr uint16) uint8 {
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

func (m *MBC1) Write(addr uint16, val uint8) {
	switch {
	case addr < 0x2000: // Enable RAM
		m.ramEnabled = val&0xF == 0xA
		mbc1log.Debug("Changed Ram Enable to %v", m.ramEnabled)
	case addr >= 0x2000 && addr < 0x4000: // Select ROM Bank
		m.activeRomBank &= 0x60
		m.activeRamBank |= int(val & 0x1F)
		mbc1log.Debug("Changed Rom Bank to %d", m.activeRomBank)
	case addr >= 0x4000 && addr < 0x5FFF:
		if m.ramRomMode {
			m.activeRamBank = int(val & 0x3)
			mbc1log.Debug("Changed Ram Bank to %d", m.activeRamBank)
		} else {
			m.activeRomBank &= 0x1F
			m.activeRomBank |= int(val&0x3) << 5
			mbc1log.Debug("Changed Rom Bank to %d", m.activeRomBank)
		}
	case addr >= 0x6000 && addr <= 0x7FFF:
		m.ramRomMode = val > 0
		mbc1log.Debug("Changed Rom/Ram Mode %v", m.ramRomMode)
	case addr >= 0xA000 && addr <= 0xBFFF: // Catridge RAM
		if m.ramEnabled {
			m.ramBanks[m.activeRamBank][addr-0xA000] = val
		}
	}
}
