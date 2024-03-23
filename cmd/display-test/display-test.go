package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"time"

	"image/color"
	"image/draw"

	"github.com/ericogr/go-display-1306-test/pkg/drawbasic"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"

	"golang.org/x/image/math/fixed"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/devices/v3/ssd1306/image1bit"
	"periph.io/x/host/v3"
)

func main() {
	fmt.Println("Starting...")

	// external parameters
	i2cBus := flag.String("bus", "/dev/i2c-3", "i2c bus, for ex. /dev/i2c-3")
	displayWidth := flag.Int("with", 128, "i2c display width, for ex. 128")
	displayHeight := flag.Int("height", 32, "i2c display width, for ex. 32")
	sequential := flag.Bool("sequential", true, "Sequential corresponds to the Sequential/Alternative COM pin configuration in the OLED panel hardware")

	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	busCloser, err := i2creg.Open(*i2cBus)
	if err != nil {
		log.Fatal(err)
	}
	defer busCloser.Close()

	var miniDisplay = ssd1306.Opts{
		W:             *displayWidth,
		H:             *displayHeight,
		Rotated:       false,
		Sequential:    *sequential,
		SwapTopBottom: false,
	}

	dev, err := ssd1306.NewI2C(busCloser, &miniDisplay)
	if err != nil {
		log.Fatalf("failed to initialize ssd1306: %v", err)
	}

	fmt.Println("Drawing...")

	img := image1bit.NewVerticalLSB(dev.Bounds())
	f := basicfont.Face7x13
	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{image1bit.On},
		Face: f,
		Dot:  fixed.P(0, 0),
	}

	var contador = 0
	for {
		clean(img)
		drawbasic.DrawBorder(img)
		drawer.Dot = fixed.Point26_6{
			// X: (fixed.I(128) - drawer.MeasureString("Contagem: 111")) / 2,
			X: fixed.I(5),
			Y: fixed.I(20),
		}
		drawer.DrawString(fmt.Sprintf("Counter: %d", contador))
		drawer.Dot.X = 0
		if err := dev.Draw(dev.Bounds(), img, image.Point{}); err != nil {
			log.Fatal(err)
		}

		time.Sleep(25 * time.Millisecond)
		contador++
	}

}

func clean(img *image1bit.VerticalLSB) {
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)
}
