package cpu

import (
	"math"
	"sync"
	"time"
)

const (
	NR10     uint16 = 0xff10
	NR11            = 0xff11
	NR12            = 0xff12
	NR13            = 0xff13
	NR14            = 0xff14
	NR21            = 0xff16
	NR22            = 0xff17
	NR23            = 0xff18
	NR24            = 0xff19
	NR30            = 0xff1a
	NR31            = 0xff1b
	NR32            = 0xff1c
	NR33            = 0xff1d
	NR34            = 0xff1e
	NR41            = 0xff20
	NR42            = 0xff21
	NR43            = 0xff22
	NR44            = 0xff23
	NR50            = 0xff24
	NR51            = 0xff25
	NR52            = 0xff26
	WAVSTART        = 0xff30
	WAVEND          = 0xff3f
)

var duty1 = []uint8{0, 1, 1, 1, 1, 1, 1, 1}
var duty2 = []uint8{0, 0, 1, 1, 1, 1, 1, 1}
var duty3 = []uint8{0, 0, 0, 0, 1, 1, 1, 1}
var duty4 = []uint8{0, 0, 0, 0, 0, 0, 1, 1}

var dutyCycles = [][]uint8{
	duty1, // 00: 12.5%
	duty2, // 01: 25%
	duty3, // 10: 50% (normal)
	duty4, // 11: 75%
}

type SoundCard struct {
	sync.Mutex
	cpu  *Core
	regs []byte

	sweepTime   uint8
	sweepSub    bool
	sweepShift  uint8
	sweepTimer  uint8
	sweepEnable bool
	shadowFreq  uint16

	wavePatternDuty1     uint8
	soundLength1         uint8
	initialVolume1       uint8
	envelopeIncrease1    bool
	numberEnvelopeSweep1 uint8
	frequency1           uint16
	stopExpire1          bool
	channel1On           bool
	f1phase              float64

	wavePatternDuty2     uint8
	soundLength2         uint8
	initialVolume2       uint8
	envelopeIncrease2    bool
	numberEnvelopeSweep2 uint8
	frequency2           uint16
	stopExpire2          bool
	channel2On           bool
	f2phase              float64

	wavTable     []uint8
	channel3On   bool
	sound3Enable bool
	soundLength3 uint16
	frequency3   uint16
	sound3volume uint8
	stopExpire3  bool
	f3phase      float64

	channel4On           bool
	sound4LSFR           uint16
	sound4Enable         bool
	soundLength4         uint8
	stopExpire4          bool
	initialVolume4       uint8
	envelopeIncrease4    bool
	numberEnvelopeSweep4 uint8
	shiftClockFrequency  uint8
	countStep15Bit       bool
	sound4Frequency      uint8
	f4phase              float64

	sampleRate   float64
	samplePeriod time.Duration
	buffer       []float32
	cycleAcc     int64
	lastUpdate   time.Time

	sound1Left  bool
	sound1Right bool
	sound2Left  bool
	sound2Right bool
	sound3Left  bool
	sound3Right bool
	sound4Left  bool
	sound4Right bool

	globalSoundEnable bool

	so1Volume uint8
	so2Volume uint8

	frameSeqPeriod int64
	frameSeqPhase  uint8
	envelopeTimer1 uint8
	envelopeTimer2 uint8
	envelopeTimer4 uint8
}

func MakeSoundCard(cpu *Core) *SoundCard {
	return &SoundCard{
		cpu:          cpu,
		regs:         make([]byte, 0xFF),
		sampleRate:   48000,
		samplePeriod: time.Second / 48000,
		lastUpdate:   time.Now(),
		wavTable:     make([]uint8, 32),
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

	so1Scale := float32(s.so1Volume) / 7
	so2Scale := float32(s.so2Volume) / 7

	for i := range out[0] {
		s1 := s.GetFrequency1Sample()
		s2 := s.GetFrequency2Sample()
		s3 := s.GetFrequency3Sample()
		s4 := s.GetFrequency4Sample()

		out[0][i] = 0
		out[1][i] = 0
		if s.sound1Left {
			out[0][i] += s1 / 4
		}
		if s.sound1Right {
			out[1][i] += s1 / 4
		}
		if s.sound2Left {
			out[0][i] += s2 / 4
		}
		if s.sound2Right {
			out[1][i] += s2 / 4
		}
		if s.sound3Left {
			out[0][i] += s3 / 4
		}
		if s.sound3Right {
			out[1][i] += s3 / 4
		}
		if s.sound4Left {
			out[0][i] += s4 / 4
		}
		if s.sound4Right {
			out[1][i] += s4 / 4
		}

		out[0][i] *= so1Scale
		out[1][i] *= so2Scale

		if out[0][i] > 1 {
			out[0][i] = 1
		}
		if out[0][i] < -1 {
			out[0][i] = -1
		}
		if out[1][i] > 1 {
			out[1][i] = 1
		}
		if out[1][i] < -1 {
			out[1][i] = -1
		}
	}

	s.Unlock()
}

func (s *SoundCard) Write(addr uint16, val uint8) {
	if addr < 0xFF40 && addr >= 0xFF00 {
		s.regs[addr-0xFF00] = val
	}

	switch addr {
	case NR10:
		s.sweepTime = (val & 0x70) >> 4
		s.sweepSub = (val & 8) > 0
		s.sweepShift = val & 7

	case NR11:
		s.wavePatternDuty1 = (val & 0xC0) >> 6
		s.soundLength1 = 64 - (val & 0x3F)

	case NR12:
		s.initialVolume1 = (val & 0xF0) >> 4
		s.envelopeIncrease1 = val&8 > 0
		s.numberEnvelopeSweep1 = val & 7

	case NR13:
		s.frequency1 &= 0x700
		s.frequency1 |= uint16(val)

	case NR14:
		s.frequency1 &= 0xFF
		s.frequency1 |= uint16(val&0x7) << 8
		s.stopExpire1 = val&0x40 > 0

		if val&0x80 > 0 {
			s.channel1On = true
			s.shadowFreq = s.frequency1
			s.f1phase = 0
			s.envelopeTimer1 = 0
			s.initialVolume1 = (s.regs[NR12-0xFF00] & 0xF0) >> 4

			s.sweepTimer = s.sweepTime
			if s.sweepTimer == 0 {
				s.sweepTimer = 8
			}
			s.sweepEnable = s.sweepTime != 0 || s.sweepShift != 0

			if s.sweepEnable && s.sweepShift > 0 {
				offset := int(s.shadowFreq >> s.sweepShift)
				newFreq := int(s.shadowFreq) + offset
				if s.sweepSub {
					newFreq = int(s.shadowFreq) - offset
				}
				if newFreq > 2047 || newFreq < 0 {
					s.channel1On = false
				}
			}

			if s.stopExpire1 && s.soundLength1 == 0 {
				s.soundLength1 = 63
			}
		}

	case NR21:
		s.wavePatternDuty2 = (val & 0xC0) >> 6
		s.soundLength2 = 64 - (val & 0x3F)

	case NR22:
		s.initialVolume2 = (val & 0xF0) >> 4
		s.envelopeIncrease2 = val&8 > 0
		s.numberEnvelopeSweep2 = val & 7

	case NR23:
		s.frequency2 &= 0x700
		s.frequency2 |= uint16(val)

	case NR24:
		s.frequency2 &= 0xFF
		s.frequency2 |= uint16(val&0x7) << 8
		s.stopExpire2 = val&0x40 > 0

		if val&0x80 > 0 {
			s.channel2On = true
			s.f2phase = 0
			s.envelopeTimer2 = 0
			s.initialVolume2 = (s.regs[NR22-0xFF00] & 0xF0) >> 4
			if s.stopExpire2 && s.soundLength2 == 0 {
				s.soundLength2 = 63
			}
		}

	case NR30:
		s.sound3Enable = val&0x80 > 0
	case NR31:
		s.soundLength3 = 256 - uint16(val)
	case NR32:
		s.sound3volume = (val & 0x60) >> 5
	case NR33:
		s.frequency3 &= 0x700
		s.frequency3 |= uint16(val)
	case NR34:
		s.frequency3 &= 0xFF
		s.frequency3 |= uint16(val&0x7) << 8
		s.stopExpire3 = val&0x40 > 0

		if val&0x80 > 0 {
			s.channel3On = true
			s.f3phase = 0
			if s.stopExpire3 && s.soundLength3 == 0 {
				s.soundLength3 = 255
			}
		}

	case NR41:
		s.soundLength4 = 64 - (val & 0x3F)
	case NR42:
		s.initialVolume4 = (val & 0xF0) >> 4
		s.envelopeIncrease4 = val&8 > 0
		s.numberEnvelopeSweep4 = val & 7
		s.sound4Enable = s.initialVolume4 > 0 || s.envelopeIncrease4
	case NR43:
		s.shiftClockFrequency = (val & 0xF0) >> 4
		s.countStep15Bit = (val & 0x08) == 0
		s.sound4Frequency = val & 7
	case NR44:
		s.stopExpire4 = val&0x40 > 0
		if val&0x80 > 0 {
			s.channel4On = true
			s.f4phase = 0
			s.sound4LSFR = 0x7FFF
			s.envelopeTimer4 = 0
			s.initialVolume4 = (s.regs[NR42-0xFF00] & 0xF0) >> 4
			if s.stopExpire4 && s.soundLength4 == 0 {
				s.soundLength4 = 63
			}
		}
	case NR51:
		s.sound1Left = val&0x80 > 0
		s.sound2Left = val&0x40 > 0
		s.sound3Left = val&0x20 > 0
		s.sound4Left = val&0x10 > 0

		s.sound1Right = val&0x08 > 0
		s.sound2Right = val&0x04 > 0
		s.sound3Right = val&0x02 > 0
		s.sound4Right = val&0x01 > 0
	case NR50:
		s.so1Volume = val & 7
		s.so2Volume = (val >> 4) & 7
	case NR52:
		s.globalSoundEnable = val&0x80 > 0
	}

	if addr >= WAVSTART && addr <= WAVEND {
		sample0 := (val & 0xF0) >> 4
		sample1 := val & 0xF
		i := (addr - WAVSTART) * 2
		s.wavTable[i+0] = sample0
		s.wavTable[i+1] = sample1
	}
}

func getFreq(val int) float32 {
	if val > 2047 {
		val = 2047
	}
	return 4194304 / (32 * (2048 - float32(val)))
}

func getWavFreq(val int) float32 {
	if val > 2047 {
		val = 2047
	}
	return 4194304 / (64 * (2048 - float32(val)))
}

func getNoiseFreq(r, s int) float32 {
	rf := float64(r)
	sf := float64(s)
	if r == 0 {
		rf = 0.5
	}
	return float32(524288 / rf / math.Pow(2, sf+1))
}

func (s *SoundCard) GetFrequency1Sample() float32 {
	if !s.channel1On || !s.globalSoundEnable {
		return 0
	}

	freq := float64(getFreq(int(s.frequency1)))
	s.f1phase += freq / s.sampleRate
	for s.f1phase >= 1.0 {
		s.f1phase -= 1.0
	}

	vol := float32(s.initialVolume1) / 15
	pos := int(s.f1phase*8) % 8
	dutyPat := dutyCycles[s.wavePatternDuty1]
	return (float32(dutyPat[pos])*2 - 1) * vol
}

func (s *SoundCard) GetFrequency2Sample() float32 {
	if !s.channel2On || !s.globalSoundEnable {
		return 0
	}

	freq := float64(getFreq(int(s.frequency2)))
	s.f2phase += freq / s.sampleRate
	for s.f2phase >= 1.0 {
		s.f2phase -= 1.0
	}

	vol := float32(s.initialVolume2) / 15
	pos := int(s.f2phase*8) % 8
	dutyPat := dutyCycles[s.wavePatternDuty2]
	return (float32(dutyPat[pos])*2 - 1) * vol
}

func (s *SoundCard) GetFrequency3Sample() float32 {
	if !(s.channel3On && s.sound3Enable) {
		return 0
	}

	freq := float64(getWavFreq(int(s.frequency3)))
	s.f3phase += freq / s.sampleRate
	for s.f3phase >= 1.0 {
		s.f3phase -= 1.0
	}

	pos := int(s.f3phase*32) % 32
	tableSample := float32(s.wavTable[pos])
	switch s.sound3volume {
	case 0:
		tableSample = 0
	case 2:
		tableSample /= 2
	case 3:
		tableSample /= 4
	}

	return (tableSample/7.5 - 1)
}

func (s *SoundCard) GetFrequency4Sample() float32 {
	if !s.channel4On || !s.globalSoundEnable || !s.sound4Enable {
		return 0
	}

	noiseFreq := float64(getNoiseFreq(int(s.sound4Frequency), int(s.shiftClockFrequency)))
	s.f4phase += noiseFreq / s.sampleRate

	for s.f4phase >= 1.0 {
		s.f4phase -= 1.0
		newBit := (s.sound4LSFR & 1) ^ ((s.sound4LSFR >> 1) & 1)
		s.sound4LSFR >>= 1
		if !s.countStep15Bit {
			s.sound4LSFR = (s.sound4LSFR & 0x3F) | (newBit << 6)
		} else {
			s.sound4LSFR |= (newBit << 14)
		}
	}

	vol := float32(s.initialVolume4) / 15
	if s.sound4LSFR&1 == 1 {
		return vol
	}
	return -vol
}

func (s *SoundCard) Cycle(clocks int) {
	if !s.globalSoundEnable {
		return
	}

	s.frameSeqPeriod += int64(clocks)

	for s.frameSeqPeriod >= 2048 {
		s.frameSeqPeriod -= 2048

		phase := s.frameSeqPhase
		s.frameSeqPhase++
		if s.frameSeqPhase > 7 {
			s.frameSeqPhase = 0
		}

		if phase == 0 || phase == 2 || phase == 4 || phase == 6 {
			if s.stopExpire1 && s.soundLength1 > 0 {
				s.soundLength1--
				if s.soundLength1 == 0 {
					s.channel1On = false
				}
			}
			if s.stopExpire2 && s.soundLength2 > 0 {
				s.soundLength2--
				if s.soundLength2 == 0 {
					s.channel2On = false
				}
			}
			if s.stopExpire3 && s.soundLength3 > 0 {
				s.soundLength3--
				if s.soundLength3 == 0 {
					s.channel3On = false
				}
			}
			if s.stopExpire4 && s.soundLength4 > 0 {
				s.soundLength4--
				if s.soundLength4 == 0 {
					s.channel4On = false
				}
			}
		}

		if phase == 2 || phase == 6 {
			s.tickSweep()
		}

		if phase == 7 {
			s.tickEnvelope1()
			s.tickEnvelope2()
			s.tickEnvelope4()
		}
	}
}

func (s *SoundCard) tickEnvelope1() {
	if s.numberEnvelopeSweep1 == 0 {
		return
	}
	s.envelopeTimer1++
	if s.envelopeTimer1 >= s.numberEnvelopeSweep1 {
		s.envelopeTimer1 = 0
		if s.envelopeIncrease1 && s.initialVolume1 < 15 {
			s.initialVolume1++
		} else if !s.envelopeIncrease1 && s.initialVolume1 > 0 {
			s.initialVolume1--
		}
	}
}

func (s *SoundCard) tickEnvelope2() {
	if s.numberEnvelopeSweep2 == 0 {
		return
	}
	s.envelopeTimer2++
	if s.envelopeTimer2 >= s.numberEnvelopeSweep2 {
		s.envelopeTimer2 = 0
		if s.envelopeIncrease2 && s.initialVolume2 < 15 {
			s.initialVolume2++
		} else if !s.envelopeIncrease2 && s.initialVolume2 > 0 {
			s.initialVolume2--
		}
	}
}

func (s *SoundCard) tickEnvelope4() {
	if s.numberEnvelopeSweep4 == 0 {
		return
	}
	s.envelopeTimer4++
	if s.envelopeTimer4 >= s.numberEnvelopeSweep4 {
		s.envelopeTimer4 = 0
		if s.envelopeIncrease4 && s.initialVolume4 < 15 {
			s.initialVolume4++
		} else if !s.envelopeIncrease4 && s.initialVolume4 > 0 {
			s.initialVolume4--
		}
	}
}

func (s *SoundCard) tickSweep() {
	if !s.channel1On {
		return
	}

	if s.sweepTimer > 0 {
		s.sweepTimer--
	}

	if s.sweepTimer == 0 {
		s.sweepTimer = s.sweepTime
		if s.sweepTimer == 0 {
			s.sweepTimer = 8
		}

		if s.sweepEnable && s.sweepShift > 0 {
			offset := int(s.shadowFreq >> s.sweepShift)
			var newFreq int
			if s.sweepSub {
				newFreq = int(s.shadowFreq) - offset
			} else {
				newFreq = int(s.shadowFreq) + offset
			}

			if newFreq > 2047 || newFreq < 0 {
				s.channel1On = false
				return
			}

			s.shadowFreq = uint16(newFreq)
			s.frequency1 = s.shadowFreq

			var checkFreq int
			if s.sweepSub {
				checkFreq = int(s.shadowFreq) - int(s.shadowFreq>>s.sweepShift)
			} else {
				checkFreq = int(s.shadowFreq) + int(s.shadowFreq>>s.sweepShift)
			}
			if checkFreq > 2047 || checkFreq < 0 {
				s.channel1On = false
			}
		}
	}
}

func (s *SoundCard) Read(addr uint16) byte {
	if addr < 0xFF00 || addr > 0xFF40 {
		return 0xFF
	}

	v := s.regs[addr-0xFF00]

	if addr >= 0xFF30 && addr <= 0xFF3F {
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
		v := uint8(0)
		if s.channel1On {
			v |= 1
		}
		if s.channel2On {
			v |= 2
		}
		if s.channel3On && s.sound3Enable {
			v |= 4
		}
		if s.channel4On {
			v |= 8
		}
		if s.globalSoundEnable {
			v |= 0x80
		}
		return v | 0x70
	default:
		return 0xFF
	}
}
