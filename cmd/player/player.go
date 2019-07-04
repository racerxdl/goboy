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
|  H   |    {{.H}} | -------- {{.HB}}  |
|  L   |    {{.L}} | -------- {{.LB}}  |
|------|-------|--------------------|
|  A   |    {{.A}} | -------- {{.AB}}  |
|  B   |    {{.B}} | -------- {{.BB}}  |
|  C   |    {{.C}} | -------- {{.CB}}  |
|  D   |    {{.D}} | -------- {{.DB}}  |
|  E   |    {{.E}} | -------- {{.EB}}  |
|------|-------|--------------------|
|      |                  ZSHC----  |
|  F   |    {{.F}} | -------- {{.FB}}  |
\-----------------------------------/

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

func MoveAndScaleTo(p pixel.Picture, x, y, s float64) pixel.Matrix {
	return pixel.IM.
		Moved(pixel.V(p.Bounds().W()/2+x/s, p.Bounds().H()/2+y/s)).
		Scaled(pixel.V(0, 0), s).
		Chained(screenOrigin)
}

func RefreshDisasm() {
	offset, page := z80.GetCurrentPage()
	dis := cpu.Disasm(int(offset), page)

	addr := z80.Registers.PC

	if len(dis) > 32 {
		s := 0
		l := 32

		o := addr - offset

		// Make sure the current instruction is displayed
		for int(o) >= l {
			s += 32
			o += 32
		}

		dis = dis[:32]
	}

	disasmText.Clear()
	disasmText.Color = colornames.Black
	fmt.Fprintf(disasmText, "Disassembler: \n\n")
	for _, d := range dis {
		disasmText.Color = colornames.Black

		if d.Address == int(z80.Registers.PC) {
			disasmText.Color = colornames.Blue
		}
		fmt.Fprintf(disasmText, "\t%04X: ", d.Address)

		disasmText.Color = colornames.Black

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
	game, err := ioutil.ReadFile("./tetris.gb")

	if err != nil {
		panic(err)
	}

	z80.Memory.LoadRom(game)

	cfg := pixelgl.WindowConfig{
		Title:  "GameBoy Emulator",
		Bounds: pixel.R(0, 0, 1024, 768),
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

	for !win.Closed() {
		vframe := z80.Memory.GetVideoFrame()
		vram := z80.GPU.GetVRAM()
		tilebuff := z80.GPU.GetTileBuffer()

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
			Draw(win, MoveAndScaleTo(tilebuff, 800, 470, 1))
		tileBufferText.Draw(win, pixel.IM.Moved(pixel.V(w-280, h-460)))
		// endregion

		debugger.Clear()
		debugTemplate.Execute(debugger, z80.GetDebugData())
		debugger.Draw(win, pixel.IM.Moved(pixel.V(w-280, h-10)))

		disasmText.Draw(win, pixel.IM.Moved(pixel.V(w-650, h-25)))

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

		if z80.IsPaused() {
			RefreshDisasm()
		}

		z80.Keys.Update(win)
		z80.GPU.UpdateVRAM()

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
