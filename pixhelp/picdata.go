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

func DrawSquare(p *pixel.PictureData, rect pixel.Rect, c color.Color) {
	nc := ToRGBA(c)
	b := p.Bounds()

	l := b.H()

	sx := rect.Min.X
	sy := rect.Min.Y
	mx := rect.Max.X
	my := rect.Max.Y

	if sx < b.Min.X {
		sx = b.Min.X
	}

	if sy < b.Min.Y {
		sy = b.Min.Y
	}

	if mx > b.Max.X {
		mx = b.Max.X
	}

	if my > b.Max.Y {
		my = b.Max.Y
	}

	for x := sx; x < mx; x++ {
		for y := sy; y < my; y++ {
			idx := int(y*l + x)
			p.Pix[idx] = nc
		}
	}
}
