package gendata

func BuildInterruptCalls() string {
	return `
// gbRSTXX Triggers the Interrupt at XX
func gbRSTXX(cpu *Core, addr uint16) {
    cpu.Registers.SaveRegisters()
    cpu.Registers.SP -= 2
    cpu.Memory.WriteWord(cpu.Registers.SP, cpu.Registers.PC)
    cpu.Registers.PC = addr
    
    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}
`
}
