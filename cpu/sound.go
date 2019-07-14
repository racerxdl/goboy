package cpu

const (
	NR10 uint16 = 0xff10
	NR11        = 0xff11
	NR12        = 0xff12
	NR13        = 0xff13
	NR14        = 0xff14
	NR21        = 0xff16
	NR22        = 0xff17
	NR23        = 0xff18
	NR24        = 0xff19
	NR30        = 0xff1a
	NR31        = 0xff1b
	NR32        = 0xff1c
	NR33        = 0xff1d
	NR34        = 0xff1e
	NR41        = 0xff20
	NR42        = 0xff21
	NR43        = 0xff22
	NR44        = 0xff23
	NR50        = 0xff24
	NR51        = 0xff25
	NR52        = 0xff26
)

var duty1 = []uint8{0, 1, 1, 1, 1, 1, 1, 1}
var duty2 = []uint8{0, 0, 1, 1, 1, 1, 1, 1}
var duty3 = []uint8{0, 0, 0, 0, 1, 1, 1, 1}
var duty4 = []uint8{0, 0, 0, 0, 0, 0, 1, 1}

var dutyCycles = [][]uint8{
	duty1, // 00: 12.5% ( _-------_-------_------- )
	duty2, // 01: 25%   ( __------__------__------ )
	duty3, // 10: 50%   ( ____----____----____---- ) (normal)
	duty4, // 11: 75%   ( ______--______--______-- )
}

type SoundCard struct {
	cpu  *Core
	regs []byte

	sweepTime            uint8
	sweepTimeCalc        float32
	sweepSub             bool
	sweepShift           uint8
	wavePatternDuty1     uint8
	soundLength1         uint8
	soundLength1calc     float32
	initialVolume1       uint8
	envelopeIncrease1    bool
	numberEnvelopeSweep1 uint8
	frequency1           uint16
	frequency1calc       float32
	stopExpire1          bool

	wavePatternDuty2     uint8
	soundLength2         uint8
	soundLength2calc     float32
	initialVolume2       uint8
	envelopeIncrease2    bool
	numberEnvelopeSweep2 uint8
	frequency2           uint16
	frequency2calc       float32
	stopExpire2          bool
	wavTable             []uint8
}

func MakeSoundCard(cpu *Core) *SoundCard {
	return &SoundCard{
		cpu:  cpu,
		regs: make([]byte, 0x3F),
	}
}

func (s *SoundCard) Write(addr uint16, val uint8) {
	if addr < 0xFF40 && addr >= 0xFF00 {
		s.regs[addr-0xFF00] = val
	}

	switch addr {
	case NR10: // Channel 1 Sweep register (R/W)
		s.sweepTime = (val & 0x70) >> 4
		s.sweepSub = (val & 8) > 0
		s.sweepShift = val & 7

		/*
		   Bit 6-4 - Sweep Time
		   Bit 3   - Sweep Increase/Decrease
		            0: Addition    (frequency increases)
		            1: Subtraction (frequency decreases)
		   Bit 2-0 - Number of sweep shift (n: 0-7)
		   Sweep Time:
		     000: sweep off - no freq change
		     001: 7.8 ms  (1/128Hz)
		     010: 15.6 ms (2/128Hz)
		     011: 23.4 ms (3/128Hz)
		     100: 31.3 ms (4/128Hz)
		     101: 39.1 ms (5/128Hz)
		     110: 46.9 ms (6/128Hz)
		     111: 54.7 ms (7/128Hz)

		   The change of frequency (NR13,NR14) at each shift is calculated by the following formula where X(0) is initial freq & X(t-1) is last freq:
		     X(t) = X(t-1) +/- X(t-1)/2^n

		*/

	case NR11: // Channel 1 Sound length/Wave pattern duty (R/W)
		s.wavePatternDuty1 = (val & 0xC0) >> 6
		s.soundLength1 = 64 - (val & 0x3F)
		/*

		     Bit 7-6 - Wave Pattern Duty (Read/Write)
		     Bit 5-0 - Sound length data (Write Only) (t1: 0-63)
		   Wave Duty:
		     00: 12.5% ( _-------_-------_------- )
		     01: 25%   ( __------__------__------ )
		     10: 50%   ( ____----____----____---- ) (normal)
		     11: 75%   ( ______--______--______-- )
		   Sound Length = (64-t1)*(1/256) seconds
		   The Length value is used only if Bit 6 in NR14 is set.
		*/

	case NR12: // Channel 1 Volume Envelope (R/W)
		s.initialVolume1 = (val & 0xF0) >> 4
		s.envelopeIncrease1 = val&8 > 0
		s.numberEnvelopeSweep1 = val & 7
		/*

		   Bit 7-4 - Initial Volume of envelope (0-0Fh) (0=No Sound)
		   Bit 3   - Envelope Direction (0=Decrease, 1=Increase)
		   Bit 2-0 - Number of envelope sweep (n: 0-7)
		             (If zero, stop envelope operation.)
		   Length of 1 step = n*(1/64) seconds
		*/
	case NR13: //  Channel 1 Frequency lo (Write Only)
		s.frequency1 &= 0x700
		s.frequency1 |= uint16(val)
		/*
		   Lower 8 bits of 11 bit frequency (x).
		   Next 3 bit are in NR14 ($FF14)
		*/
	case NR14: //  Channel 1 Frequency hi (R/W)
		s.frequency1 &= 0xFF
		s.frequency1 |= uint16(val&0x3) << 8

		//restartSound := val&0x80 > 0
		s.stopExpire1 = val&0x40 > 0
		/*
		   Bit 7   - Initial (1=Restart Sound)     (Write Only)
		   Bit 6   - Counter/consecutive selection (Read/Write)
		               (1=Stop output when length in NR11 expires)
		   Bit 2-0 - Frequency's higher 3 bits (x) (Write Only)
		   Frequency = 131072/(2048-x) Hz
		*/
	case NR21: // Channel 2 Sound Length/Wave Pattern Duty (R/W)
		s.wavePatternDuty2 = (val & 0xC0) >> 6
		s.soundLength2 = 64 - (val & 0x3F)
		/*

		     Bit 7-6 - Wave Pattern Duty (Read/Write)
		     Bit 5-0 - Sound length data (Write Only) (t1: 0-63)
		   Wave Duty:
		     00: 12.5% ( _-------_-------_------- )
		     01: 25%   ( __------__------__------ )
		     10: 50%   ( ____----____----____---- ) (normal)
		     11: 75%   ( ______--______--______-- )
		   Sound Length = (64-t1)*(1/256) seconds
		   The Length value is used only if Bit 6 in NR14 is set.
		*/
	case NR22: // Channel 2 Volume Envelope (R/W)
		s.initialVolume2 = (val & 0xF0) >> 4
		s.envelopeIncrease2 = val&8 > 0
		s.numberEnvelopeSweep2 = val & 7
		/*

		   Bit 7-4 - Initial Volume of envelope (0-0Fh) (0=No Sound)
		   Bit 3   - Envelope Direction (0=Decrease, 1=Increase)
		   Bit 2-0 - Number of envelope sweep (n: 0-7)
		             (If zero, stop envelope operation.)
		   Length of 1 step = n*(1/64) seconds
		*/
	case NR23: // Channel 2 Frequency lo data (W)
		s.frequency2 &= 0x700
		s.frequency2 |= uint16(val)
		/*
		   Lower 8 bits of 11 bit frequency (x).
		   Next 3 bit are in NR14 ($FF19   )
		*/

	case NR24: // Channel 2 Frequency hi data (R/W)
		s.frequency2 &= 0xFF
		s.frequency2 |= uint16(val&0x3) << 8

		//restartSound := val&0x80 > 0
		s.stopExpire2 = val&0x40 > 0
		/*
		   Bit 7   - Initial (1=Restart Sound)     (Write Only)
		   Bit 6   - Counter/consecutive selection (Read/Write)
		               (1=Stop output when length in NR11 expires)
		   Bit 2-0 - Frequency's higher 3 bits (x) (Write Only)
		   Frequency = 131072/(2048-x) Hz
		*/
	}

	s.refreshRegs()
}

func (s *SoundCard) refreshRegs() {
	s.frequency1calc = 131072 / (2048 - float32(s.frequency1))
	s.frequency2calc = 131072 / (2048 - float32(s.frequency2))
	s.sweepTimeCalc = float32(s.sweepTime) / 128
	s.soundLength1calc = float32(s.soundLength1) / 256
	s.soundLength2calc = float32(s.soundLength2) / 256
}

func (s *SoundCard) Cycle(clocks int) {

}

func (s *SoundCard) Read(addr uint16) byte {
	if addr < 0xFF00 || addr > 0xFF40 {
		return 0xFF
	}

	/*
	   When register is read back, the stored value is or'ed by the following values
	        NRx0 NRx1 NRx2 NRx3 NRx4
	       ---------------------------
	   NR1x  $80  $3F $00  $FF  $BF
	   NR2x  $FF  $3F $00  $FF  $BF
	   NR3x  $7F  $FF $9F  $FF  $BF
	   NR4x  $FF  $FF $00  $00  $BF
	   NR5x  $00  $00 $70

	   $FF27-$FF2F always read back as $FF
	*/

	v := s.regs[addr-0xFF00]

	if addr >= 0xFF30 && addr <= 0xFF3F { // Wave Table
		return v
	}

	switch addr {
	case NR10:
		return v | 0x80
	case NR11:
		return v | 0x3F
	case NR12:
		return v
	case NR14:
		return v | 0xBF
	case NR21:
		return v | 0x3F
	case NR22:
		return v
	case NR24:
		return v | 0xBF
	case NR30:
		return v | 0x7F
	case NR32:
		return v | 0x9F
	case NR34:
		return v | 0xBF
	case NR42:
		return v
	case NR43:
		return v
	case NR44:
		return v | 0xBF
	case NR50:
		return v
	case NR51:
		return v
	case NR52:
		return v | 0x70
	default: // NR13, NR20, NR23, NR31, NR33, NR40, NR41
		return 0xFF
	}
}
