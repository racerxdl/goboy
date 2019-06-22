package testdata

import "html/template"

var testCPUTemplate = template.Must(template.New("Test Template").Parse(`
func TestOpcode{{.OPCODE}}(t *testing.T) {
    cpu := MakeCore()
    cpu.Reset()
    cpu.Registers.Randomize()
    cpu.Memory.Randomize()

    regBefore := cpu.Registers.Clone()
    GBInstructions[0x{{.OPCODE}}](cpu)
    regAfter := cpu.Registers.Clone()

    {{.CHECKS}}
}`))

var cbTestTemplate = template.Must(template.New("CB Test Template").Parse(`
func TestCBOpcode{{.OPCODE}}(t *testing.T) {
    cpu := MakeCore()
    cpu.Reset()
    cpu.Registers.Randomize()
    cpu.Memory.Randomize()

    regBefore := cpu.Registers.Clone()
    CBInstructions[0x{{.OPCODE}}](cpu)
    regAfter := cpu.Registers.Clone()

    {{.CHECKS}}
}`))

var cycleTestTemplate = template.Must(template.New("Cycle Test Template").Parse(`
    // region Test Cycles
    if regAfter.LastClockT != {{.LASTCLOCKT}} {
        t.Errorf("Expected LastClockT to be %d but got %d", {{.LASTCLOCKT}}, regAfter.LastClockT)
    }
    if regAfter.LastClockM != {{.LASTCLOCKM}} {
        t.Errorf("Expected LastClockM to be %d but got %d", {{.LASTCLOCKM}}, regAfter.LastClockM)
    }
    // endregion
`))