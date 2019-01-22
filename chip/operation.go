package chip

import (
	"fmt"

	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

// Masks the opCode and matches pattern to identify the operation to apply
func startOperation() {
	if opCode&0xFFFF == 0x00E0 {
		operation00E0()
	} else if opCode&0xFFFF == 0x00EE {
		operation00EE()
	} else if opCode&0xF000 == 0x1000 {
		operation1NNN()
	} else if opCode&0xF000 == 0x2000 {
		operation2NNN()
	} else if opCode&0xF000 == 0x3000 {
		operation3XKK()
	} else if opCode&0xF000 == 0x4000 {
		operation4XKK()
	} else if opCode&0xF000 == 0x5000 {
		operation5XY0()
	} else if opCode&0xF000 == 0x6000 {
		operation6XKK()
	} else if opCode&0xF000 == 0x7000 {
		operation7XKK()
	} else if opCode&0xF00F == 0x8000 {
		operation8XY0()
	} else if opCode&0xF00F == 0x8001 {
		operation8XY1()
	} else if opCode&0xF00F == 0x8002 {
		operation8XY2()
	} else if opCode&0xF00F == 0x8003 {
		operation8XY3()
	} else if opCode&0xF00F == 0x8004 {
		operation8XY4()
	} else if opCode&0xF00F == 0x8005 {
		operation8XY5()
	} else if opCode&0xF00F == 0x8006 {
		operation8XY6()
	} else if opCode&0xF00F == 0x8007 {
		operation8XY7()
	} else if opCode&0xF00F == 0x800E {
		operation8XYE()
	} else if opCode&0xF00F == 0x9000 {
		operation9XY0()
	} else if opCode&0xF000 == 0xA000 {
		operationANNN()
	} else if opCode&0xF000 == 0xB000 {
		operationBNNN()
	} else if opCode&0xF000 == 0xC000 {
		operationCXKK()
	} else if opCode&0xF000 == 0xD000 {
		operationDXYN()
	} else if opCode&0xF0FF == 0xE09E {
		operationEX9E()
	} else if opCode&0xF0FF == 0xE0A1 {
		operationEXA1()
	} else if opCode&0xF0FF == 0xF007 {
		operationFX07()
	} else if opCode&0xF0FF == 0xF00A {
		operationFX0A()
	} else if opCode&0xF0FF == 0xF015 {
		operationFX15()
	} else if opCode&0xF0FF == 0xF018 {
		operationFX18()
	} else if opCode&0xF0FF == 0xF01E {
		operationFX1E()
	} else if opCode&0xF0FF == 0xF029 {
		operationFX29()
	} else if opCode&0xF0FF == 0xF033 {
		operationFX33()
	} else if opCode&0xF0FF == 0xF055 {
		operationFX55()
	} else if opCode&0xF0FF == 0xF065 {
		operationFX65()
	} else {
		fmt.Println("wrong operation code :(")
		running = false
	}
}

/*
Each func performs the individual operation depending upon the architecture
*/

func operation00E0() {
	for i := 0; i < len(pixel); i++ {
		for j := 0; j < len(pixel[i]); j++ {
			pixel[i][j] = false
		}
	}
	redraw = true
}

func operation00EE() {
	programCounter = stack[stackPointer]
	stackPointer--
}

func operation1NNN() {
	programCounter = opCode & 0x0FFF
}

func operation2NNN() {
	stackPointer++
	stack[stackPointer] = programCounter
	programCounter = opCode & 0x0FFF
}

func operation3XKK() {
	if register[(opCode&0x0F00)>>8] == byte(opCode&0x00FF) {
		programCounter += 2
	}
}
func operation4XKK() {
	if register[(opCode&0x0F00)>>8] != byte(opCode&0x00FF) {
		programCounter += 2
	}
}
func operation5XY0() {
	if register[(opCode&0x0F00)>>8] == register[(opCode&0x00F0)>>4] {
		programCounter += 2
	}
}
func operation6XKK() {
	register[(opCode&0x0F00)>>8] = byte(opCode & 0x00FF)
}
func operation7XKK() {
	register[(opCode&0x0F00)>>8] = register[(opCode&0x0F00)>>8] + byte(opCode&0x00FF)
}
func operation8XY0() {
	register[(opCode&0x0F00)>>8] = register[(opCode&0x00F0)>>4]
}
func operation8XY1() {
	register[(opCode&0x0F00)>>8] |= register[(opCode&0x00F0)>>4]
}
func operation8XY2() {
	register[(opCode&0x0F00)>>8] &= register[(opCode&0x00F0)>>4]
}
func operation8XY3() {
	register[(opCode&0x0F00)>>8] ^= register[(opCode&0x00F0)>>4]
}
func operation8XY4() {
	result := uint16(register[(opCode&0x0F00)>>8]) + uint16(register[(opCode&0x00F0)>>4])
	if result > 0xFF {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
	register[(opCode&0x0F00)>>8] = byte(result & 0xFF)
}
func operation8XY5() {
	if register[(opCode&0x0F00)>>8] > register[(opCode&0x00F0)>>4] {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
	register[(opCode&0x0F00)>>8] -= register[(opCode&0x00F0)>>4]
}
func operation8XY6() {
	register[0xF] = register[(opCode&0x0F00)>>8] & 0x1
	register[(opCode&0x0F00)>>8] >>= 1
}
func operation8XY7() {
	if register[(opCode&0x0F00)>>8] < register[(opCode&0x00F0)>>4] {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
	register[(opCode&0x0F00)>>8] = register[(opCode&0x00F0)>>4] - register[(opCode&0x0F00)>>8]
}
func operation8XYE() {
	register[0xF] = (register[(opCode&0x0F00)>>8] & (1 << 7)) >> 7
	register[(opCode&0x0F00)>>8] <<= 1
}
func operation9XY0() {
	if register[(opCode&0x0F00)>>8] != register[(opCode&0x00F0)>>4] {
		programCounter += 2
	}
}
func operationANNN() {
	index = opCode & 0x0FFF
}
func operationBNNN() {
	programCounter = (opCode&0x0FFF + uint16(register[0x0])) & 0xFFF
}
func operationCXKK() {
	var generatedNumber byte = byte(rand.Uint32()%256) & byte(opCode&0x00FF)
	register[(opCode&0x0F00)>>8] = generatedNumber
}
func operationDXYN() {
	x := register[(opCode&0x0F00)>>8]
	y := register[(opCode&0x00F0)>>4]
	n := uint16((opCode & 0x000F))
	collision := byte(0)
	for j := uint16(0); j < n; j++ {
		value := memory[index+uint16(j)]
		for i := uint16(0); i < 8; i++ {
			xx := (uint16(x) + i) % 64
			yy := (uint16(y) + j) % 32
			writeValue := ((value & (byte(1) << (7 - i))) != 0)
			prevValue := pixel[xx][yy]
			if prevValue != writeValue {
				pixel[xx][yy] = true
			} else {
				pixel[xx][yy] = false
			}
			if pixel[xx][yy] != prevValue && !pixel[xx][yy] {
				collision = byte(1)
			}
		}
	}
	register[0xF] = collision
	redraw = true
}
func operationEX9E() {
	if keyState[register[(opCode&0x0F00)>>8]] {
		programCounter += 2
	}
}
func operationEXA1() {
	if !keyState[register[(opCode&0x0F00)>>8]] {
		programCounter += 2
	}
}
func operationFX07() {
	register[(opCode&0x0F00)>>8] = delayTimer
}
func operationFX0A() {
	running = true
	keyPressed := false
	keyPressedIndex := byte(0)
	for !keyPressed && running {
		for event := sdl.PollEvent(); event != nil && running; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				for i := byte(0); i < byte(len(keyCodes)); i++ {
					if e.Keysym.Scancode == keyCodes[i] && e.State == 1 {
						keyState[i] = true
						keyPressed = true
						keyPressedIndex = i
					}
				}
				break
			}
		}
	}
	register[(opCode&0x0F00)>>8] = keyPressedIndex
}
func operationFX15() {
	delayTimer = register[(opCode&0x0F00)>>8]
}
func operationFX18() {
	soundTimer = register[(opCode&0x0F00)>>8]
}

func operationFX1E() {
	result := (index + uint16(register[(opCode&0x0F00)>>8]))

	if result > 0xFFF {
		register[0xF] = 1
	} else {
		register[0xF] = 0
	}
	index = result & 0xFFF
}
func operationFX29() {
	index = 5 * uint16(register[(opCode&0x0F00)>>8])
}
func operationFX33() {
	value := register[(opCode&0x0F00)>>8]
	memory[index+2] = value % 10
	value /= 10
	memory[index+1] = value % 10
	value /= 10
	memory[index] = value % 10
}
func operationFX55() {
	x := (opCode & 0x0F00) >> 8
	for i := uint16(0); i <= uint16(x); i++ {
		memory[i+index] = register[i]
	}
}
func operationFX65() {
	x := (opCode & 0x0F00) >> 8
	for i := uint16(0); i <= uint16(x); i++ {
		register[i] = memory[i+index]
	}
}
