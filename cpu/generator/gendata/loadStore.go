package gendata

import (
	"bytes"
	"github.com/racerxdl/goboy/cpu"
	"text/template"
)

// region Templates
var ldrrTemplate = template.Must(template.New("LDrrIO").Parse(`
// gbLDrr{{.I}}{{.O}} Sets Register {{.O}} to the value in {{.I}}
func gbLDrr{{.I}}{{.O}}(cpu *Core) {
    cpu.Registers.{{.I}} = cpu.Registers.{{.O}}
    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

var ldrHLmTemplate = template.Must(template.New("LDrHLmO").Parse(`
// gbLDrHLm{{.O}} Loads to Register {{.O}} the value in memory pointed by HL
func gbLDrHLm{{.O}}(cpu *Core) {
    cpu.Registers.{{.O}} = cpu.Memory.ReadByte(cpu.Registers.HL())
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var ldHLmrTemplate = template.Must(template.New("LDrHLmO").Parse(`
// gbLDHLmr{{.I}} Writes register {{.I}} to memory pointed by HL
func gbLDHLmr{{.I}}(cpu *Core) {
    cpu.Memory.WriteByte(cpu.Registers.HL(), cpu.Registers.{{.I}})
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var ldrnTemplate = template.Must(template.New("LDrn").Parse(`
// gbLDrn{{.O}} Loads a byte from Program Memory into {{.O}}. Increments PC
func gbLDrn{{.O}}(cpu *Core) {
    cpu.Registers.{{.O}} = cpu.Memory.ReadByte(cpu.Registers.PC)
    cpu.Registers.PC++
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var ldrrmrTemplate = template.Must(template.New("LDrrmr").Parse(`
// gbLD{{.H}}{{.L}}m{{.I}} Writes value of register {{.I}} to memory Pointed by {{.H}} << 8 + {{.L}}
func gbLD{{.H}}{{.L}}m{{.I}}(cpu *Core) {
    hl := (uint16(cpu.Registers.{{.H}}) << 8) + uint16(cpu.Registers.{{.L}})
    cpu.Memory.WriteByte(hl, cpu.Registers.{{.I}})
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var ldmmTemplate = template.Must(template.New("LDmm").Parse(`
// gbLDmm{{.I}} Writes register {{.I}} to memory pointed by PC
func gbLDmm{{.I}}(cpu *Core) {
    cpu.Memory.WriteByte(cpu.Registers.PC, cpu.Registers.{{.I}})
    cpu.Registers.PC += 2
    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}
`))

var ldrrrmTemplate = template.Must(template.New("LDrrrm").Parse(`
// gbLD{{.H}}{{.L}}{{.O}}m  Reads from memory pointed by {{.H}} << 8 + {{.L}} and writes to {{.O}} 
func gbLD{{.H}}{{.L}}{{.O}}m(cpu *Core) {
    hl := (uint16(cpu.Registers.{{.H}}) << 8) + uint16(cpu.Registers.{{.L}})
    cpu.Registers.{{.O}} = cpu.Memory.ReadByte(hl)
    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}
`))

var ldrmmTemplate = template.Must(template.New("LDrmm").Parse(`
// gbLD{{.O}}mm Writes register {{.O}} to memory pointed by PC
func gbLD{{.O}}mm(cpu *Core) {
    addr := cpu.Memory.ReadWord(cpu.Registers.PC)
    cpu.Registers.{{.O}} = cpu.Memory.ReadByte(addr)
    cpu.Registers.PC += 2
    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}
`))

var ldrrnnTemplate = template.Must(template.New("LDrrnn").Parse(`
// gbLD{{.O1}}{{.O2}}nn Reads from memory pointed by PC to {{.O2}} and PC+1 to {{.O1}}
func gbLD{{.O1}}{{.O2}}nn(cpu *Core) {
    
    cpu.Registers.{{.O2}} = cpu.Memory.ReadByte(cpu.Registers.PC)
    cpu.Registers.PC++
    cpu.Registers.{{.O1}} = cpu.Memory.ReadByte(cpu.Registers.PC)
    cpu.Registers.PC++
    
    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
`))

// endregion
// region Builders

func BuildLDRRnn() string {
	buff := bytes.NewBuffer(nil)

	for _, O1 := range cpu.AllRegisters {
		for _, O2 := range cpu.AllRegisters {
			if O1 != O2 {
				ldrrnnTemplate.Execute(buff, struct {
					O1 string
					O2 string
				}{
					O1: O1,
					O2: O2,
				})
			}
		}
	}

	return buff.String()
}

func BuildLDrmm() string {
	//
	buff := bytes.NewBuffer(nil)

	for _, O := range cpu.AllRegisters {
		ldrmmTemplate.Execute(buff, struct {
			O string
		}{
			O: O,
		})
	}

	return buff.String()
}

func BuildLDrrrm() string {
	buff := bytes.NewBuffer(nil)

	for _, O := range cpu.AllRegisters {
		for _, H := range cpu.AllRegisters {
			for _, L := range cpu.AllRegisters {
				if H != L {
					ldrrrmTemplate.Execute(buff, struct {
						O string
						H string
						L string
					}{
						O: O,
						H: H,
						L: L,
					})
				}
			}
		}
	}

	return buff.String()
}

func BuildLDmm() string {
	//
	buff := bytes.NewBuffer(nil)

	for _, I := range cpu.AllRegisters {
		ldmmTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
	}

	return buff.String()
}

func BuildLDHLmr() string {
	//
	buff := bytes.NewBuffer(nil)

	for _, I := range cpu.AllRegisters {
		ldHLmrTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
	}

	return buff.String()
}

func BuildLDrHLm() string {
	//
	buff := bytes.NewBuffer(nil)

	for _, O := range cpu.AllRegisters {
		ldrHLmTemplate.Execute(buff, struct {
			O string
		}{
			O: O,
		})
	}

	return buff.String()
}

func BuildLDrn() string {
	//
	buff := bytes.NewBuffer(nil)

	for _, O := range cpu.AllRegisters {
		ldrnTemplate.Execute(buff, struct {
			O string
		}{
			O: O,
		})
	}

	return buff.String()
}

func BuildLDRR() string {
	buff := bytes.NewBuffer(nil)

	for _, I := range cpu.AllRegisters {
		for _, O := range cpu.AllRegisters {
			if I != O {
				ldrrTemplate.Execute(buff, struct {
					I string
					O string
				}{
					I: I,
					O: O,
				})
			}
		}
	}

	return buff.String()
}

func BuildLDrrmr() string {
	buff := bytes.NewBuffer(nil)

	for _, I := range cpu.AllRegisters {
		for _, H := range cpu.AllRegisters {
			for _, L := range cpu.AllRegisters {
				if H != L {
					ldrrmrTemplate.Execute(buff, struct {
						I string
						H string
						L string
					}{
						I: I,
						H: H,
						L: L,
					})
				}
			}
		}
	}

	return buff.String()
}

func BuildLDHLmn() string {
	return `
// LDHLmn Writes byte from Program Memory into Memory (H/L). Increments Program Counter
func LDHLmn(cpu *Core) {
    cpu.Memory.WriteByte(cpu.Registers.HL(), cpu.Memory.ReadByte(cpu.Registers.PC))
    cpu.Registers.PC++
    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}`
}

// endregion
