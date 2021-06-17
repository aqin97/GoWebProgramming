package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"os"
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
	ratio := bound.Dx() / newWidth
	out := image.NewNRGBA(image.Rect(bound.Min.X/ratio, bound.Min.Y/ratio, bound.Max.X/ratio, bound.Max.Y/ratio))
	for y, j := bound.Min.Y, bound.Min.Y; y < bound.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bound.Min.X, bound.Min.X; x < bound.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return *out
}

//从tiles文件夹中扫描图片，创建瓷砖图片数据库
func tileDB() map[string][3]float64 {
	fmt.Println("start populating tiles db ...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles")
	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		if err != nil {
			fmt.Println("connot open file", name, err)
		} else {
			img, _, err := image.Decode(file)
			if err != nil {
				fmt.Println("error in populating tiledb", err, name)
			} else {
				db[name] = averageColor(img)
			}
		}
		file.Close()
	}
	fmt.Println("finished populating tiles db")
	return db
}

//匹配和瓷砖最近的平均颜色
func nearest(target [3]float64, db *map[string][3]float64) string {
	var filename string
	smallest := 1000000.0
	for k, v := range *db {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(*db, filename)
	return filename
}

//求两点之间的欧几里得距离
func distance(p1, p2 [3]float64) float64 {
	return math.Sqrt(sq(p1[0]-p2[0]) + sq(p1[1]-p2[1]) + sq(p1[2]-p2[2]))
}

func sq(n float64) float64 {
	return n * n
}

var TILESDB map[string][3]float64

func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	return db
}
