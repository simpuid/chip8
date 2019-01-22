package chip

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var memory [4096]byte
var register [16]byte
var opCode uint16
var index uint16
var delayTimer byte
var soundTimer byte
var programCounter uint16
var stackPointer uint8
var stack [16]uint16
var pixel [][]bool
var running = true
var redraw = false
var ticker *time.Ticker
var lastTimeTimer time.Time
var lastTimeOperation time.Time
var timePassedTimer = int64(0)
var timePassedOperation = int64(0)
var timerFrequency = int64(60)
var operationFrequency = int64(2000)

const startPointer = 0x200

var fontData = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

// Init chip8
func Init(codes [16]sdl.Scancode) {
	for i := 0; i < 4096; i++ {
		if i < len(fontData) {
			memory[i] = fontData[i]
		} else {
			memory[i] = 0
		}
	}
	for i := 0; i < 16; i++ {
		register[i] = 0
	}
	opCode = 0
	index = 0
	delayTimer = 0
	soundTimer = 0
	programCounter = startPointer
	stackPointer = 0
	for i := 0; i < 16; i++ {
		stack[i] = 0
	}
	pixel = make([][]bool, 64)
	for i := 0; i < 64; i++ {
		pixel[i] = make([]bool, 32)
	}
	for i := 0; i < 16; i++ {
		keyCodes[i] = codes[i]
	}
	lastTimeTimer = time.Now()
	lastTimeOperation = time.Now()
}

// ReadRom reads the ROM
func ReadRom(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("file open error", err)
	}
	defer file.Close()
	data := make([]byte, 4095-startPointer)
	total, err := file.Read(data)
	if err != nil {
		log.Fatal("file read error", err)
	}
	for i := 0; i < total; i++ {
		memory[startPointer+i] = data[i]
	}
	fmt.Println("Rom loaded ", total, " bytes")
}

//CycleEmulation emulates one cycle of operation
func CycleEmulation() bool {
	operationCycle()
	timerCycle()
	return redraw
}

// Update the timers depending on the time paseed
func timerCycle() {
	timePassedTimer += time.Now().Sub(lastTimeTimer).Nanoseconds()
	lastTimeTimer = time.Now()
	for timePassedTimer > 1000000000/timerFrequency {
		timePassedTimer -= 1000000000 / timerFrequency
		if delayTimer > 0 {
			delayTimer--
		}
		if soundTimer > 0 {
			soundTimer--
			if soundTimer == 0 {
				fmt.Println("Beep Boop !!")
			}
		}
	}
}

// Ececute opCodes depending on the time passed
func operationCycle() {
	timePassedOperation += time.Now().Sub(lastTimeOperation).Nanoseconds()
	lastTimeOperation = time.Now()
	for timePassedOperation > 1000000000/operationFrequency {
		timePassedOperation -= 1000000000 / operationFrequency
		opCode = (uint16(memory[programCounter]) << 8) | uint16(memory[programCounter+1])
		programCounter += 2
		startOperation()
	}
}

// GetPixel returns the pixel
func GetPixel() [][]bool {
	return pixel
}
