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

type Memory struct {
	workRam       []byte
	highRam       []byte
	ramBank       int
	catridge      Catridge
	saveFilename  string
	lastRamSave   time.Time
	doubleSpeed   bool
	inPrepareMode bool

	inBIOS bool
	cpu    *Core
}

func MakeMemory(cpu *Core) *Memory {
	m := &Memory{
		workRam:       make([]byte, 0x8000),
		highRam:       make([]byte, 0x7F),
		cpu:           cpu,
		ramBank:       1,
		catridge:      mbc.MakeMBC0(), // Load Default Catridge
		lastRamSave:   time.Now().Add(-time.Second * 3600),
		doubleSpeed:   false,
		inPrepareMode: false,
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

	for i := 0; i < 0x8000; i++ {
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

func (m *Memory) InDoubleSpeedMode() bool {
	return m.doubleSpeed
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
	case addr >= 0xC000 && addr <= 0xCFFF: // Work Ram Bank 0
		m.workRam[addr-0xC000] = val
	case addr >= 0xD000 && addr <= 0xDFFF: // Work Ram Bank 1 (or N in CGB)
		m.workRam[(addr-0xC000-0x2000)+uint16(m.ramBank)*0x2000] = val
	case addr >= 0xE000 && addr <= 0xFDFF: // Mirror Bank 0
		m.workRam[(addr-0xC000-0x2000)&0x1FFF] = val
	case addr >= 0xFE00 && addr <= 0xFE9F:
		m.cpu.GPU.Write(addr, val)
	case addr >= 0xFEA0 && addr <= 0xFEFF: // Not usable ... yet ...
		// Nothing
	case addr >= 0xFF00 && addr <= 0xFF7F: // I/O Ports
		switch addr {
		case 0xFF00:
			m.cpu.Keys.Write(addr, val)
			return
		case 0xFF01, 0xFF02:
			m.cpu.Serial.Write(addr, val)
			return
		case 0xFF04, 0xFF05, 0xFF06, 0xFF07:
			m.cpu.Timer.Write(addr, val)
			return
		case 0xFF0F:
			m.cpu.Registers.InterruptsFired = val
			return
		case 0xFF4D: // Prepare speed
			if m.catridge.GBC() {
				m.inPrepareMode = val&1 > 0
				memLog.Debug("Prepare for Double Speed")
				return
			}
		case 0xFF70:
			if m.catridge.GBC() {
				m.ramBank = int(val & 0x3)
				if m.ramBank == 0 {
					m.ramBank = 1
				}
				return
			}
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
			//m.cpu.Registers.A = 0x11 // CGB
		case 0x40, 0x60:
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

func (m *Memory) ReadByteForPC(addr uint16) byte {
	return m.readByte(addr, true)
}

func (m *Memory) Read(addr uint16) byte {
	return m.ReadByte(addr)
}

func (m *Memory) readByte(addr uint16, readForPC bool) byte {
	switch {
	case addr <= 0x3FFF:
		if m.inBIOS {
			//if m.catridge.GBC() {
			//    if int(addr) < len(gbcBios) {
			//        return gbcBios[addr]
			//    }
			//}

			if int(addr) < len(gbBios) {
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

	case addr >= 0xC000 && addr <= 0xCFFF: // Work Ram Bank 0
		return m.workRam[addr-0xC000]
	case addr >= 0xD000 && addr <= 0xDFFF: // Work Ram Bank 1 (or N in CGB)
		return m.workRam[addr-0xC000-0x2000+uint16(m.ramBank)*0x2000]
	case addr >= 0xE000 && addr <= 0xFDFF: // Mirror Bank 0
		return m.workRam[addr-0xC000-0x2000]

	case addr >= 0xFE00 && addr <= 0xFE9F:
		return m.cpu.GPU.Read(addr)
	case addr >= 0xFEA0 && addr <= 0xFEFF: // Not usable, ... yet ...
		// nothing
		return 0x00
	case addr >= 0xFF00 && addr <= 0xFF7F:
		switch addr {
		case 0xFF00:
			return m.cpu.Keys.Read(addr)
		case 0xFF01, 0xFF02:
			return m.cpu.Serial.Read(addr)
		case 0xFF04, 0xFF05, 0xFF06, 0xFF07:
			return m.cpu.Timer.Read(addr)
		case 0xFF0F:
			return m.cpu.Registers.InterruptsFired
		case 0xFF4D: // Prepare speed
			if m.catridge.GBC() {
				v := uint8(0)
				if m.inPrepareMode {
					v |= 1
				}
				if m.InDoubleSpeedMode() {
					v |= 7
				}
				return v
			}
			return 0x00
		case 0xFF70:
			if m.catridge.GBC() {
				return uint8(m.ramBank & 0x3)
			}
		}

		switch addr & 0x00F0 {
		case 0x00:
			return 0x00
		case 0x10, 0x20:
			return m.cpu.SoundCard.Read(addr)
		case 0x30:
			return 0x00
		case 0x40:
			return m.cpu.GPU.Read(addr)
		case 0x50, 0x60, 0x70:
			return 0xFF
		}
	case addr >= 0xFF80 && addr <= 0xFFFE:
		return m.highRam[addr-0xFF80]
	case addr == 0xFFFF:
		return m.cpu.Registers.EnabledInterrupts
	}

	return 0xFF
}

func (m *Memory) ReadWord(addr uint16) uint16 {
	return uint16(m.ReadByte(addr+1))<<8 + uint16(m.ReadByte(addr))
}

func (m *Memory) ReadWordForPC(addr uint16) uint16 {
	return uint16(m.ReadByteForPC(addr+1))<<8 + uint16(m.ReadByteForPC(addr))
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
	memLog.Debug("Is GBC Rom: %v", m.catridge.GBC())
}
