package main

import (
	"image"
	"io"

	_ "image/jpeg"
	_ "image/png"
)

// WhiteCounter подсчитывает количество пикселей белого цвета.
func WhiteCounter(r io.Reader) (int, error) {
	// image.Decode автоматически определяет формат изображения,
	// если был импортирован соответствующий пакет
	img, _, err := image.Decode(r)
	if err != nil {
		return 0, err
	}
	bounds := img.Bounds()
	var whiteCount int
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// в каждом канале применяем побитовый сдвиг, выделяем значимые байты
			// и сравниваем их с нужным значением 0xff
			if r>>8 == 0xff && g>>8 == 0xff && b>>8 == 0xff {
				whiteCount++
			}
		}
	}
	return whiteCount, nil
}
