package cpu

type Timer struct {
	div, tma, tima, tac           int
	clockMain, clockSub, clockDiv int

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
	t.div = 0
	t.tma = 0
	t.tima = 0
	t.tac = 0
	t.clockMain = 0
	t.clockSub = 0
	t.clockDiv = 0
}

func (t *Timer) Cycle() {
	t.tima++
	t.clockMain = 0
	if t.tima > 255 {
		t.tima = t.tma
		t.cpu.Registers.TriggerInterrupts |= IntTimer
	}
}

func (t *Timer) Increment() {
	t.clockSub += t.cpu.Registers.LastClockM

	if t.clockSub > 3 {
		t.clockMain++
		t.clockSub -= 4
		t.clockDiv++

		if t.clockDiv == 16 {
			t.clockDiv = 0
			t.div++
			t.div %= 256
		}
	}

	if (t.tac & 0x04) != 0x00 {
		switch t.tac & 0x03 {
		case 0:
			if t.clockMain >= 64 {
				t.Cycle()
			}
		case 1:
			if t.clockMain >= 1 {
				t.Cycle()
			}
		case 2:
			if t.clockMain >= 4 {
				t.Cycle()
			}
		case 3:
			if t.clockMain >= 16 {
				t.Cycle()
			}
		}
	}
}

func (t *Timer) Read(addr uint16) uint8 {
	switch addr {
	case 0xFF04:
		return uint8(t.div)
	case 0xFF05:
		return uint8(t.tima)
	case 0xFF06:
		return uint8(t.tma)
	case 0xFF07:
		return uint8(t.tac)
	}

	return 0x00
}

func (t *Timer) Write(addr uint16, val uint8) {
	switch addr {
	case 0xFF04:
		t.div = 0
	case 0xFF05:
		t.tima = int(val)
	case 0xFF06:
		t.tma = int(val)
	case 0xFF07:
		t.tac = int(val & 0x07)
	}
}
