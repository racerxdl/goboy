package cpu

func gbCBCall(cpu *Core) {
	ins := cpu.Memory.ReadByte(cpu.Registers.PC)
	cpu.Registers.PC++
	CBInstructions[ins](cpu)
}

// cbRLrA Rotates A to the left
func cbRLrA(cpu *Core) {
	v := cpu.Registers.A
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v << 1) | f
	cpu.Registers.A = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRLCrA Rotates A to the left with carry
func cbRLCrA(cpu *Core) {
	v := cpu.Registers.A
	c := v >> 7

	v = (v << 1) | c
	cpu.Registers.A = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRrA Rotates A to the right
func cbRRrA(cpu *Core) {
	v := cpu.Registers.A
	c := v & 1
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v >> 1) | (f << 7)
	cpu.Registers.A = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRCrA Rotates A to the right
func cbRRCrA(cpu *Core) {
	v := cpu.Registers.A
	c := v & 1

	v = (v >> 1) | (c << 7)
	cpu.Registers.A = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSLArA Shifts A to the left
func cbSLArA(cpu *Core) {
	v := cpu.Registers.A
	c := (v >> 7) > 0

	v = (v << 1)
	cpu.Registers.A = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSRArA Shifts A to the right
func cbSRArA(cpu *Core) {
	v := cpu.Registers.A
	c := (v & 1) > 0
	e := v & 0x80

	v = (v >> 1) | e
	cpu.Registers.A = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSWAPrA Swaps nibbles in register
func cbSWAPrA(cpu *Core) {
	v := cpu.Registers.A
	v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
	cpu.Registers.A = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(false)

	cpu.Registers.LastClockM = 1
	cpu.Registers.LastClockT = 4
}

// cbSRLrA Shift A right
func cbSRLrA(cpu *Core) {
	v := cpu.Registers.A
	c := (v & 1) > 0

	v = (v >> 1)
	cpu.Registers.A = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT0A Sets Flag Zero to BIT 0 from A
func cbBIT0A(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.A&(1<<0) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES0A Resets BIT 0 from A
func cbRES0A(cpu *Core) {
	cpu.Registers.A &= ^(uint8(1) << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET0A Sets BIT 0 from A
func cbSET0A(cpu *Core) {
	cpu.Registers.A |= (1 << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT1A Sets Flag Zero to BIT 1 from A
func cbBIT1A(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.A&(1<<1) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES1A Resets BIT 1 from A
func cbRES1A(cpu *Core) {
	cpu.Registers.A &= ^(uint8(1) << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET1A Sets BIT 1 from A
func cbSET1A(cpu *Core) {
	cpu.Registers.A |= (1 << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT2A Sets Flag Zero to BIT 2 from A
func cbBIT2A(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.A&(1<<2) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES2A Resets BIT 2 from A
func cbRES2A(cpu *Core) {
	cpu.Registers.A &= ^(uint8(1) << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET2A Sets BIT 2 from A
func cbSET2A(cpu *Core) {
	cpu.Registers.A |= (1 << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT3A Sets Flag Zero to BIT 3 from A
func cbBIT3A(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.A&(1<<3) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES3A Resets BIT 3 from A
func cbRES3A(cpu *Core) {
	cpu.Registers.A &= ^(uint8(1) << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET3A Sets BIT 3 from A
func cbSET3A(cpu *Core) {
	cpu.Registers.A |= (1 << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT4A Sets Flag Zero to BIT 4 from A
func cbBIT4A(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.A&(1<<4) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES4A Resets BIT 4 from A
func cbRES4A(cpu *Core) {
	cpu.Registers.A &= ^(uint8(1) << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET4A Sets BIT 4 from A
func cbSET4A(cpu *Core) {
	cpu.Registers.A |= (1 << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT5A Sets Flag Zero to BIT 5 from A
func cbBIT5A(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.A&(1<<5) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES5A Resets BIT 5 from A
func cbRES5A(cpu *Core) {
	cpu.Registers.A &= ^(uint8(1) << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET5A Sets BIT 5 from A
func cbSET5A(cpu *Core) {
	cpu.Registers.A |= (1 << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT6A Sets Flag Zero to BIT 6 from A
func cbBIT6A(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.A&(1<<6) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES6A Resets BIT 6 from A
func cbRES6A(cpu *Core) {
	cpu.Registers.A &= ^(uint8(1) << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET6A Sets BIT 6 from A
func cbSET6A(cpu *Core) {
	cpu.Registers.A |= (1 << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT7A Sets Flag Zero to BIT 7 from A
func cbBIT7A(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.A&(1<<7) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES7A Resets BIT 7 from A
func cbRES7A(cpu *Core) {
	cpu.Registers.A &= ^(uint8(1) << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET7A Sets BIT 7 from A
func cbSET7A(cpu *Core) {
	cpu.Registers.A |= (1 << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbRLrB Rotates B to the left
func cbRLrB(cpu *Core) {
	v := cpu.Registers.B
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v << 1) | f
	cpu.Registers.B = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRLCrB Rotates B to the left with carry
func cbRLCrB(cpu *Core) {
	v := cpu.Registers.B
	c := v >> 7

	v = (v << 1) | c
	cpu.Registers.B = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRrB Rotates B to the right
func cbRRrB(cpu *Core) {
	v := cpu.Registers.B
	c := v & 1
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v >> 1) | (f << 7)
	cpu.Registers.B = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRCrB Rotates B to the right
func cbRRCrB(cpu *Core) {
	v := cpu.Registers.B
	c := v & 1

	v = (v >> 1) | (c << 7)
	cpu.Registers.B = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSLArB Shifts B to the left
func cbSLArB(cpu *Core) {
	v := cpu.Registers.B
	c := (v >> 7) > 0

	v = (v << 1)
	cpu.Registers.B = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSRArB Shifts B to the right
func cbSRArB(cpu *Core) {
	v := cpu.Registers.B
	c := (v & 1) > 0
	e := v & 0x80

	v = (v >> 1) | e
	cpu.Registers.B = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSWAPrB Swaps nibbles in register
func cbSWAPrB(cpu *Core) {
	v := cpu.Registers.B
	v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
	cpu.Registers.B = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(false)

	cpu.Registers.LastClockM = 1
	cpu.Registers.LastClockT = 4
}

// cbSRLrB Shift B right
func cbSRLrB(cpu *Core) {
	v := cpu.Registers.B
	c := (v & 1) > 0

	v = (v >> 1)
	cpu.Registers.B = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT0B Sets Flag Zero to BIT 0 from B
func cbBIT0B(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.B&(1<<0) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES0B Resets BIT 0 from B
func cbRES0B(cpu *Core) {
	cpu.Registers.B &= ^(uint8(1) << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET0B Sets BIT 0 from B
func cbSET0B(cpu *Core) {
	cpu.Registers.B |= (1 << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT1B Sets Flag Zero to BIT 1 from B
func cbBIT1B(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.B&(1<<1) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES1B Resets BIT 1 from B
func cbRES1B(cpu *Core) {
	cpu.Registers.B &= ^(uint8(1) << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET1B Sets BIT 1 from B
func cbSET1B(cpu *Core) {
	cpu.Registers.B |= (1 << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT2B Sets Flag Zero to BIT 2 from B
func cbBIT2B(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.B&(1<<2) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES2B Resets BIT 2 from B
func cbRES2B(cpu *Core) {
	cpu.Registers.B &= ^(uint8(1) << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET2B Sets BIT 2 from B
func cbSET2B(cpu *Core) {
	cpu.Registers.B |= (1 << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT3B Sets Flag Zero to BIT 3 from B
func cbBIT3B(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.B&(1<<3) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES3B Resets BIT 3 from B
func cbRES3B(cpu *Core) {
	cpu.Registers.B &= ^(uint8(1) << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET3B Sets BIT 3 from B
func cbSET3B(cpu *Core) {
	cpu.Registers.B |= (1 << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT4B Sets Flag Zero to BIT 4 from B
func cbBIT4B(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.B&(1<<4) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES4B Resets BIT 4 from B
func cbRES4B(cpu *Core) {
	cpu.Registers.B &= ^(uint8(1) << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET4B Sets BIT 4 from B
func cbSET4B(cpu *Core) {
	cpu.Registers.B |= (1 << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT5B Sets Flag Zero to BIT 5 from B
func cbBIT5B(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.B&(1<<5) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES5B Resets BIT 5 from B
func cbRES5B(cpu *Core) {
	cpu.Registers.B &= ^(uint8(1) << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET5B Sets BIT 5 from B
func cbSET5B(cpu *Core) {
	cpu.Registers.B |= (1 << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT6B Sets Flag Zero to BIT 6 from B
func cbBIT6B(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.B&(1<<6) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES6B Resets BIT 6 from B
func cbRES6B(cpu *Core) {
	cpu.Registers.B &= ^(uint8(1) << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET6B Sets BIT 6 from B
func cbSET6B(cpu *Core) {
	cpu.Registers.B |= (1 << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT7B Sets Flag Zero to BIT 7 from B
func cbBIT7B(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.B&(1<<7) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES7B Resets BIT 7 from B
func cbRES7B(cpu *Core) {
	cpu.Registers.B &= ^(uint8(1) << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET7B Sets BIT 7 from B
func cbSET7B(cpu *Core) {
	cpu.Registers.B |= (1 << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbRLrC Rotates C to the left
func cbRLrC(cpu *Core) {
	v := cpu.Registers.C
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v << 1) | f
	cpu.Registers.C = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRLCrC Rotates C to the left with carry
func cbRLCrC(cpu *Core) {
	v := cpu.Registers.C
	c := v >> 7

	v = (v << 1) | c
	cpu.Registers.C = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRrC Rotates C to the right
func cbRRrC(cpu *Core) {
	v := cpu.Registers.C
	c := v & 1
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v >> 1) | (f << 7)
	cpu.Registers.C = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRCrC Rotates C to the right
func cbRRCrC(cpu *Core) {
	v := cpu.Registers.C
	c := v & 1

	v = (v >> 1) | (c << 7)
	cpu.Registers.C = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSLArC Shifts C to the left
func cbSLArC(cpu *Core) {
	v := cpu.Registers.C
	c := (v >> 7) > 0

	v = (v << 1)
	cpu.Registers.C = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSRArC Shifts C to the right
func cbSRArC(cpu *Core) {
	v := cpu.Registers.C
	c := (v & 1) > 0
	e := v & 0x80

	v = (v >> 1) | e
	cpu.Registers.C = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSWAPrC Swaps nibbles in register
func cbSWAPrC(cpu *Core) {
	v := cpu.Registers.C
	v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
	cpu.Registers.C = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(false)

	cpu.Registers.LastClockM = 1
	cpu.Registers.LastClockT = 4
}

// cbSRLrC Shift C right
func cbSRLrC(cpu *Core) {
	v := cpu.Registers.C
	c := (v & 1) > 0

	v = (v >> 1)
	cpu.Registers.C = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT0C Sets Flag Zero to BIT 0 from C
func cbBIT0C(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.C&(1<<0) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES0C Resets BIT 0 from C
func cbRES0C(cpu *Core) {
	cpu.Registers.C &= ^(uint8(1) << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET0C Sets BIT 0 from C
func cbSET0C(cpu *Core) {
	cpu.Registers.C |= (1 << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT1C Sets Flag Zero to BIT 1 from C
func cbBIT1C(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.C&(1<<1) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES1C Resets BIT 1 from C
func cbRES1C(cpu *Core) {
	cpu.Registers.C &= ^(uint8(1) << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET1C Sets BIT 1 from C
func cbSET1C(cpu *Core) {
	cpu.Registers.C |= (1 << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT2C Sets Flag Zero to BIT 2 from C
func cbBIT2C(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.C&(1<<2) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES2C Resets BIT 2 from C
func cbRES2C(cpu *Core) {
	cpu.Registers.C &= ^(uint8(1) << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET2C Sets BIT 2 from C
func cbSET2C(cpu *Core) {
	cpu.Registers.C |= (1 << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT3C Sets Flag Zero to BIT 3 from C
func cbBIT3C(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.C&(1<<3) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES3C Resets BIT 3 from C
func cbRES3C(cpu *Core) {
	cpu.Registers.C &= ^(uint8(1) << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET3C Sets BIT 3 from C
func cbSET3C(cpu *Core) {
	cpu.Registers.C |= (1 << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT4C Sets Flag Zero to BIT 4 from C
func cbBIT4C(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.C&(1<<4) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES4C Resets BIT 4 from C
func cbRES4C(cpu *Core) {
	cpu.Registers.C &= ^(uint8(1) << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET4C Sets BIT 4 from C
func cbSET4C(cpu *Core) {
	cpu.Registers.C |= (1 << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT5C Sets Flag Zero to BIT 5 from C
func cbBIT5C(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.C&(1<<5) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES5C Resets BIT 5 from C
func cbRES5C(cpu *Core) {
	cpu.Registers.C &= ^(uint8(1) << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET5C Sets BIT 5 from C
func cbSET5C(cpu *Core) {
	cpu.Registers.C |= (1 << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT6C Sets Flag Zero to BIT 6 from C
func cbBIT6C(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.C&(1<<6) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES6C Resets BIT 6 from C
func cbRES6C(cpu *Core) {
	cpu.Registers.C &= ^(uint8(1) << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET6C Sets BIT 6 from C
func cbSET6C(cpu *Core) {
	cpu.Registers.C |= (1 << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT7C Sets Flag Zero to BIT 7 from C
func cbBIT7C(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.C&(1<<7) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES7C Resets BIT 7 from C
func cbRES7C(cpu *Core) {
	cpu.Registers.C &= ^(uint8(1) << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET7C Sets BIT 7 from C
func cbSET7C(cpu *Core) {
	cpu.Registers.C |= (1 << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbRLrD Rotates D to the left
func cbRLrD(cpu *Core) {
	v := cpu.Registers.D
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v << 1) | f
	cpu.Registers.D = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRLCrD Rotates D to the left with carry
func cbRLCrD(cpu *Core) {
	v := cpu.Registers.D
	c := v >> 7

	v = (v << 1) | c
	cpu.Registers.D = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRrD Rotates D to the right
func cbRRrD(cpu *Core) {
	v := cpu.Registers.D
	c := v & 1
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v >> 1) | (f << 7)
	cpu.Registers.D = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRCrD Rotates D to the right
func cbRRCrD(cpu *Core) {
	v := cpu.Registers.D
	c := v & 1

	v = (v >> 1) | (c << 7)
	cpu.Registers.D = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSLArD Shifts D to the left
func cbSLArD(cpu *Core) {
	v := cpu.Registers.D
	c := (v >> 7) > 0

	v = (v << 1)
	cpu.Registers.D = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSRArD Shifts D to the right
func cbSRArD(cpu *Core) {
	v := cpu.Registers.D
	c := (v & 1) > 0
	e := v & 0x80

	v = (v >> 1) | e
	cpu.Registers.D = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSWAPrD Swaps nibbles in register
func cbSWAPrD(cpu *Core) {
	v := cpu.Registers.D
	v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
	cpu.Registers.D = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(false)

	cpu.Registers.LastClockM = 1
	cpu.Registers.LastClockT = 4
}

// cbSRLrD Shift D right
func cbSRLrD(cpu *Core) {
	v := cpu.Registers.D
	c := (v & 1) > 0

	v = (v >> 1)
	cpu.Registers.D = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT0D Sets Flag Zero to BIT 0 from D
func cbBIT0D(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.D&(1<<0) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES0D Resets BIT 0 from D
func cbRES0D(cpu *Core) {
	cpu.Registers.D &= ^(uint8(1) << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET0D Sets BIT 0 from D
func cbSET0D(cpu *Core) {
	cpu.Registers.D |= (1 << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT1D Sets Flag Zero to BIT 1 from D
func cbBIT1D(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.D&(1<<1) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES1D Resets BIT 1 from D
func cbRES1D(cpu *Core) {
	cpu.Registers.D &= ^(uint8(1) << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET1D Sets BIT 1 from D
func cbSET1D(cpu *Core) {
	cpu.Registers.D |= (1 << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT2D Sets Flag Zero to BIT 2 from D
func cbBIT2D(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.D&(1<<2) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES2D Resets BIT 2 from D
func cbRES2D(cpu *Core) {
	cpu.Registers.D &= ^(uint8(1) << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET2D Sets BIT 2 from D
func cbSET2D(cpu *Core) {
	cpu.Registers.D |= (1 << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT3D Sets Flag Zero to BIT 3 from D
func cbBIT3D(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.D&(1<<3) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES3D Resets BIT 3 from D
func cbRES3D(cpu *Core) {
	cpu.Registers.D &= ^(uint8(1) << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET3D Sets BIT 3 from D
func cbSET3D(cpu *Core) {
	cpu.Registers.D |= (1 << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT4D Sets Flag Zero to BIT 4 from D
func cbBIT4D(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.D&(1<<4) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES4D Resets BIT 4 from D
func cbRES4D(cpu *Core) {
	cpu.Registers.D &= ^(uint8(1) << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET4D Sets BIT 4 from D
func cbSET4D(cpu *Core) {
	cpu.Registers.D |= (1 << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT5D Sets Flag Zero to BIT 5 from D
func cbBIT5D(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.D&(1<<5) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES5D Resets BIT 5 from D
func cbRES5D(cpu *Core) {
	cpu.Registers.D &= ^(uint8(1) << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET5D Sets BIT 5 from D
func cbSET5D(cpu *Core) {
	cpu.Registers.D |= (1 << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT6D Sets Flag Zero to BIT 6 from D
func cbBIT6D(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.D&(1<<6) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES6D Resets BIT 6 from D
func cbRES6D(cpu *Core) {
	cpu.Registers.D &= ^(uint8(1) << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET6D Sets BIT 6 from D
func cbSET6D(cpu *Core) {
	cpu.Registers.D |= (1 << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT7D Sets Flag Zero to BIT 7 from D
func cbBIT7D(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.D&(1<<7) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES7D Resets BIT 7 from D
func cbRES7D(cpu *Core) {
	cpu.Registers.D &= ^(uint8(1) << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET7D Sets BIT 7 from D
func cbSET7D(cpu *Core) {
	cpu.Registers.D |= (1 << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbRLrE Rotates E to the left
func cbRLrE(cpu *Core) {
	v := cpu.Registers.E
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v << 1) | f
	cpu.Registers.E = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRLCrE Rotates E to the left with carry
func cbRLCrE(cpu *Core) {
	v := cpu.Registers.E
	c := v >> 7

	v = (v << 1) | c
	cpu.Registers.E = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRrE Rotates E to the right
func cbRRrE(cpu *Core) {
	v := cpu.Registers.E
	c := v & 1
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v >> 1) | (f << 7)
	cpu.Registers.E = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRCrE Rotates E to the right
func cbRRCrE(cpu *Core) {
	v := cpu.Registers.E
	c := v & 1

	v = (v >> 1) | (c << 7)
	cpu.Registers.E = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSLArE Shifts E to the left
func cbSLArE(cpu *Core) {
	v := cpu.Registers.E
	c := (v >> 7) > 0

	v = (v << 1)
	cpu.Registers.E = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSRArE Shifts E to the right
func cbSRArE(cpu *Core) {
	v := cpu.Registers.E
	c := (v & 1) > 0
	e := v & 0x80

	v = (v >> 1) | e
	cpu.Registers.E = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSWAPrE Swaps nibbles in register
func cbSWAPrE(cpu *Core) {
	v := cpu.Registers.E
	v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
	cpu.Registers.E = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(false)

	cpu.Registers.LastClockM = 1
	cpu.Registers.LastClockT = 4
}

// cbSRLrE Shift E right
func cbSRLrE(cpu *Core) {
	v := cpu.Registers.E
	c := (v & 1) > 0

	v = (v >> 1)
	cpu.Registers.E = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT0E Sets Flag Zero to BIT 0 from E
func cbBIT0E(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.E&(1<<0) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES0E Resets BIT 0 from E
func cbRES0E(cpu *Core) {
	cpu.Registers.E &= ^(uint8(1) << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET0E Sets BIT 0 from E
func cbSET0E(cpu *Core) {
	cpu.Registers.E |= (1 << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT1E Sets Flag Zero to BIT 1 from E
func cbBIT1E(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.E&(1<<1) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES1E Resets BIT 1 from E
func cbRES1E(cpu *Core) {
	cpu.Registers.E &= ^(uint8(1) << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET1E Sets BIT 1 from E
func cbSET1E(cpu *Core) {
	cpu.Registers.E |= (1 << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT2E Sets Flag Zero to BIT 2 from E
func cbBIT2E(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.E&(1<<2) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES2E Resets BIT 2 from E
func cbRES2E(cpu *Core) {
	cpu.Registers.E &= ^(uint8(1) << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET2E Sets BIT 2 from E
func cbSET2E(cpu *Core) {
	cpu.Registers.E |= (1 << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT3E Sets Flag Zero to BIT 3 from E
func cbBIT3E(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.E&(1<<3) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES3E Resets BIT 3 from E
func cbRES3E(cpu *Core) {
	cpu.Registers.E &= ^(uint8(1) << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET3E Sets BIT 3 from E
func cbSET3E(cpu *Core) {
	cpu.Registers.E |= (1 << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT4E Sets Flag Zero to BIT 4 from E
func cbBIT4E(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.E&(1<<4) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES4E Resets BIT 4 from E
func cbRES4E(cpu *Core) {
	cpu.Registers.E &= ^(uint8(1) << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET4E Sets BIT 4 from E
func cbSET4E(cpu *Core) {
	cpu.Registers.E |= (1 << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT5E Sets Flag Zero to BIT 5 from E
func cbBIT5E(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.E&(1<<5) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES5E Resets BIT 5 from E
func cbRES5E(cpu *Core) {
	cpu.Registers.E &= ^(uint8(1) << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET5E Sets BIT 5 from E
func cbSET5E(cpu *Core) {
	cpu.Registers.E |= (1 << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT6E Sets Flag Zero to BIT 6 from E
func cbBIT6E(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.E&(1<<6) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES6E Resets BIT 6 from E
func cbRES6E(cpu *Core) {
	cpu.Registers.E &= ^(uint8(1) << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET6E Sets BIT 6 from E
func cbSET6E(cpu *Core) {
	cpu.Registers.E |= (1 << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT7E Sets Flag Zero to BIT 7 from E
func cbBIT7E(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.E&(1<<7) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES7E Resets BIT 7 from E
func cbRES7E(cpu *Core) {
	cpu.Registers.E &= ^(uint8(1) << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET7E Sets BIT 7 from E
func cbSET7E(cpu *Core) {
	cpu.Registers.E |= (1 << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbRLrH Rotates H to the left
func cbRLrH(cpu *Core) {
	v := cpu.Registers.H
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v << 1) | f
	cpu.Registers.H = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRLCrH Rotates H to the left with carry
func cbRLCrH(cpu *Core) {
	v := cpu.Registers.H
	c := v >> 7

	v = (v << 1) | c
	cpu.Registers.H = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRrH Rotates H to the right
func cbRRrH(cpu *Core) {
	v := cpu.Registers.H
	c := v & 1
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v >> 1) | (f << 7)
	cpu.Registers.H = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRCrH Rotates H to the right
func cbRRCrH(cpu *Core) {
	v := cpu.Registers.H
	c := v & 1

	v = (v >> 1) | (c << 7)
	cpu.Registers.H = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSLArH Shifts H to the left
func cbSLArH(cpu *Core) {
	v := cpu.Registers.H
	c := (v >> 7) > 0

	v = (v << 1)
	cpu.Registers.H = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSRArH Shifts H to the right
func cbSRArH(cpu *Core) {
	v := cpu.Registers.H
	c := (v & 1) > 0
	e := v & 0x80

	v = (v >> 1) | e
	cpu.Registers.H = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSWAPrH Swaps nibbles in register
func cbSWAPrH(cpu *Core) {
	v := cpu.Registers.H
	v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
	cpu.Registers.H = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(false)

	cpu.Registers.LastClockM = 1
	cpu.Registers.LastClockT = 4
}

// cbSRLrH Shift H right
func cbSRLrH(cpu *Core) {
	v := cpu.Registers.H
	c := (v & 1) > 0

	v = (v >> 1)
	cpu.Registers.H = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT0H Sets Flag Zero to BIT 0 from H
func cbBIT0H(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.H&(1<<0) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES0H Resets BIT 0 from H
func cbRES0H(cpu *Core) {
	cpu.Registers.H &= ^(uint8(1) << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET0H Sets BIT 0 from H
func cbSET0H(cpu *Core) {
	cpu.Registers.H |= (1 << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT1H Sets Flag Zero to BIT 1 from H
func cbBIT1H(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.H&(1<<1) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES1H Resets BIT 1 from H
func cbRES1H(cpu *Core) {
	cpu.Registers.H &= ^(uint8(1) << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET1H Sets BIT 1 from H
func cbSET1H(cpu *Core) {
	cpu.Registers.H |= (1 << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT2H Sets Flag Zero to BIT 2 from H
func cbBIT2H(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.H&(1<<2) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES2H Resets BIT 2 from H
func cbRES2H(cpu *Core) {
	cpu.Registers.H &= ^(uint8(1) << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET2H Sets BIT 2 from H
func cbSET2H(cpu *Core) {
	cpu.Registers.H |= (1 << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT3H Sets Flag Zero to BIT 3 from H
func cbBIT3H(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.H&(1<<3) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES3H Resets BIT 3 from H
func cbRES3H(cpu *Core) {
	cpu.Registers.H &= ^(uint8(1) << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET3H Sets BIT 3 from H
func cbSET3H(cpu *Core) {
	cpu.Registers.H |= (1 << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT4H Sets Flag Zero to BIT 4 from H
func cbBIT4H(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.H&(1<<4) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES4H Resets BIT 4 from H
func cbRES4H(cpu *Core) {
	cpu.Registers.H &= ^(uint8(1) << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET4H Sets BIT 4 from H
func cbSET4H(cpu *Core) {
	cpu.Registers.H |= (1 << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT5H Sets Flag Zero to BIT 5 from H
func cbBIT5H(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.H&(1<<5) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES5H Resets BIT 5 from H
func cbRES5H(cpu *Core) {
	cpu.Registers.H &= ^(uint8(1) << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET5H Sets BIT 5 from H
func cbSET5H(cpu *Core) {
	cpu.Registers.H |= (1 << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT6H Sets Flag Zero to BIT 6 from H
func cbBIT6H(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.H&(1<<6) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES6H Resets BIT 6 from H
func cbRES6H(cpu *Core) {
	cpu.Registers.H &= ^(uint8(1) << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET6H Sets BIT 6 from H
func cbSET6H(cpu *Core) {
	cpu.Registers.H |= (1 << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT7H Sets Flag Zero to BIT 7 from H
func cbBIT7H(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.H&(1<<7) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES7H Resets BIT 7 from H
func cbRES7H(cpu *Core) {
	cpu.Registers.H &= ^(uint8(1) << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET7H Sets BIT 7 from H
func cbSET7H(cpu *Core) {
	cpu.Registers.H |= (1 << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbRLrL Rotates L to the left
func cbRLrL(cpu *Core) {
	v := cpu.Registers.L
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v << 1) | f
	cpu.Registers.L = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRLCrL Rotates L to the left with carry
func cbRLCrL(cpu *Core) {
	v := cpu.Registers.L
	c := v >> 7

	v = (v << 1) | c
	cpu.Registers.L = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRrL Rotates L to the right
func cbRRrL(cpu *Core) {
	v := cpu.Registers.L
	c := v & 1
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v >> 1) | (f << 7)
	cpu.Registers.L = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRCrL Rotates L to the right
func cbRRCrL(cpu *Core) {
	v := cpu.Registers.L
	c := v & 1

	v = (v >> 1) | (c << 7)
	cpu.Registers.L = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSLArL Shifts L to the left
func cbSLArL(cpu *Core) {
	v := cpu.Registers.L
	c := (v >> 7) > 0

	v = (v << 1)
	cpu.Registers.L = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSRArL Shifts L to the right
func cbSRArL(cpu *Core) {
	v := cpu.Registers.L
	c := (v & 1) > 0
	e := v & 0x80

	v = (v >> 1) | e
	cpu.Registers.L = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSWAPrL Swaps nibbles in register
func cbSWAPrL(cpu *Core) {
	v := cpu.Registers.L
	v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
	cpu.Registers.L = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(false)

	cpu.Registers.LastClockM = 1
	cpu.Registers.LastClockT = 4
}

// cbSRLrL Shift L right
func cbSRLrL(cpu *Core) {
	v := cpu.Registers.L
	c := (v & 1) > 0

	v = (v >> 1)
	cpu.Registers.L = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT0L Sets Flag Zero to BIT 0 from L
func cbBIT0L(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.L&(1<<0) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES0L Resets BIT 0 from L
func cbRES0L(cpu *Core) {
	cpu.Registers.L &= ^(uint8(1) << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET0L Sets BIT 0 from L
func cbSET0L(cpu *Core) {
	cpu.Registers.L |= (1 << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT1L Sets Flag Zero to BIT 1 from L
func cbBIT1L(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.L&(1<<1) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES1L Resets BIT 1 from L
func cbRES1L(cpu *Core) {
	cpu.Registers.L &= ^(uint8(1) << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET1L Sets BIT 1 from L
func cbSET1L(cpu *Core) {
	cpu.Registers.L |= (1 << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT2L Sets Flag Zero to BIT 2 from L
func cbBIT2L(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.L&(1<<2) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES2L Resets BIT 2 from L
func cbRES2L(cpu *Core) {
	cpu.Registers.L &= ^(uint8(1) << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET2L Sets BIT 2 from L
func cbSET2L(cpu *Core) {
	cpu.Registers.L |= (1 << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT3L Sets Flag Zero to BIT 3 from L
func cbBIT3L(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.L&(1<<3) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES3L Resets BIT 3 from L
func cbRES3L(cpu *Core) {
	cpu.Registers.L &= ^(uint8(1) << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET3L Sets BIT 3 from L
func cbSET3L(cpu *Core) {
	cpu.Registers.L |= (1 << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT4L Sets Flag Zero to BIT 4 from L
func cbBIT4L(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.L&(1<<4) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES4L Resets BIT 4 from L
func cbRES4L(cpu *Core) {
	cpu.Registers.L &= ^(uint8(1) << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET4L Sets BIT 4 from L
func cbSET4L(cpu *Core) {
	cpu.Registers.L |= (1 << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT5L Sets Flag Zero to BIT 5 from L
func cbBIT5L(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.L&(1<<5) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES5L Resets BIT 5 from L
func cbRES5L(cpu *Core) {
	cpu.Registers.L &= ^(uint8(1) << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET5L Sets BIT 5 from L
func cbSET5L(cpu *Core) {
	cpu.Registers.L |= (1 << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT6L Sets Flag Zero to BIT 6 from L
func cbBIT6L(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.L&(1<<6) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES6L Resets BIT 6 from L
func cbRES6L(cpu *Core) {
	cpu.Registers.L &= ^(uint8(1) << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET6L Sets BIT 6 from L
func cbSET6L(cpu *Core) {
	cpu.Registers.L |= (1 << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT7L Sets Flag Zero to BIT 7 from L
func cbBIT7L(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.L&(1<<7) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES7L Resets BIT 7 from L
func cbRES7L(cpu *Core) {
	cpu.Registers.L &= ^(uint8(1) << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET7L Sets BIT 7 from L
func cbSET7L(cpu *Core) {
	cpu.Registers.L |= (1 << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbRLrF Rotates F to the left
func cbRLrF(cpu *Core) {
	v := cpu.Registers.F
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v << 1) | f
	cpu.Registers.F = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRLCrF Rotates F to the left with carry
func cbRLCrF(cpu *Core) {
	v := cpu.Registers.F
	c := v >> 7

	v = (v << 1) | c
	cpu.Registers.F = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRrF Rotates F to the right
func cbRRrF(cpu *Core) {
	v := cpu.Registers.F
	c := v & 1
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v >> 1) | (f << 7)
	cpu.Registers.F = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbRRCrF Rotates F to the right
func cbRRCrF(cpu *Core) {
	v := cpu.Registers.F
	c := v & 1

	v = (v >> 1) | (c << 7)
	cpu.Registers.F = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSLArF Shifts F to the left
func cbSLArF(cpu *Core) {
	v := cpu.Registers.F
	c := (v >> 7) > 0

	v = (v << 1)
	cpu.Registers.F = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSRArF Shifts F to the right
func cbSRArF(cpu *Core) {
	v := cpu.Registers.F
	c := (v & 1) > 0
	e := v & 0x80

	v = (v >> 1) | e
	cpu.Registers.F = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// cbSWAPrF Swaps nibbles in register
func cbSWAPrF(cpu *Core) {
	v := cpu.Registers.F
	v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
	cpu.Registers.F = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(false)

	cpu.Registers.LastClockM = 1
	cpu.Registers.LastClockT = 4
}

// cbSRLrF Shift F right
func cbSRLrF(cpu *Core) {
	v := cpu.Registers.F
	c := (v & 1) > 0

	v = (v >> 1)
	cpu.Registers.F = v

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT0F Sets Flag Zero to BIT 0 from F
func cbBIT0F(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.F&(1<<0) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES0F Resets BIT 0 from F
func cbRES0F(cpu *Core) {
	cpu.Registers.F &= ^(uint8(1) << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET0F Sets BIT 0 from F
func cbSET0F(cpu *Core) {
	cpu.Registers.F |= (1 << 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT1F Sets Flag Zero to BIT 1 from F
func cbBIT1F(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.F&(1<<1) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES1F Resets BIT 1 from F
func cbRES1F(cpu *Core) {
	cpu.Registers.F &= ^(uint8(1) << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET1F Sets BIT 1 from F
func cbSET1F(cpu *Core) {
	cpu.Registers.F |= (1 << 1)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT2F Sets Flag Zero to BIT 2 from F
func cbBIT2F(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.F&(1<<2) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES2F Resets BIT 2 from F
func cbRES2F(cpu *Core) {
	cpu.Registers.F &= ^(uint8(1) << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET2F Sets BIT 2 from F
func cbSET2F(cpu *Core) {
	cpu.Registers.F |= (1 << 2)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT3F Sets Flag Zero to BIT 3 from F
func cbBIT3F(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.F&(1<<3) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES3F Resets BIT 3 from F
func cbRES3F(cpu *Core) {
	cpu.Registers.F &= ^(uint8(1) << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET3F Sets BIT 3 from F
func cbSET3F(cpu *Core) {
	cpu.Registers.F |= (1 << 3)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT4F Sets Flag Zero to BIT 4 from F
func cbBIT4F(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.F&(1<<4) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES4F Resets BIT 4 from F
func cbRES4F(cpu *Core) {
	cpu.Registers.F &= ^(uint8(1) << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET4F Sets BIT 4 from F
func cbSET4F(cpu *Core) {
	cpu.Registers.F |= (1 << 4)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT5F Sets Flag Zero to BIT 5 from F
func cbBIT5F(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.F&(1<<5) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES5F Resets BIT 5 from F
func cbRES5F(cpu *Core) {
	cpu.Registers.F &= ^(uint8(1) << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET5F Sets BIT 5 from F
func cbSET5F(cpu *Core) {
	cpu.Registers.F |= (1 << 5)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT6F Sets Flag Zero to BIT 6 from F
func cbBIT6F(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.F&(1<<6) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES6F Resets BIT 6 from F
func cbRES6F(cpu *Core) {
	cpu.Registers.F &= ^(uint8(1) << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET6F Sets BIT 6 from F
func cbSET6F(cpu *Core) {
	cpu.Registers.F |= (1 << 6)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBIT7F Sets Flag Zero to BIT 7 from F
func cbBIT7F(cpu *Core) {
	cpu.Registers.SetZero(cpu.Registers.F&(1<<7) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 4
}

// cbRES7F Resets BIT 7 from F
func cbRES7F(cpu *Core) {
	cpu.Registers.F &= ^(uint8(1) << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSET7F Sets BIT 7 from F
func cbSET7F(cpu *Core) {
	cpu.Registers.F |= (1 << 7)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbBITm0 Sets Flag Zero to BIT 0 from [HL]
func cbBITm0(cpu *Core) {
	cpu.Registers.SetZero(cpu.Memory.ReadByte(cpu.Registers.HL())&(1<<0) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbRESHL0 Resets BIT 0 from [HL]
func cbRESHL0(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) & ^(uint8(1) << 0)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbSETHL0 Sets BIT 0 from [HL]
func cbSETHL0(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) | (1 << 0)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbBITm1 Sets Flag Zero to BIT 1 from [HL]
func cbBITm1(cpu *Core) {
	cpu.Registers.SetZero(cpu.Memory.ReadByte(cpu.Registers.HL())&(1<<1) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbRESHL1 Resets BIT 1 from [HL]
func cbRESHL1(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) & ^(uint8(1) << 1)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbSETHL1 Sets BIT 1 from [HL]
func cbSETHL1(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) | (1 << 1)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbBITm2 Sets Flag Zero to BIT 2 from [HL]
func cbBITm2(cpu *Core) {
	cpu.Registers.SetZero(cpu.Memory.ReadByte(cpu.Registers.HL())&(1<<2) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbRESHL2 Resets BIT 2 from [HL]
func cbRESHL2(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) & ^(uint8(1) << 2)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbSETHL2 Sets BIT 2 from [HL]
func cbSETHL2(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) | (1 << 2)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbBITm3 Sets Flag Zero to BIT 3 from [HL]
func cbBITm3(cpu *Core) {
	cpu.Registers.SetZero(cpu.Memory.ReadByte(cpu.Registers.HL())&(1<<3) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbRESHL3 Resets BIT 3 from [HL]
func cbRESHL3(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) & ^(uint8(1) << 3)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbSETHL3 Sets BIT 3 from [HL]
func cbSETHL3(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) | (1 << 3)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbBITm4 Sets Flag Zero to BIT 4 from [HL]
func cbBITm4(cpu *Core) {
	cpu.Registers.SetZero(cpu.Memory.ReadByte(cpu.Registers.HL())&(1<<4) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbRESHL4 Resets BIT 4 from [HL]
func cbRESHL4(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) & ^(uint8(1) << 4)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbSETHL4 Sets BIT 4 from [HL]
func cbSETHL4(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) | (1 << 4)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbBITm5 Sets Flag Zero to BIT 5 from [HL]
func cbBITm5(cpu *Core) {
	cpu.Registers.SetZero(cpu.Memory.ReadByte(cpu.Registers.HL())&(1<<5) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbRESHL5 Resets BIT 5 from [HL]
func cbRESHL5(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) & ^(uint8(1) << 5)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbSETHL5 Sets BIT 5 from [HL]
func cbSETHL5(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) | (1 << 5)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbBITm6 Sets Flag Zero to BIT 6 from [HL]
func cbBITm6(cpu *Core) {
	cpu.Registers.SetZero(cpu.Memory.ReadByte(cpu.Registers.HL())&(1<<6) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbRESHL6 Resets BIT 6 from [HL]
func cbRESHL6(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) & ^(uint8(1) << 6)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbSETHL6 Sets BIT 6 from [HL]
func cbSETHL6(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) | (1 << 6)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbBITm7 Sets Flag Zero to BIT 7 from [HL]
func cbBITm7(cpu *Core) {
	cpu.Registers.SetZero(cpu.Memory.ReadByte(cpu.Registers.HL())&(1<<7) != 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbRESHL7 Resets BIT 7 from [HL]
func cbRESHL7(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) & ^(uint8(1) << 7)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbSETHL7 Sets BIT 7 from [HL]
func cbSETHL7(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL()) | (1 << 7)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// cbRLHL Rotates [HL] to the left
func cbRLHL(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL())
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v << 1) | f

	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 16
}

// cbRLHL Rotates [HL] to the left with carry
func cbRLCHL(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL())
	c := v >> 7

	v = (v << 1) | c
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 16
}

// cbRLHL Rotates [HL] to the right
func cbRRHL(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL())
	c := (v >> 7) > 0
	f := byte(0)
	if cpu.Registers.GetCarry() {
		f = 1
	}

	v = (v >> 1) | (f << 7)

	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 16
}

// cbRLHL Rotates [HL] to the right with carry
func cbRRCHL(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL())
	c := v >> 7

	v = (v >> 1) | (c << 7)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 16
}

// cbSLAHL Shifts [HL] to the left
func cbSLAHL(cpu *Core) {
	b := cpu.Memory.ReadByte(cpu.Registers.HL())
	c := b >> 7
	b = b << 1
	cpu.Memory.WriteByte(cpu.Registers.HL(), b)

	cpu.Registers.SetZero(b == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c > 0)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 16
}

// cbSRAHL Shifts [HL] to the right. Keeps bit 7 constant
func cbSRAHL(cpu *Core) {
	b := cpu.Memory.ReadByte(cpu.Registers.HL())
	c := b&1 > 0
	e := b & 0x80
	b = b>>1 | e
	cpu.Memory.WriteByte(cpu.Registers.HL(), b)

	cpu.Registers.SetZero(b == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 16
}

// cbSWAPHL Swaps nibbles in [HL]
func cbSWAPHL(cpu *Core) {
	v := cpu.Memory.ReadByte(cpu.Registers.HL())
	v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
	cpu.Memory.WriteByte(cpu.Registers.HL(), v)

	cpu.Registers.SetZero(v == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(false)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 8
}

// cbSRLHL Shifts [HL] to the right.
func cbSRLHL(cpu *Core) {
	b := cpu.Memory.ReadByte(cpu.Registers.HL())
	c := b&1 > 0
	b = b >> 1
	cpu.Memory.WriteByte(cpu.Registers.HL(), b)

	cpu.Registers.SetZero(b == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)
	cpu.Registers.SetCarry(c)

	cpu.Registers.LastClockM = 4
	cpu.Registers.LastClockT = 16
}

var CBInstructions = []GBInstruction{
	//region CB00 Group
	cbRLCrB,
	cbRLCrC,
	cbRLCrD,
	cbRLCrE,
	cbRLCrH,
	cbRLCrL,
	cbRLCHL,
	cbRLCrA,
	cbRRCrB,
	cbRRCrC,
	cbRRCrD,
	cbRRCrE,
	cbRRCrH,
	cbRRCrL,
	cbRRCHL,
	cbRRCrA,
	//endregion
	//region CB10 Group
	cbRLrB,
	cbRLrC,
	cbRLrD,
	cbRLrE,
	cbRLrH,
	cbRLrL,
	cbRLHL,
	cbRLrA,
	cbRRrB,
	cbRRrC,
	cbRRrD,
	cbRRrE,
	cbRRrH,
	cbRRrL,
	cbRRHL,
	cbRRrA,
	//endregion
	//region CB20 Group
	cbSLArB,
	cbSLArC,
	cbSLArD,
	cbSLArE,
	cbSLArH,
	cbSLArL,
	cbSLAHL,
	cbSLArA,
	cbSRArB,
	cbSRArC,
	cbSRArD,
	cbSRArE,
	cbSRArH,
	cbSRArL,
	cbSRAHL,
	cbSRArA,
	//endregion
	//region CB30 Group
	cbSWAPrB,
	cbSWAPrC,
	cbSWAPrD,
	cbSWAPrE,
	cbSWAPrH,
	cbSWAPrL,
	cbSWAPHL,
	cbSWAPrA,
	cbSRLrB,
	cbSRLrC,
	cbSRLrD,
	cbSRLrE,
	cbSRLrH,
	cbSRLrL,
	cbSRLHL,
	cbSRLrA,
	//endregion
	//region CB40 Group
	cbBIT0B,
	cbBIT0C,
	cbBIT0D,
	cbBIT0E,
	cbBIT0H,
	cbBIT0L,
	cbBITm0,
	cbBIT0A,
	cbBIT1B,
	cbBIT1C,
	cbBIT1D,
	cbBIT1E,
	cbBIT1H,
	cbBIT1L,
	cbBITm1,
	cbBIT1A,
	//endregion
	//region CB50 Group
	cbBIT2B,
	cbBIT2C,
	cbBIT2D,
	cbBIT2E,
	cbBIT2H,
	cbBIT2L,
	cbBITm2,
	cbBIT2A,
	cbBIT3B,
	cbBIT3C,
	cbBIT3D,
	cbBIT3E,
	cbBIT3H,
	cbBIT3L,
	cbBITm3,
	cbBIT3A,
	//endregion
	//region CB60 Group
	cbBIT4B,
	cbBIT4C,
	cbBIT4D,
	cbBIT4E,
	cbBIT4H,
	cbBIT4L,
	cbBITm4,
	cbBIT4A,
	cbBIT5B,
	cbBIT5C,
	cbBIT5D,
	cbBIT5E,
	cbBIT5H,
	cbBIT5L,
	cbBITm5,
	cbBIT5A,
	//endregion
	//region CB70 Group
	cbBIT6B,
	cbBIT6C,
	cbBIT6D,
	cbBIT6E,
	cbBIT6H,
	cbBIT6L,
	cbBITm6,
	cbBIT6A,
	cbBIT7B,
	cbBIT7C,
	cbBIT7D,
	cbBIT7E,
	cbBIT7H,
	cbBIT7L,
	cbBITm7,
	cbBIT7A,
	//endregion
	//region CB80 Group
	cbRES0B,
	cbRES0C,
	cbRES0D,
	cbRES0E,
	cbRES0H,
	cbRES0L,
	cbRESHL0,
	cbRES0A,
	cbRES1B,
	cbRES1C,
	cbRES1D,
	cbRES1E,
	cbRES1H,
	cbRES1L,
	cbRESHL1,
	cbRES1A,
	//endregion
	//region CB90 Group
	cbRES2B,
	cbRES2C,
	cbRES2D,
	cbRES2E,
	cbRES2H,
	cbRES2L,
	cbRESHL2,
	cbRES2A,
	cbRES3B,
	cbRES3C,
	cbRES3D,
	cbRES3E,
	cbRES3H,
	cbRES3L,
	cbRESHL3,
	cbRES3A,
	//endregion
	//region CBA0 Group
	cbRES4B,
	cbRES4C,
	cbRES4D,
	cbRES4E,
	cbRES4H,
	cbRES4L,
	cbRESHL4,
	cbRES4A,
	cbRES6B,
	cbRES6C,
	cbRES6D,
	cbRES6E,
	cbRES6H,
	cbRES6L,
	cbRESHL6,
	cbRES6A,
	//endregion
	//region CBB0 Group
	cbRES6B,
	cbRES6C,
	cbRES6D,
	cbRES6E,
	cbRES6H,
	cbRES6L,
	cbRESHL6,
	cbRES6A,
	cbRES7B,
	cbRES7C,
	cbRES7D,
	cbRES7E,
	cbRES7H,
	cbRES7L,
	cbRESHL7,
	cbRES7A,
	//endregion
	//region CBC0 Group
	cbSET0B,
	cbSET0C,
	cbSET0D,
	cbSET0E,
	cbSET0H,
	cbSET0L,
	cbSETHL0,
	cbSET0A,
	cbSET1B,
	cbSET1C,
	cbSET1D,
	cbSET1E,
	cbSET1H,
	cbSET1L,
	cbSETHL1,
	cbSET1A,
	//endregion
	//region CBD0 Group
	cbSET2B,
	cbSET2C,
	cbSET2D,
	cbSET2E,
	cbSET2H,
	cbSET2L,
	cbSETHL2,
	cbSET2A,
	cbSET3B,
	cbSET3C,
	cbSET3D,
	cbSET3E,
	cbSET3H,
	cbSET3L,
	cbSETHL3,
	cbSET3A,
	//endregion
	//region CBE0 Group
	cbSET4B,
	cbSET4C,
	cbSET4D,
	cbSET4E,
	cbSET4H,
	cbSET4L,
	cbSETHL4,
	cbSET4A,
	cbSET5B,
	cbSET5C,
	cbSET5D,
	cbSET5E,
	cbSET5H,
	cbSET5L,
	cbSETHL5,
	cbSET5A,
	//endregion
	//region CBF0 Group
	cbSET6B,
	cbSET6C,
	cbSET6D,
	cbSET6E,
	cbSET6H,
	cbSET6L,
	cbSETHL6,
	cbSET6A,
	cbSET7B,
	cbSET7C,
	cbSET7D,
	cbSET7E,
	cbSET7H,
	cbSET7L,
	cbSETHL7,
	cbSET7A,
	//endregion
}
