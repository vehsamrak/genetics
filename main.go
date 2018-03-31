package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/vehsamrak/genetics/src/graphics"
)

const screenSide = 10
const speed = 1

var worldCounter int
var pixelService *graphics.PixelService

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	pixelService = graphics.New(screenSide)

	if err := ebiten.Run(update, screenSide, screenSide, 50, "Bacterium simulation"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	pixelService.Screen = screen
	handleKeyboardInput()

	if ebiten.IsRunningSlowly() {
		return nil
	}

	if worldCounter > 10-speed {
		worldCounter = 0
	} else {
		worldCounter++
		pixelService.RefreshScreen()

		return nil
	}

	pixelService.ClearScreen()
	pixelService.DrawRandomPixel()
	pixelService.RefreshScreen()

	return nil
}

func handleKeyboardInput() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}
