package img

/*
#include <stdlib.h>
#include <SDL3_image/SDL_image.h>
*/
import "C"

const (
	VERSION int = C.SDL_IMAGE_VERSION
)

func GetVersion() int {
	return int(C.IMG_Version())
}
