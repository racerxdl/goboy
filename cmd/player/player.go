package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 160, 120))

	pic := pixel.PictureDataFromImage(img)

	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			if y == x {
				pic.Pix[y*pic.Stride+x] = colornames.Red
			} else {
				pic.Pix[y*pic.Stride+x] = colornames.Black
			}
		}
	}

	t := 0
	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		pic.Pix[t] = colornames.Red
		pixel.NewSprite(pic, pic.Bounds()).Draw(win, pixel.IM.Moved(pixel.V(100, 100)))

		win.Update()

		if t >= 160*120 {
			t = 0
		} else {
			t = t + 1
		}

	}
}

func main() {
	pixelgl.Run(run)
}
