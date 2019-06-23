package gendata

func BuildFlowControl() string {
	return `
   
func gbJPnn(cpu *Core) {
    cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)
    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}

func gbJPHL(cpu *Core) {
    cpu.Registers.PC = cpu.Registers.HL()
    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gbJPNZnn(cpu *Core) {
    if cpu.Registers.GetZero() {
        cpu.Registers.LastClockM = 3
        cpu.Registers.LastClockT = 12
        cpu.Registers.PC += 2
    } else {
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)
        cpu.Registers.LastClockM = 4
        cpu.Registers.LastClockT = 16
    }
}

func gbJPZnn(cpu *Core) {
    if !cpu.Registers.GetZero() {
        cpu.Registers.LastClockM = 3
        cpu.Registers.LastClockT = 12
        cpu.Registers.PC += 2
    } else {
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)
        cpu.Registers.LastClockM = 4
        cpu.Registers.LastClockT = 16
    }
}

func gbJPNCnn(cpu *Core) {
    if cpu.Registers.GetCarry() {
        cpu.Registers.LastClockM = 3
        cpu.Registers.LastClockT = 12
        cpu.Registers.PC += 2
    } else {
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)
        cpu.Registers.LastClockM = 4
        cpu.Registers.LastClockT = 16
    }
}

func gbJPCnn(cpu *Core) {
    if !cpu.Registers.GetCarry() {
        cpu.Registers.LastClockM = 3
        cpu.Registers.LastClockT = 12
        cpu.Registers.PC += 2
    } else {
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)
        cpu.Registers.LastClockM = 4
        cpu.Registers.LastClockT = 16
    }
}

func gbJRn(cpu *Core) {
    v := int(int8(cpu.Memory.ReadByte(cpu.Registers.PC)))
    cpu.Registers.PC++

    //if v > 127 {
    //    v = -((^v + 1) & 0xFF)
    //}

    cpu.Registers.PC = uint16(int(cpu.Registers.PC) + v)

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}

func gbJRNZn(cpu *Core) {
    v := int(int8(cpu.Memory.ReadByte(cpu.Registers.PC)))
    cpu.Registers.PC++

    //if v > 127 {
    //    v = -((^v + 1) & 0xFF)
    //}

    if cpu.Registers.GetZero() {
        cpu.Registers.LastClockM = 2
        cpu.Registers.LastClockT = 8
        return
    }

    cpu.Registers.PC = uint16(int(cpu.Registers.PC) + v)

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}

func gbJRZn(cpu *Core) {
    v := int(int8(cpu.Memory.ReadByte(cpu.Registers.PC)))
    cpu.Registers.PC++

    //if v > 127 {
    //    v = -((^v + 1) & 0xFF)
    //}

    if !cpu.Registers.GetZero() {
        cpu.Registers.LastClockM = 2
        cpu.Registers.LastClockT = 8
        return
    }

    cpu.Registers.PC = uint16(int(cpu.Registers.PC) + v)

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}

func gbJRNCn(cpu *Core) {
    v := int(int8(cpu.Memory.ReadByte(cpu.Registers.PC)))
    cpu.Registers.PC++

    //if v > 127 {
    //    v = -((^v + 1) & 0xFF)
    //}

    if cpu.Registers.GetCarry() {
        cpu.Registers.LastClockM = 2
        cpu.Registers.LastClockT = 8
        return
    }

    cpu.Registers.PC = uint16(int(cpu.Registers.PC) + v)

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}

func gbJRCn(cpu *Core) {
    v := int(int8(cpu.Memory.ReadByte(cpu.Registers.PC)))
    cpu.Registers.PC++

    //if v > 127 {
    //    v = -((^v + 1) & 0xFF)
    //}

    if cpu.Registers.GetCarry() {
        cpu.Registers.LastClockM = 2
        cpu.Registers.LastClockT = 8
        return
    }

    cpu.Registers.PC = uint16(int(cpu.Registers.PC) + v)

    cpu.Registers.LastClockM = 3
    cpu.Registers.LastClockT = 12
}

func gbStop(cpu *Core) {
    cpu.stopped = true
    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gbCALLnn(cpu *Core) {
    cpu.Registers.SP -= 2
    cpu.Memory.WriteWord(cpu.Registers.SP, cpu.Registers.PC + 2)
    cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)

    cpu.Registers.LastClockM = 6
    cpu.Registers.LastClockT = 24
}

func gbCALLNZnn(cpu *Core) {
    if cpu.Registers.GetZero() {
        cpu.Registers.PC += 2
        cpu.Registers.LastClockM = 3
        cpu.Registers.LastClockT = 12
    } else {
        cpu.Registers.SP -= 2
        cpu.Memory.WriteWord(cpu.Registers.SP, cpu.Registers.PC+2)
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)

        cpu.Registers.LastClockM = 6
        cpu.Registers.LastClockT = 24
    }
}

func gbCALLZnn(cpu *Core) {
    if !cpu.Registers.GetZero() {
        cpu.Registers.PC += 2
        cpu.Registers.LastClockM = 3
        cpu.Registers.LastClockT = 12
    } else {
        cpu.Registers.SP -= 2
        cpu.Memory.WriteWord(cpu.Registers.SP, cpu.Registers.PC+2)
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)

        cpu.Registers.LastClockM = 6
        cpu.Registers.LastClockT = 24
    }
}

func gbCALLNCnn(cpu *Core) {
    if cpu.Registers.GetCarry() {
        cpu.Registers.PC += 2
        cpu.Registers.LastClockM = 3
        cpu.Registers.LastClockT = 12
    } else {
        cpu.Registers.SP -= 2
        cpu.Memory.WriteWord(cpu.Registers.SP, cpu.Registers.PC+2)
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)

        cpu.Registers.LastClockM = 6
        cpu.Registers.LastClockT = 24
    }
}

func gbCALLCnn(cpu *Core) {
    if !cpu.Registers.GetCarry() {
        cpu.Registers.PC += 2
        cpu.Registers.LastClockM = 3
        cpu.Registers.LastClockT = 12
    } else {
        cpu.Registers.SP -= 2
        cpu.Memory.WriteWord(cpu.Registers.SP, cpu.Registers.PC+2)
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.PC)

        cpu.Registers.LastClockM = 6
        cpu.Registers.LastClockT = 24
    }
}

func gbRET(cpu *Core) {
    cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.SP)
    cpu.Registers.SP += 2

    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}

func gbRETI(cpu *Core) {
    cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.SP)
    cpu.Registers.SP += 2
    cpu.Registers.InterruptEnable = true

    cpu.Registers.LastClockM = 4
    cpu.Registers.LastClockT = 16
}

func gbRETNZ(cpu *Core) {
    if cpu.Registers.GetZero() {
        cpu.Registers.LastClockM = 2
        cpu.Registers.LastClockT = 8
    } else {
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.SP)
        cpu.Registers.SP += 2

        cpu.Registers.LastClockM = 5
        cpu.Registers.LastClockT = 20
    }
}
func gbRETZ(cpu *Core) {
    if !cpu.Registers.GetZero() {
        cpu.Registers.LastClockM = 2
        cpu.Registers.LastClockT = 8
    } else {
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.SP)
        cpu.Registers.SP += 2

        cpu.Registers.LastClockM = 5
        cpu.Registers.LastClockT = 20
    }
}

func gbRETNC(cpu *Core) {
    if cpu.Registers.GetCarry() {
        cpu.Registers.LastClockM = 2
        cpu.Registers.LastClockT = 8
    } else {
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.SP)
        cpu.Registers.SP += 2

        cpu.Registers.LastClockM = 5
        cpu.Registers.LastClockT = 20
    }
}

func gbRETC(cpu *Core) {
    if !cpu.Registers.GetCarry() {
        cpu.Registers.LastClockM = 2
        cpu.Registers.LastClockT = 8
    } else {
        cpu.Registers.PC = cpu.Memory.ReadWord(cpu.Registers.SP)
        cpu.Registers.SP += 2

        cpu.Registers.LastClockM = 5
        cpu.Registers.LastClockT = 20
    }
}

func gbDI(cpu *Core) {
    cpu.Registers.InterruptEnable = false
    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gbEI(cpu *Core) {
    cpu.Registers.InterruptEnable = true
    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gbNOP(cpu *Core) {
    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gbNOPWARN(cpu *Core, opcode int) {
    cpuLog.Warn("Opcode not implemented: 0x%02x at 0x%04x", opcode, cpu.Registers.PC-1)
    cpu.Registers.LastClockM = 0
    cpu.Registers.LastClockT = 0
	cpu.paused = true
}

func gbHALT(cpu *Core) {
    cpu.halted = true
    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
 
`
}
