package encoders

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

// PNG Encoder
type pngEncoder struct {
	path  string
	frame int
}

// NewPNGEncoder creates a new PNG encoder
func NewPNGEncoder() Encoder {
	return &pngEncoder{}
}

// Initialize initializes the encoder
func (p *pngEncoder) Initialize(path string) error {
	p.path = path + "%d.png"
	return nil
}

// Encode encodes an image
func (p *pngEncoder) Encode(img image.Image) error {
	file, err := os.OpenFile(fmt.Sprintf(p.path, p.frame), os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	p.frame++
	return nil
}

// Cleanup cleans up the encoder
func (p *pngEncoder) Cleanup() error {
	return nil
}
