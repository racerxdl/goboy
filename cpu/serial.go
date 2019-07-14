package cpu

import (
	"github.com/racerxdl/goboy/gameboy"
	"time"
)

type Serial struct {
	buff          uint8
	startTransfer bool
	clockSpeed    bool
	externalClock bool

	lastUpdate time.Time
	cpu        *Core
	cycledBits int
}

func MakeSerial(cpu *Core) *Serial {
	return &Serial{
		lastUpdate: time.Now(),
		cpu:        cpu,
		cycledBits: 0,
	}
}

func (s *Serial) Write(addr uint16, val uint8) {
	switch addr {
	case 0xFF01: // Serial Data Output
		s.buff = val
	case 0xFF02:
		s.startTransfer = val&0x80 > 0
		s.clockSpeed = val&0x02 > 0
		s.externalClock = val&0x01 > 0
	}
}

func (s *Serial) Read(addr uint16) byte {
	switch addr {
	case 0xFF01: // Serial Data Input
		return 0xFF
	case 0xFF02:
		v := uint8(0)
		if s.startTransfer {
			v |= 0x80
		}
		if s.clockSpeed {
			v |= 0x02
		}
		if s.externalClock {
			v |= 0x01
		}

		return v
	}

	return 0xFF
}

func (s *Serial) Cycle(clocks int) {
	if !s.startTransfer {
		return
	}
	/*
	     8192Hz -  1KB/s - Bit 1 cleared, Normal
	    16384Hz -  2KB/s - Bit 1 cleared, Double Speed Mode
	   262144Hz - 32KB/s - Bit 1 set,     Normal
	   524288Hz - 64KB/s - Bit 1 set,     Double Speed Mode
	*/
	var period time.Duration

	if s.cpu.colorMode {
		if s.clockSpeed {
			period = time.Second / 524288
		} else {
			period = time.Second / 16384
		}
	} else {
		if s.clockSpeed {
			period = time.Second / 262144
		} else {
			period = time.Second / 8192
		}
	}

	if time.Since(s.lastUpdate) > period {
		s.cycledBits++
		s.lastUpdate = time.Now()
	}

	if s.cycledBits == 8 {
		s.startTransfer = false
		s.cpu.Registers.InterruptsFired |= gameboy.IntSerial
		s.cycledBits = 0
	}
}
