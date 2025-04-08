# img2cube

画像から Cube 形式の LUT ファイルを生成するツール。

## 必要なもの

- [Go](https://golang.org/) 1.18 以上

## 使い方

1. `go run makesrc.go` を実行し、 `src.png` を生成する
2. `src.png` にカラーグレーディングを行い、 `output.png` に保存する
3. `go run img2cube.go output.png` を実行し、 `output.cube` を生成する

### Affinity Photo の場合

調整レイヤー「LUT」を作成し、LUT ファイルを読み込む

- [LUT調整 - Affinity Designerヘルプ](https://affinity.help/designer/ja.lproj/index.html?page=pages/Adjustments/adjustment_3dLut.html)

### Adobe Photoshop の場合

調整レイヤー「カラールックアップ」を作成し、LUT ファイルを読み込む

## 参考

- [Cube LUT Specification (Adobe)](https://wwwimages2.adobe.com/content/dam/acom/en/products/speedgrade/cc/pdfs/cube-lut-specification-1.0.pdf)
	- Photoshop でよく使われる *.cube ファイル
	- 一般的な Cube LUT ファイルは RGB 各 32 段階に分割し、値域は [0, 1]
	- 分割数は 2 から 256 までの任意の値を取れるが、256 ではメモリを大量に消費するため 32 が一般的
	- 32 分割で *.cube ファイルは 880 KB 程度になる（有効桁数 6 桁）
	- 32 分割の場合、各チャンネルは (step * 255 + 16) / 31 に対応する値を記述すればよい
