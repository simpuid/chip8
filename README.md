# Chip8
A chip8 interpreter written in GO lang
## Building the project
This project depends on go-sdl2 `github.com/veandco/go-sdl2`, so it is required to build the project.  
Get the project through `$ go get github.com/simpukr/chip8`. This will create a directory named `chip8` in the directory `$GOPATH/src/github.com/simpukr/chip8`  
Change to the newly created directory and execute `$ go build -o <yourfilename>` where `<yourfilename>` indicates the name of the executable the project generates. This will create an executable with name <yourfilename> in the same directory.  
## Executing
  The program can be execute by `$ ./<yourFileName> <relativeRomPath> <relativeConfigurationPath>`  
  The first argument `<relativeRomPath>` should indicate to the respective ROM path the program should load.  
  The second argument `<relativeConfigurationPath>` should point to the configuration file containing the configuration parameters. This argument is optional.
  ## Configuration
  The configuration file is a simple text file containing the values for different parameters (they must be integer).  
  Those parameters are (in the exact order):
  ```
  Chip8 timer frequency, The frequency used by the delay timer and sound timer. Default is 60.
  ```
  ```
  Chip8 operation frequency, no of operations the interpreter should execute in one second. Default is 2000
  ```
  ```
  Xscale of the window. Default is 1. Increase it to make window stretch horizontally.
  ```
  ```
  Yscale of the window. Default is 1. Increase it to make window stretch vertically.
  ```
  ```
  SDL_ScanCode for each key (0, 1, 2, 3, 4, 5, 6, 7, 8, 9, A, B, C, D, E, F) seperated by new lines
  ```
  The interpreter takes input through 16 buttons indexed in hexadecimal [0 to F] and the SDL_ScanCodes are used to indicate the key used for input. 
  Scan codes can be found at https://wiki.libsdl.org/SDLScancodeLookup , just put the decimal number in the respective position to bind that key for input.  
  Example configuration file:
  ```
60
2000
5
5
98
95
96
97
92
93
94
89
90
92
99
88
87
86
85
84
  ```
Timer frequency is 60. Operation frequency is 2000. Xscale is 5 and same for Yscale. 0th key is binded to key at ScanCode 98, 1st to 95, 2nd to 96 and in similar fashion 15th (F) to 84.
