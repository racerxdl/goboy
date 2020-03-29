package conv

import (
	"github.com/faiface/pixel"
	"github.com/racerxdl/goboy/localimg"
)

func LocalImgToPixelPicture(data *localimg.PictureData) *pixel.PictureData {
	return &pixel.PictureData{
		Pix:    data.Pix,
		Stride: data.Stride,
		Rect:   pixel.R(data.Rect.Min.X, data.Rect.Min.Y, data.Rect.Max.X, data.Rect.Max.Y),
	}
}

func GetSprite(data *localimg.PictureData) *pixel.Sprite {
	pic := LocalImgToPixelPicture(data)
	return pixel.NewSprite(pic, pic.Bounds())
}
