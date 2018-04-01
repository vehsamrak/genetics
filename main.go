package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/vehsamrak/genetics/src/graphics"
)

const screenEnlargeFactor = 5
const screenHeight = 7 * screenEnlargeFactor
const screenWidth = 4 * screenEnlargeFactor
const speed = 100

var worldCounter int
var pixelService *graphics.PixelService

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	pixelService = graphics.New(screenHeight, screenWidth)

	ebiten.SetFullscreen(true)

	if err := ebiten.Run(update, screenHeight, screenWidth, 1, "Bacterium simulation"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	pixelService.Screen = screen
	handleKeyboardInput()

	if ebiten.IsRunningSlowly() {
		return nil
	}

	if worldCounter > 120-speed {
		worldCounter = 0
	} else {
		worldCounter++
		pixelService.RefreshScreen()

		return nil
	}

	//pixelService.ClearScreen()
	pixelService.DrawRandomPixel()
	pixelService.RefreshScreen()

	return nil
}

func handleKeyboardInput() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}
