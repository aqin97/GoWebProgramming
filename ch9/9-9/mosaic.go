package main

import (
	"image"
)

//求取平均颜色的函数
func averageColor(img image.Image) [3]float64 {
	bound := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bound.Max.X * bound.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

//缩放到指定大小
func resize(in image.Image, newWidth int) image.NRGBA {
	bound := in.Bounds()
	ratio := bound.Dx()/newWidth
	out := image.NewNRGBA(image.Rect(bound.Min.X/ratio, bound.Min.Y/ratio, bound.Max.X/ratio, bound.Max.Y/ratio))
	for y, j := bound.Min.Y
}