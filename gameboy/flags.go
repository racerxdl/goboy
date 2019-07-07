package gameboy

//region Interrupts
const (
	IntVblank  = 0x01
	IntLcdstat = 0x02
	IntTimer   = 0x04
	IntSerial  = 0x08
	IntJoypad  = 0x10
)

//endregion
//region LCDSTAT
const (
	FlagLycLy      = 0x40
	FlagOamMode    = 0x20
	FlagVblankMode = 0x10
	FlagHblankMode = 0x08
)

//endregion
//region Registers
const (
	FlagCarry        = 0x10
	FlagHalfCarry    = 0x20
	FlagSub          = 0x40
	FlagZero         = 0x80
	InvFlagCarry     = 0xEF
	InvFlagHalfCarry = 0xDF
	InvFlagSub       = 0xBF
	InvFlagZero      = 0x7F
)

//endregion
