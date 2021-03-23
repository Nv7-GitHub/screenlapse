package encoders

import (
	"bytes"
	"image"
	"image/jpeg"
	"path/filepath"

	"github.com/icza/mjpeg"
)

// NewAVIEncoder creates a new AVI encoder
func NewAVIEncoder() Encoder {
	return &aviEncoder{}
}

type aviEncoder struct {
	aw mjpeg.AviWriter
}

func (a *aviEncoder) Initialize(path string, framerate, w, h int) error {
	ext := filepath.Ext(path)
	if ext == "" {
		path += ".avi"
	}
	var err error
	a.aw, err = mjpeg.New(path, int32(w), int32(h), int32(framerate))
	return err
}

func (a *aviEncoder) Encode(frame image.Image) error {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := jpeg.Encode(buf, frame, nil)
	if err != nil {
		return err
	}
	err = a.aw.AddFrame(buf.Bytes())
	return err
}

func (a *aviEncoder) Cleanup() error {
	return a.aw.Close()
}
