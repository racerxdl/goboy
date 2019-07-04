package cpu

type Timer struct {
	mainTime, subTime, divTime int
	divReg, tacReg             uint8
	timaReg, tmaReg            int

	cpu *Core
}

func MakeTimer(cpu *Core) *Timer {
	t := &Timer{
		cpu: cpu,
	}

	t.Reset()

	return t
}

func (t *Timer) Reset() {
	t.divReg = 0
	t.tmaReg = 0
	t.timaReg = 0
	t.tacReg = 0
	t.mainTime = 0
	t.subTime = 0
	t.divTime = 0
}

func (t *Timer) Cycle() {
	t.mainTime = 0
	t.timaReg++

	if t.timaReg > 0xFF {
		t.timaReg = t.tmaReg
		t.cpu.Registers.TriggerInterrupts |= IntTimer
	}

	t.timaReg &= 0xFF
}

func (t *Timer) Increment() {
	t.subTime += t.cpu.Registers.LastClockM

	if t.subTime >= 4 {
		t.mainTime++
		t.subTime -= 4
		t.divTime++

		if t.divTime == 16 {
			t.divReg++
			t.divTime = 0
		}
	}

	if (t.tacReg & 0x04) > 0 {
		switch t.tacReg & 0x03 {
		case 0:
			if t.mainTime >= 64 { // 4k
				t.Cycle()
			}
		case 1:
			if t.mainTime >= 1 { // 256k
				t.Cycle()
			}
		case 2:
			if t.mainTime >= 4 { // 64k
				t.Cycle()
			}
		case 3:
			if t.mainTime >= 16 { // 16k
				t.Cycle()
			}
		}
	}
}

func (t *Timer) Read(addr uint16) uint8 {
	switch addr {
	case 0xFF04:
		return uint8(t.divReg)
	case 0xFF05:
		return uint8(t.timaReg)
	case 0xFF06:
		return uint8(t.tmaReg)
	case 0xFF07:
		return uint8(t.tacReg)
	}

	return 0x00
}

func (t *Timer) Write(addr uint16, val uint8) {
	switch addr {
	case 0xFF04:
		t.divReg = 0
	case 0xFF05:
		t.timaReg = int(val)
	case 0xFF06:
		t.tmaReg = int(val)
	case 0xFF07:
		t.tacReg = uint8(val) & 0x07
		/*
		   Bit  2   - Timer Enable
		   Bits 1-0 - Input Clock Select
		              00: CPU Clock / 1024 (DMG, CGB:   4096 Hz, SGB:   ~4194 Hz)
		              01: CPU Clock / 16   (DMG, CGB: 262144 Hz, SGB: ~268400 Hz)
		              10: CPU Clock / 64   (DMG, CGB:  65536 Hz, SGB:  ~67110 Hz)
		              11: CPU Clock / 256  (DMG, CGB:  16384 Hz, SGB:  ~16780 Hz)

		   Note: The "Timer Enable" bit only affects the timer, the divider is ALWAYS counting.
		*/
	}
}
