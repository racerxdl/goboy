package cpu

type SoundCard struct {
	cpu *Core
}

func MakeSoundCard(cpu *Core) *SoundCard {
	return &SoundCard{
		cpu: cpu,
	}
}

func (s *SoundCard) Write(addr uint16, val uint8) {

}

func (s *SoundCard) Read(addr uint16) byte {

	return 0xFF
}
