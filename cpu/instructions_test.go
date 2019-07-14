package cpu

import (
	"math/rand"
	"testing"
)

const RunCycles = 10

// region 0x00 Test NOP

func TestNOP(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x00) \"NOP\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x00](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x01 Test LDBCnn

func TestLDBCnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x01) \"LDBCnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16((0xA0 << 8) + rand.Intn(0xFFF))
		var var0 = uint8(rand.Intn(0xFF))
		var var1 = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, var0)
		cpu.Memory.WriteByte(cpu.Registers.PC+1, var1)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x01](cpu)
		RegAfter := cpu.Registers.Clone()

		if (var0) != (RegAfter.C) {
			t.Errorf("Expected var0 to be %v but got %v", var0, RegAfter.C)
		}
		if (var1) != (RegAfter.B) {
			t.Errorf("Expected var1 to be %v but got %v", var1, RegAfter.B)
		}
		if (RegBefore.PC + 2) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x02 Test LDBCmA

func TestLDBCmA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x02) \"LDBCmA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.B = 0xA0
		cpu.Registers.C = uint8(rand.Intn(0xFF))

		var hl = (uint16(cpu.Registers.B) << 8) + uint16(cpu.Registers.C)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x02](cpu)
		RegAfter := cpu.Registers.Clone()

		if (RegBefore.A) != (cpu.Memory.ReadByte(hl)) {
			t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, cpu.Memory.ReadByte(hl))
		}
		if (RegBefore.A) != (RegAfter.A) {
			t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, RegAfter.A)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x03 Test INCBC

func TestINCBC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x03) \"INCBC\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x03](cpu)
		RegAfter := cpu.Registers.Clone()

		var valA = RegBefore.C + 1

		valB := RegBefore.B + 1
		if valA != 0 {
			valB = RegBefore.B
		}

		if (valA) != (RegAfter.C) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.C)
		}
		if (valB) != (RegAfter.B) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.B)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x04 Test INCr_b

func TestINCrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x04) \"INCr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x04](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.B + 1
		var halfCarry = (RegBefore.B&0xF)+1 > 0xF

		if (val) != (RegAfter.B) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.B)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x05 Test DECr_b

func TestDECrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x05) \"DECr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x05](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.B - 1
		var halfCarry = (RegBefore.B & 0xF) == 0x00

		if (val) != (RegAfter.B) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.B)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x06 Test LDrn_b

func TestLDrnB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x06) \"LDrn_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		cpu.Registers.PC = cpu.Registers.HL() // Put PC in High Ram random value

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x06](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (val) != (RegAfter.B) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.B)
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x07 Test RLCA

func TestRLCA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x07) \"RLCA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x07](cpu)
		RegAfter := cpu.Registers.Clone()

		var c = (RegBefore.A >> 7) & 0x1
		var val = (RegBefore.A << 1) | c

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (c > 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected c > 0 to be %v but got %v", c > 0, RegAfter.GetCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() {
			t.Errorf("Expected Flag Zero to be zero")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x08 Test LDmmSP

func TestLDmmSP(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x08) \"LDmmSP\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFFF)))
		var addr = uint16(((0xA0 << 8) + rand.Intn(0xFFF)))

		cpu.Memory.WriteWord(cpu.Registers.PC, addr)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x08](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadWord(addr)) != (RegAfter.SP) {
			t.Errorf("Expected cpu.Memory.ReadWord(addr) to be %v but got %v", cpu.Memory.ReadWord(addr), RegAfter.SP)
		}
		if (RegBefore.PC + 2) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 20 {
			t.Errorf("Expected LastClockT to be %d but got %d", 20, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 5 {
			t.Errorf("Expected LastClockM to be %d but got %d", 5, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x09 Test ADDHLBC

func TestADDHLBC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x09) \"ADDHLBC\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x09](cpu)
		RegAfter := cpu.Registers.Clone()

		var ab = int(RegBefore.B)<<8 + int(RegBefore.C)
		var sum = int(RegBefore.HL()) + ab
		var halfCarry = (int(RegBefore.HL())&0xFFF)+(ab&0xFFF) > 0xFFF

		if (sum & 0xFFFF) != int(RegAfter.HL()) {
			t.Errorf("Expected sum & 0xFFFF to be %v but got %v", sum&0xFFFF, RegAfter.HL())
		}
		if (sum > 65535) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 65535 to be %v but got %v", sum > 65535, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x0A Test LDABCm

func TestLDABCm(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x0A) \"LDABCm\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.B = 0xA0
		cpu.Registers.C = uint8(rand.Intn(0xFF))

		var hl = uint16(cpu.Registers.B)<<8 + uint16(cpu.Registers.C)
		var val = uint8(rand.Intn(0xFF))
		cpu.Memory.WriteByte(hl, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x0A](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x0B Test DECBC

func TestDECBC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x0B) \"DECBC\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x0B](cpu)
		RegAfter := cpu.Registers.Clone()

		var valA = RegBefore.C - 1

		valB := RegBefore.B - 1
		if valA != 255 {
			valB = RegBefore.B
		}

		if (valA) != (RegAfter.C) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.C)
		}
		if (valB) != (RegAfter.B) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.B)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x0C Test INCr_c

func TestINCrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x0C) \"INCr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x0C](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.C + 1
		var halfCarry = (RegBefore.C&0xF)+1 > 0xF

		if (val) != (RegAfter.C) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.C)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x0D Test DECr_c

func TestDECrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x0D) \"DECr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x0D](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.C - 1
		var halfCarry = (RegBefore.C & 0xF) == 0x00

		if (val) != (RegAfter.C) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.C)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x0E Test LDrn_c

func TestLDrnC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x0E) \"LDrn_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		cpu.Registers.PC = cpu.Registers.HL() // Put PC in High Ram random value

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x0E](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (val) != (RegAfter.C) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.C)
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x0F Test RRCA

func TestRRCA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x0F) \"RRCA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x0F](cpu)
		RegAfter := cpu.Registers.Clone()

		var c = RegBefore.A & 1
		var val = (RegBefore.A >> 1) | (c << 7)

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (c > 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected c > 0 to be %v but got %v", c > 0, RegAfter.GetCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() {
			t.Errorf("Expected Flag Zero to be zero")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x10 Test STOP

func TestStop(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x10) \"STOP\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x10](cpu)
		RegAfter := cpu.Registers.Clone()

		if cpu.stopped != true {
			t.Errorf("Expected cpu to be stopped")
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x11 Test LDDEnn

func TestLDDEnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x11) \"LDDEnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16((0xA0 << 8) + rand.Intn(0xFFF))
		var var0 = uint8(rand.Intn(0xFF))
		var var1 = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, var0)
		cpu.Memory.WriteByte(cpu.Registers.PC+1, var1)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x11](cpu)
		RegAfter := cpu.Registers.Clone()

		if (var0) != (RegAfter.E) {
			t.Errorf("Expected var0 to be %v but got %v", var0, RegAfter.E)
		}
		if (var1) != (RegAfter.D) {
			t.Errorf("Expected var1 to be %v but got %v", var1, RegAfter.D)
		}
		if (RegBefore.PC + 2) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x12 Test LDDEmA

func TestLDDEmA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x12) \"LDDEmA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.D = 0xA0
		cpu.Registers.E = uint8(rand.Intn(0xFF))

		var hl = (uint16(cpu.Registers.D) << 8) + uint16(cpu.Registers.E)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x12](cpu)
		RegAfter := cpu.Registers.Clone()

		if (RegBefore.A) != (cpu.Memory.ReadByte(hl)) {
			t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, cpu.Memory.ReadByte(hl))
		}
		if (RegBefore.A) != (RegAfter.A) {
			t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, RegAfter.A)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x13 Test INCDE

func TestINCDE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x13) \"INCDE\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x13](cpu)
		RegAfter := cpu.Registers.Clone()

		var valA = RegBefore.E + 1

		valB := RegBefore.D + 1
		if valA != 0 {
			valB = RegBefore.D
		}

		if (valA) != (RegAfter.E) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.E)
		}
		if (valB) != (RegAfter.D) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.D)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x14 Test INCr_d

func TestINCrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x14) \"INCr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x14](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.D + 1
		var halfCarry = (RegBefore.D&0xF)+1 > 0xF

		if (val) != (RegAfter.D) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.D)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x15 Test DECr_d

func TestDECrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x15) \"DECr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x15](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.D - 1
		var halfCarry = (RegBefore.D & 0xF) == 0x00

		if (val) != (RegAfter.D) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.D)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x16 Test LDrn_d

func TestLDrnD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x16) \"LDrn_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		cpu.Registers.PC = cpu.Registers.HL() // Put PC in High Ram random value

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x16](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (val) != (RegAfter.D) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.D)
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x17 Test RLA

func TestRLA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x17) \"RLA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x17](cpu)
		RegAfter := cpu.Registers.Clone()

		c := (RegBefore.A >> 7) > 0
		f := uint8(0)
		if RegBefore.GetCarry() {
			f = 1
		}

		var val = (RegBefore.A << 1) | f

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (c) != (RegAfter.GetCarry()) {
			t.Errorf("Expected c to be %v but got %v", c, RegAfter.GetCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() {
			t.Errorf("Expected Flag Zero to be zero")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x18 Test JRn

func TestJRn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x18) \"JRn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(0xFF) - 128
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x18](cpu)
		RegAfter := cpu.Registers.Clone()

		if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
			t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", uint16(int(RegBefore.PC)+signedV+1)&0xFFFF, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(127)
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x18](cpu)
		RegAfter := cpu.Registers.Clone()

		if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
			t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", uint16(int(RegBefore.PC)+signedV+1)&0xFFFF, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x19 Test ADDHLDE

func TestADDHLDE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x19) \"ADDHLDE\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x19](cpu)
		RegAfter := cpu.Registers.Clone()

		var ab = int(RegBefore.D)<<8 + int(RegBefore.E)
		var sum = int(RegBefore.HL()) + ab
		var halfCarry = (int(RegBefore.HL())&0xFFF)+(ab&0xFFF) > 0xFFF

		if (sum & 0xFFFF) != int(RegAfter.HL()) {
			t.Errorf("Expected sum & 0xFFFF to be %v but got %v", sum&0xFFFF, RegAfter.HL())
		}
		if (sum > 65535) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 65535 to be %v but got %v", sum > 65535, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x1A Test LDADEm

func TestLDADEm(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x1A) \"LDADEm\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.D = 0xA0
		cpu.Registers.E = uint8(rand.Intn(0xFF))

		var hl = uint16(cpu.Registers.D)<<8 + uint16(cpu.Registers.E)
		var val = uint8(rand.Intn(0xFF))
		cpu.Memory.WriteByte(hl, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x1A](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x1B Test DECDE

func TestDECDE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x1B) \"DECDE\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x1B](cpu)
		RegAfter := cpu.Registers.Clone()

		var valA = RegBefore.E - 1

		valB := RegBefore.D - 1
		if valA != 255 {
			valB = RegBefore.D
		}

		if (valA) != (RegAfter.E) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.E)
		}
		if (valB) != (RegAfter.D) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.D)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x1C Test INCr_e

func TestINCrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x1C) \"INCr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x1C](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.E + 1
		var halfCarry = (RegBefore.E&0xF)+1 > 0xF

		if (val) != (RegAfter.E) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.E)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x1D Test DECr_e

func TestDECrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x1D) \"DECr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x1D](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.E - 1
		var halfCarry = (RegBefore.E & 0xF) == 0x00

		if (val) != (RegAfter.E) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.E)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x1E Test LDrn_e

func TestLDrnE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x1E) \"LDrn_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		cpu.Registers.PC = cpu.Registers.HL() // Put PC in High Ram random value

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x1E](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (val) != (RegAfter.E) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.E)
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x1F Test RRA

func TestRRA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x1F) \"RRA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x1F](cpu)
		RegAfter := cpu.Registers.Clone()

		var c = RegBefore.A & 1

		var f = uint8(0)
		if RegBefore.GetCarry() {
			f = 1
		}

		var val = (RegBefore.A >> 1) | (f << 7)

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (c > 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected c > 0 to be %v but got %v", c > 0, RegAfter.GetCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() {
			t.Errorf("Expected Flag Zero to be zero")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x20 Test JRNZn

func TestJRNZn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x20) \"JRNZn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(128) - 128
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x20](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetZero() {

			if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
				t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", uint16(int(RegBefore.PC)+signedV+1)&0xFFFF, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		} else {

			if (RegBefore.PC + 1) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (8) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 8)
			}
			if (RegAfter.LastClockM) != (8 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 8/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(127)
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x20](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetZero() {

			if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
				t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", uint16(int(RegBefore.PC)+signedV+1)&0xFFFF, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		} else {

			if (RegBefore.PC + 1) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (8) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 8)
			}
			if (RegAfter.LastClockM) != (8 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 8/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x21 Test LDHLnn

func TestLDHLnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x21) \"LDHLnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16((0xA0 << 8) + rand.Intn(0xFFF))
		var var0 = uint8(rand.Intn(0xFF))
		var var1 = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, var0)
		cpu.Memory.WriteByte(cpu.Registers.PC+1, var1)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x21](cpu)
		RegAfter := cpu.Registers.Clone()

		if (var0) != (RegAfter.L) {
			t.Errorf("Expected var0 to be %v but got %v", var0, RegAfter.L)
		}
		if (var1) != (RegAfter.H) {
			t.Errorf("Expected var1 to be %v but got %v", var1, RegAfter.H)
		}
		if (RegBefore.PC + 2) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x22 Test LDHLIA

func TestLDHLIA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x22) \"LDHLIA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x22](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(RegBefore.HL())) != (RegBefore.A) {
			t.Errorf("Expected cpu.Memory.ReadByte(RegBefore.HL()) to be %v but got %v", cpu.Memory.ReadByte(RegBefore.HL()), RegBefore.A)
		}
		if (RegBefore.HL() + 1) != (RegAfter.HL()) {
			t.Errorf("Expected RegBefore.HL() + 1 to be %v but got %v", RegBefore.HL()+1, RegAfter.HL())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x23 Test INCHL

func TestINCHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x23) \"INCHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x23](cpu)
		RegAfter := cpu.Registers.Clone()

		var valA = RegBefore.L + 1

		valB := RegBefore.H + 1
		if valA != 0 {
			valB = RegBefore.H
		}

		if (valA) != (RegAfter.L) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.L)
		}
		if (valB) != (RegAfter.H) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.H)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x24 Test INCr_h

func TestINCrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x24) \"INCr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x24](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.H + 1
		var halfCarry = (RegBefore.H&0xF)+1 > 0xF

		if (val) != (RegAfter.H) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.H)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x25 Test DECr_h

func TestDECrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x25) \"DECr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x25](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.H - 1
		var halfCarry = (RegBefore.H & 0xF) == 0x00

		if (val) != (RegAfter.H) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.H)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x26 Test LDrn_h

func TestLDrnH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x26) \"LDrn_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		cpu.Registers.PC = cpu.Registers.HL() // Put PC in High Ram random value

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x26](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (val) != (RegAfter.H) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.H)
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x27 Test DAA

func TestDAA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x27) \"DAA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x27](cpu)
		RegAfter := cpu.Registers.Clone()

		correction := int(0)
		a := int(RegBefore.A)

		if RegBefore.GetCarry() {
			correction = 0x60
		}

		if RegBefore.GetHalfCarry() || (!RegBefore.GetSub() && ((a & 0x0F) > 9)) {
			correction |= 0x06
		}

		if RegBefore.GetCarry() || (!RegBefore.GetSub() && (a > 0x99)) {
			correction |= 0x60
		}

		if RegBefore.GetSub() {
			a -= correction
		} else {
			a += correction
		}

		expectedA := uint8(a)
		expectedZero := expectedA == 0
		expectedCarry := RegBefore.GetCarry()

		if (correction<<2)&0x100 != 0 {
			expectedCarry = true
		}

		if (expectedCarry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", expectedCarry, RegAfter.GetCarry())
		}
		if (expectedZero) != (RegAfter.GetZero()) {
			t.Errorf("Expected zero to be %v but got %v", expectedZero, RegAfter.GetZero())
		}
		if expectedA != (RegAfter.A) {
			t.Errorf("Expected a & 0xFF to be %v but got %v", expectedA&0xFF, RegAfter.A)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x28 Test JRZn

func TestJRZn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x28) \"JRZn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(127) - 128
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x28](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetZero() {

			if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
				t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", uint16(int(RegBefore.PC)+signedV+1)&0xFFFF, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		} else {

			if (RegBefore.PC + 1) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (8) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 8)
			}
			if (RegAfter.LastClockM) != (8 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 8/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(127)
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x28](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetZero() {

			if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
				t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", uint16(int(RegBefore.PC)+signedV+1)&0xFFFF, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		} else {

			if (RegBefore.PC + 1) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (8) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 8)
			}
			if (RegAfter.LastClockM) != (8 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 8/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x29 Test ADDHLHL

func TestADDHLHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x29) \"ADDHLHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x29](cpu)
		RegAfter := cpu.Registers.Clone()

		var ab = int(RegBefore.H)<<8 + int(RegBefore.L)
		var sum = int(RegBefore.HL()) + ab
		var halfCarry = (int(RegBefore.HL())&0xFFF)+(ab&0xFFF) > 0xFFF

		if (sum & 0xFFFF) != int(RegAfter.HL()) {
			t.Errorf("Expected sum & 0xFFFF to be %v but got %v", sum&0xFFFF, RegAfter.HL())
		}
		if (sum > 65535) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 65535 to be %v but got %v", sum > 65535, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x2A Test LDAHLI

func TestLDAHLI(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x2A) \"LDAHLI\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x2A](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (RegBefore.HL() + 1) != (RegAfter.HL()) {
			t.Errorf("Expected RegBefore.HL + 1 to be %v but got %v", RegBefore.HL()+1, RegAfter.HL())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x2B Test DECHL

func TestDECHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x2B) \"DECHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x2B](cpu)
		RegAfter := cpu.Registers.Clone()

		var valA = RegBefore.L - 1

		valB := RegBefore.H - 1
		if valA != 255 {
			valB = RegBefore.H
		}

		if (valA) != (RegAfter.L) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.L)
		}
		if (valB) != (RegAfter.H) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.H)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x2C Test INCr_l

func TestINCrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x2C) \"INCr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x2C](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.L + 1
		var halfCarry = (RegBefore.L&0xF)+1 > 0xF

		if (val) != (RegAfter.L) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.L)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x2D Test DECr_l

func TestDECrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x2D) \"DECr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x2D](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.L - 1
		var halfCarry = (RegBefore.L & 0xF) == 0x00

		if (val) != (RegAfter.L) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.L)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x2E Test LDrn_l

func TestLDrnL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x2E) \"LDrn_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		cpu.Registers.PC = cpu.Registers.HL() // Put PC in High Ram random value

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x2E](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (val) != (RegAfter.L) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.L)
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x2F Test CPL

func TestCPL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x2F) \"CPL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x2F](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = ^RegBefore.A

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x30 Test JRNCn

func TestJRNCn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x30) \"JRNCn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(127) - 128
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x30](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetCarry() {

			if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
				t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF), RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		} else {

			if (RegBefore.PC + 1) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (8) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 8)
			}
			if (RegAfter.LastClockM) != (8 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 8/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(127)
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x30](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetCarry() {

			if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
				t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF), RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		} else {

			if (RegBefore.PC + 1) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (8) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 8)
			}
			if (RegAfter.LastClockM) != (8 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 8/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x31 Test LDSPnn

func TestLDSPnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x31) \"LDSPnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFFF)))
		var var0 = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, var0)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x31](cpu)
		RegAfter := cpu.Registers.Clone()

		if (var0) != (RegAfter.SP) {
			t.Errorf("Expected var0 to be %v but got %v", var0, RegAfter.SP)
		}
		if (RegBefore.PC + 2) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x32 Test LDHLDA

func TestLDHLDA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x32) \"LDHLDA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x32](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(RegBefore.HL())) != (RegBefore.A) {
			t.Errorf("Expected cpu.Memory.ReadByte(RegBefore.HL) to be %v but got %v", cpu.Memory.ReadByte(RegBefore.HL()), RegBefore.A)
		}
		if (RegBefore.HL() - 1) != (RegAfter.HL()) {
			t.Errorf("Expected RegBefore.HL - 1 to be %v but got %v", RegBefore.HL()-1, RegAfter.HL())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x33 Test INCSP

func TestINCSP(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x33) \"INCSP\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x33](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.SP + 1

		if (val) != (RegAfter.SP) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.SP)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x34 Test INCHLm

func TestINCHLm(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x34) \"INCHLm\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x34](cpu)
		RegAfter := cpu.Registers.Clone()

		var valAfter = cpu.Memory.ReadByte(RegBefore.HL())

		var newVal = uint8(val + 1)
		var halfCarry = (val&0xF)+1 > 0xF

		if (newVal) != (valAfter) {
			t.Errorf("Expected newVal to be %v but got %v", newVal, valAfter)
		}
		if (newVal == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected newVal == 0 to be %v but got %v", newVal == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x35 Test DECHLm

func TestDECHLm(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x35) \"DECHLm\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x35](cpu)
		RegAfter := cpu.Registers.Clone()

		var valAfter = cpu.Memory.ReadByte(RegBefore.HL())

		var newVal = uint8(val - 1)
		var halfCarry = (val & 0xF) == 0x00

		if (newVal) != (valAfter) {
			t.Errorf("Expected newVal to be %v but got %v", newVal, valAfter)
		}
		if (newVal == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected newVal == 0 to be %v but got %v", newVal == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x36 Test LDHLmn

func TestLDHLmn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x36) \"LDHLmn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		cpu.Registers.PC = cpu.Registers.HL() // Put PC in High Ram random value

		var val = uint8(rand.Intn(0x50))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		cpu.Registers.H = 0xFF
		cpu.Registers.L = uint8((0x80 + rand.Intn(0x50)))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x36](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(cpu.Registers.HL())) != (val) {
			t.Errorf("Expected cpu.Memory.ReadByte(cpu.Registers.HL()) to be %v but got %v", cpu.Memory.ReadByte(cpu.Registers.HL()), val)
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x37 Test SCF

func TestSCF(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x37) \"SCF\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x37](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if !RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be one")
		}
		// endregion

	}
}

// endregion
// region 0x38 Test JRCn

func TestJRCn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x38) \"JRCn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(127) - 128
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x38](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetCarry() {
			if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
				t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF), RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		} else {
			if (RegBefore.PC + 1) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (8) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 8)
			}
			if (RegAfter.LastClockM) != (8 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 8/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var signedV = rand.Intn(127)
		var v = uint8(signedV)

		cpu.Memory.WriteByte(cpu.Registers.PC, v)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x38](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetCarry() {

			if (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF) != (RegAfter.PC) {
				t.Errorf("Expected (RegBefore.PC + signedV + 1) & 0xFFFF to be %v but got %v", (uint16(int(RegBefore.PC)+signedV+1) & 0xFFFF), RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		} else {

			if (RegBefore.PC + 1) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (8) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 8)
			}
			if (RegAfter.LastClockM) != (8 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 8/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x39 Test ADDHLSP

func TestADDHLSP(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x39) \"ADDHLSP\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x39](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.HL()) + int(RegBefore.SP)
		var halfCarry = (RegBefore.HL()&0xFFF)+(RegBefore.SP&0xFFF) > 0xFFF

		if (sum & 0xFFFF) != int(RegAfter.HL()) {
			t.Errorf("Expected sum & 0xFFFF to be %v but got %v", sum&0xFFFF, RegAfter.HL())
		}
		if (sum > 65535) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 65535 to be %v but got %v", sum > 65535, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x3A Test LDAHLD

func TestLDAHLD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x3A) \"LDAHLD\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x3A](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (RegBefore.HL() - 1) != (RegAfter.HL()) {
			t.Errorf("Expected RegBefore.HL - 1 to be %v but got %v", RegBefore.HL()-1, RegAfter.HL())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x3B Test DECSP

func TestDECSP(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x3B) \"DECSP\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x3B](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = uint16(RegBefore.SP - 1)

		if (val) != (RegAfter.SP) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.SP)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x3C Test INCr_a

func TestINCrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x3C) \"INCr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x3C](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A + 1
		var halfCarry = (RegBefore.A&0xF)+1 > 0xF

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x3D Test DECr_a

func TestDECrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x3D) \"DECr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x3D](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A - 1
		var halfCarry = (RegBefore.A & 0xF) == 0x00

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x3E Test LDrn_a

func TestLDrnA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x3E) \"LDrn_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		cpu.Registers.PC = cpu.Registers.HL() // Put PC in High Ram random value

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x3E](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x3F Test CCF

func TestCCF(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x3F) \"CCF\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x3F](cpu)
		RegAfter := cpu.Registers.Clone()

		if (!RegBefore.GetCarry()) != (RegAfter.GetCarry()) {
			t.Errorf("Expected !RegBefore.GetCarry() to be %v but got %v", !RegBefore.GetCarry(), RegAfter.GetCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x40 Test LDrr_bb

func TestLDrrBB(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x40) \"LDrr_bb\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x40](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.B) != (RegBefore.B) {
			t.Errorf("Expected RegAfter.B to be %v but got %v", RegAfter.B, RegBefore.B)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x41 Test LDrr_bc

func TestLDrrBC(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x41) \"LDrr_bc\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x41](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.B) != (RegBefore.C) {
			t.Errorf("Expected RegAfter.B to be %v but got %v", RegAfter.B, RegBefore.C)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x42 Test LDrr_bd

func TestLDrrBD(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x42) \"LDrr_bd\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x42](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.B) != (RegBefore.D) {
			t.Errorf("Expected RegAfter.B to be %v but got %v", RegAfter.B, RegBefore.D)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x43 Test LDrr_be

func TestLDrrBE(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x43) \"LDrr_be\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x43](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.B) != (RegBefore.E) {
			t.Errorf("Expected RegAfter.B to be %v but got %v", RegAfter.B, RegBefore.E)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x44 Test LDrr_bh

func TestLDrrBH(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x44) \"LDrr_bh\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x44](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.B) != (RegBefore.H) {
			t.Errorf("Expected RegAfter.B to be %v but got %v", RegAfter.B, RegBefore.H)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x45 Test LDrr_bl

func TestLDrrBL(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x45) \"LDrr_bl\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x45](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.B) != (RegBefore.L) {
			t.Errorf("Expected RegAfter.B to be %v but got %v", RegAfter.B, RegBefore.L)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x46 Test LDrHLm_b

func TestLDrHLmB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x46) \"LDrHLm_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x46](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.B) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.B)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x47 Test LDrr_ba

func TestLDrrBA(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x47) \"LDrr_ba\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x47](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.B) != (RegBefore.A) {
			t.Errorf("Expected RegAfter.B to be %v but got %v", RegAfter.B, RegBefore.A)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x48 Test LDrr_cb

func TestLDrrCB(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x48) \"LDrr_cb\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x48](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.C) != (RegBefore.B) {
			t.Errorf("Expected RegAfter.C to be %v but got %v", RegAfter.C, RegBefore.B)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x49 Test LDrr_cc

func TestLDrrCC(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x49) \"LDrr_cc\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x49](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.C) != (RegBefore.C) {
			t.Errorf("Expected RegAfter.C to be %v but got %v", RegAfter.C, RegBefore.C)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x4A Test LDrr_cd

func TestLDrrCD(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x4A) \"LDrr_cd\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x4A](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.C) != (RegBefore.D) {
			t.Errorf("Expected RegAfter.C to be %v but got %v", RegAfter.C, RegBefore.D)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x4B Test LDrr_ce

func TestLDrrCE(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x4B) \"LDrr_ce\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x4B](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.C) != (RegBefore.E) {
			t.Errorf("Expected RegAfter.C to be %v but got %v", RegAfter.C, RegBefore.E)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x4C Test LDrr_ch

func TestLDrrCH(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x4C) \"LDrr_ch\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x4C](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.C) != (RegBefore.H) {
			t.Errorf("Expected RegAfter.C to be %v but got %v", RegAfter.C, RegBefore.H)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x4D Test LDrr_cl

func TestLDrrCL(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x4D) \"LDrr_cl\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x4D](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.C) != (RegBefore.L) {
			t.Errorf("Expected RegAfter.C to be %v but got %v", RegAfter.C, RegBefore.L)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x4E Test LDrHLm_c

func TestLDrHLmC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x4E) \"LDrHLm_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x4E](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.C) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.C)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x4F Test LDrr_ca

func TestLDrrCA(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x4F) \"LDrr_ca\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x4F](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.C) != (RegBefore.A) {
			t.Errorf("Expected RegAfter.C to be %v but got %v", RegAfter.C, RegBefore.A)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x50 Test LDrr_db

func TestLDrrDB(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x50) \"LDrr_db\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x50](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.D) != (RegBefore.B) {
			t.Errorf("Expected RegAfter.D to be %v but got %v", RegAfter.D, RegBefore.B)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x51 Test LDrr_dc

func TestLDrrDC(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x51) \"LDrr_dc\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x51](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.D) != (RegBefore.C) {
			t.Errorf("Expected RegAfter.D to be %v but got %v", RegAfter.D, RegBefore.C)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x52 Test LDrr_dd

func TestLDrrDD(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x52) \"LDrr_dd\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x52](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.D) != (RegBefore.D) {
			t.Errorf("Expected RegAfter.D to be %v but got %v", RegAfter.D, RegBefore.D)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x53 Test LDrr_de

func TestLDrrDE(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x53) \"LDrr_de\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x53](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.D) != (RegBefore.E) {
			t.Errorf("Expected RegAfter.D to be %v but got %v", RegAfter.D, RegBefore.E)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x54 Test LDrr_dh

func TestLDrrDH(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x54) \"LDrr_dh\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x54](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.D) != (RegBefore.H) {
			t.Errorf("Expected RegAfter.D to be %v but got %v", RegAfter.D, RegBefore.H)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x55 Test LDrr_dl

func TestLDrrDL(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x55) \"LDrr_dl\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x55](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.D) != (RegBefore.L) {
			t.Errorf("Expected RegAfter.D to be %v but got %v", RegAfter.D, RegBefore.L)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x56 Test LDrHLm_d

func TestLDrHLmD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x56) \"LDrHLm_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x56](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.D) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.D)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x57 Test LDrr_da

func TestLDrrDA(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x57) \"LDrr_da\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x57](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.D) != (RegBefore.A) {
			t.Errorf("Expected RegAfter.D to be %v but got %v", RegAfter.D, RegBefore.A)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x58 Test LDrr_eb

func TestLDrrEB(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x58) \"LDrr_eb\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x58](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.E) != (RegBefore.B) {
			t.Errorf("Expected RegAfter.E to be %v but got %v", RegAfter.E, RegBefore.B)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x59 Test LDrr_ec

func TestLDrrEC(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x59) \"LDrr_ec\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x59](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.E) != (RegBefore.C) {
			t.Errorf("Expected RegAfter.E to be %v but got %v", RegAfter.E, RegBefore.C)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x5A Test LDrr_ed

func TestLDrrED(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x5A) \"LDrr_ed\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x5A](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.E) != (RegBefore.D) {
			t.Errorf("Expected RegAfter.E to be %v but got %v", RegAfter.E, RegBefore.D)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x5B Test LDrr_ee

func TestLDrrEE(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x5B) \"LDrr_ee\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x5B](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.E) != (RegBefore.E) {
			t.Errorf("Expected RegAfter.E to be %v but got %v", RegAfter.E, RegBefore.E)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x5C Test LDrr_eh

func TestLDrrEH(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x5C) \"LDrr_eh\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x5C](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.E) != (RegBefore.H) {
			t.Errorf("Expected RegAfter.E to be %v but got %v", RegAfter.E, RegBefore.H)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x5D Test LDrr_el

func TestLDrrEL(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x5D) \"LDrr_el\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x5D](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.E) != (RegBefore.L) {
			t.Errorf("Expected RegAfter.E to be %v but got %v", RegAfter.E, RegBefore.L)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x5E Test LDrHLm_e

func TestLDrHLmE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x5E) \"LDrHLm_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x5E](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.E) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.E)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x5F Test LDrr_ea

func TestLDrrEA(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x5F) \"LDrr_ea\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x5F](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.E) != (RegBefore.A) {
			t.Errorf("Expected RegAfter.E to be %v but got %v", RegAfter.E, RegBefore.A)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x60 Test LDrr_hb

func TestLDrrHB(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x60) \"LDrr_hb\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x60](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.H) != (RegBefore.B) {
			t.Errorf("Expected RegAfter.H to be %v but got %v", RegAfter.H, RegBefore.B)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x61 Test LDrr_hc

func TestLDrrHC(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x61) \"LDrr_hc\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x61](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.H) != (RegBefore.C) {
			t.Errorf("Expected RegAfter.H to be %v but got %v", RegAfter.H, RegBefore.C)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x62 Test LDrr_hd

func TestLDrrHD(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x62) \"LDrr_hd\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x62](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.H) != (RegBefore.D) {
			t.Errorf("Expected RegAfter.H to be %v but got %v", RegAfter.H, RegBefore.D)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x63 Test LDrr_he

func TestLDrrHE(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x63) \"LDrr_he\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x63](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.H) != (RegBefore.E) {
			t.Errorf("Expected RegAfter.H to be %v but got %v", RegAfter.H, RegBefore.E)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x64 Test LDrr_hh

func TestLDrrHH(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x64) \"LDrr_hh\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x64](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.H) != (RegBefore.H) {
			t.Errorf("Expected RegAfter.H to be %v but got %v", RegAfter.H, RegBefore.H)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x65 Test LDrr_hl

func TestLDrrHL(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x65) \"LDrr_hl\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x65](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.H) != (RegBefore.L) {
			t.Errorf("Expected RegAfter.H to be %v but got %v", RegAfter.H, RegBefore.L)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x66 Test LDrHLm_h

func TestLDrHLmH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x66) \"LDrHLm_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x66](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.H) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.H)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x67 Test LDrr_ha

func TestLDrrHA(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x67) \"LDrr_ha\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x67](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.H) != (RegBefore.A) {
			t.Errorf("Expected RegAfter.H to be %v but got %v", RegAfter.H, RegBefore.A)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x68 Test LDrr_lb

func TestLDrrLB(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x68) \"LDrr_lb\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x68](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.L) != (RegBefore.B) {
			t.Errorf("Expected RegAfter.L to be %v but got %v", RegAfter.L, RegBefore.B)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x69 Test LDrr_lc

func TestLDrrLC(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x69) \"LDrr_lc\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x69](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.L) != (RegBefore.C) {
			t.Errorf("Expected RegAfter.L to be %v but got %v", RegAfter.L, RegBefore.C)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x6A Test LDrr_ld

func TestLDrrLD(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x6A) \"LDrr_ld\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x6A](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.L) != (RegBefore.D) {
			t.Errorf("Expected RegAfter.L to be %v but got %v", RegAfter.L, RegBefore.D)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x6B Test LDrr_le

func TestLDrrLE(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x6B) \"LDrr_le\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x6B](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.L) != (RegBefore.E) {
			t.Errorf("Expected RegAfter.L to be %v but got %v", RegAfter.L, RegBefore.E)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x6C Test LDrr_lh

func TestLDrrLH(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x6C) \"LDrr_lh\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x6C](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.L) != (RegBefore.H) {
			t.Errorf("Expected RegAfter.L to be %v but got %v", RegAfter.L, RegBefore.H)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x6D Test LDrr_ll

func TestLDrrLL(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x6D) \"LDrr_ll\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x6D](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.L) != (RegBefore.L) {
			t.Errorf("Expected RegAfter.L to be %v but got %v", RegAfter.L, RegBefore.L)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x6E Test LDrHLm_l

func TestLDrHLmL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x6E) \"LDrHLm_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x6E](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.L) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.L)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x6F Test LDrr_la

func TestLDrrLA(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x6F) \"LDrr_la\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x6F](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.L) != (RegBefore.A) {
			t.Errorf("Expected RegAfter.L to be %v but got %v", RegAfter.L, RegBefore.A)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x70 Test LDHLmr_b

func TestLDHLmrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x70) \"LDHLmr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x70](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(cpu.Registers.HL())) != (RegAfter.B) {
			t.Errorf("Expected cpu.Memory.ReadByte(cpu.Registers.HL()) to be %v but got %v", cpu.Memory.ReadByte(cpu.Registers.HL()), RegAfter.B)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x71 Test LDHLmr_c

func TestLDHLmrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x71) \"LDHLmr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x71](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(cpu.Registers.HL())) != (RegAfter.C) {
			t.Errorf("Expected cpu.Memory.ReadByte(cpu.Registers.HL()) to be %v but got %v", cpu.Memory.ReadByte(cpu.Registers.HL()), RegAfter.C)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x72 Test LDHLmr_d

func TestLDHLmrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x72) \"LDHLmr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x72](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(cpu.Registers.HL())) != (RegAfter.D) {
			t.Errorf("Expected cpu.Memory.ReadByte(cpu.Registers.HL()) to be %v but got %v", cpu.Memory.ReadByte(cpu.Registers.HL()), RegAfter.D)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x73 Test LDHLmr_e

func TestLDHLmrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x73) \"LDHLmr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x73](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(cpu.Registers.HL())) != (RegAfter.E) {
			t.Errorf("Expected cpu.Memory.ReadByte(cpu.Registers.HL()) to be %v but got %v", cpu.Memory.ReadByte(cpu.Registers.HL()), RegAfter.E)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x74 Test LDHLmr_h

func TestLDHLmrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x74) \"LDHLmr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x74](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(cpu.Registers.HL())) != (RegAfter.H) {
			t.Errorf("Expected cpu.Memory.ReadByte(cpu.Registers.HL()) to be %v but got %v", cpu.Memory.ReadByte(cpu.Registers.HL()), RegAfter.H)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x75 Test LDHLmr_l

func TestLDHLmrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x75) \"LDHLmr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x75](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(cpu.Registers.HL())) != (RegAfter.L) {
			t.Errorf("Expected cpu.Memory.ReadByte(cpu.Registers.HL()) to be %v but got %v", cpu.Memory.ReadByte(cpu.Registers.HL()), RegAfter.L)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x76 Test HALT

func TestHALT(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x76) \"HALT\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x76](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.InterruptEnable && cpu.halted {
			t.Errorf("Expected cpu not to be halted when interrupts are disabled")
		}

		if !cpu.halted && RegBefore.InterruptEnable {
			t.Errorf("Expected cpu to be halted")
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x77 Test LDHLmr_a

func TestLDHLmrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x77) \"LDHLmr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x77](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(cpu.Registers.HL())) != (RegAfter.A) {
			t.Errorf("Expected cpu.Memory.ReadByte(cpu.Registers.HL()) to be %v but got %v", cpu.Memory.ReadByte(cpu.Registers.HL()), RegAfter.A)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x78 Test LDrr_ab

func TestLDrrAB(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x78) \"LDrr_ab\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x78](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.A) != (RegBefore.B) {
			t.Errorf("Expected RegAfter.A to be %v but got %v", RegAfter.A, RegBefore.B)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x79 Test LDrr_ac

func TestLDrrAC(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x79) \"LDrr_ac\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x79](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.A) != (RegBefore.C) {
			t.Errorf("Expected RegAfter.A to be %v but got %v", RegAfter.A, RegBefore.C)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x7A Test LDrr_ad

func TestLDrrAD(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x7A) \"LDrr_ad\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x7A](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.A) != (RegBefore.D) {
			t.Errorf("Expected RegAfter.A to be %v but got %v", RegAfter.A, RegBefore.D)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x7B Test LDrr_ae

func TestLDrrAE(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x7B) \"LDrr_ae\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x7B](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.A) != (RegBefore.E) {
			t.Errorf("Expected RegAfter.A to be %v but got %v", RegAfter.A, RegBefore.E)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x7C Test LDrr_ah

func TestLDrrAH(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x7C) \"LDrr_ah\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x7C](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.A) != (RegBefore.H) {
			t.Errorf("Expected RegAfter.A to be %v but got %v", RegAfter.A, RegBefore.H)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x7D Test LDrr_al

func TestLDrrAL(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x7D) \"LDrr_al\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x7D](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.A) != (RegBefore.L) {
			t.Errorf("Expected RegAfter.A to be %v but got %v", RegAfter.A, RegBefore.L)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x7E Test LDrHLm_a

func TestLDrHLmA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x7E) \"LDrHLm_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x7E](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x7F Test LDrr_aa

func TestLDrrAA(t *testing.T) {
	cpu := MakeCore()
	// Console.WriteLine("Testing (0x7F) \"LDrr_aa\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x7F](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Difference

		if (RegAfter.A) != (RegBefore.A) {
			t.Errorf("Expected RegAfter.A to be %v but got %v", RegAfter.A, RegBefore.A)
		} // endregion

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0x80 Test ADDr_b

func TestADDrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x80) \"ADDr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x80](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = uint16(RegBefore.A) + uint16(RegBefore.B)
		var halfCarry = (uint8(sum) & 0xF) < (RegBefore.A & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "B" != "A" {
			if RegBefore.B != RegAfter.B {
				t.Errorf("Expected RegBefore.B != RegAfter.B to be %v but got %v", false, uint8(sum) != RegAfter.A)
			}
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x81 Test ADDr_c

func TestADDrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x81) \"ADDr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x81](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = uint16(RegBefore.A) + uint16(RegBefore.C)
		var halfCarry = (uint8(sum) & 0xF) < (RegBefore.A & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "C" != "A" {
			if RegBefore.C != RegAfter.C {
				t.Errorf("Expected RegBefore.C != RegAfter.C to be %v but got %v", false, uint8(sum) != RegAfter.A)
			}
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x82 Test ADDr_d

func TestADDrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x82) \"ADDr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x82](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = uint16(RegBefore.A) + uint16(RegBefore.D)
		var halfCarry = (uint8(sum) & 0xF) < (RegBefore.A & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "D" != "A" {
			if RegBefore.D != RegAfter.D {
				t.Errorf("Expected RegBefore.D != RegAfter.D to be %v but got %v", false, uint8(sum) != RegAfter.A)
			}
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x83 Test ADDr_e

func TestADDrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x83) \"ADDr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x83](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = uint16(RegBefore.A) + uint16(RegBefore.E)
		var halfCarry = (uint8(sum) & 0xF) < (RegBefore.A & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "E" != "A" {
			if RegBefore.E != RegAfter.E {
				t.Errorf("Expected RegBefore.E != RegAfter.E to be %v but got %v", false, uint8(sum) != RegAfter.A)
			}
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x84 Test ADDr_h

func TestADDrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x84) \"ADDr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x84](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = uint16(RegBefore.A) + uint16(RegBefore.H)
		var halfCarry = (uint8(sum) & 0xF) < (RegBefore.A & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "H" != "A" {
			if RegBefore.H != RegAfter.H {
				t.Errorf("Expected RegBefore.H != RegAfter.H to be %v but got %v", false, uint8(sum) != RegAfter.A)
			}
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x85 Test ADDr_l

func TestADDrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x85) \"ADDr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x85](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = uint16(RegBefore.A) + uint16(RegBefore.L)
		var halfCarry = (uint8(sum) & 0xF) < (RegBefore.A & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "L" != "A" {
			if RegBefore.L != RegAfter.L {
				t.Errorf("Expected RegBefore.L != RegAfter.L to be %v but got %v", false, uint8(sum) != RegAfter.A)
			}
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x86 Test ADDHL

func TestADDHLm(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x86) \"ADDHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x86](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = uint16(RegBefore.A) + uint16(val)
		var halfCarry = (RegBefore.A&0xF)+(val&0xF) > 0xF

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x87 Test ADDr_a

func TestADDrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x87) \"ADDr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x87](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = uint16(RegBefore.A) + uint16(RegBefore.A)
		var halfCarry = (uint8(sum) & 0xF) < (RegBefore.A & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "A" != "A" {
			if RegBefore.A != RegAfter.A {
				t.Errorf("Expected RegBefore.A != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
			}
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x88 Test ADCr_b

func TestADCrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x88) \"ADCr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x88](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = uint16(RegBefore.A) + uint16(RegBefore.B) + uint16(f)
		var halfCarry = (RegBefore.A&0xF)+(RegBefore.B&0xF)+f > 0xF

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "B" != "A" {

			if (RegBefore.B) != (RegAfter.B) {
				t.Errorf("Expected RegBefore.B to be %v but got %v", RegBefore.B, RegAfter.B)
			}
		}

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x89 Test ADCr_c

func TestADCrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x89) \"ADCr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x89](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = uint16(RegBefore.A) + uint16(RegBefore.C) + uint16(f)
		var halfCarry = (RegBefore.A&0xF)+(RegBefore.C&0xF)+f > 0xF

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "C" != "A" {

			if (RegBefore.C) != (RegAfter.C) {
				t.Errorf("Expected RegBefore.C to be %v but got %v", RegBefore.C, RegAfter.C)
			}
		}

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x8A Test ADCr_d

func TestADCrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x8A) \"ADCr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x8A](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = uint16(RegBefore.A) + uint16(RegBefore.D) + uint16(f)
		var halfCarry = (RegBefore.A&0xF)+(RegBefore.D&0xF)+f > 0xF

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "D" != "A" {

			if (RegBefore.D) != (RegAfter.D) {
				t.Errorf("Expected RegBefore.D to be %v but got %v", RegBefore.D, RegAfter.D)
			}
		}

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x8B Test ADCr_e

func TestADCrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x8B) \"ADCr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x8B](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = uint16(RegBefore.A) + uint16(RegBefore.E) + uint16(f)
		var halfCarry = (RegBefore.A&0xF)+(RegBefore.E&0xF)+f > 0xF

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "E" != "A" {

			if (RegBefore.E) != (RegAfter.E) {
				t.Errorf("Expected RegBefore.E to be %v but got %v", RegBefore.E, RegAfter.E)
			}
		}

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x8C Test ADCr_h

func TestADCrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x8C) \"ADCr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x8C](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = uint16(RegBefore.A) + uint16(RegBefore.H) + uint16(f)
		var halfCarry = (RegBefore.A&0xF)+(RegBefore.H&0xF)+f > 0xF

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "H" != "A" {

			if (RegBefore.H) != (RegAfter.H) {
				t.Errorf("Expected RegBefore.H to be %v but got %v", RegBefore.H, RegAfter.H)
			}
		}

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x8D Test ADCr_l

func TestADCrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x8D) \"ADCr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x8D](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = uint16(RegBefore.A) + uint16(RegBefore.L) + uint16(f)
		var halfCarry = (RegBefore.A&0xF)+(RegBefore.L&0xF)+f > 0xF

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "L" != "A" {

			if (RegBefore.L) != (RegAfter.L) {
				t.Errorf("Expected RegBefore.L to be %v but got %v", RegBefore.L, RegAfter.L)
			}
		}

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x8E Test ADCHL

func TestADCHLm(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x8E) \"ADCHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x8E](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = uint16(RegBefore.A) + uint16(val) + uint16(f)
		var halfCarry = (RegBefore.A&0xF)+(val&0xF)+f > 0xF

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x8F Test ADCr_a

func TestADCrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x8F) \"ADCr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x8F](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = uint16(RegBefore.A) + uint16(RegBefore.A) + uint16(f)
		var halfCarry = (RegBefore.A&0xF)+(RegBefore.A&0xF)+f > 0xF

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "A" != "A" {

			if (RegBefore.A) != (RegAfter.A) {
				t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, RegAfter.A)
			}
		}

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}
		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0x90 Test SUBr_b

func TestSUBrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x90) \"SUBr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x90](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.A) - int(RegBefore.B)
		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.B & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "B" != "A" {
			if (RegBefore.B) != (RegAfter.B) {
				t.Errorf("Expected RegBefore.B to be %v but got %v", RegBefore.B, RegAfter.B)
			}
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x91 Test SUBr_c

func TestSUBrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x91) \"SUBr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x91](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.A) - int(RegBefore.C)
		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.C & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "C" != "A" {
			if (RegBefore.C) != (RegAfter.C) {
				t.Errorf("Expected RegBefore.C to be %v but got %v", RegBefore.C, RegAfter.C)
			}
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x92 Test SUBr_d

func TestSUBrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x92) \"SUBr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x92](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.A) - int(RegBefore.D)
		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.D & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "D" != "A" {
			if (RegBefore.D) != (RegAfter.D) {
				t.Errorf("Expected RegBefore.D to be %v but got %v", RegBefore.D, RegAfter.D)
			}
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x93 Test SUBr_e

func TestSUBrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x93) \"SUBr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x93](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.A) - int(RegBefore.E)
		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.E & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "E" != "A" {
			if (RegBefore.E) != (RegAfter.E) {
				t.Errorf("Expected RegBefore.E to be %v but got %v", RegBefore.E, RegAfter.E)
			}
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x94 Test SUBr_h

func TestSUBrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x94) \"SUBr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x94](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.A) - int(RegBefore.H)
		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.H & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "H" != "A" {
			if (RegBefore.H) != (RegAfter.H) {
				t.Errorf("Expected RegBefore.H to be %v but got %v", RegBefore.H, RegAfter.H)
			}
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x95 Test SUBr_l

func TestSUBrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x95) \"SUBr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x95](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.A) - int(RegBefore.L)
		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.L & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "L" != "A" {
			if (RegBefore.L) != (RegAfter.L) {
				t.Errorf("Expected RegBefore.L to be %v but got %v", RegBefore.L, RegAfter.L)
			}
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x96 Test SUBHL

func TestSUBHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x96) \"SUBHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x96](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.A) - int(val)
		var halfCarry = (RegBefore.A & 0xF) < (val & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x97 Test SUBr_a

func TestSUBrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x97) \"SUBr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x97](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.A) - int(RegBefore.A)
		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.A & 0xF)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "A" != "A" {
			if (RegBefore.A) != (RegAfter.A) {
				t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, RegAfter.A)
			}
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x98 Test SBCr_b

func TestSBCrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x98) \"SBCr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x98](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = int(RegBefore.A) - int(RegBefore.B) - int(f)
		var halfCarry = (RegBefore.A & 0xF) < ((RegBefore.B & 0xF) + f)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "B" != "A" {
			if (RegBefore.B) != (RegAfter.B) {
				t.Errorf("Expected RegBefore.B to be %v but got %v", RegBefore.B, RegAfter.B)
			}
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x99 Test SBCr_c

func TestSBCrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x99) \"SBCr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x99](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = int(RegBefore.A) - int(RegBefore.C) - int(f)
		var halfCarry = (RegBefore.A & 0xF) < ((RegBefore.C & 0xF) + f)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "C" != "A" {
			if (RegBefore.C) != (RegAfter.C) {
				t.Errorf("Expected RegBefore.C to be %v but got %v", RegBefore.C, RegAfter.C)
			}
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x9A Test SBCr_d

func TestSBCrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x9A) \"SBCr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x9A](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = int(RegBefore.A) - int(RegBefore.D) - int(f)
		var halfCarry = (RegBefore.A & 0xF) < ((RegBefore.D & 0xF) + f)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "D" != "A" {
			if (RegBefore.D) != (RegAfter.D) {
				t.Errorf("Expected RegBefore.D to be %v but got %v", RegBefore.D, RegAfter.D)
			}
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x9B Test SBCr_e

func TestSBCrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x9B) \"SBCr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x9B](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = int(RegBefore.A) - int(RegBefore.E) - int(f)
		var halfCarry = (RegBefore.A & 0xF) < ((RegBefore.E & 0xF) + f)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "E" != "A" {
			if (RegBefore.E) != (RegAfter.E) {
				t.Errorf("Expected RegBefore.E to be %v but got %v", RegBefore.E, RegAfter.E)
			}
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x9C Test SBCr_h

func TestSBCrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x9C) \"SBCr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x9C](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = int(RegBefore.A) - int(RegBefore.H) - int(f)
		var halfCarry = (RegBefore.A & 0xF) < ((RegBefore.H & 0xF) + f)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "H" != "A" {
			if (RegBefore.H) != (RegAfter.H) {
				t.Errorf("Expected RegBefore.H to be %v but got %v", RegBefore.H, RegAfter.H)
			}
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x9D Test SBCr_l

func TestSBCrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x9D) \"SBCr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x9D](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = int(RegBefore.A) - int(RegBefore.L) - int(f)
		var halfCarry = (RegBefore.A & 0xF) < ((RegBefore.L & 0xF) + f)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "L" != "A" {
			if (RegBefore.L) != (RegAfter.L) {
				t.Errorf("Expected RegBefore.L to be %v but got %v", RegBefore.L, RegAfter.L)
			}
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x9E Test SBCHL

func TestSBCHLm(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x9E) \"SBCHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x9E](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = int(RegBefore.A) - int(val) - int(f)
		var halfCarry = (RegBefore.A & 0xF) < ((val & 0xF) + f)

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0x9F Test SBCr_a

func TestSBCrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0x9F) \"SBCr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0x9F](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = int(RegBefore.A) - int(RegBefore.A) - int(f)
		var halfCarry = (RegBefore.A & 0xF) < ((RegBefore.A & 0xF) + f)

		if uint8(sum) != RegAfter.A {
			t.Errorf("Expected uint8(sum) != RegAfter.A to be %v but got %v", false, uint8(sum) != RegAfter.A)
		}

		if "A" != "A" {
			if (RegBefore.A) != (RegAfter.A) {
				t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, RegAfter.A)
			}
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xA0 Test ANDr_b

func TestANDrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA0) \"ANDr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA0](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A & RegBefore.B

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xA1 Test ANDr_c

func TestANDrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA1) \"ANDr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA1](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A & RegBefore.C

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xA2 Test ANDr_d

func TestANDrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA2) \"ANDr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA2](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A & RegBefore.D

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xA3 Test ANDr_e

func TestANDrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA3) \"ANDr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA3](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A & RegBefore.E

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xA4 Test ANDr_h

func TestANDrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA4) \"ANDr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA4](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A & RegBefore.H

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xA5 Test ANDr_l

func TestANDrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA5) \"ANDr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA5](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A & RegBefore.L

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xA6 Test ANDHL

func TestANDHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA6) \"ANDHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA6](cpu)
		RegAfter := cpu.Registers.Clone()

		val = RegBefore.A & val

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xA7 Test ANDr_a

func TestANDrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA7) \"ANDr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA7](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A & RegBefore.A

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xA8 Test XORr_b

func TestXORrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA8) \"XORr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA8](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A ^ RegBefore.B

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xA9 Test XORr_c

func TestXORrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xA9) \"XORr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xA9](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A ^ RegBefore.C

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xAA Test XORr_d

func TestXORrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xAA) \"XORr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xAA](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A ^ RegBefore.D

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xAB Test XORr_e

func TestXORrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xAB) \"XORr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xAB](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A ^ RegBefore.E

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xAC Test XORr_h

func TestXORrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xAC) \"XORr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xAC](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A ^ RegBefore.H

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xAD Test XORr_l

func TestXORrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xAD) \"XORr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xAD](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A ^ RegBefore.L

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xAE Test XORHL

func TestXORHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xAE) \"XORHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xAE](cpu)
		RegAfter := cpu.Registers.Clone()

		val = RegBefore.A ^ val

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xAF Test XORr_a

func TestXORrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xAF) \"XORr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xAF](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A ^ RegBefore.A

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xB0 Test ORr_b

func TestORrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB0) \"ORr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB0](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A | RegBefore.B

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xB1 Test ORr_c

func TestORrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB1) \"ORr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB1](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A | RegBefore.C

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xB2 Test ORr_d

func TestORrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB2) \"ORr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB2](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A | RegBefore.D

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xB3 Test ORr_e

func TestORrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB3) \"ORr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB3](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A | RegBefore.E

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xB4 Test ORr_h

func TestORrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB4) \"ORr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB4](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A | RegBefore.H

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xB5 Test ORr_l

func TestORrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB5) \"ORr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB5](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A | RegBefore.L

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xB6 Test ORHL

func TestORHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB6) \"ORHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB6](cpu)
		RegAfter := cpu.Registers.Clone()

		val = RegBefore.A | val

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xB7 Test ORr_a

func TestORrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB7) \"ORr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB7](cpu)
		RegAfter := cpu.Registers.Clone()

		var val = RegBefore.A | RegBefore.A

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xB8 Test CPr_b

func TestCPrB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB8) \"CPr_b\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB8](cpu)
		RegAfter := cpu.Registers.Clone()

		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.B & 0xF)
		var carry = RegBefore.A < RegBefore.B

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}
		if (RegBefore.A == RegBefore.B) != (RegAfter.GetZero()) {
			t.Errorf("Expected RegBefore.A == RegBefore.B to be %v but got %v", RegBefore.A == RegBefore.B, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xB9 Test CPr_c

func TestCPrC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xB9) \"CPr_c\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xB9](cpu)
		RegAfter := cpu.Registers.Clone()

		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.C & 0xF)
		var carry = RegBefore.A < RegBefore.C

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}
		if (RegBefore.A == RegBefore.C) != (RegAfter.GetZero()) {
			t.Errorf("Expected RegBefore.A == RegBefore.C to be %v but got %v", RegBefore.A == RegBefore.C, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xBA Test CPr_d

func TestCPrD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xBA) \"CPr_d\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xBA](cpu)
		RegAfter := cpu.Registers.Clone()

		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.D & 0xF)
		var carry = RegBefore.A < RegBefore.D

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}
		if (RegBefore.A == RegBefore.D) != (RegAfter.GetZero()) {
			t.Errorf("Expected RegBefore.A == RegBefore.D to be %v but got %v", RegBefore.A == RegBefore.D, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xBB Test CPr_e

func TestCPrE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xBB) \"CPr_e\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xBB](cpu)
		RegAfter := cpu.Registers.Clone()

		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.E & 0xF)
		var carry = RegBefore.A < RegBefore.E

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}
		if (RegBefore.A == RegBefore.E) != (RegAfter.GetZero()) {
			t.Errorf("Expected RegBefore.A == RegBefore.E to be %v but got %v", RegBefore.A == RegBefore.E, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xBC Test CPr_h

func TestCPrH(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xBC) \"CPr_h\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xBC](cpu)
		RegAfter := cpu.Registers.Clone()

		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.H & 0xF)
		var carry = RegBefore.A < RegBefore.H

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}
		if (RegBefore.A == RegBefore.H) != (RegAfter.GetZero()) {
			t.Errorf("Expected RegBefore.A == RegBefore.H to be %v but got %v", RegBefore.A == RegBefore.H, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xBD Test CPr_l

func TestCPrL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xBD) \"CPr_l\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xBD](cpu)
		RegAfter := cpu.Registers.Clone()

		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.L & 0xF)
		var carry = RegBefore.A < RegBefore.L

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}
		if (RegBefore.A == RegBefore.L) != (RegAfter.GetZero()) {
			t.Errorf("Expected RegBefore.A == RegBefore.L to be %v but got %v", RegBefore.A == RegBefore.L, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xBE Test CPHL

func TestCPHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xBE) \"CPHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.H = 0xA0
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xBE](cpu)
		RegAfter := cpu.Registers.Clone()

		var halfCarry = (RegBefore.A & 0xF) < (val & 0xF)
		var carry = RegBefore.A < val

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}
		if (RegBefore.A == val) != (RegAfter.GetZero()) {
			t.Errorf("Expected RegBefore.A == val to be %v but got %v", RegBefore.A == val, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xBF Test CPr_a

func TestCPrA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xBF) \"CPr_a\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xBF](cpu)
		RegAfter := cpu.Registers.Clone()

		var halfCarry = (RegBefore.A & 0xF) < (RegBefore.A & 0xF)
		var carry = RegBefore.A < RegBefore.A

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}
		if (RegBefore.A == RegBefore.A) != (RegAfter.GetZero()) {
			t.Errorf("Expected RegBefore.A == RegBefore.A to be %v but got %v", RegBefore.A == RegBefore.A, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xC0 Test RETNZ

func TestRETNZ(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xC0) \"RETNZ\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.SP, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xC0](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetZero() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (20) != (RegAfter.LastClockT) {
				t.Errorf("Expected 20 to be %v but got %v", 20, RegAfter.LastClockT)
			}
			if (20 / 4) != (RegAfter.LastClockM) {
				t.Errorf("Expected 20 / 4 to be %v but got %v", 20/4, RegAfter.LastClockM)
			}
			if (RegBefore.SP + 2) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
			}
		} else {

			if (8) != (RegAfter.LastClockT) {
				t.Errorf("Expected 8 to be %v but got %v", 8, RegAfter.LastClockT)
			}
			if (8 / 4) != (RegAfter.LastClockM) {
				t.Errorf("Expected 8 / 4 to be %v but got %v", 8/4, RegAfter.LastClockM)
			}
			if (RegBefore.SP) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP to be %v but got %v", RegBefore.SP, RegAfter.SP)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xC1 Test POPBC

func TestPOPBC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xC1) \"POPBC\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint8(rand.Intn(0xFF))
		var valB = uint8(rand.Intn(0xFF))

		cpu.Registers.SP--
		cpu.Memory.WriteByte(cpu.Registers.SP, valA)
		cpu.Registers.SP--
		cpu.Memory.WriteByte(cpu.Registers.SP, valB)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xC1](cpu)
		RegAfter := cpu.Registers.Clone()

		if "C" == "F" {
			valB &= 0xF0
		}

		if (RegBefore.SP + 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
		}
		if (valA) != (RegAfter.B) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.B)
		}
		if (valB) != (RegAfter.C) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.C)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xC2 Test JPNZnn

func TestJPNZnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xC2) \"JPNZnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xC2](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetZero() {
			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (16) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 16)
			}
			if (RegAfter.LastClockM) != (16 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 16/4)
			}
		} else {
			if (RegBefore.PC + 2) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xC3 Test JPnn

func TestJPnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xC3) \"JPnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xC3](cpu)
		RegAfter := cpu.Registers.Clone()

		if (valA) != (RegAfter.PC) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xC4 Test CALLNZnn

func TestCALLNZnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xC4) \"CALLNZnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xC4](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetZero() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (24) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 24)
			}
			if (RegAfter.LastClockM) != (24 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 24/4)
			}
			if (RegBefore.SP - 2) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP - 2 to be %v but got %v", RegBefore.SP-2, RegAfter.SP)
			}
		} else {

			if (RegBefore.PC + 2) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
			if (RegBefore.SP) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP to be %v but got %v", RegBefore.SP, RegAfter.SP)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xC5 Test PUSHBC

func TestPUSHBC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xC5) \"PUSHBC\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xC5](cpu)
		RegAfter := cpu.Registers.Clone()

		var valB = cpu.Memory.ReadByte(RegAfter.SP)
		var valA = cpu.Memory.ReadByte(RegAfter.SP + 1)

		if (RegBefore.SP - 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP - 2 to be %v but got %v", RegBefore.SP-2, RegAfter.SP)
		}
		if (RegBefore.B) != (valA) {
			t.Errorf("Expected RegBefore.B to be %v but got %v", RegBefore.B, valA)
		}
		if (RegBefore.C) != (valB) {
			t.Errorf("Expected RegBefore.C to be %v but got %v", RegBefore.C, valB)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xC6 Test ADDn

func TestADDn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xC6) \"ADDn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFF)))
		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xC6](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = uint16(RegBefore.A) + uint16(val)
		var halfCarry = (RegBefore.A&0xF)+(val&0xF) > 0xF

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xC8 Test RETZ

func TestRETZ(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xC8) \"RETZ\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.SP, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xC8](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetZero() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (20) != (RegAfter.LastClockT) {
				t.Errorf("Expected 20 to be %v but got %v", 20, RegAfter.LastClockT)
			}
			if (20 / 4) != (RegAfter.LastClockM) {
				t.Errorf("Expected 20 / 4 to be %v but got %v", 20/4, RegAfter.LastClockM)
			}
			if (RegBefore.SP + 2) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
			}
		} else {

			if (8) != (RegAfter.LastClockT) {
				t.Errorf("Expected 8 to be %v but got %v", 8, RegAfter.LastClockT)
			}
			if (8 / 4) != (RegAfter.LastClockM) {
				t.Errorf("Expected 8 / 4 to be %v but got %v", 8/4, RegAfter.LastClockM)
			}
			if (RegBefore.SP) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP to be %v but got %v", RegBefore.SP, RegAfter.SP)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xC9 Test RET

func TestRET(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xC9) \"RET\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.SP, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xC9](cpu)
		RegAfter := cpu.Registers.Clone()

		if (valA) != (RegAfter.PC) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
		}
		if (RegBefore.SP + 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xCA Test JPZnn

func TestJPZnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xCA) \"JPZnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xCA](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetZero() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (16) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 16)
			}
			if (RegAfter.LastClockM) != (16 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 16/4)
			}
		} else {

			if (RegBefore.PC + 2) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xCC Test CALLZnn

func TestCALLZnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xCC) \"CALLZnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xCC](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetZero() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (24) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 24)
			}
			if (RegAfter.LastClockM) != (24 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 24/4)
			}
			if (RegBefore.SP - 2) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP - 2 to be %v but got %v", RegBefore.SP-2, RegAfter.SP)
			}
		} else {

			if (RegBefore.PC + 2) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
			if (RegBefore.SP) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP to be %v but got %v", RegBefore.SP, RegAfter.SP)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xCD Test CALLnn

func TestCALLnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xCD) \"CALLnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xCD](cpu)
		RegAfter := cpu.Registers.Clone()

		if (valA) != (RegAfter.PC) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
		}
		if (RegBefore.SP - 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP - 2 to be %v but got %v", RegBefore.SP-2, RegAfter.SP)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 24 {
			t.Errorf("Expected LastClockT to be %d but got %d", 24, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 6 {
			t.Errorf("Expected LastClockM to be %d but got %d", 6, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xCE Test ADCn

func TestADCn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xCE) \"ADCn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFF)))
		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xCE](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = uint16(RegBefore.A) + uint16(val) + uint16(f)
		var halfCarry = (RegBefore.A&0xF)+(val&0xF)+f > 0xF

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}

		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if (sum > 255) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum > 255 to be %v but got %v", sum > 255, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xD0 Test RETNC

func TestRETNC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xD0) \"RETNC\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.SP, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xD0](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetCarry() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (20) != (RegAfter.LastClockT) {
				t.Errorf("Expected 20 to be %v but got %v", 20, RegAfter.LastClockT)
			}
			if (20 / 4) != (RegAfter.LastClockM) {
				t.Errorf("Expected 20 / 4 to be %v but got %v", 20/4, RegAfter.LastClockM)
			}
			if (RegBefore.SP + 2) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
			}
		} else {

			if (8) != (RegAfter.LastClockT) {
				t.Errorf("Expected 8 to be %v but got %v", 8, RegAfter.LastClockT)
			}
			if (8 / 4) != (RegAfter.LastClockM) {
				t.Errorf("Expected 8 / 4 to be %v but got %v", 8/4, RegAfter.LastClockM)
			}
			if (RegBefore.SP) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP to be %v but got %v", RegBefore.SP, RegAfter.SP)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xD1 Test POPDE

func TestPOPDE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xD1) \"POPDE\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint8(rand.Intn(0xFF))
		var valB = uint8(rand.Intn(0xFF))

		cpu.Registers.SP--
		cpu.Memory.WriteByte(cpu.Registers.SP, valA)
		cpu.Registers.SP--
		cpu.Memory.WriteByte(cpu.Registers.SP, valB)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xD1](cpu)
		RegAfter := cpu.Registers.Clone()

		if "E" == "F" {
			valB &= 0xF0
		}

		if (RegBefore.SP + 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
		}
		if (valA) != (RegAfter.D) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.D)
		}
		if (valB) != (RegAfter.E) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.E)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xD2 Test JPNCnn

func TestJPNCnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xD2) \"JPNCnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xD2](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetCarry() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (16) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 16)
			}
			if (RegAfter.LastClockM) != (16 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 16/4)
			}
		} else {

			if (RegBefore.PC + 2) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xD3 Test XX

func TestNOPWARN_D3(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xD3) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xD3](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xD4 Test CALLNCnn

func TestCALLNCnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xD4) \"CALLNCnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xD4](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegBefore.GetCarry() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (24) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 24)
			}
			if (RegAfter.LastClockM) != (24 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 24/4)
			}
			if (RegBefore.SP - 2) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP - 2 to be %v but got %v", RegBefore.SP-2, RegAfter.SP)
			}
		} else {

			if (RegBefore.PC + 2) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
			if (RegBefore.SP) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP to be %v but got %v", RegBefore.SP, RegAfter.SP)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xD5 Test PUSHDE

func TestPUSHDE(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xD5) \"PUSHDE\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xD5](cpu)
		RegAfter := cpu.Registers.Clone()

		var valB = cpu.Memory.ReadByte(RegAfter.SP)
		var valA = cpu.Memory.ReadByte(RegAfter.SP + 1)

		if (RegBefore.SP - 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP - 2 to be %v but got %v", RegBefore.SP-2, RegAfter.SP)
		}
		if (RegBefore.D) != (valA) {
			t.Errorf("Expected RegBefore.D to be %v but got %v", RegBefore.D, valA)
		}
		if (RegBefore.E) != (valB) {
			t.Errorf("Expected RegBefore.E to be %v but got %v", RegBefore.E, valB)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xD6 Test SUBn

func TestSUBn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xD6) \"SUBn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFF)))
		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xD6](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.A) - int(val)
		var halfCarry = (RegBefore.A & 0xF) < (val & 0xF)

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}
		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xD8 Test RETC

func TestRETC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xD8) \"RETC\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.SP, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xD8](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetCarry() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (20) != (RegAfter.LastClockT) {
				t.Errorf("Expected 20 to be %v but got %v", 20, RegAfter.LastClockT)
			}
			if (20 / 4) != (RegAfter.LastClockM) {
				t.Errorf("Expected 20 / 4 to be %v but got %v", 20/4, RegAfter.LastClockM)
			}
			if (RegBefore.SP + 2) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
			}
		} else {

			if (8) != (RegAfter.LastClockT) {
				t.Errorf("Expected 8 to be %v but got %v", 8, RegAfter.LastClockT)
			}
			if (8 / 4) != (RegAfter.LastClockM) {
				t.Errorf("Expected 8 / 4 to be %v but got %v", 8/4, RegAfter.LastClockM)
			}
			if (RegBefore.SP) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP to be %v but got %v", RegBefore.SP, RegAfter.SP)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xD9 Test RETI

func TestRETI(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xD9) \"RETI\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.SP, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xD9](cpu)
		RegAfter := cpu.Registers.Clone()

		if (valA) != (RegAfter.PC) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
		}
		if (RegBefore.SP + 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
		}

		if !RegAfter.InterruptEnable {
			t.Errorf("Expected RegAfter.InterruptEnable to be true got false")
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xDA Test JPCnn

func TestJPCnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xDA) \"JPCnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xDA](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetCarry() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (16) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 16)
			}
			if (RegAfter.LastClockM) != (16 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 16/4)
			}
		} else {

			if (RegBefore.PC + 2) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xDB Test XX

func TestNOPWARN_DB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xDB) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xDB](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xDC Test CALLCnn

func TestCALLCnn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xDC) \"CALLCnn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint16(rand.Intn(0xFFFF))

		cpu.Memory.WriteWord(cpu.Registers.PC, valA)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xDC](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegBefore.GetCarry() {

			if (valA) != (RegAfter.PC) {
				t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (24) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 24)
			}
			if (RegAfter.LastClockM) != (24 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 24/4)
			}
			if (RegBefore.SP - 2) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP - 2 to be %v but got %v", RegBefore.SP-2, RegAfter.SP)
			}
		} else {

			if (RegBefore.PC + 2) != (RegAfter.PC) {
				t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
			}
			if (RegAfter.LastClockT) != (12) {
				t.Errorf("Expected RegAfter.LastClockT to be %v but got %v", RegAfter.LastClockT, 12)
			}
			if (RegAfter.LastClockM) != (12 / 4) {
				t.Errorf("Expected RegAfter.LastClockM to be %v but got %v", RegAfter.LastClockM, 12/4)
			}
			if (RegBefore.SP) != (RegAfter.SP) {
				t.Errorf("Expected RegBefore.SP to be %v but got %v", RegBefore.SP, RegAfter.SP)
			}
		}

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xDD Test XX

func TestNOPWARN_DD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xDD) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xDD](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xDE Test SBCn

func TestSBCn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xDE) \"SBCn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFF)))
		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xDE](cpu)
		RegAfter := cpu.Registers.Clone()

		f := uint8(0)

		if RegBefore.GetCarry() {
			f = 1
		}

		var sum = int(RegBefore.A) - int(val) - int(f)
		var halfCarry = (RegBefore.A & 0xF) < (val&0xF)+f

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}
		if uint8(sum&0xFF) != (RegAfter.A) {
			t.Errorf("Expected sum & 0xFF to be %v but got %v", sum&0xFF, RegAfter.A)
		}

		if (sum < 0) != (RegAfter.GetCarry()) {
			t.Errorf("Expected sum < 0 to be %v but got %v", sum < 0, RegAfter.GetCarry())
		}

		if ((sum & 0xFF) == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected (sum & 0xFF) == 0 to be %v but got %v", (sum&0xFF) == 0, RegAfter.GetZero())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
// region 0xE0 Test LDIOnA

func TestLDIOnA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xE0) \"LDIOnA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + uint16(rand.Intn(0xFF))))

		cpu.Memory.WriteByte(cpu.Registers.PC, 0x80)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE0](cpu)
		RegAfter := cpu.Registers.Clone()

		if (RegBefore.A) != (cpu.Memory.ReadByte(0xFF80)) {
			t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, cpu.Memory.ReadByte(0xFF80))
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xE1 Test POPHL

func TestPOPHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xE1) \"POPHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint8(rand.Intn(0xFF))
		var valB = uint8(rand.Intn(0xFF))

		cpu.Registers.SP--
		cpu.Memory.WriteByte(cpu.Registers.SP, valA)
		cpu.Registers.SP--
		cpu.Memory.WriteByte(cpu.Registers.SP, valB)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE1](cpu)
		RegAfter := cpu.Registers.Clone()

		if "L" == "F" {
			valB &= 0xF0
		}

		if (RegBefore.SP + 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
		}
		if (valA) != (RegAfter.H) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.H)
		}
		if (valB) != (RegAfter.L) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.L)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xE2 Test LDIOCA

func TestLDIOCA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xE2) \"LDIOCA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.C = 0x80

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE2](cpu)
		RegAfter := cpu.Registers.Clone()

		if (RegBefore.A) != (cpu.Memory.ReadByte(0xFF80)) {
			t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, cpu.Memory.ReadByte(0xFF80))
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xE3 Test XX

func TestNOPWARN_E3(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xE3) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE3](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xE4 Test XX

func TestNOPWARN_E4(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xE4) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE4](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xE5 Test PUSHHL

func TestPUSHHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xE5) \"PUSHHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE5](cpu)
		RegAfter := cpu.Registers.Clone()

		var valB = cpu.Memory.ReadByte(RegAfter.SP)
		var valA = cpu.Memory.ReadByte(RegAfter.SP + 1)

		if (RegBefore.SP - 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP - 2 to be %v but got %v", RegBefore.SP-2, RegAfter.SP)
		}
		if (RegBefore.H) != (valA) {
			t.Errorf("Expected RegBefore.H to be %v but got %v", RegBefore.H, valA)
		}
		if (RegBefore.L) != (valB) {
			t.Errorf("Expected RegBefore.L to be %v but got %v", RegBefore.L, valB)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xE6 Test ANDn

func TestANDn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xE6) \"ANDn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFF)))
		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE6](cpu)
		RegAfter := cpu.Registers.Clone()

		val = RegBefore.A & val

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}
		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if !RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be one")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xE8 Test ADDSPn

func TestADDSPn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xE8) \"ADDSPn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16((0xA0 << 8) + rand.Intn(0xFF))
		var signedV = rand.Intn(127) - 128

		cpu.Memory.WriteByte(cpu.Registers.PC, uint8(signedV))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE8](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.SP) + signedV
		var halfCarry = int(RegBefore.SP&0xF)+(signedV&0xF) > 0xF
		var carry = int(RegBefore.SP&0xFF)+(signedV&0xFF) > 0xFF

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}

		if uint16(sum&0xFFFF) != (RegAfter.SP) {
			t.Errorf("Expected sum & 0xFFFF to be %v but got %v", sum&0xFFFF, RegAfter.SP)
		}

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() {
			t.Errorf("Expected Flag Zero to be zero")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16((0xA0 << 8) + rand.Intn(0xFF))
		var signedV = rand.Intn(127)

		cpu.Memory.WriteByte(cpu.Registers.PC, uint8(signedV))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE8](cpu)
		RegAfter := cpu.Registers.Clone()

		var sum = int(RegBefore.SP) + signedV
		var halfCarry = int(RegBefore.SP&0xF)+(signedV&0xF) > 0xF
		var carry = int(RegBefore.SP&0xFF)+(signedV&0xFF) > 0xFF

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}

		if uint16(sum&0xFFFF) != (RegAfter.SP) {
			t.Errorf("Expected sum & 0xFFFF to be %v but got %v", sum&0xFFFF, RegAfter.SP)
		}

		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}

		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() {
			t.Errorf("Expected Flag Zero to be zero")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xE9 Test JPHL

func TestJPHL(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xE9) \"JPHL\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.H = uint8(rand.Intn(0xFF))
		cpu.Registers.L = uint8(rand.Intn(0xFF))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xE9](cpu)
		RegAfter := cpu.Registers.Clone()

		if (RegBefore.HL()) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.HL to be %v but got %v", RegBefore.HL(), RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xEA Test LDmmA

func TestLDmmA(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xEA) \"LDmmA\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to High Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFFF)))

		var addr = uint16(((0xA0 << 8) + rand.Intn(0xFFF)))

		cpu.Memory.WriteWord(cpu.Registers.PC, addr)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xEA](cpu)
		RegAfter := cpu.Registers.Clone()

		if (cpu.Memory.ReadByte(addr)) != (RegBefore.A) {
			t.Errorf("Expected cpu.Memory.ReadByte(addr) to be %v but got %v", cpu.Memory.ReadByte(addr), RegBefore.A)
		}
		if (RegBefore.PC + 2) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xEB Test XX

func TestNOPWARN_EB(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xEB) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xEB](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xEC Test XX

func TestNOPWARN_EC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xEC) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xEC](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xED Test XX

func TestNOPWARN_ED(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xED) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xED](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xEE Test XORn

func TestXORn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xEE) \"XORn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFF)))
		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xEE](cpu)
		RegAfter := cpu.Registers.Clone()

		val = RegBefore.A ^ val

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}
		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xF0 Test LDAIOn

func TestLDAIOn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xF0) \"LDAIOn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + uint16(rand.Intn(0xFF))))

		var val = uint8(rand.Intn(0x10))

		cpu.Memory.WriteByte(cpu.Registers.PC, 0x80)
		cpu.Memory.WriteByte(0xFF80, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF0](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xF1 Test POPAF

func TestPOPAF(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xF1) \"POPAF\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		var valA = uint8(rand.Intn(0xFF))
		var valB = uint8(rand.Intn(0xFF))

		cpu.Registers.SP--
		cpu.Memory.WriteByte(cpu.Registers.SP, valA)
		cpu.Registers.SP--
		cpu.Memory.WriteByte(cpu.Registers.SP, valB)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF1](cpu)
		RegAfter := cpu.Registers.Clone()

		if "F" == "F" {
			valB &= 0xF0
		}

		if (RegBefore.SP + 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP + 2 to be %v but got %v", RegBefore.SP+2, RegAfter.SP)
		}
		if (valA) != (RegAfter.A) {
			t.Errorf("Expected valA to be %v but got %v", valA, RegAfter.A)
		}
		if (valB) != (RegAfter.F) {
			t.Errorf("Expected valB to be %v but got %v", valB, RegAfter.F)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		// endregion

	}
}

// endregion
// region 0xF2 Test LDAIOC

func TestLDAIOC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xF2) \"LDAIOC\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.C = 0x80
		var val = uint8(rand.Intn(0x10))

		cpu.Memory.WriteByte(0xFF80, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF2](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xF3 Test DI

func TestDI(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xF3) \"DI\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF3](cpu)
		RegAfter := cpu.Registers.Clone()

		if RegAfter.InterruptEnable {
			t.Errorf("Expected Interrupt Enable to be disabled")
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xF4 Test XX

func TestNOPWARN_F4(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xF4) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF4](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xF5 Test PUSHAF

func TestPUSHAF(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xF5) \"PUSHAF\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.SP = uint16(((0xA1 << 8) + rand.Intn(0xF0)))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF5](cpu)
		RegAfter := cpu.Registers.Clone()

		var valB = cpu.Memory.ReadByte(RegAfter.SP)
		var valA = cpu.Memory.ReadByte(RegAfter.SP + 1)

		if (RegBefore.SP - 2) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.SP - 2 to be %v but got %v", RegBefore.SP-2, RegAfter.SP)
		}
		if (RegBefore.A) != (valA) {
			t.Errorf("Expected RegBefore.A to be %v but got %v", RegBefore.A, valA)
		}
		if (RegBefore.F) != (valB) {
			t.Errorf("Expected RegBefore.F to be %v but got %v", RegBefore.F, valB)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xF6 Test ORn

func TestORn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xF6) \"ORn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFF)))
		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF6](cpu)
		RegAfter := cpu.Registers.Clone()

		val = RegBefore.A | val

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}
		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (val == 0) != (RegAfter.GetZero()) {
			t.Errorf("Expected val == 0 to be %v but got %v", val == 0, RegAfter.GetZero())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		if RegAfter.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to be zero")
		}
		if RegAfter.GetCarry() {
			t.Errorf("Expected Flag Carry to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xF8 Test LDHLSPn

func TestLDHLSPn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xF8) \"LDHLSPn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFFF)))

		var signedV = rand.Intn(127) - 128

		cpu.Memory.WriteByte(cpu.Registers.PC, uint8(signedV))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF8](cpu)
		RegAfter := cpu.Registers.Clone()

		if uint16(int(RegBefore.SP)+signedV) != (RegAfter.HL()) {
			t.Errorf("Expected RegBefore.SP + signedV to be %v but got %v", int(RegBefore.SP)+signedV, RegAfter.HL())
		}

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}

		if ((int(RegBefore.SP)&0xF)+(signedV&0xF) > 0xF) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected (RegBefore.SP & 0xF) + (signedV & 0xF) > 0xF to be %v but got %v", (int(RegBefore.SP)&0xF)+(signedV&0xF) > 0xF, RegAfter.GetHalfCarry())
		}

		if ((int(RegBefore.SP)&0xFF)+(signedV&0xFF) > 0xFF) != (RegAfter.GetCarry()) {
			t.Errorf("Expected (RegBefore.SP & 0xFF) + (signedV & 0xFF) > 0xFF to be %v but got %v", (int(RegBefore.SP)&0xFF)+(signedV&0xFF) > 0xFF, RegAfter.GetCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() {
			t.Errorf("Expected Flag Zero to be zero")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFFF)))

		var signedV = rand.Intn(127)

		cpu.Memory.WriteByte(cpu.Registers.PC, uint8(signedV))

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF8](cpu)
		RegAfter := cpu.Registers.Clone()

		if uint16(int(RegBefore.SP)+signedV) != (RegAfter.HL()) {
			t.Errorf("Expected RegBefore.SP + signedV to be %v but got %v", int(RegBefore.SP)+signedV, RegAfter.HL())
		}

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}

		if ((int(RegBefore.SP)&0xF)+(signedV&0xF) > 0xF) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected (RegBefore.SP & 0xF) + (signedV & 0xF) > 0xF to be %v but got %v", (int(RegBefore.SP)&0xF)+(signedV&0xF) > 0xF, RegAfter.GetHalfCarry())
		}

		if ((int(RegBefore.SP)&0xFF)+(signedV&0xFF) > 0xFF) != (RegAfter.GetCarry()) {
			t.Errorf("Expected (RegBefore.SP & 0xFF) + (signedV & 0xFF) > 0xFF to be %v but got %v", (int(RegBefore.SP)&0xFF)+(signedV&0xFF) > 0xFF, RegAfter.GetCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 12 {
			t.Errorf("Expected LastClockT to be %d but got %d", 12, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 3 {
			t.Errorf("Expected LastClockM to be %d but got %d", 3, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() {
			t.Errorf("Expected Flag Zero to be zero")
		}
		if RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be zero")
		}
		// endregion

	}
}

// endregion
// region 0xF9 Test LDSPHLr

func TestLDHLSPr(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xF9) \"LDSPHLr\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xF9](cpu)
		RegAfter := cpu.Registers.Clone()

		if (RegBefore.HL()) != (RegAfter.SP) {
			t.Errorf("Expected RegBefore.HL to be %v but got %v", RegBefore.HL(), RegAfter.SP)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xFA Test LDAmm

func TestLDAmm(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xFA) \"LDAmm\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		// Force write to Catridge Ram Random Address (avoid writting to non writeable addresses)
		cpu.Registers.PC = uint16(((0xA0 << 8) + rand.Intn(0xFFF)))

		var addr = uint16(((0xA0 << 8) + rand.Intn(0xFFF)))
		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(addr, val)
		cpu.Memory.WriteWord(cpu.Registers.PC, addr)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xFA](cpu)
		RegAfter := cpu.Registers.Clone()

		if (val) != (RegAfter.A) {
			t.Errorf("Expected val to be %v but got %v", val, RegAfter.A)
		}
		if (RegBefore.PC + 2) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 2 to be %v but got %v", RegBefore.PC+2, RegAfter.PC)
		}

		// region Test Cycles
		if RegAfter.LastClockT != 16 {
			t.Errorf("Expected LastClockT to be %d but got %d", 16, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 4 {
			t.Errorf("Expected LastClockM to be %d but got %d", 4, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xFB Test EI

func TestEI(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xFB) \"EI\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xFB](cpu)
		RegAfter := cpu.Registers.Clone()

		if !RegAfter.InterruptEnable {
			t.Errorf("Expected Interrupt Enable to be enabled")
		}

		// region Test Cycles
		if RegAfter.LastClockT != 4 {
			t.Errorf("Expected LastClockT to be %d but got %d", 4, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 1 {
			t.Errorf("Expected LastClockM to be %d but got %d", 1, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xFC Test XX

func TestNOPWARN_FC(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xFC) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xFC](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xFD Test XX

func TestNOPWARN_FD(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xFD) \"XX\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xFD](cpu)
		RegAfter := cpu.Registers.Clone()

		// region Test Cycles
		if RegAfter.LastClockT != 0 {
			t.Errorf("Expected LastClockT to be %d but got %d", 0, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 0 {
			t.Errorf("Expected LastClockM to be %d but got %d", 0, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if RegAfter.GetZero() != RegBefore.GetZero() {
			t.Errorf("Expected Flag Zero to not change")
		}
		if RegAfter.GetSub() != RegBefore.GetSub() {
			t.Errorf("Expected Flag Sub to not change")
		}
		if RegAfter.GetHalfCarry() != RegBefore.GetHalfCarry() {
			t.Errorf("Expected Flag Half Carry to not change")
		}
		if RegAfter.GetCarry() != RegBefore.GetCarry() {
			t.Errorf("Expected Flag Carry to not change")
		}
		// endregion

	}
}

// endregion
// region 0xFE Test CPn

func TestCPn(t *testing.T) {
	cpu := MakeCore()

	// Console.WriteLine("Testing (0xFE) \"CPn\"")
	for i := 0; i < RunCycles; i++ {
		cpu.Reset()
		cpu.Registers.Randomize()
		cpu.Memory.Randomize()

		cpu.Registers.PC = uint16((0xA0 << 8) + rand.Intn(0xFF))
		var val = uint8(rand.Intn(0xFF))

		cpu.Memory.WriteByte(cpu.Registers.PC, val)

		cpu.Memory.WriteByte(cpu.Registers.HL(), val)

		RegBefore := cpu.Registers.Clone()
		GBInstructions[0xFE](cpu)
		RegAfter := cpu.Registers.Clone()

		var halfCarry = (RegBefore.A & 0xF) < (val & 0xF)
		var carry = RegBefore.A < val

		if (RegBefore.PC + 1) != (RegAfter.PC) {
			t.Errorf("Expected RegBefore.PC + 1 to be %v but got %v", RegBefore.PC+1, RegAfter.PC)
		}
		if (carry) != (RegAfter.GetCarry()) {
			t.Errorf("Expected carry to be %v but got %v", carry, RegAfter.GetCarry())
		}
		if (RegBefore.A == val) != (RegAfter.GetZero()) {
			t.Errorf("Expected RegBefore.A == val to be %v but got %v", RegBefore.A == val, RegAfter.GetZero())
		}
		if (halfCarry) != (RegAfter.GetHalfCarry()) {
			t.Errorf("Expected halfCarry to be %v but got %v", halfCarry, RegAfter.GetHalfCarry())
		}

		// region Test Cycles
		if RegAfter.LastClockT != 8 {
			t.Errorf("Expected LastClockT to be %d but got %d", 8, RegAfter.LastClockT)
		}
		if RegAfter.LastClockM != 2 {
			t.Errorf("Expected LastClockM to be %d but got %d", 2, RegAfter.LastClockM)
		}
		// endregion

		// region Test Flags
		if !RegAfter.GetSub() {
			t.Errorf("Expected Flag Sub to be one")
		}
		// endregion

	}
}

// endregion
