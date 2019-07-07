package cpu

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/racerxdl/goboy/gameboy"
)

type GBKeys struct {
	cpu           *Core
	direction     uint8
	keys          uint8
	selectedInput uint8
}

func MakeGBKeys(cpu *Core) *GBKeys {
	return &GBKeys{
		cpu:           cpu,
		direction:     0xF,
		keys:          0xF,
		selectedInput: 0,
	}
}

func (k *GBKeys) Write(addr uint16, val uint8) {
	k.selectedInput = val & 0x30
}

func (k *GBKeys) Read(addr uint16) uint8 {
	switch k.selectedInput {
	case 0x10:
		return k.keys | k.selectedInput
	case 0x20:
		return k.direction | k.selectedInput
	case 0x30: // That shouldn't happen, but thats what the hardware will return
		return k.keys | k.direction | k.selectedInput
	default:
		return 0xF | k.selectedInput
	}
}

func (k *GBKeys) SetDirectionBit(bit int, val bool) {
	if val {
		k.direction |= (byte)(1 << uint(bit))
	} else {
		k.direction &= (byte)(^(1 << uint(bit)))
	}

	k.direction &= 0xF
}

func (k *GBKeys) SetKeysBit(bit int, val bool) {
	if val {
		k.keys |= (byte)(1 << uint(bit))
	} else {
		k.keys &= (byte)(^(1 << uint(bit)))
	}

	k.keys &= 0xF
}

func (k *GBKeys) Update(win *pixelgl.Window) {
	lastK := k.keys
	lastD := k.direction

	k.SetDirectionBit(0, !win.Pressed(pixelgl.KeyRight))
	k.SetDirectionBit(1, !win.Pressed(pixelgl.KeyLeft))
	k.SetDirectionBit(2, !win.Pressed(pixelgl.KeyUp))
	k.SetDirectionBit(3, !win.Pressed(pixelgl.KeyDown))

	k.SetKeysBit(0, !win.Pressed(pixelgl.KeyZ))
	k.SetKeysBit(1, !win.Pressed(pixelgl.KeyX))
	k.SetKeysBit(2, !win.Pressed(pixelgl.KeySpace))
	k.SetKeysBit(3, !win.Pressed(pixelgl.KeyEnter))

	if lastK != k.keys || lastD != k.direction {
		k.cpu.Registers.TriggerInterrupts |= gameboy.IntJoypad
	}
}
