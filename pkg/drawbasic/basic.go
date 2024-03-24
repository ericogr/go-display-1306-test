package drawbasic

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/fogleman/gg"
	"periph.io/x/devices/v3/ssd1306"
)

func Draw(dev *ssd1306.Dev, img image.Image) {
	if err := dev.Draw(dev.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}
}

func DrawProgressBar(displayContext *gg.Context, width, height, value, max float64) {
	displayContext.SetColor(color.Black)
	displayContext.Clear()
	displayContext.SetColor(color.White)
	displayContext.DrawRoundedRectangle(0, 0, float64(width), float64(height), 5)
	displayContext.Stroke()
	displayContext.DrawString(
		fmt.Sprintf("Progress: %d%%", int(value)),
		3,
		11,
	)
	w := (width - 5) / 100 * (max / 100 * value)
	displayContext.DrawRectangle(3, 18, w, 10)
	displayContext.SetLineWidth(1)
	displayContext.Fill()
}
