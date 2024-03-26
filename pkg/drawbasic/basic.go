package drawbasic

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"periph.io/x/devices/v3/ssd1306"
)

func Draw(dev *ssd1306.Dev, img image.Image) {
	if err := dev.Draw(dev.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}
}

func Clear(displayContext *gg.Context) {
	displayContext.SetColor(color.Black)
	displayContext.Clear()
}

func DrawProgressBar(displayContext *gg.Context, width, height, value, max float64) {
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
	displayContext.Fill()
	displayContext.SetLineWidth(1)
	displayContext.DrawRectangle(3, 18, 122, 10)
	displayContext.Stroke()
}

func DrawResources(displayContext *gg.Context, sleep time.Duration, width, height float64) {
	before, err := cpu.Get()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(sleep)
	after, err := cpu.Get()
	if err != nil {
		log.Fatal(err)
	}
	total := float64(after.Total - before.Total)
	userCpu := float64(after.User-before.User) / total * 100
	sysCpu := float64(after.System-before.System) / total * 100

	displayContext.SetColor(color.White)
	displayContext.DrawRoundedRectangle(0, 0, float64(width), float64(height), 5)
	displayContext.Stroke()
	displayContext.DrawString(
		fmt.Sprintf("cpu: %.1f%%\n", userCpu+sysCpu),
		3,
		12,
	)

	memory, err := memory.Get()
	if err != nil {
		log.Fatal(err)
	}
	displayContext.DrawString(
		fmt.Sprintf("mem free: %d Mb", (memory.Free/1024/1024)),
		3,
		24,
	)
}
