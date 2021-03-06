package gendata

import (
	"bytes"
	"text/template"
)

// region Templates
var ldrrTemplate = template.Must(template.New("LDrr").Parse(`
// gbLDrr{{.I}}{{.O}} Sets Register {{.O}} to the value in {{.I}}
func gbLDrr{{.I}}{{.O}}(cpu *Core) {
    cpu.Registers.{{.I}} = cpu.Registers.{{.O}}
    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
`))

var ldrHLmTemplate = template.Must(template.New("LDrHLm").Parse(`
// gbLDrHLm{{.O}} Sets {{.O}} to [HL]
func gbLDrHLm{{.O}}(cpu *Core) {
    cpu.Registers.{{.O}} = cpu.Memory.ReadByte(cpu.Registers.HL())
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var ldHLmrTemplate = template.Must(template.New("LDHLmr").Parse(`
// gbLDHLmr{{.I}} Sets [HL] to {{.I}}
func gbLDHLmr{{.I}}(cpu *Core) {
    cpu.Memory.WriteByte(cpu.Registers.HL(), cpu.Registers.{{.I}})
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var ldrnTemplate = template.Must(template.New("LDrn").Parse(`
// gbLDrn{{.O}} Loads a byte from Program Memory into {{.O}}. Increments PC
func gbLDrn{{.O}}(cpu *Core) {
    cpu.Registers.{{.O}} = cpu.Memory.ReadByteForPC(cpu.Registers.PC)
    cpu.Registers.PC++
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var ldrrmrTemplate = template.Must(template.New("LDrrmr").Parse(`
// gbLD{{.H}}{{.L}}m{{.I}} Writes value of register {{.I}} to [{{.H}} << 8 + {{.L}}]
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
    cpu.Memory.WriteByte(cpu.Memory.ReadWordForPC(cpu.Registers.PC), cpu.Registers.{{.I}})
    cpu.Registers.PC += 2
    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}
`))

var ldrrrmTemplate = template.Must(template.New("LDrrrm").Parse(`
// gbLD{{.O}}{{.H}}{{.L}}m  Sets {{.O}} to [{{.H}} << 8 + {{.L}}]
func gbLD{{.O}}{{.H}}{{.L}}m(cpu *Core) {
    hl := (uint16(cpu.Registers.{{.H}}) << 8) + uint16(cpu.Registers.{{.L}})
    cpu.Registers.{{.O}} = cpu.Memory.ReadByte(hl)
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var ldrmmTemplate = template.Must(template.New("LDrmm").Parse(`
// gbLD{{.O}}mm Writes register {{.O}} to memory pointed by PC
func gbLD{{.O}}mm(cpu *Core) {
    addr := cpu.Memory.ReadWordForPC(cpu.Registers.PC)
    cpu.Registers.{{.O}} = cpu.Memory.ReadByte(addr)
    cpu.Registers.PC += 2
    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}
`))

var ldrrnnTemplate = template.Must(template.New("LDrrnn").Parse(`
// gbLD{{.O1}}{{.O2}}nn Reads [PC] to {{.O2}} and [PC+1] to {{.O1}}
func gbLD{{.O1}}{{.O2}}nn(cpu *Core) {

    cpu.Registers.{{.O2}} = cpu.Memory.ReadByteForPC(cpu.Registers.PC)
    cpu.Registers.PC++
    cpu.Registers.{{.O1}} = cpu.Memory.ReadByteForPC(cpu.Registers.PC)
    cpu.Registers.PC++

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
`))

var ldhliTemplate = template.Must(template.New("LDHLI").Parse(`
// gbLDHLI{{.I}} Sets {{.I}} to [HL] and increments HL.
func gbLDHLI{{.I}}(cpu *Core) {
    cpu.Memory.WriteByte(cpu.Registers.HL(), cpu.Registers.{{.I}})
    cpu.Registers.L++
    if cpu.Registers.L == 0 {
        cpu.Registers.H++
    }
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}
`))

var ldrIOnTemplate = template.Must(template.New("LDRIOn").Parse(`
// gbLD{{.O}}IOn
func gbLD{{.O}}IOn(cpu *Core) {
    addr := cpu.Memory.ReadByteForPC(cpu.Registers.PC)
    cpu.Registers.PC++
    cpu.Registers.{{.O}} = cpu.Memory.ReadByte(0xFF00 + uint16(addr))

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
`))

var ldIOnrTemplate = template.Must(template.New("LDIOnR").Parse(`
// gbLDIOn{{.O}}
func gbLDIOn{{.O}}(cpu *Core) {
    addr := cpu.Memory.ReadByteForPC(cpu.Registers.PC)
    cpu.Registers.PC++
	cpu.Memory.WriteByte(0xFF00 + uint16(addr), cpu.Registers.{{.O}})

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
`))

// endregion
// region Builders

func BuildLDrIOn() string {
	//
	buff := bytes.NewBuffer(nil)

	regs := []string{
		"A",
	}

	for _, O := range regs {
		ldrIOnTemplate.Execute(buff, struct {
			O string
		}{
			O: O,
		})
		ldIOnrTemplate.Execute(buff, struct {
			O string
		}{
			O: O,
		})
	}

	return buff.String()
}

func BuildLDHLI() string {
	//
	buff := bytes.NewBuffer(nil)

	regs := []string{
		"A",
	}

	for _, I := range regs {
		ldhliTemplate.Execute(buff, struct {
			I string
		}{
			I: I,
		})
	}

	return buff.String()
}

func BuildLDRRnn() string {
	buff := bytes.NewBuffer(nil)

	regs := [][2]string{
		{"H", "L"},
		{"B", "C"},
		{"D", "E"},
	}

	for _, O := range regs {
		ldrrnnTemplate.Execute(buff, struct {
			O1 string
			O2 string
		}{
			O1: O[0],
			O2: O[1],
		})
	}

	//for _, O1 := range AllRegisters {
	//    for _, O2 := range AllRegisters {
	//        if O1 != O2 {
	//            ldrrnnTemplate.Execute(buff, struct {
	//                O1 string
	//                O2 string
	//            }{
	//                O1: O1,
	//                O2: O2,
	//            })
	//        }
	//    }
	//}

	return buff.String()
}

func BuildLDrmm() string {
	//
	buff := bytes.NewBuffer(nil)

	regs := []string{
		"A",
	}

	for _, O := range regs {
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

	regs := [][3]string{
		{"A", "B", "C"},
		{"A", "D", "E"},
	}

	for _, I := range regs {
		ldrrrmTemplate.Execute(buff, struct {
			O string
			H string
			L string
		}{
			O: I[0],
			H: I[1],
			L: I[2],
		})
	}

	//
	//for _, O := range AllRegisters {
	//    for _, H := range AllRegisters {
	//        for _, L := range AllRegisters {
	//            if H != L && H != O && L != O {
	//                ldrrrmTemplate.Execute(buff, struct {
	//                    O string
	//                    H string
	//                    L string
	//                }{
	//                    O: O,
	//                    H: H,
	//                    L: L,
	//                })
	//            }
	//        }
	//    }
	//}

	return buff.String()
}

func BuildLDmm() string {
	buff := bytes.NewBuffer(nil)

	regs := []string{
		"A",
	}

	for _, I := range regs {
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

	for _, I := range AllRegisters {
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

	for _, O := range AllRegisters {
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

	for _, O := range AllRegisters {
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

	for _, I := range AllRegisters {
		for _, O := range AllRegisters {
			ldrrTemplate.Execute(buff, struct {
				I string
				O string
			}{
				I: I,
				O: O,
			})
		}
	}

	return buff.String()
}

func BuildLDrrmr() string {
	buff := bytes.NewBuffer(nil)

	regs := [][3]string{
		{"B", "C", "A"}, // LDBCmA
		{"D", "E", "A"}, // LDDEmA
	}

	for _, I := range regs {
		ldrrmrTemplate.Execute(buff, struct {
			I string
			H string
			L string
		}{
			I: I[2],
			H: I[0],
			L: I[1],
		})
	}

	//for _, I := range AllRegisters {
	//	for _, H := range AllRegisters {
	//		for _, L := range AllRegisters {
	//			if H != L && H != I && L != I && // Avoid repeated procs
	//				L != "H" && H != "L" { // LHmX does not exists
	//				ldrrmrTemplate.Execute(buff, struct {
	//					I string
	//					H string
	//					L string
	//				}{
	//					I: I,
	//					H: H,
	//					L: L,
	//				})
	//			}
	//		}
	//	}
	//}

	return buff.String()
}

func BuildLSSingles() string {
	return `
// LDHLmn Writes byte from Program Memory into Memory (H/L). Increments Program Counter
func gbLDHLmn(cpu *Core) {
    cpu.Memory.WriteByte(cpu.Registers.HL(), cpu.Memory.ReadByteForPC(cpu.Registers.PC))
    cpu.Registers.PC++
    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
// LDSPnn Reads word from Program Counter and stores in SP
func gbLDSPnn(cpu *Core) {
    cpu.Registers.SP = cpu.Memory.ReadWordForPC(cpu.Registers.PC)
    cpu.Registers.PC += 2
    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
// LDmmSP
func gbLDmmSP(cpu *Core) {
    addr := cpu.Memory.ReadWordForPC(cpu.Registers.PC)
    cpu.Memory.WriteWord(addr, cpu.Registers.SP)
    cpu.Registers.PC += 2
    cpu.Registers.LastClockM = 5
    cpu.Registers.LastClockT = 20
}
// LDAIOC
func gbLDAIOC(cpu *Core) {
    cpu.Registers.A = cpu.Memory.ReadByte(0xFF00 + uint16(cpu.Registers.C))
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// LDIOCA
func gbLDIOCA(cpu *Core) {
    cpu.Memory.WriteByte(0xFF00 + uint16(cpu.Registers.C), cpu.Registers.A)
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// LDHLSPn
func gbLDHLSPn(cpu *Core) {
    v := int(int8(cpu.Memory.ReadByteForPC(cpu.Registers.PC)))
    cpu.Registers.PC++

    cpu.Registers.SetZero(false)
    cpu.Registers.SetSub(false)
	cpu.Registers.SetHalfCarry(int(cpu.Registers.SP&0xF) + (v & 0xF) > 0xF)
	cpu.Registers.SetCarry(int(cpu.Registers.SP&0xFF) + (v & 0xFF) > 0xFF)

    v += int(cpu.Registers.SP)

    cpu.Registers.H = uint8(uint(v) >> 8)
    cpu.Registers.L = uint8(uint(v) & 0xFF)

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}

// LDHLSPr
func gbLDHLSPr(cpu *Core) {
    cpu.Registers.SP = cpu.Registers.HL()
    cpu.Registers.LastClockM = 2
    cpu.Registers.LastClockT = 8
}

// gbLDHLDA Sets A to [HL] and decrements HL.
func gbLDHLDA(cpu *Core) {
  cpu.Memory.WriteByte(cpu.Registers.HL(), cpu.Registers.A)
  cpu.Registers.L--
  if cpu.Registers.L == 255 {
    cpu.Registers.H--
  }

  cpu.Registers.LastClockM = 2
  cpu.Registers.LastClockT = 8
}

// gbLDAHLI Reads byte from [HL] into A and increments H/L
func gbLDAHLI(cpu *Core) {
  cpu.Registers.A = cpu.Memory.ReadByte(cpu.Registers.HL())
  cpu.Registers.L++
  if cpu.Registers.L == 0 {
    cpu.Registers.H++
  }

  cpu.Registers.LastClockM = 2
  cpu.Registers.LastClockT = 8
}

// gbLDAHLD Reads byte from [HL] into A and decrements H/L
func gbLDAHLD(cpu *Core) {
  cpu.Registers.A = cpu.Memory.ReadByte(cpu.Registers.HL())
  cpu.Registers.L--
  if cpu.Registers.L == 255 {
    cpu.Registers.H--
  }

  cpu.Registers.LastClockM = 2
  cpu.Registers.LastClockT = 8
}
`
}

// endregion
