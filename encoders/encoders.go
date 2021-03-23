package encoders

import "image"

// Encoder contains all the methods required for an encoder.
type Encoder interface {
	Initialize(path string, framerate, w, h int) error
	Encode(img image.Image) error
	Cleanup() error
}
