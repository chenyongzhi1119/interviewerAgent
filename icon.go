package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

// trayIconPNG 生成菜单栏图标（22×22，蓝色圆形背景 + 白色大写 I）
func trayIconPNG() []byte {
	return renderIcon(22,
		6, 4, 15, 7,   // 顶横：x0,y0,x1,y1
		9, 7, 13, 15,  // 竖
		6, 15, 15, 18, // 底横
	)
}

// appIconPNG 生成 512×512 Finder 图标
func appIconPNG() []byte {
	return renderIcon(512,
		136, 90, 376, 170,   // 顶横
		216, 170, 296, 342,  // 竖
		136, 342, 376, 422,  // 底横
	)
}

// renderIcon 在 size×size 的蓝色圆形上画白色大写 I。
func renderIcon(size, tx0, ty0, tx1, ty1, vx0, vy0, vx1, vy1, bx0, by0, bx1, by1 int) []byte {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	blue := color.RGBA{0, 82, 217, 255}
	white := color.RGBA{255, 255, 255, 255}
	transp := color.RGBA{0, 0, 0, 0}

	cx := float64(size) / 2
	r := cx - 0.5

	// 蓝色圆形背景
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

	// 白色大写 I（三段矩形）
	fill(img, tx0, ty0, tx1, ty1, white) // 顶横
	fill(img, vx0, vy0, vx1, vy1, white) // 竖
	fill(img, bx0, by0, bx1, by1, white) // 底横

	var buf bytes.Buffer
	png.Encode(&buf, img)
	return buf.Bytes()
}

func fill(img *image.RGBA, x0, y0, x1, y1 int, c color.RGBA) {
	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			img.Set(x, y, c)
		}
	}
}
