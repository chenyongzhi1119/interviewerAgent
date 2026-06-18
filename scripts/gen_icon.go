//go:build ignore

// 生成 512×512 应用图标 PNG（蓝色圆形 + 白色大写 I）
// 用法：go run scripts/gen_icon.go assets/app_icon_512.png
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	out := "assets/app_icon_512.png"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}

	const size = 512
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	blue := color.RGBA{0, 82, 217, 255}
	white := color.RGBA{255, 255, 255, 255}
	transp := color.RGBA{0, 0, 0, 0}

	cx := float64(size) / 2
	r := cx - 4

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dx := float64(x) + 0.5 - cx
			dy := float64(y) + 0.5 - cx
			if dx*dx+dy*dy > r*r {
				img.Set(x, y, transp)
			} else {
				img.Set(x, y, blue)
			}
		}
	}

	// 白色大写 I（衬线风格，三段矩形）
	fill := func(x0, y0, x1, y1 int) {
		for y := y0; y < y1; y++ {
			for x := x0; x < x1; x++ {
				img.Set(x, y, white)
			}
		}
	}
	fill(136, 90, 376, 170)  // 顶横
	fill(216, 170, 296, 342) // 竖
	fill(136, 342, 376, 422) // 底横

	os.MkdirAll("assets", 0o755)
	f, _ := os.Create(out)
	defer f.Close()
	png.Encode(f, img)
}
