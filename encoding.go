package main

import "image"

var encoders map[string]Encoder = map[string]Encoder{"Test Encoder": nil}

// Encoder contains all the methods required for an encoder.
type Encoder interface {
	Initialize(path string) error
	Encode(img image.Image) error
	Cleanup() error
}
