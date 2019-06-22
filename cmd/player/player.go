package main

import (
    "fmt"
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "github.com/faiface/pixel/text"
    "github.com/racerxdl/goboy/cpu"
    "golang.org/x/image/colornames"
    "golang.org/x/image/font/basicfont"
    "html/template"
    "io/ioutil"
)

var z80 = cpu.MakeCore()

var debugTemplate = template.Must(template.New("debugger").Parse(`
CPU DATA
    PC: {{.PC}} SP {{.SPX}} 
    HL {{.HL}}  H: {{.C}} L: {{.L}}
    
    A: {{.A}} B: {{.B}} C: {{.C}} D: {{.D}}
    E: {{.A}} E: {{.B}}

    F: {{.F}}

GPU DATA
    Scroll X:       {{.GPUSCROLLX}}
    Scroll Y:       {{.GPUSCROLLY}}
    Window X:       {{.GPUWINX}}
    Window Y:       {{.GPUWINY}}
    Mode Clocks:    {{.GPUMODECLOCKS}}
    Line:           {{.GPULINE}}
`))

var screenOrigin pixel.Matrix

func MoveAndScaleTo(p *pixel.PictureData, x, y, s float64) pixel.Matrix {
    return pixel.IM.
        Moved(pixel.V(p.Bounds().W() / 2 + x / s, p.Bounds().H() / 2 + y / s)).
        Scaled(pixel.V(0,0), s).
        Chained(screenOrigin)
}

func run() {

    game, err := ioutil.ReadFile("./opus5.gb")

    if err != nil {
        panic(err)
    }

    z80.Memory.LoadRom(game)

    cfg := pixelgl.WindowConfig{
        Title:  "Pixel Rocks!",
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

    r := win.Bounds()
    w := r.Max.X
    h := r.Max.Y

    for !win.Closed() {
        vframe := z80.Memory.GetVideoFrame()
        vram := z80.GPU.GetVRAM()

        win.Clear(colornames.Skyblue)

        pixel.NewSprite(vframe, vframe.Bounds()).
            Draw(win, MoveAndScaleTo(vframe,10,10, 2))

        pixel.NewSprite(vram, vram.Bounds()).
           Draw(win, MoveAndScaleTo(vram, 10, 310, 1))

        debugger.Clear()
        debugTemplate.Execute(debugger, z80.GetDebugData())
        debugger.Draw(win, pixel.IM.Moved(pixel.V(w - 200, h - 50)))

        if win.JustPressed(pixelgl.KeyZ) {
            z80.Reset()
        }

        if win.JustPressed(pixelgl.KeyC) {
            z80.Continue()
        }

        if win.JustPressed(pixelgl.KeyP) {
            z80.Pause()
        }

        win.Update()
    }
}

func main() {
    pixelgl.Run(run)
}
