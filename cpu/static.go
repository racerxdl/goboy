package cpu

var lsfr15table []uint8
var lsfr8table []uint8

func init() {
	lsfr15table = make([]uint8, 0x8000)
	LSFR := uint16(0x7FFF)
	LSFRShifted := uint16(0x3FFF)
	randomFactor := uint8(1)
	// Initialize LUT for LSFR15

	for i := 0; i < 0x8000; i++ {
		randomFactor = uint8(1 - (LSFR & 1))

		lsfr15table[i] = randomFactor

		LSFRShifted = LSFR >> 1
		LSFR = LSFRShifted | (((LSFRShifted ^ LSFR) & 0x1) << 14)
	}

	lsfr8table = make([]uint8, 0x80)

	LSFR = 0x7F //Seed value has all its bits set.

	for i := 0; i < 0x80; i++ {
		randomFactor = uint8(1 - (LSFR & 1))

		lsfr8table[i] = randomFactor

		LSFRShifted = LSFR >> 1
		LSFR = LSFRShifted | (((LSFRShifted ^ LSFR) & 0x1) << 6)
	}
}
