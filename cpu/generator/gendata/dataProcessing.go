package gendata

import (
	"bytes"
	"text/template"
)

// region Templates
// region ADD
var addTemplate = template.Must(template.New("ADDr").Parse(`
// gbADDr{{.I}} Adds {{.I}} to A
func gbADDr{{.I}}(cpu *Core) {
    sum := uint16(cpu.Registers.A) + uint16(cpu.Registers.{{.I}})

    cpu.Registers.SetCarry(sum > 255)
    cpu.Registers.SetZero(sum & 0xFF == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry((cpu.Registers.A & 0xF) + (cpu.Registers.{{.I}} & 0xF) > 0xF)

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))
var addHLrrTemplate = template.Must(template.New("ADDHLrr").Parse(`
// gbADDHL{{.I0}}{{.I1}} Adds ({{.I0}} << 8) + {{.I1}} to HL
func gbADDHL{{.I0}}{{.I1}}(cpu *Core) {
    {{.I0}}{{.I1}} := uint16(cpu.Registers.{{.I0}}) << 8 + uint16(cpu.Registers.{{.I1}})
    sum := {{.I0}}{{.I1}} + cpu.Registers.HL()

    cpu.Registers.SetCarry(sum > 65535)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry((({{.I0}}{{.I1}} & 0xFFF) + (cpu.Registers.HL() & 0xFFF)) > 0xFFF)

    cpu.Registers.H = uint8(sum >> 8)
    cpu.Registers.L = uint8(sum & 0xFF)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var adcrTemplate = template.Must(template.New("ADCr").Parse(`
// gbADCr{{.I}} Sets A to A + {{.I}} + FlagCarry
func gbADCr{{.I}}(cpu *Core) {
    b := int(cpu.Registers.{{.I}})
    f := 0
    if cpu.Registers.GetCarry() {
        f = 1
    }

    sum := int(cpu.Registers.A) + b + f

    cpu.Registers.SetZero(sum & 0xFF == 0)
    cpu.Registers.SetCarry(sum > 255)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(int(cpu.Registers.A & 0xF) + (b & 0xF) + f > 0xF)

    cpu.Registers.A = uint8(sum & 0xFF)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

// endregion
// region SUB
var subTemplate = template.Must(template.New("SUBr").Parse(`
// gbSUBr{{.I}} Subtracts {{.I}} to A
func gbSUBr{{.I}}(cpu *Core) {
    sum := int16(cpu.Registers.A) - int16(cpu.Registers.{{.I}})

    cpu.Registers.SetCarry(sum < 0)
    cpu.Registers.SetZero(sum & 0xFF == 0)
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry((cpu.Registers.A & 0xF) < (cpu.Registers.{{.I}} & 0xF))

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))
var subHLrrTemplate = template.Must(template.New("SUBHLrr").Parse(`
// gbSUBHL{{.I0}}{{.I1}} HL from ({{.I0}} << 8) + {{.I1}}
func gbSUBHL{{.I0}}{{.I1}}(cpu *Core) {
    {{.I0}}{{.I1}} := int(cpu.Registers.{{.I0}}) << 8 + int(cpu.Registers.{{.I1}})
    sum := {{.I0}}{{.I1}} - int(cpu.Registers.HL())

    cpu.Registers.SetCarry(sum < 0)
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry((({{.I0}}{{.I1}} & 0xFFF) > int(cpu.Registers.HL() & 0xFFF)))

    cpu.Registers.H = uint8(sum >> 8)
    cpu.Registers.L = uint8(sum & 0xFF)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var sbcrTemplate = template.Must(template.New("SBCr").Parse(`
// gbSBCr{{.I}} Sets A to A - {{.I}} - FlagCarry
func gbSBCr{{.I}}(cpu *Core) {
    b := int(cpu.Registers.{{.I}})
    f := 0
    if cpu.Registers.GetCarry() {
        f = 1
    }

    sum := int(cpu.Registers.A) + b + f

    cpu.Registers.SetZero(sum & 0xFF == 0)
    cpu.Registers.SetCarry(sum < 0)
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry(int(cpu.Registers.A & 0xF) < (b & 0xF) + f)

    cpu.Registers.A = uint8(sum & 0xFF)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

// endregion
// region CP
var cpTemplate = template.Must(template.New("CPr").Parse(`
// gbCPr{{.I}} Checks if {{.I}} == A
func gbCPr{{.I}}(cpu *Core) {
    cpu.Registers.SetCarry(cpu.Registers.A < cpu.Registers.{{.I}})
    cpu.Registers.SetZero(cpu.Registers.A == cpu.Registers.{{.I}})
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry((cpu.Registers.A & 0xF) < (cpu.Registers.{{.I}} & 0xF))

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

// endregion
// region Operators
var andTemplate = template.Must(template.New("ANDr").Parse(`
// gbANDr{{.I}} Sets A to A & {{.I}}
func gbANDr{{.I}}(cpu *Core) {
	cpu.Registers.A &= cpu.Registers.{{.I}}

    cpu.Registers.SetCarry(false)
    cpu.Registers.SetZero(cpu.Registers.A == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(true)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

var orTemplate = template.Must(template.New("ORr").Parse(`
// gbORr{{.I}} Sets A to A | {{.I}}
func gbORr{{.I}}(cpu *Core) {
	cpu.Registers.A |= cpu.Registers.{{.I}}

    cpu.Registers.SetCarry(false)
    cpu.Registers.SetZero(cpu.Registers.A == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

var xorTemplate = template.Must(template.New("XORr").Parse(`
// gbXORr{{.I}} Sets A to A | {{.I}}
func gbXORr{{.I}}(cpu *Core) {
	cpu.Registers.A ^= cpu.Registers.{{.I}}

    cpu.Registers.SetCarry(false)
    cpu.Registers.SetZero(cpu.Registers.A == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(false)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

// endregion
// region INC/DEC
var incTemplate = template.Must(template.New("INCr").Parse(`
// gbINCr{{.I}} Sets {{.I}} to {{.I}} + 1
func gbINCr{{.I}}(cpu *Core) {
	v := cpu.Registers.{{.I}}
	cpu.Registers.{{.I}}++

    //cpu.Registers.SetCarry(int(v) + 1 > 255)
    cpu.Registers.SetZero(cpu.Registers.{{.I}} == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry((v & 0xF) + 1 > 0xF)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

var decTemplate = template.Must(template.New("DECr").Parse(`
// gbDECr{{.I}} Sets {{.I}} to {{.I}} - 1
func gbDECr{{.I}}(cpu *Core) {
	v := cpu.Registers.{{.I}}
	cpu.Registers.{{.I}}--

    //cpu.Registers.SetCarry(int(v) - 1 < 0)
    cpu.Registers.SetZero(cpu.Registers.{{.I}} == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry((v & 0xF) == 0)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

var incrrTemplate = template.Must(template.New("INCrr").Parse(`
// gbINC{{.I0}}{{.I1}} Sets ( {{.I0}} << 8 + {{.I1}} ) to (  {{.I0}} << 8 + {{.I1}} ) + 1
func gbINC{{.I0}}{{.I1}}(cpu *Core) {
	cpu.Registers.{{.I1}}++

	if cpu.Registers.{{.I1}} == 0 {
		cpu.Registers.{{.I0}}++
	}

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var decrrTemplate = template.Must(template.New("DECrr").Parse(`
// gbDEC{{.I0}}{{.I1}} Sets ( {{.I0}} << 8 + {{.I1}} ) to ( {{.I0}} << 8 + {{.I1}} ) - 1
func gbDEC{{.I0}}{{.I1}}(cpu *Core) {
	cpu.Registers.{{.I1}}--

	if cpu.Registers.{{.I1}} == 0 {
		cpu.Registers.{{.I0}}--
	}

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

// endregion
// endregion
// region Builders

func BuildADD() string {
	//
	buff := bytes.NewBuffer(nil)

	for _, I := range AllRegisters {
		addTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		adcrTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
	}

	for _, I0 := range AllRegisters {
		for _, I1 := range AllRegisters {
			if I0 != I1 {
				addHLrrTemplate.Execute(buff, struct {
					I0 string
					I1 string
				}{
					I0: I0,
					I1: I1,
				})
			}
		}
	}

	return buff.String() + `
// gbADDHL Adds byte from [HL] to A
func gbADDHL(cpu *Core) {
    z := cpu.Memory.ReadByte(cpu.Registers.HL())
    sum := uint16(cpu.Registers.A) + uint16(z)

    cpu.Registers.SetCarry(sum > 255)
    cpu.Registers.SetZero(sum & 0xFF == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry((cpu.Registers.A & 0xF) + (z & 0xF) > 0xF)

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
// gbADDn Adds byte from [PC] to A
func gbADDn(cpu *Core) {
    z := cpu.Memory.ReadByte(cpu.Registers.PC)
    sum := uint16(cpu.Registers.A) + uint16(z)

    cpu.Registers.SetCarry(sum > 255)
    cpu.Registers.SetZero(sum & 0xFF == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry((cpu.Registers.A & 0xF) + (z & 0xF) > 0xF)

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// gbADDHLSP Adds SP to HL
func gbADDHLSP(cpu *Core) {
    sum := cpu.Registers.HL() + cpu.Registers.SP
    cpu.Registers.SetCarry(sum > 65535)
    //cpu.Registers.SetZero(su & 0xFF == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(((cpu.Registers.SP & 0xFFF) + (cpu.Registers.HL() & 0xFFF)) > 0xFFF)

    cpu.Registers.H = uint8(sum >> 8)
    cpu.Registers.L = uint8(sum & 0xFF)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// gbADDSPn Reads a signed byte from [PC] and adds to SP
func gbADDSPn(cpu *Core) {
    a := int(cpu.Memory.ReadByte(cpu.Registers.PC))
    cpu.Registers.PC++
    a = (a << 24) >> 24 // Convert unsigned byte to signed

    cpu.Registers.SetZero(false)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetCarry(int(cpu.Registers.SP & 0xFF) + (a & 0xFF) > 0xFF)
    cpu.Registers.SetHalfCarry(int(cpu.Registers.SP & 0xF) + (a & 0xF) > 0xF)

    cpu.Registers.SP = uint16(int(cpu.Registers.SP) + a)


    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}

// gbADCHL Sets A to A + [HL] + FlagCarry
func gbADCHL(cpu *Core) {
    a := int(cpu.Registers.A)
    b := int(cpu.Memory.ReadByte(cpu.Registers.HL()))

    f := 0
    if cpu.Registers.GetCarry() {
        f = 1
    }

    sum := a + b + f

    cpu.Registers.SetZero(sum == 0)
    cpu.Registers.SetCarry(sum > 255)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(int(a & 0xF) + (b & 0xF) + f > 0xF)

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// gbADCn Sets A to A + [PC] + FlagCarry
func gbADCn(cpu *Core) {
    a := int(cpu.Registers.A)
    b := int(cpu.Memory.ReadByte(cpu.Registers.PC))

    f := 0
    if cpu.Registers.GetCarry() {
        f = 1
    }
    cpu.Registers.PC++

    sum := a + b + f

    cpu.Registers.SetZero(sum == 0)
    cpu.Registers.SetCarry(sum > 255)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(int(a & 0xF) + (b & 0xF) + f > 0xF)

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`
}
func BuildSUB() string {
	//
	buff := bytes.NewBuffer(nil)

	for _, I := range AllRegisters {
		subTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		sbcrTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
	}

	for _, I0 := range AllRegisters {
		for _, I1 := range AllRegisters {
			if I0 != I1 {
				subHLrrTemplate.Execute(buff, struct {
					I0 string
					I1 string
				}{
					I0: I0,
					I1: I1,
				})
			}
		}
	}

	return buff.String() + `
// gbSUBHL Subtracts byte from [HL] to A
func gbSUBHL(cpu *Core) {
    z := cpu.Memory.ReadByte(cpu.Registers.HL())
    sum := int16(cpu.Registers.A) - int16(z)

    cpu.Registers.SetCarry(sum < 0)
    cpu.Registers.SetZero(sum & 0xFF == 0)
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry((cpu.Registers.A & 0xF) < (z & 0xF))

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
// gbSUBn Subtracts byte from [PC] to A
func gbSUBn(cpu *Core) {
    z := cpu.Memory.ReadByte(cpu.Registers.PC)
    sum := int16(cpu.Registers.A) - int16(z)

    cpu.Registers.SetCarry(sum < 0)
    cpu.Registers.SetZero(sum & 0xFF == 0)
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry((cpu.Registers.A & 0xF) < (z & 0xF) )

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// gbSUBHLSP Subtracts SP from HL
func gbSUBHLSP(cpu *Core) {
    sum := int(cpu.Registers.HL()) - int(cpu.Registers.SP)
    cpu.Registers.SetCarry(sum < 0)
    cpu.Registers.SetZero(sum & 0xFFFF == 0)
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry(((cpu.Registers.SP & 0xFFF) < (cpu.Registers.HL() & 0xFFF)))

    cpu.Registers.H = uint8(sum >> 8)
    cpu.Registers.L = uint8(sum & 0xFF)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// gbSBCHL Sets A to A - [HL] - FlagCarry
func gbSBCHL(cpu *Core) {
    a := int(cpu.Registers.A)
    b := int(cpu.Memory.ReadByte(cpu.Registers.HL()))

    f := 0
    if cpu.Registers.GetCarry() {
        f = 1
    }

    sum := a - b - f

    cpu.Registers.SetZero(sum == 0)
    cpu.Registers.SetCarry(sum < 0)
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry(int(a & 0xF) < (b & 0xF) + f)

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// gbSBCn Sets A to A - [PC] - FlagCarry
func gbSBCn(cpu *Core) {
    a := int(cpu.Registers.A)
    b := int(cpu.Memory.ReadByte(cpu.Registers.PC))

    f := 0
    if cpu.Registers.GetCarry() {
        f = 1
    }
    cpu.Registers.PC++

    sum := a - b - f

    cpu.Registers.SetZero(sum == 0)
    cpu.Registers.SetCarry(sum < 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(int(a & 0xF) < (b & 0xF) + f)

    cpu.Registers.A = uint8(sum)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`
}
func BuildCP() string {
	//
	buff := bytes.NewBuffer(nil)

	for _, I := range AllRegisters {
		cpTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
	}

	return buff.String() + `
// gbCPHL Compares byte from [HL] to A
func gbCPHL(cpu *Core) {
    b := cpu.Memory.ReadByte(cpu.Registers.HL())

    cpu.Registers.SetCarry(cpu.Registers.A < b)
    cpu.Registers.SetZero(cpu.Registers.A == b)
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry((cpu.Registers.A & 0xF) < (b & 0xF))

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

// gbCPn Compares byte from [PC] to A
func gbCPn(cpu *Core) {
    b := cpu.Memory.ReadByte(cpu.Registers.PC)
    cpu.Registers.PC++

    cpu.Registers.SetCarry(cpu.Registers.A < b)
    cpu.Registers.SetZero(cpu.Registers.A == b)
    cpu.Registers.SetSub(true)
    cpu.Registers.SetHalfCarry((cpu.Registers.A & 0xF) < (b & 0xF))

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

`
}
func BuildOperators() string {
	buff := bytes.NewBuffer(nil)

	for _, I := range AllRegisters {
		andTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		orTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		xorTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
	}

	return buff.String() + `

// gbDAA
func gbDAA(cpu *Core) {
	a := int(cpu.Registers.A)

	if cpu.Registers.GetSub() {
		if cpu.Registers.GetHalfCarry() {
			a -= 0x6
		} else {
			a -= 0x60
		}
	} else {
		if cpu.Registers.GetHalfCarry() || (a & 0xF) > 0x9 {
			a += 0x6
		} else {
			a += 0x60
		}
	}

	cpu.Registers.A = uint8(a)
	

	cpu.Registers.SetZero(a == 0)
	cpu.Registers.SetHalfCarry(false)
	
	if a & 0x100 == 0x100 {
		cpu.Registers.SetCarry(true)
	}

	cpu.Registers.LastClockM = 1
	cpu.Registers.LastClockT = 4
}

// gbANDHL Sets A to A & [HL]
func gbANDHL(cpu *Core) {
	cpu.Registers.A &= cpu.Memory.ReadByte(cpu.Registers.HL())

	cpu.Registers.SetCarry(false)
	cpu.Registers.SetZero(cpu.Registers.A == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(true)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// gbANDn Sets A to A & [PC]
func gbANDn(cpu *Core) {
	cpu.Registers.A &= cpu.Memory.ReadByte(cpu.Registers.PC)
	cpu.Registers.PC++

	cpu.Registers.SetCarry(false)
	cpu.Registers.SetZero(cpu.Registers.A == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(true)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// gbORHL Sets A to A | [HL]
func gbORHL(cpu *Core) {
	cpu.Registers.A |= cpu.Memory.ReadByte(cpu.Registers.HL())

	cpu.Registers.SetCarry(false)
	cpu.Registers.SetZero(cpu.Registers.A == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// gbORn Sets A to A | [PC]
func gbORn(cpu *Core) {
	cpu.Registers.A |= cpu.Memory.ReadByte(cpu.Registers.PC)
	cpu.Registers.PC++

	cpu.Registers.SetCarry(false)
	cpu.Registers.SetZero(cpu.Registers.A == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// gbXORHL Sets A to A ^ [HL]
func gbXORHL(cpu *Core) {
	cpu.Registers.A ^= cpu.Memory.ReadByte(cpu.Registers.HL())

	cpu.Registers.SetCarry(false)
	cpu.Registers.SetZero(cpu.Registers.A == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}

// gbXORn Sets A to A ^ [PC]
func gbXORn(cpu *Core) {
	cpu.Registers.A ^= cpu.Memory.ReadByte(cpu.Registers.PC)
	cpu.Registers.PC++

	cpu.Registers.SetCarry(false)
	cpu.Registers.SetZero(cpu.Registers.A == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(false)

	cpu.Registers.LastClockM = 2
	cpu.Registers.LastClockT = 8
}
`
}
func BuildIncDec() string {
	buff := bytes.NewBuffer(nil)

	for _, I := range AllRegisters {
		incTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
		decTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
	}

	for _, I0 := range AllRegisters {
		for _, I1 := range AllRegisters {
			incrrTemplate.Execute(buff, struct {
				I0 string
				I1 string
			}{
				I0: I0,
				I1: I1,
			})
			decrrTemplate.Execute(buff, struct {
				I0 string
				I1 string
			}{
				I0: I0,
				I1: I1,
			})
		}
	}

	return buff.String() + `

// gbINCHLm Sets [HL] to [HL] + 1
func gbINCHLm(cpu *Core) {
	v := int(cpu.Memory.ReadByte(cpu.Registers.HL()))
	cpu.Memory.WriteByte(cpu.Registers.HL(), (byte)(v+1))

	cpu.Registers.SetCarry((v+1) > 255)
	cpu.Registers.SetZero((v+1) & 0xFF == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry((v & 0xF) + 1 > 0xF)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// gbDECHLm Sets [HL] to [HL] - 1
func gbDECHLm(cpu *Core) {
	v := int(cpu.Memory.ReadByte(cpu.Registers.HL()))
	cpu.Memory.WriteByte(cpu.Registers.HL(), (byte)(v-1))

	cpu.Registers.SetCarry((v-1) < 0)
	cpu.Registers.SetZero((v-1) & 0xFF == 0)
	cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry((v & 0xF) == 0)

	cpu.Registers.LastClockM = 3
	cpu.Registers.LastClockT = 12
}

// gbDECSP Sets SP = SP - 1
func gbDECSP(cpu *Core) {
    cpu.Registers.SP--

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
// gbINCSP Sets SP = SP + 1
func gbINCSP(cpu *Core) {
    cpu.Registers.SP++

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

`
}

// endregion
