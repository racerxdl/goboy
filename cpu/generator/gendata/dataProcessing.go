package gendata

import (
	"bytes"
	"github.com/racerxdl/goboy/cpu"
	"text/template"
)

// region Templates
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

// endregion
// region Builders

func BuildADD() string {
	//
	buff := bytes.NewBuffer(nil)

	for _, I := range cpu.AllRegisters {
		addTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
	}

	for _, I0 := range cpu.AllRegisters {
		for _, I1 := range cpu.AllRegisters {
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
// gbADDHL Adds byte from memory pointed by HL to A
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
// gbADDn Adds byte from memory pointed by PC to A
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
    //cpu.Registers.SetZero(sum & 0xFF == 0)
    cpu.Registers.SetSub(false)
    cpu.Registers.SetHalfCarry(((cpu.Registers.SP & 0xFFF) + (cpu.Registers.HL() & 0xFFF)) > 0xFFF)

    cpu.Registers.H = uint8(sum >> 8)
    cpu.Registers.L = uint8(sum & 0xFF)

    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// gbADDSPn Reads a signed byte from memory pointed by PC and adds to SP
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

`
}

// endregion
