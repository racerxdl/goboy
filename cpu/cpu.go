package cpu

import (
	"fmt"
	"github.com/quan-to/slog"
	"github.com/racerxdl/goboy/gameboy"
	"runtime"
	"sync"
	"time"
)

// We need to generate the instructions
//go:generate go run generator/gen.go

var cpuLog = slog.Scope("CPU")

type Core struct {
	Registers *Registers
	Memory    *Memory
	GPU       *GPU
	Timer     *Timer
	Keys      *GBKeys
	SoundCard *SoundCard
	Serial    *Serial

	running        bool
	paused         bool
	lastUpdate     time.Time
	lastCycleTime  int64
	clockT, clockM int
	halted         bool
	stopped        bool
	step           bool
	speedMul       float64
	baseClock      time.Duration
	colorMode      bool

	l sync.Mutex
}

type DebugData struct {
	PC, SP, A, B, C, D, E, F, H, L, HL       string
	AB, BB, CB, DB, EB, FB, LB, HB           string
	BC, DE, EI, EIB, IF, IFB                 string
	PCX, SPX                                 string
	IME                                      string
	GPUSCROLLX, GPUSCROLLY, GPUWINX, GPUWINY string
	GPUMODECLOCKS, GPULINE                   string
	HALTED                                   string
	RamBank                                  string
	RomBank                                  string
}

func MakeCore() *Core {
	c := &Core{
		Registers: MakeRegisters(),
		running:   false,
		stopped:   false,
		paused:    true,
		speedMul:  1,
		baseClock: ColorModePeriod,
		colorMode: false,
	}
	c.Memory = MakeMemory(c)
	c.GPU = MakeGPU(c)
	c.Timer = MakeTimer(c)
	c.Keys = MakeGBKeys(c)
	c.SoundCard = MakeSoundCard(c)
	c.Serial = MakeSerial(c)
	c.Reset()
	return c
}

func (c *Core) Start() {
	c.l.Lock()
	if !c.running {
		c.running = true
		c.lastUpdate = time.Now()
		go c.loop()
	}
	c.l.Unlock()
}

func (c *Core) Stop() {
	c.l.Lock()

	if c.running {
		c.running = false
	}

	c.l.Unlock()
}

func (c *Core) SetSpeedHack(mul float64) {
	cpuLog.Info("Set speed to %0.2f", mul)
	c.speedMul = mul
}

func (c *Core) loop() {
	cpuLog.Info("CPU Loop Started")
	for c.running {
		if !c.paused {
			delta := time.Since(c.lastUpdate)
			c.lastCycleTime = delta.Nanoseconds()
			c.cycle()
			c.lastUpdate = time.Now()
			if c.step {
				c.step = false
				c.paused = true
			}
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}
	cpuLog.Info("CPU Loop Stopped")
}

func (c *Core) GetCurrentPage() (uint16, []byte) {
	a := (c.Registers.PC / 256) * 256
	buff := make([]byte, 256)
	for i := 0; i < 256; i++ {
		buff[i] = c.Memory.ReadByte(a + uint16(i))
	}

	return a, buff
}

func (c *Core) GetStack() (uint16, []byte) {
	e := int(c.Registers.SP) + 2
	s := e - 512

	if s < 0 {
		s = 0
	}

	buff := make([]byte, e-s)
	for i := 0; i < e-s; i++ {
		buff[i] = c.Memory.ReadByte(uint16(s + i))
	}

	return uint16(s), buff
}

func (c *Core) GetDebugData() DebugData {
	return DebugData{
		PCX: fmt.Sprintf("%04X", c.Registers.PC),
		SPX: fmt.Sprintf("%04X", c.Registers.SP),
		PC:  fmt.Sprintf("%05d", c.Registers.PC),
		SP:  fmt.Sprintf("%05d", c.Registers.SP),

		A:  fmt.Sprintf("%02X", c.Registers.A),
		B:  fmt.Sprintf("%02X", c.Registers.B),
		C:  fmt.Sprintf("%02X", c.Registers.C),
		D:  fmt.Sprintf("%02X", c.Registers.D),
		E:  fmt.Sprintf("%02X", c.Registers.E),
		H:  fmt.Sprintf("%02X", c.Registers.H),
		L:  fmt.Sprintf("%02X", c.Registers.L),
		F:  fmt.Sprintf("%02X", c.Registers.F),
		EI: fmt.Sprintf("%02X", c.Registers.EnabledInterrupts),
		IF: fmt.Sprintf("%02X", c.Registers.InterruptsFired),

		AB:  fmt.Sprintf("%08b", c.Registers.A),
		BB:  fmt.Sprintf("%08b", c.Registers.B),
		CB:  fmt.Sprintf("%08b", c.Registers.C),
		DB:  fmt.Sprintf("%08b", c.Registers.D),
		EB:  fmt.Sprintf("%08b", c.Registers.E),
		HB:  fmt.Sprintf("%08b", c.Registers.H),
		LB:  fmt.Sprintf("%08b", c.Registers.L),
		FB:  fmt.Sprintf("%08b", c.Registers.F),
		EIB: fmt.Sprintf("%08b", c.Registers.EnabledInterrupts),
		IFB: fmt.Sprintf("%08b", c.Registers.InterruptsFired),

		HL:  fmt.Sprintf("%04X", c.Registers.HL()),
		BC:  fmt.Sprintf("%04X", c.Registers.BC()),
		DE:  fmt.Sprintf("%04X", c.Registers.DE()),
		IME: fmt.Sprintf("%5v", c.Registers.InterruptEnable),

		GPUSCROLLX:    fmt.Sprintf("%4d", c.GPU.scrollX),
		GPUSCROLLY:    fmt.Sprintf("%4d", c.GPU.scrollY),
		GPUWINX:       fmt.Sprintf("%4d", c.GPU.winX),
		GPUWINY:       fmt.Sprintf("%4d", c.GPU.winY),
		GPUMODECLOCKS: fmt.Sprintf("%4d", c.GPU.modeClocks),
		GPULINE:       fmt.Sprintf("%4d", c.GPU.line),

		HALTED:  fmt.Sprintf("%v", c.halted),
		RamBank: fmt.Sprintf("%d", c.Memory.ramBank),
		RomBank: fmt.Sprintf("%d", c.Memory.catridge.RomBank()),
	}
}

func (c *Core) IsPaused() bool {
	return c.paused
}

const instructionsInterp = 1

var lastPrint = time.Now()

func (c *Core) cycle() {
	c.l.Lock()

	x := time.Now()
	waitingClockT := c.clockT

	// Normal Cycle
	for cc := 0; cc < instructionsInterp; cc++ {
		totalClockM := 0
		totalClockT := 0
		c.Registers.CycleCount++
		if !c.stopped {
			if c.halted {
				totalClockM += 1
				totalClockT += 4
			} else {
				op := c.Memory.ReadByte(c.Registers.PC)
				c.Registers.PC++
				GBInstructions[op](c)

				totalClockM += c.Registers.LastClockM
				totalClockT += c.Registers.LastClockT
			}

			// Check Interrupts
			if c.Registers.InterruptEnable && (c.Registers.EnabledInterrupts&c.Registers.InterruptsFired) > 0 {
				curHalt := c.halted
				c.halted = false
				c.Registers.InterruptEnable = false
				interruptsFired := c.Registers.EnabledInterrupts & c.Registers.InterruptsFired

				switch {
				case (interruptsFired & gameboy.IntVblank) > 0:
					c.Registers.InterruptsFired &= ^uint8(gameboy.IntVblank)
					gbRSTXX(c, AddrIntVblank) // V-Blank
					totalClockM += c.Registers.LastClockM
					totalClockT += c.Registers.LastClockT
				case (interruptsFired & gameboy.IntLcdstat) > 0:
					//cpuLog.Debug("(INT) [LCDSTAT]")
					c.Registers.InterruptsFired &= ^uint8(gameboy.IntLcdstat)
					gbRSTXX(c, AddrIntLcdstat) // LCD Stat
					totalClockM += c.Registers.LastClockM
					totalClockT += c.Registers.LastClockT
				case (interruptsFired & gameboy.IntTimer) > 0:
					//cpuLog.Debug("(INT) [TIMER]")
					c.Registers.InterruptsFired &= ^uint8(gameboy.IntTimer)
					gbRSTXX(c, AddrIntTimer) // Timer
					totalClockM += c.Registers.LastClockM
					totalClockT += c.Registers.LastClockT
				case (interruptsFired & gameboy.IntSerial) > 0:
					//cpuLog.Debug("(INT) [SERIAL]")
					c.Registers.InterruptsFired &= ^uint8(gameboy.IntSerial)
					gbRSTXX(c, AddrIntSerial) // Serial
					totalClockM += c.Registers.LastClockM
					totalClockT += c.Registers.LastClockT
				case (interruptsFired & gameboy.IntJoypad) > 0:
					c.Registers.InterruptsFired &= ^uint8(gameboy.IntJoypad)
					gbRSTXX(c, AddrIntJoypad) // Joypad Interrupt
					totalClockM += c.Registers.LastClockM
					totalClockT += c.Registers.LastClockT
				default:
					c.Registers.InterruptEnable = true
					c.halted = curHalt
				}
			}

			c.clockM += totalClockM
			c.clockT += totalClockT

			// Sound Flow
			c.SoundCard.Cycle(totalClockM)

			// GPU Flow
			c.GPU.Cycle(totalClockM)

			// Timer Flow
			c.Timer.Increment(totalClockT)

			// Serial Flow
			c.Serial.Cycle(totalClockM)

		}
	}
	waitingClockT = c.clockT - waitingClockT
	c.l.Unlock()

	cycleDuration := time.Duration(int64(waitingClockT)) * time.Duration(float64(c.baseClock)/c.speedMul)
	if time.Since(lastPrint) > time.Second/4 {
		fmt.Println("Cycle Duration", cycleDuration, waitingClockT, c.baseClock)
		lastPrint = time.Now()
	}
	//if time.Since(x) - cycleDuration > time.Millisecond * 10 {
	//	time.Sleep(time.Millisecond)
	//}
	// Sleep is not precise enough, so we will do a busy loop
	for time.Since(x) < time.Duration(float64(cycleDuration)*1.5) {
		runtime.Gosched()
	}

	if c.stopped && c.Memory.inPrepareMode {
		if !c.Memory.doubleSpeed {
			cpuLog.Info("Switching to Double Speed Mode")
			c.Memory.doubleSpeed = true
			c.baseClock = ColorModePeriod
			//c.paused = true
		}
		c.stopped = false
		c.Memory.inPrepareMode = false
	}
}

func (c *Core) Reset() {
	c.l.Lock()
	c.halted = false
	c.clockT = 0
	c.clockM = 0
	c.Registers.Reset()
	c.Memory.Reset()
	c.GPU.Reset()
	c.Timer.Reset()
	c.l.Unlock()
}

func (c *Core) Pause() {
	c.l.Lock()
	c.paused = true
	c.l.Unlock()
}

func (c *Core) Continue() {
	c.l.Lock()
	c.paused = false
	c.l.Unlock()
}

func (c *Core) Step() {
	c.l.Lock()
	c.step = true
	c.paused = false
	c.l.Unlock()
}
