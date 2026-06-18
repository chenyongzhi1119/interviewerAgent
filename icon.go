package main

import "encoding/base64"

// trayIconPNG 是菜单栏图标（22×22 蓝色圆形，编译时内嵌）
func trayIconPNG() []byte {
	data, _ := base64.StdEncoding.DecodeString(`iVBORw0KGgoAAAANSUhEUgAAABYAAAAWCAYAAADEtGw7AAAAR0lEQVR42mNgGBQg6OZ/6hlECFPdQJItIMdQgoZTYihew2liMDUMxWr4qMFUN5Smho9GHm0zCX3LCpqVbjQtj2lag1ChzgMAQwSuIUS8cDQAAAAASUVORK5CYII=`)
	return data
}
