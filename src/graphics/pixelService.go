package graphics

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

const (
	ColorRed = iota
	ColorGreen
	ColorBlue
	ColorBlack
	ColorWhite
	ColorCyan
	ColorMagenta
	ColorYellow
	ColorPink
	colorCount
)

// PixelService service for working with ebiten pixels
type PixelService struct {
	Screen       *ebiten.Image
	Pixels       []byte
	ScreenHeight int
	ScreenWidth  int
}

// New pixel service constructor
func New(screenHeight int, screenWidth int) (service *PixelService) {
	service = &PixelService{
		ScreenHeight: screenHeight,
		ScreenWidth:  screenWidth,
		Pixels:       make([]byte, screenHeight*screenWidth*4),
	}

	return
}

// FindRandomPixel returns random start pixel of the screen
func (service *PixelService) FindRandomPixel() int {
	x := rand.Intn(service.ScreenHeight)
	y := rand.Intn(service.ScreenWidth)

	return service.FindPixelNumberByXY(x, y)
}

// FindPixelNumberByXY returns start pixel of the given coordinates
func (service *PixelService) FindPixelNumberByXY(x int, y int) int {
	return (y*service.ScreenHeight + x) * 4
}

// ClearScreen clears entire screen
func (service *PixelService) ClearScreen() {
	service.Pixels = make([]byte, service.ScreenHeight*service.ScreenWidth*4)
}

// DrawRandomPixel draws pixel selected randomly
func (service *PixelService) DrawRandomPixel() {
	randomPixelStartPoint := service.FindRandomPixel()
	randomColor := service.getRandomColor()
	colorBytes := service.getColorBytes(randomColor)

	service.Pixels[randomPixelStartPoint] = colorBytes[0]
	service.Pixels[randomPixelStartPoint+1] = colorBytes[1]
	service.Pixels[randomPixelStartPoint+2] = colorBytes[2]
	service.Pixels[randomPixelStartPoint+3] = colorBytes[3]
}

// RefreshScreen must be called to redraw screen with nex pixels
func (service *PixelService) RefreshScreen() {
	service.Screen.ReplacePixels(service.Pixels)
}

func (service *PixelService) getColorBytes(color int) []byte {
	var red, green, blue byte
	var colorBytes = make([]byte, 4)

	switch color {
	case ColorBlack:
		red = 0x00
		green = 0x00
		blue = 0x00
	case ColorRed:
		red = 0xff
		green = 0x00
		blue = 0x00
	case ColorGreen:
		red = 0x00
		green = 0xff
		blue = 0x00
	case ColorBlue:
		red = 0x00
		green = 0x00
		blue = 0xff
	case ColorYellow:
		red = 0xff
		green = 0xff
		blue = 0x00
	case ColorMagenta:
		red = 0xff
		green = 0x00
		blue = 0xff
	case ColorCyan:
		red = 0x00
		green = 0xff
		blue = 0xff
	case ColorWhite:
		red = 0xff
		green = 0xff
		blue = 0xff
	case ColorPink:
		red = 0xe6
		green = 0x34
		blue = 0x62
	}

	colorBytes[0] = red
	colorBytes[1] = green
	colorBytes[2] = blue
	colorBytes[3] = 0x00 + byte(rand.Intn(255))

	return colorBytes
}

func (service *PixelService) getRandomColor() int {
	return rand.Intn(colorCount)
}
