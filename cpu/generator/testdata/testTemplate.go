package testdata

import (
    "bytes"
    "text/template"
)

var testCPUTemplate = template.Must(template.New("Test Template").Parse(`
func TestOpcode{{.OPCODE}}(t *testing.T) {
    cpu := MakeCore()
    cpu.Reset()
    cpu.Registers.Randomize()
    cpu.Memory.Randomize()

    RegBefore := cpu.Registers.Clone()
    GBInstructions[0x{{.OPCODE}}](cpu)
    RegAfter := cpu.Registers.Clone()

    {{.CHECKS}}
}`))

var cbTestTemplate = template.Must(template.New("CB Test Template").Parse(`
func TestCBOpcode{{.OPCODE}}(t *testing.T) {
    cpu := MakeCore()
    cpu.Reset()
    cpu.Registers.Randomize()
    cpu.Memory.Randomize()

    RegBefore := cpu.Registers.Clone()
    CBInstructions[0x{{.OPCODE}}](cpu)
    RegAfter := cpu.Registers.Clone()

    {{.CHECKS}}
}`))

var cycleTestTemplate = template.Must(template.New("Cycle Test Template").Parse(`
    // region Test Cycles
    if RegAfter.LastClockT != {{.LASTCLOCKT}} {
        t.Errorf("Expected LastClockT to be %d but got %d", {{.LASTCLOCKT}}, RegAfter.LastClockT)
    }
    if RegAfter.LastClockM != {{.LASTCLOCKM}} {
        t.Errorf("Expected LastClockM to be %d but got %d", {{.LASTCLOCKM}}, RegAfter.LastClockM)
    }
    // endregion
`))

func GenCycleTest(cycles int) string {
    b := bytes.NewBuffer(nil)

    cycleTestTemplate.Execute(b, struct {
        LASTCLOCKT int
        LASTCLOCKM int
    }{
        LASTCLOCKT: cycles,
        LASTCLOCKM: cycles / 4,
    })

    return b.String()
}

func GenFlagTest(zshc string) string {
    test := `
    // region Test Flags`

    // region Flag Zero
    switch zshc[0] {
    case '-': // Does not change
        test += `
    if RegAfter.GetZero() != RegBefore.GetZero() {
        t.Errorf("Expected Flag Zero to not change")
    }`
    case '0': // Always exit as 0
        test += `
    if RegAfter.GetZero()  {
        t.Errorf("Expected Flag Zero to be zero")
    }`
    case '1': // Always exit as 1
        test += `
    if !RegAfter.GetZero(){
        t.Errorf("Expected Flag Zero to be one")
    }`
    }
    // endregion
    // region Flag Sub
    switch zshc[1] {
    case '-': // Does not change
        test += `
    if RegAfter.GetSub() != RegBefore.GetSub() {
        t.Errorf("Expected Flag Sub to not change")
    }`
    case '0': // Always exit as 0
        test += `
    if RegAfter.GetSub()  {
        t.Errorf("Expected Flag Sub to be zero")
    }`
    case '1': // Always exit as 1
        test += `
    if !RegAfter.GetSub(){
        t.Errorf("Expected Flag Sub to be one")
    }`
    }
    // endregion
    // region Flag Half Carry
    switch zshc[2] {
    case '-': // Does not change
        test += `
    if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
        t.Errorf("Expected Flag Half Carry to not change")
    }`
    case '0': // Always exit as 0
        test += `
    if RegAfter.GetHalfCarry()  {
        t.Errorf("Expected Flag Half Carry to be zero")
    }`
    case '1': // Always exit as 1
        test += `
    if !RegAfter.GetHalfCarry(){
        t.Errorf("Expected Flag Half Carry to be one")
    }`
    }
    // endregion
    // region Flag Carry
    switch zshc[3] {
    case '-': // Does not change
        test += `
    if RegAfter.GetCarry() != RegBefore.GetCarry() {
        t.Errorf("Expected Flag Carry to not change")
    }`
    case '0': // Always exit as 0
        test += `
    if RegAfter.GetCarry()  {
        t.Errorf("Expected Flag Carry to be zero")
    }`
    case '1': // Always exit as 1
        test += `
    if !RegAfter.GetCarry(){
        t.Errorf("Expected Flag Carry to be one")
    }`
    }
    // endregion

    test += `
    // endregion
`

    return test
}
