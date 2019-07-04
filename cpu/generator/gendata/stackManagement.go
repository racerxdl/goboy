package gendata

import (
	"bytes"
	"text/template"
)

func BuildStackManagement() string {
	b := bytes.NewBuffer(nil)

	regs := [][2]string{
		{"A", "F"},
		{"B", "C"},
		{"D", "E"},
		{"H", "L"},
	}

	for _, I := range regs {
		pushTemplate.Execute(b, struct {
			I0 string
			I1 string
		}{
			I0: I[0],
			I1: I[1],
		})
		popTemplate.Execute(b, struct {
			I0 string
			I1 string
		}{
			I0: I[0],
			I1: I[1],
		})
	}

	//AllRegisters := []string{
	//	"A", "B", "C", "D", "E", "H", "L", "F",
	//}
	//
	//for _, I0 := range AllRegisters {
	//	for _, I1 := range AllRegisters {
	//		if I0 != I1 {
	//			pushTemplate.Execute(b, struct {
	//				I0 string
	//				I1 string
	//			}{
	//				I0: I0,
	//				I1: I1,
	//			})
	//			popTemplate.Execute(b, struct {
	//				I0 string
	//				I1 string
	//			}{
	//				I0: I0,
	//				I1: I1,
	//			})
	//		}
	//	}
	//}

	return b.String()
}

var pushTemplate = template.Must(template.New("PUSH").Parse(`
// gbPUSH{{.I0}}{{.I1}} Writes {{.I0}} and {{.I1}} to the Stack
func gbPUSH{{.I0}}{{.I1}}(cpu *Core) {
    cpu.Registers.SP--
    cpu.Memory.WriteByte(cpu.Registers.SP, cpu.Registers.{{.I0}})
    cpu.Registers.SP--
    cpu.Memory.WriteByte(cpu.Registers.SP, cpu.Registers.{{.I1}})

    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}
`))
var popTemplate = template.Must(template.New("POP").Parse(`
// gbPOP{{.I0}}{{.I1}} Reads {{.I0}} and {{.I1}} to the Stack
func gbPOP{{.I0}}{{.I1}}(cpu *Core) {
    
    cpu.Registers.{{.I1}} = cpu.Memory.ReadByte(cpu.Registers.SP)
    cpu.Registers.SP++
    cpu.Registers.{{.I0}} = cpu.Memory.ReadByte(cpu.Registers.SP)
    cpu.Registers.SP++

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}
`))
