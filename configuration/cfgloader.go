package configuration

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

// TimerFrequency frequency of delayTimer and soundTimer
var TimerFrequency int64

// OperationFrequency number of operations executed per second
var OperationFrequency int64

// XScale xscale of the window
var XScale int32

// YScale of the window
var YScale int32

// KeyCodes stores the keymask to compare againt to update the key state
var KeyCodes = [16]sdl.Scancode{
	sdl.SCANCODE_0, sdl.SCANCODE_1,
	sdl.SCANCODE_2, sdl.SCANCODE_3,
	sdl.SCANCODE_4, sdl.SCANCODE_5,
	sdl.SCANCODE_6, sdl.SCANCODE_7,
	sdl.SCANCODE_8, sdl.SCANCODE_9,
	sdl.SCANCODE_A, sdl.SCANCODE_B,
	sdl.SCANCODE_C, sdl.SCANCODE_D,
	sdl.SCANCODE_E, sdl.SCANCODE_F,
}

// Load parses the file
func Load(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("can't open cfg file : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	TimerFrequency = int64(read(scanner))
	OperationFrequency = int64(read(scanner))
	XScale = int32(read(scanner))
	YScale = int32(read(scanner))
	for i := 0; i < 16; i++ {
		KeyCodes[i] = sdl.Scancode(read(scanner))
	}
}

// Reset the value of all important parameters
func Reset() {
	TimerFrequency = 60
	OperationFrequency = 2000
	XScale = 1
	YScale = 1
	KeyCodes = [16]sdl.Scancode{
		sdl.SCANCODE_0, sdl.SCANCODE_1,
		sdl.SCANCODE_2, sdl.SCANCODE_3,
		sdl.SCANCODE_4, sdl.SCANCODE_5,
		sdl.SCANCODE_6, sdl.SCANCODE_7,
		sdl.SCANCODE_8, sdl.SCANCODE_9,
		sdl.SCANCODE_A, sdl.SCANCODE_B,
		sdl.SCANCODE_C, sdl.SCANCODE_D,
		sdl.SCANCODE_E, sdl.SCANCODE_F,
	}
}

// Reads an integer from the scanner
func read(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		log.Fatal("cfg read error")
	}
	input, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("please write an integer")
	}
	return input
}
