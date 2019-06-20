package cpu

import (
	"github.com/quan-to/slog"
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

	l sync.Mutex
}

func MakeCore() *Core {
	c := &Core{
		Registers: MakeRegisters(),
		running:   false,
		stopped:   false,
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
			if delta.Nanoseconds() < CpuPeriodMs*1000 {
				continue
			}
			c.lastCycleTime = delta.Nanoseconds()
			c.cycle()
			c.lastUpdate = time.Now()
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}
	cpuLog.Info("CPU Loop Stopped")
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
