package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ericogr/go-display-1306-test/pkg/drawbasic"
	"github.com/fogleman/gg"

	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/host/v3"
)

func main() {
	fmt.Println("Initializing...")

	// external parameters
	i2cBus := flag.String("bus", "/dev/i2c-3", "i2c bus, for ex. /dev/i2c-3")
	displayWidth := flag.Int("with", 128, "i2c display width, for ex. 128")
	displayHeight := flag.Int("height", 32, "i2c display width, for ex. 32")
	sequential := flag.Bool("sequential", true, "Sequential corresponds to the Sequential/Alternative COM pin configuration in the OLED panel hardware")
	flag.Parse()

	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Opening display %s size %dx%d\n", *i2cBus, *displayWidth, *displayHeight)
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

	fmt.Println("Start drawing...")
	displayContext := gg.NewContext(*displayWidth, *displayHeight)

	counter := 0
	for {
		drawbasic.DrawProgressBar(displayContext, float64(*displayWidth), float64(*displayHeight), float64(counter%100+1), 100)
		drawbasic.Draw(dev, displayContext.Image())
		time.Sleep(50 * time.Millisecond)
		counter++
	}
}
