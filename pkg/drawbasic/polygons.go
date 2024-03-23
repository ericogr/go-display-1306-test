package drawbasic

import "periph.io/x/devices/v3/ssd1306/image1bit"

func DrawBorder(img *image1bit.VerticalLSB) {
	x1 := 0
	y1 := 0
	x2 := 127
	y2 := 31

	for x := x1; x <= x2; x++ {
		img.Set(x, y1, image1bit.On)
		img.Set(x, y2, image1bit.On)
	}

	for y := y1; y <= y2; y++ {
		img.Set(x1, y, image1bit.On)
		img.Set(x2, y, image1bit.On)
	}
}
