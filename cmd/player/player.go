package main

import (
	"encoding/binary"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/racerxdl/goboy/cpu"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"html/template"
	"io/ioutil"
	"strings"
)

var z80 = cpu.MakeCore()

var debugTemplate = template.Must(template.New("debugger").Parse(`
              CPU DATA
       /---------------------\
       |      |  DEC  | HEX  |
       |  PC  | {{.PC}} | {{.PCX}} |
       |  SP  | {{.SP}} | {{.SPX}} |
       |---------------------|
      /                       \
/-----------------------------------\
|      |  HEX  |       BINARY       |
|  HL  |  {{.HL}} | {{.HB}} {{.LB}}  |
|  BC  |  {{.BC}} | {{.BB}} {{.CB}}  |
|  DE  |  {{.DE}} | {{.DB}} {{.EB}}  |
|  H   |    {{.H}} | -------- {{.HB}}  |
|  L   |    {{.L}} | -------- {{.LB}}  |
|------|-------|--------------------|
|  A   |    {{.A}} | -------- {{.AB}}  |
|  B   |    {{.B}} | -------- {{.BB}}  |
|  C   |    {{.C}} | -------- {{.CB}}  |
|  D   |    {{.D}} | -------- {{.DB}}  |
|  E   |    {{.E}} | -------- {{.EB}}  |
|------|-------|--------------------|
|  EI  |    {{.EI}} | -------- {{.EIB}}  |
|  IF  |    {{.IF}} | -------- {{.IFB}}  |
|  IME | {{.IME}} | -------- --------  |
|      |                  ZSHC----  |
|  F   |    {{.F}} | -------- {{.FB}}  |
\-----------------------------------/

       Halted: {{.HALTED}}

               GPU DATA
      /-----------------------\
      | Scroll X       | {{.GPUSCROLLX}} |
      | Scroll Y       | {{.GPUSCROLLY}} |
      | Window X       | {{.GPUWINX}} |
      | Window Y       | {{.GPUWINY}} |
      | Mode Clocks    | {{.GPUMODECLOCKS}} |
      | Line           | {{.GPULINE}} |
      \-----------------------/
`))

var screenOrigin pixel.Matrix
var disasmText *text.Text
var stackText *text.Text

func MoveAndScaleTo(p pixel.Picture, x, y, s float64) pixel.Matrix {
	return pixel.IM.
		Moved(pixel.V(p.Bounds().W()/2+x/s, p.Bounds().H()/2+y/s)).
		Scaled(pixel.V(0, 0), s).
		Chained(screenOrigin)
}

func OverObject(v pixel.Vec, p pixel.Picture, matrix pixel.Matrix) bool {
	b := p.Bounds().Size() // matrix.Unproject(p.Bounds().Size())
	w := b.X
	h := b.Y
	v = matrix.Unproject(v)
	v.X += w / 2
	v.Y += h / 2

	return v.X > 0 && v.X < w && v.Y > 0 && v.Y < h
}

const maxDisasmLines = 32
const maxStackLines = 16

func RefreshStack() {
	stackText.Clear()
	stackText.Color = colornames.Black
	fmt.Fprintf(stackText, "Stack: \n\n")

	offset, page := z80.GetStack()
	l := len(page)
	if l == 0 {
		return
	}

	b := maxStackLines * 2

	for i := 0; i < l; i += 2 {
		if i < l-b {
			continue
		}

		addr := offset + uint16(i)
		stackText.Color = colornames.Black

		v := uint16(page[i+1]<<8) + uint16(page[i])

		if addr == z80.Registers.SP {
			stackText.Color = colornames.Blue
		}
		fmt.Fprintf(stackText, "\t%04X: %02X %02X (%06d)\n", addr, page[i+1], page[i], v)
	}
}

func RefreshDisasm() {
	offset, page := z80.GetCurrentPage()
	dis := cpu.Disasm(int(offset), page)

	if len(dis) > maxDisasmLines {
		o := 0

		for i, v := range dis { // Find where PC is in disasm
			if uint16(v.Address) == z80.Registers.PC {
				o = i
				break
			}
		}

		for o > maxDisasmLines { // Move disasm lines up until PC is visible
			dis = dis[maxDisasmLines:]
			o -= maxDisasmLines
		}
		l := maxDisasmLines
		if len(dis) < l {
			l = len(dis)
		}
		dis = dis[:l]
	}

	disasmText.Clear()
	disasmText.Color = colornames.Black
	fmt.Fprintf(disasmText, "Disassembler: \n\n")
	for i, d := range dis {

		if i > 0 && strings.Contains(dis[i-1].Instruction, "RET") {
			fmt.Fprintf(disasmText, "\n") // empty line on RET
		}

		disasmText.Color = colornames.Black

		if d.Address == int(z80.Registers.PC) {
			disasmText.Color = colornames.Blue
		}
		fmt.Fprintf(disasmText, "\t%04X: ", d.Address)

		disasmText.Color = colornames.Black

		if strings.Contains(d.Instruction, "RET") {
			disasmText.Color = colornames.Blueviolet
		}

		if strings.Contains(d.Instruction, "CALL") {
			disasmText.Color = colornames.Blueviolet
		}

		v := d.Instruction
		if strings.Contains(d.Instruction, "d8") { // Imediated 8 bit unsigned
			v = strings.ReplaceAll(v, "d8", fmt.Sprintf("$%02X", d.Argument[0]))
		} else if strings.Contains(d.Instruction, "d16") { // Imediated 16 bit unsigned
			v = strings.ReplaceAll(v, "d16", fmt.Sprintf("$%04X", binary.LittleEndian.Uint16(d.Argument)))
		} else if strings.Contains(d.Instruction, "a8") { // 8 bit unsigned relative to $FF00
			v = strings.ReplaceAll(v, "a8", fmt.Sprintf("%d", d.Argument[0]))
		} else if strings.Contains(d.Instruction, "a16") { // 16 bit unsigned address
			v = strings.ReplaceAll(v, "a16", fmt.Sprintf("$%04X", binary.LittleEndian.Uint16(d.Argument)))
		} else if strings.Contains(d.Instruction, "r8") { // 8 bit signed
			if strings.Contains(d.Instruction, "J") { // Jump Relative
				o := uint16(int(d.Address)+int(int8(d.Argument[0]))) + 2
				v = strings.ReplaceAll(v, "r8", fmt.Sprintf("$%04X", o))
			} else {
				v = strings.ReplaceAll(v, "r8", fmt.Sprintf("%d", int8(d.Argument[0])))
			}
		}

		fmt.Fprintf(disasmText, "%s\n", v)
	}
}

func run() {
	//game, err := ioutil.ReadFile("./opus5.gb")
	//game, err := ioutil.ReadFile("./tetris.gb")
	game, err := ioutil.ReadFile("./cpu_instrs.gb")
	//game, err := ioutil.ReadFile("/home/lucas/Pokemon - Blue Version (UE) [S][!].gb")
	//game, err := ioutil.ReadFile("/home/lucas/Legend of Zelda, The - Link's Awakening (U) (V1.2) [!].gb")
	//game, err := ioutil.ReadFile("/home/lucas/Works/GBxCart-RW/Interface_Programs/GBxCart_RW_Console_Flasher_v1.19/ZELDA-DX.GB")
	//game, err := ioutil.ReadFile("/home/lucas/Works/gb-test-roms/cpu_instrs/individual/02-interrupts.gb")
	//game, err := ioutil.ReadFile("/home/lucas/Works/gb-test-roms/instr_timing/instr_timing.gb")
	//game, err := ioutil.ReadFile("/home/lucas/Works/gb-test-roms/interrupt_time/interrupt_time.gb")
	//game, err := ioutil.ReadFile("/home/lucas/Works/gb-test-roms/dmg_sound/rom_singles/02-len ctr.gb")
	if err != nil {
		panic(err)
	}

	z80.Memory.LoadRom(game)
	z80.Memory.SetSaveFile(fmt.Sprintf("%s.sav", z80.Memory.RomName()))
	z80.Memory.LoadCatridgeRAMData()

	cfg := pixelgl.WindowConfig{
		Title:  "GameBoy Emulator",
		Bounds: pixel.R(0, 0, 1280, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	screenOrigin = pixel.IM.
		ScaledXY(pixel.V(0, 0), pixel.V(1, -1)).
		Moved(pixel.V(0, win.Bounds().H()))

	fmt.Println(screenOrigin)

	atlas := text.NewAtlas(
		basicfont.Face7x13,
		text.ASCII,
	)

	z80.Start()

	debugger := text.New(pixel.V(0, 0), atlas)
	debugger.Color = colornames.Black

	disasmText = text.New(pixel.V(0, 0), atlas)
	stackText = text.New(pixel.V(0, 0), atlas)

	lcdText := text.New(pixel.V(0, 0), atlas)
	lcdText.Color = colornames.Black
	lcdText.WriteString("LCD")

	vramText := text.New(pixel.V(0, 0), atlas)
	vramText.Color = colornames.Black
	vramText.WriteString("Video RAM")

	tileBufferText := text.New(pixel.V(0, 0), atlas)
	tileBufferText.Color = colornames.Black
	tileBufferText.WriteString("Tile Buffer")

	r := win.Bounds()
	w := r.Max.X
	h := r.Max.Y

	RefreshDisasm()

	win.SetTitle(fmt.Sprintf("GameBoy Emulator - %s", z80.Memory.RomName()))
	win.SetSmooth(false)

	hh := true

	for !win.Closed() {
		vframe := z80.Memory.GetVideoFrame()
		vram := z80.GPU.GetBGRam()
		if !hh {
			vram = z80.GPU.GetWinRam()
		}
		tilebuff := z80.GPU.GetTileBuffer()

		mp := win.MousePosition()

		if OverObject(mp, vram, MoveAndScaleTo(vram, 10, 350, 1.25)) {
			// Highlight Tile
			n := MoveAndScaleTo(vram, 10, 350, 1.25).Unproject(mp)
			n.X += vram.Bounds().Size().X / 2
			n.Y += vram.Bounds().Size().Y / 2
			z80.GPU.SetHighlightTile(n.X, n.Y)
		} else {
			z80.GPU.SetHighlightTile(-1, -1)
		}
		win.Clear(colornames.Skyblue)

		// region LCD
		pixel.NewSprite(vframe, vframe.Bounds()).
			Draw(win, MoveAndScaleTo(vframe, 10, 30, 2))
		lcdText.Draw(win, pixel.IM.Moved(pixel.V(155, h-20)))
		// endregion
		// region Video Ram
		pixel.NewSprite(vram, vram.Bounds()).
			Draw(win, MoveAndScaleTo(vram, 10, 350, 1.25))
		vramText.Draw(win, pixel.IM.Moved(pixel.V(140, h-340)))
		// endregion
		// region Tile Buffer
		pixel.NewSprite(tilebuff, tilebuff.Bounds()).
			Draw(win, MoveAndScaleTo(tilebuff, 350, 30, 1))
		tileBufferText.Draw(win, pixel.IM.Moved(pixel.V(w-900, h-20)))
		// endregion

		debugger.Clear()
		debugTemplate.Execute(debugger, z80.GetDebugData())
		debugger.Draw(win, pixel.IM.Moved(pixel.V(w-520, h-10)))

		disasmText.Draw(win, pixel.IM.Moved(pixel.V(w-780, h-25)))

		stackText.Draw(win, pixel.IM.Moved(pixel.V(w-500, h-550)))

		if win.JustPressed(pixelgl.KeyR) {
			z80.Reset()
		}

		if win.JustPressed(pixelgl.KeyC) {
			z80.Continue()
		}

		if win.JustPressed(pixelgl.KeyS) {
			z80.Step()
		}

		if win.JustPressed(pixelgl.KeyP) {
			z80.Pause()
		}

		if win.JustPressed(pixelgl.KeyU) {
			hh = !hh
			z80.GPU.SetHighlightBG(hh)
		}

		// region SpeedHack
		if win.JustPressed(pixelgl.KeyF1) {
			z80.SetSpeedHack(8)
		}
		if win.JustPressed(pixelgl.KeyF2) {
			z80.SetSpeedHack(4)
		}
		if win.JustPressed(pixelgl.KeyF3) {
			z80.SetSpeedHack(2)
		}
		if win.JustPressed(pixelgl.KeyF4) {
			z80.SetSpeedHack(1)
		}
		if win.JustPressed(pixelgl.KeyF5) {
			z80.SetSpeedHack(1.0 / 2)
		}
		if win.JustPressed(pixelgl.KeyF6) {
			z80.SetSpeedHack(1.0 / 4)
		}
		if win.JustPressed(pixelgl.KeyF7) {
			z80.SetSpeedHack(1.0 / 256)
		}
		// endregion

		if z80.IsPaused() {
			RefreshDisasm()
			RefreshStack()
		}

		z80.Keys.Update(win)
		z80.GPU.UpdateVRAM()

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
