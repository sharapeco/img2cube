package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const SIZE = 64
var SPV = [SIZE]uint8{ 0, 4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 130, 134, 138, 142, 146, 150, 154, 158, 162, 166, 170, 174, 178, 182, 186, 190, 194, 198, 202, 206, 210, 215, 219, 223, 227, 231, 235, 239, 243, 247, 251, 255 }

func main() {
	const UNIT = 8
	const W = 512
	const H = 512

	// 点 (0, 0) から W × H の画像を作成する
	img := image.NewRGBA(image.Rect(0, 0, W * UNIT, H * UNIT))

	// 画素の色
	var r, g, b = 0, 0, 0

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			for iy := 0; iy < UNIT; iy++ {
				for ix := 0; ix < UNIT; ix++ {
					img.Set(UNIT*x + ix, UNIT*y + iy, color.RGBA{SPV[r], SPV[g], SPV[b], 255})
				}
			}

			// 次の色
			r++
			if r >= SIZE {
				r = 0
				g++
				if g >= SIZE {
					g = 0
					b++
					if b >= SIZE {
						b = 0
					}
				}
			}
		}
	}

	// lutsrc.png に保存する
	f, _ := os.OpenFile("src.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
