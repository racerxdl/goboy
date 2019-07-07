package cpu

import (
	"github.com/racerxdl/goboy/gameboy"
	"math/rand"
	"strings"
)

var AllRegisters = []string{
	"A", "B", "C", "D", "E", "H", "L", "F",
}

type Registers struct {
	// Exposed Registers
	A, B, C, D, E, H, L, F byte

	PC, SP uint16

	// Stored Registers
	a, b, c, d, e, h, l, f byte

	InterruptEnable        bool
	EnabledInterrupts      byte
	TriggerInterrupts      byte
	CycleCount             int
	LastClockM, LastClockT int
}

// region Register Helpers
func (r *Registers) HL() uint16 {
	return uint16(r.H)<<8 + uint16(r.L)
}

func (r *Registers) BC() uint16 {
	return uint16(r.B)<<8 + uint16(r.C)
}

func (r *Registers) DE() uint16 {
	return uint16(r.D)<<8 + uint16(r.E)
}

func (r *Registers) GetZero() bool {
	return r.F&gameboy.FlagZero > 0
}

func (r *Registers) SetZero(v bool) {
	if v {
		r.F |= gameboy.FlagZero
	} else {
		r.F &= gameboy.InvFlagZero
	}
}

func (r *Registers) GetSub() bool {
	return r.F&gameboy.FlagSub > 0
}

func (r *Registers) SetSub(v bool) {
	if v {
		r.F |= gameboy.FlagSub
	} else {
		r.F &= gameboy.InvFlagSub
	}
}

func (r *Registers) GetHalfCarry() bool {
	return r.F&gameboy.FlagHalfCarry > 0
}

func (r *Registers) SetHalfCarry(v bool) {
	if v {
		r.F |= gameboy.FlagHalfCarry
	} else {
		r.F &= gameboy.InvFlagHalfCarry
	}
}

func (r *Registers) GetCarry() bool {
	return r.F&gameboy.FlagCarry > 0
}

func (r *Registers) SetCarry(v bool) {
	if v {
		r.F |= gameboy.FlagCarry
	} else {
		r.F &= gameboy.InvFlagCarry
	}
}

// endregion

func MakeRegisters() *Registers {
	r := &Registers{}
	r.Reset()
	return r
}

func (r *Registers) Clone() Registers {
	return *r
}

func (r *Registers) SaveRegisters() {
	r.a = r.A
	r.b = r.B
	r.c = r.C
	r.d = r.D
	r.e = r.E
	r.h = r.H
	r.l = r.L
	r.f = r.F
}

func (r *Registers) LoadRegisters() {
	r.A = r.a
	r.B = r.b
	r.C = r.c
	r.D = r.d
	r.E = r.e
	r.H = r.h
	r.L = r.l
	r.F = r.f
}

func (r *Registers) Reset() {
	r.CycleCount = 0
	r.A = 0
	r.B = 0
	r.C = 0
	r.D = 0
	r.E = 0
	r.H = 0
	r.L = 0
	r.F = 0
	r.PC = 0
	r.SP = 0
	r.InterruptEnable = true
	r.TriggerInterrupts = 0
	r.LastClockM = 0
	r.LastClockT = 0
	r.EnabledInterrupts = 0
}

func (r *Registers) GetRegister(reg string) byte {
	reg = strings.ToUpper(reg)
	switch reg {
	case "A":
		return r.A
	case "B":
		return r.B
	case "C":
		return r.C
	case "D":
		return r.D
	case "E":
		return r.E
	case "H":
		return r.H
	case "L":
		return r.L
	case "F":
		return r.F
	case "_A":
		return r.a
	case "_B":
		return r.b
	case "_C":
		return r.c
	case "_D":
		return r.d
	case "_E":
		return r.e
	case "_H":
		return r.h
	case "_L":
		return r.l
	case "_F":
		return r.f
	}

	return 0
}

func (r *Registers) SetRegister(reg string, v byte) {
	reg = strings.ToUpper(reg)
	switch reg {
	case "A":
		r.A = v
	case "B":
		r.B = v
	case "C":
		r.C = v
	case "D":
		r.D = v
	case "E":
		r.E = v
	case "H":
		r.H = v
	case "L":
		r.L = v
	case "F":
		r.F = v
	case "_A":
		r.a = v
	case "_B":
		r.b = v
	case "_C":
		r.c = v
	case "_D":
		r.d = v
	case "_E":
		r.e = v
	case "_H":
		r.h = v
	case "_L":
		r.l = v
	case "_F":
		r.f = v
	}
}

func (r *Registers) Randomize() {
	for _, v := range AllRegisters {
		r.SetRegister(v, byte(rand.Intn(255)))
	}

	r.PC = uint16(rand.Int31n(0xFFFF))
	r.SP = uint16(rand.Int31n(0xFFFF))

	r.InterruptEnable = rand.Float32() > 0.5
	r.EnabledInterrupts = byte(rand.Int31n(255))
	r.TriggerInterrupts = byte(rand.Int31n(255))
}
