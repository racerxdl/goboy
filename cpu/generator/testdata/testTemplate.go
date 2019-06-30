package testdata

import "text/template"

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

func GenFlagTest(zshc string) string {
    test := `
    // region Test Flags`

    // region Flag Zero
    switch zshc[0] {
    case '-': // Does not change
        test += `
    if regAfter.GetZero() != regBefore.GetZero() {
        t.Errorf("Expected Flag Zero to not change")
    }`
    case '0': // Always exit as 0
        test += `
    if regAfter.GetZero()  {
        t.Errorf("Expected Flag Zero to be zero")
    }`
    case '1': // Always exit as 1
        test += `
    if !regAfter.GetZero(){
        t.Errorf("Expected Flag Zero to be one")
    }`
    }
    // endregion
    // region Flag Sub
    switch zshc[1] {
    case '-': // Does not change
        test += `
    if regAfter.GetSub() != regBefore.GetSub() {
        t.Errorf("Expected Flag Sub to not change")
    }`
    case '0': // Always exit as 0
        test += `
    if regAfter.GetSub()  {
        t.Errorf("Expected Flag Sub to be zero")
    }`
    case '1': // Always exit as 1
        test += `
    if !regAfter.GetSub(){
        t.Errorf("Expected Flag Sub to be one")
    }`
    }
    // endregion
    // region Flag Half Carry
    switch zshc[2] {
    case '-': // Does not change
        test += `
    if regAfter.GetHalfCarry() != regBefore.GetHalfCarry() {
        t.Errorf("Expected Flag Half Carry to not change")
    }`
    case '0': // Always exit as 0
        test += `
    if regAfter.GetHalfCarry()  {
        t.Errorf("Expected Flag Half Carry to be zero")
    }`
    case '1': // Always exit as 1
        test += `
    if !regAfter.GetHalfCarry(){
        t.Errorf("Expected Flag Half Carry to be one")
    }`
    }
    // endregion
    // region Flag Carry
    switch zshc[3] {
    case '-': // Does not change
        test += `
    if regAfter.GetCarry() != regBefore.GetCarry() {
        t.Errorf("Expected Flag Carry to not change")
    }`
    case '0': // Always exit as 0
        test += `
    if regAfter.GetCarry()  {
        t.Errorf("Expected Flag Carry to be zero")
    }`
    case '1': // Always exit as 1
        test += `
    if !regAfter.GetCarry(){
        t.Errorf("Expected Flag Carry to be one")
    }`
    }
    // endregion

    test += `
    // endregion
`

    return test
}
