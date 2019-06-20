package gendata

func BuildBitManipulation() string {
	return `

func gbRLA(cpu *Core) {
    c := (cpu.Registers.A >> 7) > 0
    f := byte(0)
    if cpu.Registers.GetCarry() {
        f = 1
    }

    cpu.Registers.A = cpu.Registers.A << 1 | f

    cpu.Registers.SetCarry(c)
    cpu.Registers.SetZero(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetSub(false)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gnRLCA(cpu *Core) {
    c := (cpu.Registers.A >> 7) & 0x1
    cpu.Registers.A = cpu.Registers.A << 1 | c

    cpu.Registers.SetCarry(c > 0)
    cpu.Registers.SetZero(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetSub(false)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gnRRA(cpu *Core) {
    c := cpu.Registers.A & 0x1
    f := byte(0)
    if cpu.Registers.GetCarry() {
        f = 1
    }

    cpu.Registers.A = (cpu.Registers.A >> 1) | (f << 7)

    cpu.Registers.SetCarry(c > 0)
    cpu.Registers.SetZero(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetSub(false)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gnRRCA(cpu *Core) {
    c := cpu.Registers.A & 0x1

    cpu.Registers.A = (cpu.Registers.A >> 1) | (c << 7)

    cpu.Registers.SetCarry(c > 0)
    cpu.Registers.SetZero(false)
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetSub(false)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gnCPL(cpu *Core) {
    cpu.Registers.A = ^cpu.Registers.A

    cpu.Registers.SetHalfCarry(true)
    cpu.Registers.SetSub(true)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}
func gnCCF(cpu *Core) {
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(!cpu.Registers.GetCarry())
    cpu.Registers.SetSub(false)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

func gnSCF(cpu *Core) {
    cpu.Registers.SetHalfCarry(false)
    cpu.Registers.SetCarry(true)
    cpu.Registers.SetSub(false)

    cpu.Registers.LastClockM = 1
    cpu.Registers.LastClockT = 4
}

`
}
