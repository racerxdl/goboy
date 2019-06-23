package cpu

import (
	"fmt"
	"github.com/quan-to/slog"
	"strings"
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

	running        bool
	paused         bool
	lastUpdate     time.Time
	lastCycleTime  int64
	clockT, clockM int
	halted         bool
	stopped        bool
	step           bool

	l sync.Mutex
}

type DebugData struct {
	PC, SP, A, B, C, D, E, F, H, L, HL       string
	AB, BB, CB, DB, EB, FB, LB, HB           string
	PCX, SPX                                 string
	GPUSCROLLX, GPUSCROLLY, GPUWINX, GPUWINY string
	GPUMODECLOCKS, GPULINE                   string
}

func MakeCore() *Core {
	c := &Core{
		Registers: MakeRegisters(),
		running:   false,
		stopped:   false,
		paused:    true,
	}
	c.Memory = MakeMemory(c)
	c.GPU = MakeGPU(c)
	c.Timer = MakeTimer(c)
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
	a := (c.Registers.PC / 64) * 64
	buff := make([]byte, 64)
	for i := 0; i < 64; i++ {
		buff[i] = c.Memory.ReadByteNoSideEffect(a + uint16(i))
	}

	return a, buff
}

func (c *Core) GetDebugData() DebugData {
	return DebugData{
		PCX: strings.ToUpper(fmt.Sprintf("%04x", c.Registers.PC)),
		SPX: strings.ToUpper(fmt.Sprintf("%04x", c.Registers.SP)),
		PC:  fmt.Sprintf("%05d", c.Registers.PC),
		SP:  fmt.Sprintf("%05d", c.Registers.SP),

		A: strings.ToUpper(fmt.Sprintf("%02x", c.Registers.A)),
		B: strings.ToUpper(fmt.Sprintf("%02x", c.Registers.B)),
		C: strings.ToUpper(fmt.Sprintf("%02x", c.Registers.C)),
		D: strings.ToUpper(fmt.Sprintf("%02x", c.Registers.D)),
		E: strings.ToUpper(fmt.Sprintf("%02x", c.Registers.E)),
		H: strings.ToUpper(fmt.Sprintf("%02x", c.Registers.H)),
		L: strings.ToUpper(fmt.Sprintf("%02x", c.Registers.L)),
		F: strings.ToUpper(fmt.Sprintf("%02x", c.Registers.F)),

		AB: fmt.Sprintf("%08b", c.Registers.A),
		BB: fmt.Sprintf("%08b", c.Registers.B),
		CB: fmt.Sprintf("%08b", c.Registers.C),
		DB: fmt.Sprintf("%08b", c.Registers.D),
		EB: fmt.Sprintf("%08b", c.Registers.E),
		HB: fmt.Sprintf("%08b", c.Registers.H),
		LB: fmt.Sprintf("%08b", c.Registers.L),
		FB: fmt.Sprintf("%08b", c.Registers.F),

		HL: strings.ToUpper(fmt.Sprintf("%04x", c.Registers.HL())),

		GPUSCROLLX:    fmt.Sprintf("%4d", c.GPU.scrollX),
		GPUSCROLLY:    fmt.Sprintf("%4d", c.GPU.scrollY),
		GPUWINX:       fmt.Sprintf("%4d", c.GPU.winX),
		GPUWINY:       fmt.Sprintf("%4d", c.GPU.winX),
		GPUMODECLOCKS: fmt.Sprintf("%4d", c.GPU.modeClocks),
		GPULINE:       fmt.Sprintf("%4d", c.GPU.line),
	}
}

func (c *Core) IsPaused() bool {
	return c.paused
}

func (c *Core) cycle() {
	c.l.Lock()
	// Normal Cycle
	c.Registers.CycleCount++

	totalClockM := 0
	totalClockT := 0

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
	if c.Registers.InterruptEnable && c.Registers.EnabledInterrupts != 0 && c.Registers.TriggerInterrupts != 0 {
		c.halted = false
		c.Registers.InterruptEnable = false
		interruptsFired := c.Registers.EnabledInterrupts & c.Registers.TriggerInterrupts

		switch {
		case (interruptsFired & IntVblank) > 0:
			c.Registers.TriggerInterrupts &^= IntVblank
			gbRSTXX(c, AddrIntVblank) // V-Blank
			totalClockM += c.Registers.LastClockM
			totalClockT += c.Registers.LastClockT
		case (interruptsFired & IntLcdstat) > 0:
			cpuLog.Debug("(INT) [LCDSTAT]")
			c.Registers.TriggerInterrupts &^= IntLcdstat
			gbRSTXX(c, AddrIntLcdstat) // LCD Stat
			totalClockM += c.Registers.LastClockM
			totalClockT += c.Registers.LastClockT
		case (interruptsFired & IntTimer) > 0:
			cpuLog.Debug("(INT) [TIMER]")
			c.Registers.TriggerInterrupts &^= IntTimer
			gbRSTXX(c, AddrIntTimer) // Timer
			totalClockM += c.Registers.LastClockM
			totalClockT += c.Registers.LastClockT
		case (interruptsFired & IntSerial) > 0:
			cpuLog.Debug("(INT) [SERIAL]")
			c.Registers.TriggerInterrupts &^= IntSerial
			gbRSTXX(c, AddrIntSerial) // Serial
			totalClockM += c.Registers.LastClockM
			totalClockT += c.Registers.LastClockT
		case (interruptsFired & IntJoypad) > 0:
			c.Registers.TriggerInterrupts &^= IntJoypad
			gbRSTXX(c, AddrIntJoypad) // Joypad Interrupt
			totalClockM += c.Registers.LastClockM
			totalClockT += c.Registers.LastClockT
		default:
			c.Registers.InterruptEnable = true
		}
	}

	c.clockM += totalClockM
	c.clockT += totalClockT

	if !c.stopped {
		// GPU Flow
		c.GPU.Cycle()

		// Timer Flow
		c.Timer.Increment()
	}

	c.l.Unlock()

	cycleDuration := time.Second * time.Duration(int64(totalClockM)) / CpuClock
	time.Sleep(cycleDuration)
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
