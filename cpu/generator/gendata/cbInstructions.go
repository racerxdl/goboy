package gendata

import (
	"bytes"
	"text/template"
)

var rlrTemplate = template.Must(template.New("RLr").Parse(`
// cbRLr{{.I}} Rotates {{.I}} to the left
func cbRLr{{.I}}(cpu *Core) {
    v := cpu.Registers.{{.I}}
    c := (v >> 7) > 0
    f := byte(0)
    if cpu.Registers.GetCarry() {
        f = 1
    }

    v = (v << 1) | f
    cpu.Registers.{{.I}} = v

    cpu.Registers.SetZero(v == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(c)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var rlcrTemplate = template.Must(template.New("RLCr").Parse(`
// cbRLCr{{.I}} Rotates {{.I}} to the left with carry
func cbRLCr{{.I}}(cpu *Core) {
    v := cpu.Registers.{{.I}}
    c := v >> 7

    v = (v << 1) | c
    cpu.Registers.{{.I}} = v

    cpu.Registers.SetZero(v == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(c > 0)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var rrrTemplate = template.Must(template.New("RRr").Parse(`
// cbRRr{{.I}} Rotates {{.I}} to the right
func cbRRr{{.I}}(cpu *Core) {
    v := cpu.Registers.{{.I}}
    c := v & 1
    f := byte(0)
    if cpu.Registers.GetCarry() {
        f = 1
    }

    v = (v >> 1) | (f << 7)
    cpu.Registers.{{.I}} = v

    cpu.Registers.SetZero(v == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(c > 0)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var rrcrTemplate = template.Must(template.New("RRCr").Parse(`
// cbRRCr{{.I}} Rotates {{.I}} to the right
func cbRRCr{{.I}}(cpu *Core) {
    v := cpu.Registers.{{.I}}
    c := v & 1

    v = (v >> 1) | (c << 7)
    cpu.Registers.{{.I}} = v

    cpu.Registers.SetZero(v == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(c > 0)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var slarTemplate = template.Must(template.New("SLAr").Parse(`
// cbSLAr{{.I}} Shifts {{.I}} to the left
func cbSLAr{{.I}}(cpu *Core) {
    v := cpu.Registers.{{.I}}
    c := (v >> 7) > 0

    v = (v << 1)
    cpu.Registers.{{.I}} = v

    cpu.Registers.SetZero(v == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(c)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var srarTemplate = template.Must(template.New("SRAr").Parse(`
// cbSRAr{{.I}} Shifts {{.I}} to the right
func cbSRAr{{.I}}(cpu *Core) {
    v := cpu.Registers.{{.I}}
    c := (v & 1) > 0
    e := v & 0x80

    v = (v >> 1) | e
    cpu.Registers.{{.I}} = v

    cpu.Registers.SetZero(v == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(c)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var swaprTemplate = template.Must(template.New("SWAPr").Parse(`
// cbSWAPr{{.I}} Swaps nibbles in register
func cbSWAPr{{.I}}(cpu *Core) {
    v := cpu.Registers.{{.I}}
    v = ((v & 0x0F) << 4) | ((v & 0xF0) >> 4)
    cpu.Registers.{{.I}} = v

    cpu.Registers.SetZero(v == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(false)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

var srlrTemplate = template.Must(template.New("SRLr").Parse(`
// cbSRLr{{.I}} Shift {{.I}} right
func cbSRLr{{.I}}(cpu *Core) {
    v := cpu.Registers.{{.I}}
    c := (v & 1) > 0

    v = (v >> 1) 
    cpu.Registers.{{.I}} = v

    cpu.Registers.SetZero(v == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(c)

    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 8
}
`))

var bitTemplate = template.Must(template.New("BIT").Parse(`
// cbBIT{{.N}}{{.I}} Sets Flag Zero to BIT {{.N}} from {{.I}}
func cbBIT{{.N}}{{.I}}(cpu *Core) {
    cpu.Registers.SetZero(cpu.Registers.{{.I}} & ( 1 << {{.N}}) != 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 4
}
`))

var bitmTemplate = template.Must(template.New("BITm").Parse(`
// cbBITm{{.N}} Sets Flag Zero to BIT {{.N}} from [HL]
func cbBITm{{.N}}(cpu *Core) {
    cpu.Registers.SetZero(cpu.Memory.ReadByte(cpu.Registers.HL()) & ( 1 << {{.N}}) != 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
`))

var resTemplate = template.Must(template.New("RES").Parse(`
// cbRES{{.N}}{{.I}} Resets BIT {{.N}} from {{.I}}
func cbRES{{.N}}{{.I}}(cpu *Core) {
    cpu.Registers.{{.I}} &= ^(uint8(1) << {{.N}})

    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 8
}
`))

var reshlTemplate = template.Must(template.New("RESHL").Parse(`
// cbRESHL{{.N}} Resets BIT {{.N}} from [HL]
func cbRESHL{{.N}}(cpu *Core) {
    v := cpu.Memory.ReadByte(cpu.Registers.HL()) & ^(uint8(1) << {{.N}})
    cpu.Memory.WriteByte(cpu.Registers.HL(), v)

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
`))

var setTemplate = template.Must(template.New("SET").Parse(`
// cbSET{{.N}}{{.I}} Sets BIT {{.N}} from {{.I}}
func cbSET{{.N}}{{.I}}(cpu *Core) {
    cpu.Registers.{{.I}} |= (1 << {{.N}})

    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 8
}
`))

var sethlTemplate = template.Must(template.New("SETHL").Parse(`
// cbSETHL{{.N}} Sets BIT {{.N}} from [HL]
func cbSETHL{{.N}}(cpu *Core) {
    v := cpu.Memory.ReadByte(cpu.Registers.HL()) | (1 << {{.N}})
    cpu.Memory.WriteByte(cpu.Registers.HL(), v)

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
`))

func BuildCB() string {
	buff := bytes.NewBuffer(nil)

	for _, I := range AllRegisters {
		rlrTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		rlcrTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		rrrTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		rrcrTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		slarTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		srarTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		swaprTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		srlrTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		for n := 0; n < 8; n++ {
			bitTemplate.Execute(buff, struct {
				I string
				N int
			}{
				I: I,
				N: n,
			})
			resTemplate.Execute(buff, struct {
				I string
				N int
			}{
				I: I,
				N: n,
			})
			setTemplate.Execute(buff, struct {
				I string
				N int
			}{
				I: I,
				N: n,
			})
		}
	}

	for n := 0; n < 8; n++ {
		bitmTemplate.Execute(buff, struct {
			N int
		}{
			N: n,
		})
		reshlTemplate.Execute(buff, struct {
			N int
		}{
			N: n,
		})
		sethlTemplate.Execute(buff, struct {
			N int
		}{
			N: n,
		})
	}

	return buff.String() + `
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
    c := b & 1 > 0
    e := b & 0x80
    b = b >> 1 | e
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
    c := b & 1 > 0
    b = b >> 1
    cpu.Memory.WriteByte(cpu.Registers.HL(), b)

    cpu.Registers.SetZero(b == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(c)

    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}
`
}
