package util

import (
	"image"
	"image/color"
)

func GetImageColors(img image.Image) map[color.Color]struct{} {
	colors := make(map[color.Color]struct{})
	var empty struct{}

	for i := 0; i < img.Bounds().Dy(); i++ {
		for j := 0; j < img.Bounds().Dx(); j++ {
			colors[img.At(j, i)] = empty
		}
	}

	return colors
}
