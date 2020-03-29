package localimg

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

type PictureData struct {
	Pix    []color.RGBA
	Stride int
	Rect   Rect
}

// MakePictureData creates a zero-initialized PictureData covering the given rectangle.
func MakePictureData(rect Rect) *PictureData {
	w := int(math.Ceil(rect.Max.X)) - int(math.Floor(rect.Min.X))
	h := int(math.Ceil(rect.Max.Y)) - int(math.Floor(rect.Min.Y))
	pd := &PictureData{
		Stride: w,
		Rect:   rect,
	}
	pd.Pix = make([]color.RGBA, w*h)
	return pd
}

func verticalFlip(rgba *image.RGBA) {
	bounds := rgba.Bounds()
	width := bounds.Dx()

	tmpRow := make([]uint8, width*4)
	for i, j := 0, bounds.Dy()-1; i < j; i, j = i+1, j-1 {
		iRow := rgba.Pix[i*rgba.Stride : i*rgba.Stride+width*4]
		jRow := rgba.Pix[j*rgba.Stride : j*rgba.Stride+width*4]

		copy(tmpRow, iRow)
		copy(iRow, jRow)
		copy(jRow, tmpRow)
	}
}

// PictureDataFromImage converts an image.Image into PictureData.
//
// The resulting PictureData's Bounds will be the equivalent of the supplied image.Image's Bounds.
func PictureDataFromImage(img image.Image) *PictureData {
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)

	verticalFlip(rgba)

	pd := MakePictureData(R(
		float64(rgba.Bounds().Min.X),
		float64(rgba.Bounds().Min.Y),
		float64(rgba.Bounds().Max.X),
		float64(rgba.Bounds().Max.Y),
	))

	for i := range pd.Pix {
		pd.Pix[i].R = rgba.Pix[i*4+0]
		pd.Pix[i].G = rgba.Pix[i*4+1]
		pd.Pix[i].B = rgba.Pix[i*4+2]
		pd.Pix[i].A = rgba.Pix[i*4+3]
	}

	return pd
}

// Image converts PictureData into an image.RGBA.
//
// The resulting image.RGBA's Bounds will be equivalent of the PictureData's Bounds.
func (pd *PictureData) Image() *image.RGBA {
	bounds := image.Rect(
		int(math.Floor(pd.Rect.Min.X)),
		int(math.Floor(pd.Rect.Min.Y)),
		int(math.Ceil(pd.Rect.Max.X)),
		int(math.Ceil(pd.Rect.Max.Y)),
	)
	rgba := image.NewRGBA(bounds)

	i := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			off := pd.Index(V(float64(x), float64(y)))
			rgba.Pix[i*4+0] = pd.Pix[off].R
			rgba.Pix[i*4+1] = pd.Pix[off].G
			rgba.Pix[i*4+2] = pd.Pix[off].B
			rgba.Pix[i*4+3] = pd.Pix[off].A
			i++
		}
	}

	verticalFlip(rgba)

	return rgba
}

// Index returns the index of the pixel at the specified position inside the Pix slice.
func (pd *PictureData) Index(at Vec) int {
	at = at.Sub(pd.Rect.Min.Map(math.Floor))
	x, y := int(at.X), int(at.Y)
	return y*pd.Stride + x
}

// Bounds returns the bounds of this PictureData.
func (pd *PictureData) Bounds() Rect {
	return pd.Rect
}

// Color returns the color located at the given position.
//func (pd *PictureData) Color(at Vec) RGBA {
//	if !pd.Rect.Contains(at) {
//		return RGBA{0, 0, 0, 0}
//	}
//	return ToRGBA(pd.Pix[pd.Index(at)])
//}
