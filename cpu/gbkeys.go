package cpu

import "github.com/faiface/pixel/pixelgl"

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

func (k *GBKeys) Write(val uint8) {
	k.selectedInput = val & 0x30
}

func (k *GBKeys) Read() uint8 {
	switch k.selectedInput {
	case 0x10:
		return k.keys | 0x10
	case 0x20:
		return k.direction | 0x20
	//case 0x30:
	//    return k.keys | k.direction | 0x20 | 0x10
	default:
		return k.selectedInput & 0xF
	}
}

func (k *GBKeys) SetDirectionBit(bit int, val bool) {
	if val {
		k.direction |= (byte)(1 << uint(bit))
	} else {
		k.direction &= (byte)((^(1 << uint(bit))) & 0xF)
	}
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

	k.SetDirectionBit(0, win.Pressed(pixelgl.KeyRight))
	k.SetDirectionBit(1, win.Pressed(pixelgl.KeyLeft))
	k.SetDirectionBit(2, win.Pressed(pixelgl.KeyUp))
	k.SetDirectionBit(3, win.Pressed(pixelgl.KeyDown))

	k.SetKeysBit(0, win.Pressed(pixelgl.KeyZ))
	k.SetKeysBit(1, win.Pressed(pixelgl.KeyX))
	k.SetKeysBit(2, win.Pressed(pixelgl.KeySpace))
	k.SetKeysBit(3, win.Pressed(pixelgl.KeyEnter))

	if lastK != k.keys || lastD != k.direction {
		k.cpu.Registers.TriggerInterrupts |= IntJoypad
	}
}
