package main

import (
	"log"
	"os"

	"github.com/simpukr/chip8/render"

	"github.com/simpukr/chip8/configuration"

	"github.com/simpukr/chip8/chip"
)

// Reads arguments and do accordingly
func main() {
	arguments := os.Args[1:]
	configuration.Reset()
	if len(arguments) < 1 {
		log.Fatal("rom path not supplied :(")
	} else if len(arguments) == 2 {
		configuration.Load(arguments[1])
	}
	render.Init(configuration.XScale, configuration.YScale)
	defer render.End()
	chip.Init(configuration.KeyCodes)
	chip.ReadRom(arguments[0])
	for chip.PollEvents() {
		if chip.CycleEmulation() {
			render.Draw(chip.GetPixel())
		}
	}
}
