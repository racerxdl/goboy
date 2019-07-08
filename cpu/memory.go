package cpu

import (
	"github.com/faiface/pixel"
	"github.com/quan-to/slog"
	"github.com/racerxdl/goboy/cpu/mbc"
	"github.com/racerxdl/goboy/gameboy"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var memLog = slog.Scope("Memory")

var gbBios = []byte{
	0x31, 0xFE, 0xFF, 0xAF, 0x21, 0xFF, 0x9F, 0x32, 0xCB, 0x7C, 0x20, 0xFB, 0x21, 0x26, 0xFF, 0x0E,
	0x11, 0x3E, 0x80, 0x32, 0xE2, 0x0C, 0x3E, 0xF3, 0xE2, 0x32, 0x3E, 0x77, 0x77, 0x3E, 0xFC, 0xE0,
	0x47, 0x11, 0x04, 0x01, 0x21, 0x10, 0x80, 0x1A, 0xCD, 0x95, 0x00, 0xCD, 0x96, 0x00, 0x13, 0x7B,
	0xFE, 0x34, 0x20, 0xF3, 0x11, 0xD8, 0x00, 0x06, 0x08, 0x1A, 0x13, 0x22, 0x23, 0x05, 0x20, 0xF9,
	0x3E, 0x19, 0xEA, 0x10, 0x99, 0x21, 0x2F, 0x99, 0x0E, 0x0C, 0x3D, 0x28, 0x08, 0x32, 0x0D, 0x20,
	0xF9, 0x2E, 0x0F, 0x18, 0xF3, 0x67, 0x3E, 0x64, 0x57, 0xE0, 0x42, 0x3E, 0x91, 0xE0, 0x40, 0x04,
	0x1E, 0x02, 0x0E, 0x0C, 0xF0, 0x44, 0xFE, 0x90, 0x20, 0xFA, 0x0D, 0x20, 0xF7, 0x1D, 0x20, 0xF2,
	0x0E, 0x13, 0x24, 0x7C, 0x1E, 0x83, 0xFE, 0x62, 0x28, 0x06, 0x1E, 0xC1, 0xFE, 0x64, 0x20, 0x06,
	0x7B, 0xE2, 0x0C, 0x3E, 0x87, 0xF2, 0xF0, 0x42, 0x90, 0xE0, 0x42, 0x15, 0x20, 0xD2, 0x05, 0x20,
	0x4F, 0x16, 0x20, 0x18, 0xCB, 0x4F, 0x06, 0x04, 0xC5, 0xCB, 0x11, 0x17, 0xC1, 0xCB, 0x11, 0x17,
	0x05, 0x20, 0xF5, 0x22, 0x23, 0x22, 0x23, 0xC9, 0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B,
	0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D, 0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E,
	0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99, 0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC,
	0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E, 0x3c, 0x42, 0xB9, 0xA5, 0xB9, 0xA5, 0x42, 0x4C,
	0x21, 0x04, 0x01, 0x11, 0xA8, 0x00, 0x1A, 0x13, 0xBE, 0x20, 0xFE, 0x23, 0x7D, 0xFE, 0x34, 0x20,
	0xF5, 0x06, 0x19, 0x78, 0x86, 0x23, 0x05, 0x20, 0xFB, 0x86, 0x20, 0xFE, 0x3E, 0x01, 0xE0, 0x50,
}

type Memory struct {
	workRam      []byte
	highRam      []byte
	catridge     Catridge
	saveFilename string
	lastRamSave  time.Time

	inBIOS bool
	cpu    *Core
}

func MakeMemory(cpu *Core) *Memory {
	m := &Memory{
		workRam:     make([]byte, 0x2000),
		highRam:     make([]byte, 0x7F),
		cpu:         cpu,
		catridge:    mbc.MakeMBC0(), // Load Default Catridge
		lastRamSave: time.Now().Add(-time.Second * 3600),
	}

	m.Reset()

	return m
}

func (m *Memory) SetSaveFile(filename string) {
	m.saveFilename = filename
}

func (m *Memory) SaveCatridgeRAMData() {
	if time.Since(m.lastRamSave) > time.Second*5 {
		_ = ioutil.WriteFile(m.saveFilename, m.catridge.DumpRam(), os.ModePerm)
		m.lastRamSave = time.Now()
	}
}

func (m *Memory) LoadCatridgeRAMData() {
	d, err := ioutil.ReadFile(m.saveFilename)
	if err != nil {
		memLog.Error("Error loading battery at %s: %s", m.saveFilename, err)
		return
	}

	m.catridge.LoadRam(d)
}

func (m *Memory) Reset() {
	m.catridge.Reset()

	for i := 0; i < 0x2000; i++ {
		m.workRam[i] = 0x00
	}

	for i := 0; i < 0x7F; i++ {
		m.highRam[i] = 0x00
	}

	m.inBIOS = true
}

func (m *Memory) Randomize() {
	m.catridge.Randomize()

	for i := 0; i < 0x2000; i++ {
		m.workRam[i] = byte(rand.Int31n(255))
	}

	for i := 0; i < 0x7F; i++ {
		m.highRam[i] = byte(rand.Int31n(255))
	}
}

func (m *Memory) GetVideoSprite() *pixel.Sprite {
	return pixel.NewSprite(m.cpu.GPU.GetLCDBuffer(), m.cpu.GPU.GetLCDBuffer().Bounds())
}

func (m *Memory) GetVideoFrame() *pixel.PictureData {
	return m.cpu.GPU.GetLCDBuffer()
}

func (m *Memory) WriteByte(addr uint16, val byte) {
	switch {
	case addr < 0x3FFF && addr <= 0x7FFF: // Catridge Bank N
		m.catridge.Write(addr, val)
	case addr >= 0x8000 && addr <= 0x9FFF: // Video RAM
		m.cpu.GPU.Write(addr, val)
	case addr >= 0xA000 && addr <= 0xBFFF: // Catridge RAM
		m.catridge.Write(addr, val)
		m.SaveCatridgeRAMData()
	case addr >= 0xC000 && addr <= 0xEFFF: // Work Ram
		m.workRam[addr&0x1FFF] = val
	case addr >= 0xFE00 && addr <= 0xFE9F:
		m.cpu.GPU.Write(addr, val)
	case addr >= 0xFEA0 && addr <= 0xFEFF: // Not usable ... yet ...
		// Nothing
	case addr >= 0xFF00 && addr <= 0xFF7F: // I/O Ports
		switch addr {
		case 0xFF00:
			m.cpu.Keys.Write(addr, val)
		case 0xFF04, 0xFF05, 0xFF06, 0xFF07:
			m.cpu.Timer.Write(addr, val)
		case 0xFF0F:
			m.cpu.Registers.TriggerInterrupts = val
		}

		baseAddr := addr - 0xFF00

		switch baseAddr & 0x00F0 {
		case 0x00:
		case 0x10, 0x20:
			m.cpu.SoundCard.Write(addr, val)
		case 0x30:
			// TODO
		case 0x50:
			cpuLog.Info("Disabling Internal BIOS")
			m.inBIOS = false
		case 0x40, 0x60, 0x70:
			m.cpu.GPU.Write(addr, val)
		}
	case addr >= 0xFF80 && addr <= 0xFFFE:
		m.highRam[addr-0xFF80] = val
	case addr == 0xFFFF:
		m.cpu.Registers.EnabledInterrupts = val
	}
}

func (m *Memory) ReadByte(addr uint16) byte {
	return m.readByte(addr, false)
}

func (m *Memory) Read(addr uint16) byte {
	return m.ReadByte(addr)
}

func (m *Memory) ReadByteNoSideEffect(addr uint16) byte {
	return m.readByte(addr, true)
}

func (m *Memory) readByte(addr uint16, noSideEffects bool) byte {
	switch {
	case addr <= 0x3FFF:
		if m.inBIOS {
			if addr < 0x100 {
				return gbBios[addr]
			}
		}
		return m.catridge.Read(addr)
	case addr >= 0x4000 && addr <= 0x7FFF:
		return m.catridge.Read(addr)
	case addr >= 0x8000 && addr <= 0x9FFF:
		return m.cpu.GPU.Read(addr)
	case addr >= 0xA000 && addr <= 0xBFFF:
		return m.catridge.Read(addr)
	case addr >= 0xC000 && addr <= 0xEFFF:
		return m.workRam[addr&0x1FFF]
	case addr >= 0xFE00 && addr <= 0xFE9F:
		return m.cpu.GPU.Read(addr)
	case addr >= 0xFEA0 && addr <= 0xFEFF: // Not usable, ... yet ...
		// nothing
		return 0x00
	case addr >= 0xFF00 && addr <= 0xFF7F:
		switch addr {
		case 0xFF00:
			return m.cpu.Keys.Read(addr)
		case 0xFF04, 0xFF05, 0xFF06, 0xFF07:
			return m.cpu.Timer.Read(addr)
		case 0xFF0F:
			return m.cpu.Registers.TriggerInterrupts
		}

		switch addr & 0x00F0 {
		case 0x00:
			return 0x00
		case 0x10, 0x20:
			return m.cpu.SoundCard.Read(addr)
		case 0x30:
			return 0x00
		case 0x40, 0x50, 0x60, 0x70:
			return m.cpu.GPU.Read(addr)
		}
	case addr >= 0xFF80 && addr <= 0xFFFE:
		return m.highRam[addr-0xFF80]
	case addr == 0xFFFF:
		return m.cpu.Registers.EnabledInterrupts
	}

	return 0x00
}

func (m *Memory) ReadWord(addr uint16) uint16 {
	return uint16(m.ReadByte(addr+1))<<8 + uint16(m.ReadByte(addr))
}

func (m *Memory) ReadBytes(addr uint16, length int) []byte {
	b := make([]byte, length)

	for i := 0; i < length; i++ {
		b[i] = m.ReadByte(addr + uint16(i))
	}

	return b
}

func (m *Memory) WriteWord(addr uint16, val uint16) {
	b0 := byte(val >> 8)
	b1 := byte(val & 0xFF)

	m.WriteByte(addr, b1)
	m.WriteByte(addr+1, b0)
}

func (m *Memory) RomName() string {
	return m.catridge.RomName()
}

func (m *Memory) RomSize() gameboy.RomSize {
	return m.catridge.RomSize()
}

func (m *Memory) MBCType() gameboy.MBCType {
	return m.catridge.MBCType()
}

func (m *Memory) CatridgeRamSize() gameboy.RamSize {
	return m.catridge.CatridgeRamSize()
}

func (m *Memory) LoadRom(data []byte) {
	mbcType := gameboy.CatridgeTypeToMBC(data[0x147])

	switch mbcType {
	case gameboy.MBCNone:
		m.catridge = mbc.MakeMBC0()
	case gameboy.MBC1:
		m.catridge = mbc.MakeMBC1()
	case gameboy.MBC3:
		m.catridge = mbc.MakeMBC3()
	case gameboy.MBC5:
		m.catridge = mbc.MakeMBC5()
	default:
		memLog.Warn("MBC Type %s not implemented!", mbcType)
	}

	m.catridge.LoadRom(data)

	memLog.Debug("Loaded %s", m.RomName())
	memLog.Debug("MBC Type: %s", mbcType)
	memLog.Debug("Rom Size: %s", m.RomSize())
	memLog.Debug("Ram Size: %s", m.CatridgeRamSize())
}
