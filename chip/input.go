package chip

import (
	"github.com/veandco/go-sdl2/sdl"
)

var keyState = [16]bool{}
var keyCodes = [16]sdl.Scancode{
	sdl.SCANCODE_0, sdl.SCANCODE_1,
	sdl.SCANCODE_2, sdl.SCANCODE_3,
	sdl.SCANCODE_4, sdl.SCANCODE_5,
	sdl.SCANCODE_6, sdl.SCANCODE_7,
	sdl.SCANCODE_8, sdl.SCANCODE_9,
	sdl.SCANCODE_A, sdl.SCANCODE_B,
	sdl.SCANCODE_C, sdl.SCANCODE_D,
	sdl.SCANCODE_E, sdl.SCANCODE_F,
}

// PollEvents polls the event list and return application running status
func PollEvents() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			running = false
			break
		case *sdl.KeyboardEvent:
			for i := 0; i < len(keyCodes); i++ {
				if e.Keysym.Scancode == keyCodes[i] {
					if e.State == 0 {
						keyState[i] = false
					} else {
						keyState[i] = true
					}
				}
			}
			break
		}
	}
	return running
}
