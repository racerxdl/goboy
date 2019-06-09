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
`
}

/*
   private static void ADDn(CPU cpu) {
       var reg = cpu.reg;
       var z = (int) cpu.memory.ReadByte(reg.PC);
       reg.PC++;
       var sum = reg.A + z;

       reg.FlagCarry = sum > 255;
       reg.FlagZero = (sum & 0xFF) == 0;
       reg.FlagSub = false;
       reg.FlagHalfCarry = ((reg.A & 0xF) + (z & 0xF)) > 0xF;

       reg.A = (byte) sum;

       reg.lastClockM = 2;
       reg.lastClockT = 8;
   }
*/

// endregion
