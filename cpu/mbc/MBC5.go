package mbc

import (
	"bytes"
	"github.com/quan-to/slog"
	"github.com/racerxdl/goboy/gameboy"
	"math/rand"
)

var mbc5log = slog.Scope("MBC5")

// No MBC
type MBC5 struct {
	romBanks      [][0x4000]byte
	ramBanks      [][0x2000]byte
	activeRomBank int
	activeRamBank int
	ramEnabled    bool
}

func MakeMBC5() *MBC5 {
	return &MBC5{
		romBanks:      make([][0x4000]byte, 512),
		ramBanks:      make([][0x2000]byte, 16),
		activeRomBank: 0,
		activeRamBank: 0,
		ramEnabled:    false,
	}
}

func (m *MBC5) Reset() {
	m.activeRomBank = 0
	m.activeRamBank = 0
}

func (m *MBC5) Randomize() {
	for i := 0; i < len(m.romBanks[0]); i++ {
		m.romBanks[0][i] = byte(rand.Int31n(255))
	}

	for i := 0; i < len(m.ramBanks); i++ {
		for z := 0; z < len(m.ramBanks[i]); z++ {
			m.ramBanks[i][z] = byte(rand.Int31n(255))
		}
	}
}

func (m *MBC5) LoadRom(data []byte) {
	copy(m.romBanks[0][:], data) // Copy first rom bank
	data = data[0x4000:]
	n := 1

	for len(data) > 0 {
		copy(m.romBanks[n][:], data)
		data = data[0x4000:]
		n++
	}

	mbc5log.Debug("Loaded %d banks", n)
}

func (m *MBC5) RomName() string {
	o := m.romBanks[0][0x134 : 0x134+0xE]
	n := bytes.Index(o, []byte{0x00})
	if n != -1 {
		o = o[:n]
	}
	return string(o)
}

func (m *MBC5) CatridgeRamSize() gameboy.RamSize {
	return gameboy.RamSize(m.romBanks[0][0x149])
}

func (m *MBC5) RomSize() gameboy.RomSize {
	return gameboy.RomSize(m.romBanks[0][0x148])
}

func (m *MBC5) MBCType() gameboy.MBCType {
	return gameboy.MBC5
}

func (m *MBC5) Read(addr uint16) uint8 {
	switch {
	case addr <= 0x3FFF:
		return m.romBanks[0][addr]
	case addr >= 0x4000 && addr <= 0x7FFF:
		return m.romBanks[m.activeRomBank][addr&0x3FFF]
	case addr >= 0xA000 && addr <= 0xBFFF:
		return m.ramBanks[m.activeRamBank][addr-0xA000]
	}

	return 0x00
}

func (m *MBC5) Write(addr uint16, val uint8) {
	switch {
	case addr < 0x2000: // Enable RAM		// RAM enable
		m.ramEnabled = val&0xA == 0xA
		//mbc5log.Debug("Changed Ram Enable to %v", m.ramEnabled)
	case addr >= 0x2000 && addr < 0x3000: // ROM Bank Low Bits
		//m.activeRomBank &= 0x100
		m.activeRomBank = (m.activeRomBank & 0x100) | int(val)
		//mbc5log.Debug("Changed Rom Bank to %d", m.activeRomBank)
	case addr >= 0x3000 && addr < 0x4000: // ROM Bank High Bits
		m.activeRomBank = (m.activeRomBank & 0xFF) | (int(val&1) << 8)
		//mbc5log.Debug("Changed Rom Bank to %d", m.activeRomBank)
	case addr >= 0x4000 && addr <= 0x5FFF:
		m.activeRamBank = int(val & 0xF)
		//mbc5log.Debug("Changed Ram Bank to %d", m.activeRamBank)
	case addr >= 0xA000 && addr <= 0xBFFF: // Catridge RAM
		if m.ramEnabled {
			m.ramBanks[m.activeRamBank][addr-0xA000] = val
		}
	}
}

func (m *MBC5) LoadRam(data []byte) {
	copy(m.ramBanks[0][:], data) // Copy first rom bank
	data = data[0x2000:]
	n := 1

	for len(data) > 0 {
		copy(m.ramBanks[n][:], data)
		data = data[0x2000:]
		n++
	}

	mbc5log.Debug("Loaded %d ram banks", n)
}

func (m *MBC5) DumpRam() []byte {
	c := make([]byte, 0x2000*len(m.ramBanks))

	for i, v := range m.ramBanks {
		copy(c[0x2000*i:], v[:])
	}

	return c
}

func (m *MBC5) GBC() bool {
	return m.romBanks[0][0x143] == 0x80 || m.romBanks[0][0x143] == 0xC0
}
