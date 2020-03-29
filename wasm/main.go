// +build js

package main

import (
	"bytes"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
	"github.com/racerxdl/goboy/cpu"
	"html/template"
	"image/color"
	"syscall/js"
)

var z80 = cpu.MakeCore()

var cvs *canvas.Canvas2d
var done chan struct{}

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

var debugWindow js.Value

func updateDebugger() {
	buff := &bytes.Buffer{}
	debugTemplate.Execute(buff, z80.GetDebugData())
	data := buff.String()

	debugWindow.Set("innerHTML", data)
}

func main() {
	cvs, _ = canvas.NewCanvas2d(false)
	//cvs.Create(int(js.Global().Get("innerWidth").Float()*0.9), int(js.Global().Get("innerHeight").Float()*0.9)) // Make Canvas 90% of window size.  For testing rendering canvas smaller than full windows

	screenCanvas := js.Global().Get("document").Call("getElementById", "screen")
	debugWindow = js.Global().Get("document").Call("getElementById", "debug")

	cvs.Set(screenCanvas, 180, 164)

	cvs.Start(30, Render)

	rom, _ := Asset("assets/tetris.gb")

	z80.Memory.LoadRom(rom)

	z80.Start()
	updateDebugger()
	z80.Continue()

	<-done
	z80.Stop()
}
func Render(gc *draw2dimg.GraphicContext) bool {
	vframe := z80.Memory.GetVideoFrame()
	gc.SetFillColor(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
	gc.Clear()
	gc.Save()
	gc.Translate(10, 10)
	gc.DrawImage(vframe.Image())
	gc.Restore()
	z80.GPU.UpdateVRAM()
	updateDebugger()

	return true
}
