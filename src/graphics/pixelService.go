package graphics

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

// PixelService service for working with ebiten pixels
type PixelService struct {
	Screen     *ebiten.Image
	Pixels     []byte
	ScreenSide int
}

// New pixel service constructor
func New(screenSide int) (service *PixelService) {
	service = &PixelService{
		ScreenSide: screenSide,
		Pixels:     make([]byte, screenSide*screenSide*4),
	}

	return
}

// FindRandomPixel returns random start pixel of the screen
func (service *PixelService) FindRandomPixel() int {
	x := rand.Intn(service.ScreenSide)
	y := rand.Intn(service.ScreenSide)

	return service.FindPixelNumberByXY(x, y)
}

// FindPixelNumberByXY returns start pixel of the given coordinates
func (service *PixelService) FindPixelNumberByXY(x int, y int) int {
	return (y*service.ScreenSide + x) * 4
}

// ClearScreen clears entire screen
func (service *PixelService) ClearScreen() {
	service.Pixels = make([]byte, service.ScreenSide*service.ScreenSide*4)
}

// DrawRandomPixel draws pixel selected randomly
func (service *PixelService) DrawRandomPixel() {
	randomPixelStartPoint := service.FindRandomPixel()

	service.Pixels[randomPixelStartPoint] = 0xff
	service.Pixels[randomPixelStartPoint+1] = 0xff
	service.Pixels[randomPixelStartPoint+2] = 0xff
	service.Pixels[randomPixelStartPoint+3] = 0xff
}

// RefreshScreen must be called to redraw screen with nex pixels
func (service *PixelService) RefreshScreen() {
	service.Screen.ReplacePixels(service.Pixels)
}
