package main

import (
	"fmt"

	IMG "github.com/smack0007/sdl3-go/img"
	SDL "github.com/smack0007/sdl3-go/sdl"
)

func main() {
	fmt.Printf("=== Compiled ===\n")
	fmt.Printf("SDL Version: %d.%d.%d\n", SDL.VERSIONNUM_MAJOR(SDL.VERSION), SDL.VERSIONNUM_MINOR(SDL.VERSION), SDL.VERSIONNUM_MICRO(SDL.VERSION))
	fmt.Printf("SDL Revision: %s\n", SDL.REVISION)
	fmt.Printf("SDL_image Version: %d.%d.%d\n", SDL.VERSIONNUM_MAJOR(IMG.VERSION), SDL.VERSIONNUM_MINOR(IMG.VERSION), SDL.VERSIONNUM_MICRO(IMG.VERSION))

	fmt.Printf("\n")

	fmt.Printf("=== Linked ===\n")
	sdlVersion := SDL.GetVersion()
	fmt.Printf("SDL Version: %d.%d.%d\n", SDL.VERSIONNUM_MAJOR(sdlVersion), SDL.VERSIONNUM_MINOR(sdlVersion), SDL.VERSIONNUM_MICRO(sdlVersion))
	fmt.Printf("SDL Revision: %s\n", SDL.GetRevision())
	imgVersion := IMG.GetVersion()
	fmt.Printf("SDL_image Version: %d.%d.%d\n", SDL.VERSIONNUM_MAJOR(imgVersion), SDL.VERSIONNUM_MINOR(imgVersion), SDL.VERSIONNUM_MICRO(imgVersion))
}
