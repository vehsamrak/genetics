package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const screenWidth = 10
const screenHeight = 10
const speed = 1

var worldCounter int
var pixels = make([]byte, screenWidth*screenHeight*4)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 50, "Bacterium simulation"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if worldCounter > 10-speed {
		worldCounter = 0
	} else {
		worldCounter++
		screen.ReplacePixels(pixels)

		return nil
	}

	pixels = make([]byte, screenWidth*screenHeight*4)

	for x := range pixels {
		// only 4-th pixels are visible, while others are RGB colours
		if x%4 != 0 {
			continue
		}

		if rand.Intn(10) == 0 {
			pixels[x] = 0xff
			pixels[x+1] = 0xff
			pixels[x+2] = 0xff
			pixels[x+3] = 0xff
		}
	}

	screen.ReplacePixels(pixels)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	return nil
}
