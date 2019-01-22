package render

import (
	"github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window
var surface *sdl.Surface
var xScale int32
var yScale int32

// Init the sdl
func Init(x int32, y int32) {
	xScale = x
	yScale = y
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	win, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 64*xScale, 32*yScale, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	window = win
	surf, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface = surf
}

// Draw refreshes the screen to the supplied pixelMap
func Draw(pixelMap [][]bool) {
	background := uint32(0xFF220022)
	foreground := uint32(0xFF55FF55)
	surface.FillRect(nil, background)
	rect := sdl.Rect{X: 0, Y: 0, W: xScale, H: yScale}
	for i := 0; i < len(pixelMap); i++ {
		for j := 0; j < len(pixelMap[i]); j++ {
			if pixelMap[i][j] {
				rect.X = int32(i) * xScale
				rect.Y = int32(j) * yScale
				if pixelMap[i][j] {
					surface.FillRect(&rect, foreground)
				}
			}
		}
	}
	window.UpdateSurface()
}

// End clean ups
func End() {
	defer window.Destroy()
	defer sdl.Quit()
}
