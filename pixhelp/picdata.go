package pixhelp

import (
	"github.com/faiface/pixel"
	"image/color"
)

func ClearPictureData(p *pixel.PictureData, c color.Color) {
	t := p.Bounds()
	t.Size()
	ts := t.Min
	te := t.Max
	nc := ToRGBA(c)
	for x := ts.X; x < te.X; x++ {
		for y := ts.Y; y < te.Y; y++ {
			p.Pix[int(x)+int(y)*p.Stride] = nc
		}
	}
}

func ToRGBA(c color.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}
