package main

import (
	"fmt"
	"image"
	"log"
	"time"

	"image/color"

	"golang.org/x/image/draw"
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

	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Use i2creg I²C bus registry to find the first available I²C bus.
	b, err := i2creg.Open("/dev/i2c-3")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	var miniDisplay = ssd1306.Opts{
		W:             128,
		H:             32,
		Rotated:       false,
		Sequential:    true,
		SwapTopBottom: false,
	}

	dev, err := ssd1306.NewI2C(b, &miniDisplay)
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
		Dot:  fixed.P(0, 9),
	}

	var contador = 0
	for {
		draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)
		drawer.DrawString(fmt.Sprintf("contador %d", contador))
		drawer.Dot.X = 0
		if err := dev.Draw(dev.Bounds(), img, image.Point{}); err != nil {
			log.Fatal(err)
		}

		time.Sleep(25 * time.Millisecond)
		contador++
	}

}
