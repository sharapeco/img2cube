package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

const SIZE = 64

var SPV = [SIZE]uint8{0, 4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 130, 134, 138, 142, 146, 150, 154, 158, 162, 166, 170, 174, 178, 182, 186, 190, 194, 198, 202, 206, 210, 215, 219, 223, 227, 231, 235, 239, 243, 247, 251, 255}

func main() {
	const UNIT = 8
	const W = 512
	const H = 512

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: go run img2cube.go input.png")
		os.Exit(1)
	}
	inputFilePath := args[0]

	baseName := getBaseName(inputFilePath)
	outputName := baseName + ".cube"

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Could not open file.\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	// 画像を image 型として読み込む
	inputImage, format, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("error:decode\n", format, err)
		os.Exit(1)
	}

	// 画像サイズをチェック
	rect := inputImage.Bounds()
	if rect.Min.X != 0 || rect.Min.Y != 0 || rect.Max.X != UNIT*W || rect.Max.Y != UNIT*H {
		fmt.Println("Invalid image size")
		os.Exit(1)
	}

	// 出力ファイル
	outputFile, err := os.OpenFile(outputName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		fmt.Println("error:write\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// ヘッダ
	fmt.Fprintln(outputFile, `# Created by mklut64.go`)
	fmt.Fprintln(outputFile, `TITLE "`+baseName+`"`)
	fmt.Fprintf(outputFile, "LUT_3D_SIZE %d\n", SIZE)
	fmt.Fprintln(outputFile, `DOMAIN_MIN 0.0 0.0 0.0`)
	fmt.Fprintln(outputFile, `DOMAIN_MAX 1.0 1.0 1.0`)
	fmt.Fprintln(outputFile, ``)

	// 画素の色を調べて LUT テーブルを作成する
	var amp float64 = 1.0 / 255 / UNIT / UNIT
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			// 8 × 8 の画素の平均値を取る
			var r, g, b float64 = 0, 0, 0
			for iy := 0; iy < UNIT; iy++ {
				for ix := 0; ix < UNIT; ix++ {
					c := color.RGBAModel.Convert(inputImage.At(UNIT*x+ix, UNIT*y+iy)).(color.RGBA)
					r += float64(c.R)
					g += float64(c.G)
					b += float64(c.B)
				}
			}
			fmt.Fprintf(outputFile, "%.6f %.6f %.6f\n", amp*r, amp*g, amp*b)
		}
	}

	fmt.Println("Maybe OK", outputName)
}

// ファイル名から拡張子を除いた名前を返す（ディレクトリはそのまま）
func getBaseName(path string) string {
	return path[:len(path)-len(filepath.Ext(path))]
}
