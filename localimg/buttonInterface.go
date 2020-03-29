package localimg

type Button int

type IButton interface {
	Pressed(button Button) bool
}

// List of all keyboard buttons.
const (
	KeyUnknown Button = iota
	KeySpace
	KeyApostrophe
	KeyComma
	KeyMinus
	KeyPeriod
	KeySlash
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeySemicolon
	KeyEqual
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyLeftBracket
	KeyBackslash
	KeyRightBracket
	KeyGraveAccent
	KeyWorld1
	KeyWorld2
	KeyEscape
	KeyEnter
	KeyTab
	KeyBackspace
	KeyInsert
	KeyDelete
	KeyRight
	KeyLeft
	KeyDown
	KeyUp
	KeyPageUp
	KeyPageDown
	KeyHome
	KeyEnd
	KeyCapsLock
	KeyScrollLock
	KeyNumLock
	KeyPrintScreen
	KeyPause
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyKP0
	KeyKP1
	KeyKP2
	KeyKP3
	KeyKP4
	KeyKP5
	KeyKP6
	KeyKP7
	KeyKP8
	KeyKP9
	KeyKPDecimal
	KeyKPDivide
	KeyKPMultiply
	KeyKPSubtract
	KeyKPAdd
	KeyKPEnter
	KeyKPEqual
	KeyLeftShift
	KeyLeftControl
	KeyLeftAlt
	KeyLeftSuper
	KeyRightShift
	KeyRightControl
	KeyRightAlt
	KeyRightSuper
	KeyMenu
	KeyLast
)

/*

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
		k.cpu.Registers.InterruptsFired |= gameboy.IntJoypad
	}
}
*/
