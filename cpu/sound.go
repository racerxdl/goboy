package cpu

import (
	"math"
	"sync"
	"time"
)

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
	sync.Mutex
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
	stopExpire1          bool
	channel1On           bool

	wavePatternDuty2     uint8
	soundLength2         uint8
	soundLength2calc     float32
	initialVolume2       uint8
	envelopeIncrease2    bool
	numberEnvelopeSweep2 uint8
	frequency2           uint16
	stopExpire2          bool
	wavTable             []uint8
	channel2On           bool

	sampleRate   float64
	samplePeriod time.Duration
	//phaseTest float64
	buffer     []float32
	cycleAcc   int64
	lastUpdate time.Time

	f1timerSampleCount int64
	lastSweep1Freq     int
	lastSweep1Sample   int64

	f2timerSampleCount int64
	lastSweep2Freq     int
	lastSweep2Sample   int64

	sound1Left  bool
	sound1Right bool
	sound2Left  bool
	sound2Right bool
	sound3Left  bool
	sound3Right bool
	sound4Left  bool
	sound4Right bool

	globalSoundEnable bool
}

func MakeSoundCard(cpu *Core) *SoundCard {
	return &SoundCard{
		cpu:          cpu,
		regs:         make([]byte, 0xFF),
		sampleRate:   48000,
		samplePeriod: time.Second / 48000,
		lastUpdate:   time.Now(),
	}
}

func (s *SoundCard) SetSampleRate(sampleRate float64) {
	s.Lock()
	s.sampleRate = sampleRate
	s.samplePeriod = time.Duration(float64(time.Second) / sampleRate)
	s.Unlock()
}

func (s *SoundCard) ProcessAudio(out [][]float32) {
	s.Lock()

	for i := range out[0] {
		out[0][i] = 0
		out[1][i] = 0
		if s.sound1Left {
			out[0][i] += s.GetFrequency1Sample()
		}
		if s.sound1Right {
			out[1][i] += s.GetFrequency1Sample()
		}
		if s.sound2Left {
			out[0][i] += s.GetFrequency2Sample()
		}
		if s.sound2Right {
			out[1][i] += s.GetFrequency2Sample()
		}

		if out[0][i] > 1 {
			out[0][i] = 1
		}

		if out[1][i] > 1 {
			out[1][i] = 1
		}
	}

	s.Unlock()
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
		if s.sweepTime != 0 {
			s.lastSweep1Freq = -1
			s.lastSweep1Sample = 0
			s.lastSweep2Sample = 0
		}

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
		s.frequency1 |= uint16(val&0x7) << 8

		restartSound := val&0x80 > 0
		s.stopExpire1 = val&0x40 > 0
		if restartSound {
			s.channel1On = true
			s.lastSweep1Freq = int(s.frequency1)
			s.f1timerSampleCount = 0
			s.lastSweep1Sample = 0
		}
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
		s.frequency2 |= uint16(val&0x7) << 8

		restartSound := val&0x80 > 0
		if restartSound {
			s.channel2On = true
			s.f2timerSampleCount = 0
			s.lastSweep2Freq = int(s.frequency2)
			s.lastSweep2Sample = 0
		}
		s.stopExpire2 = val&0x40 > 0
		/*
		   Bit 7   - Initial (1=Restart Sound)     (Write Only)
		   Bit 6   - Counter/consecutive selection (Read/Write)
		               (1=Stop output when length in NR11 expires)
		   Bit 2-0 - Frequency's higher 3 bits (x) (Write Only)
		   Frequency = 131072/(2048-x) Hz
		*/
	case NR51:
		/*
		 Bit 7 - Output sound 4 to SO2 terminal Left
		 Bit 6 - Output sound 3 to SO2 terminal Left
		 Bit 5 - Output sound 2 to SO2 terminal Left
		 Bit 4 - Output sound 1 to SO2 terminal Left

		 Bit 3 - Output sound 4 to SO1 terminal Right
		 Bit 2 - Output sound 3 to SO1 terminal Right
		 Bit 1 - Output sound 2 to SO1 terminal Right
		 Bit 0 - Output sound 1 to SO1 terminal Right
		*/
		s.sound1Left = val&0x80 > 0
		s.sound2Left = val&0x40 > 0
		s.sound3Left = val&0x20 > 0
		s.sound4Left = val&0x10 > 0

		s.sound1Right = val&0x08 > 0
		s.sound2Right = val&0x04 > 0
		s.sound3Right = val&0x02 > 0
		s.sound4Right = val&0x01 > 0
	case NR52:
		s.globalSoundEnable = val&0x80 > 0
	}

	s.refreshRegs()
}

func getFreq(val int) float32 {
	// = 4194304 4 2 2048 X Hz
	if val > 2047 {
		val = 2047
	}
	return 4194304 / (4 * 2 * 2 * 2 * (2048 - float32(val))) / 2
}

func (s *SoundCard) refreshRegs() {
	s.sweepTimeCalc = float32(s.sweepTime) / 128
	s.soundLength1calc = float32(s.soundLength1) / 256
	s.soundLength2calc = float32(s.soundLength2) / 256
}

func (s *SoundCard) GetFrequency1Sample() float32 {
	if !s.channel1On || !s.globalSoundEnable {
		s.f1timerSampleCount = 0
		s.lastSweep1Sample = 0
		return 0
	}

	f1period := float64(1e6 / getFreq(int(s.frequency1)))
	samplePeriodMicros := 1e6 / s.sampleRate
	s.f1timerSampleCount++

	// Calculate sweep if needed
	if s.sweepTime != 0 {
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
		sweepPeriodMicros := (float64(s.sweepTime) * 1e6) / 128
		sweepPeriodSamples := int64(sweepPeriodMicros / samplePeriodMicros)
		if s.f1timerSampleCount-s.lastSweep1Sample > sweepPeriodSamples {
			// Shift
			s.lastSweep1Sample = s.f1timerSampleCount
			shiftFactor := int(math.Pow(2, float64(s.sweepShift)))
			offset := s.lastSweep1Freq / shiftFactor
			if s.sweepSub {
				s.lastSweep1Freq -= offset
			} else {
				s.lastSweep1Freq += offset
			}

			if s.lastSweep1Freq <= 0 || s.lastSweep1Freq >= 2048 {
				s.channel1On = false
				s.lastSweep1Freq = 0
				return 0
			}
		}
		f1period = float64(1e6 / getFreq(s.lastSweep1Freq))
	}

	f1periodNumSamples := int64(f1period / samplePeriodMicros)

	sampleLengthMicros := float64(s.soundLength1calc * 1e6)
	microsSoundPassed := float64(s.f1timerSampleCount) * samplePeriodMicros
	if s.stopExpire1 && microsSoundPassed > sampleLengthMicros {
		s.channel1On = false
		s.f1timerSampleCount = 0
		return 0
	}

	vol := float32(s.initialVolume1)

	// Calculate envelope
	if s.numberEnvelopeSweep1 != 0 {
		stepLengthMicros := (float32(s.numberEnvelopeSweep1) * 1e6) / 64
		currentStep := int(float32(microsSoundPassed) / stepLengthMicros)
		if currentStep > 15 {
			currentStep = 15
		}
		if s.envelopeIncrease1 {
			vol += float32(currentStep)
		} else {
			vol -= float32(currentStep)
		}
		if vol > 15 {
			vol = 15
		}
		if vol <= 0 {
			return 0
		}
	}

	vol /= 15

	dutyPat := dutyCycles[s.wavePatternDuty1]
	samplesPerPoint := int(f1periodNumSamples) / len(dutyPat)
	if samplesPerPoint == 0 {
		return 0
	}
	p := int((s.f1timerSampleCount / int64(samplesPerPoint)) % int64(len(dutyPat)))
	currentSample := float32(dutyPat[p]) * vol

	return currentSample
}

func (s *SoundCard) GetFrequency2Sample() float32 {
	if !s.channel2On || !s.globalSoundEnable {
		s.f2timerSampleCount = 0
		s.lastSweep2Sample = 0
		return 0
	}

	f2period := float64(1e6 / getFreq(int(s.frequency2)))
	samplePeriodMicros := 1e6 / s.sampleRate
	s.f2timerSampleCount++

	// Calculate sweep if needed
	if s.sweepTime != 0 {
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
		sweepPeriodMicros := (float64(s.sweepTime) * 1e6) / 128
		sweepPeriodSamples := int64(sweepPeriodMicros / samplePeriodMicros)
		if s.f2timerSampleCount-s.lastSweep2Sample > sweepPeriodSamples {
			// Shift
			s.lastSweep2Sample = s.f2timerSampleCount
			shiftFactor := int(math.Pow(2, float64(s.sweepShift)))
			offset := s.lastSweep2Freq / shiftFactor
			if s.sweepSub {
				s.lastSweep2Freq -= offset
			} else {
				s.lastSweep2Freq += offset
			}

			if s.lastSweep2Freq <= 0 || s.lastSweep2Freq >= 2048 {
				s.channel2On = false
				s.lastSweep2Freq = 0
				return 0
			}
		}
		f2period = float64(1e6 / getFreq(s.lastSweep2Freq))
	}

	f2periodNumSamples := int64(f2period / samplePeriodMicros)

	sampleLengthMicros := float64(s.soundLength2calc * 1e6)
	microsSoundPassed := float64(s.f2timerSampleCount) * samplePeriodMicros
	if s.stopExpire2 && microsSoundPassed > sampleLengthMicros {
		s.channel2On = false
		s.f2timerSampleCount = 0
		return 0
	}

	vol := float32(s.initialVolume2)

	// Calculate envelope
	if s.numberEnvelopeSweep2 != 0 {
		stepLengthMicros := (float32(s.numberEnvelopeSweep2) * 1e6) / 64
		currentStep := int(float32(microsSoundPassed) / stepLengthMicros)
		if currentStep > 15 {
			currentStep = 15
		}
		if s.envelopeIncrease2 {
			vol += float32(currentStep)
		} else {
			vol -= float32(currentStep)
		}
		if vol > 15 {
			vol = 15
		}
		if vol <= 0 {
			return 0
		}
	}

	vol /= 15

	dutyPat := dutyCycles[s.wavePatternDuty2]
	samplesPerPoint := int(f2periodNumSamples) / len(dutyPat)
	if samplesPerPoint == 0 {
		return 0
	}
	p := int((s.f2timerSampleCount / int64(samplesPerPoint)) % int64(len(dutyPat)))
	currentSample := float32(dutyPat[p]) * vol

	return currentSample
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
		return v // v | 0x7F
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
		v := uint8(0)
		if s.channel1On {
			v |= 1
		}
		if s.channel2On {
			v |= 2
		}
		if s.globalSoundEnable {
			v |= 0x80
		}
		return v //| 0x70
	default: // NR13, NR20, NR23, NR31, NR33, NR40, NR41
		return 0xFF
	}
}
