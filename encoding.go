package main

import (
	"image"
	"time"

	"github.com/Nv7-Github/screenlapse/encoders"
	"github.com/kbinani/screenshot"
)

var recording = false
var enc encoders.Encoder

var encs map[string]func() encoders.Encoder = map[string]func() encoders.Encoder{
	"PNG": encoders.NewPNGEncoder,
	"AVI": encoders.NewAVIEncoder,
}

func record(path string, speedup int, framerate int, encoder string, display int) {
	recording = true
	enc = encs[encoder]()
	rct := screenshot.GetDisplayBounds(display)
	err := enc.Initialize(path, framerate, rct.Dx(), rct.Dy())
	handle(err)
	go func() {
		var scr image.Image
		var err error
		rct := screenshot.GetDisplayBounds(display)
		for recording {
			time.Sleep(time.Duration(1/framerate) * time.Second)
			scr, err = screenshot.CaptureRect(rct)
			handle(err)
			if recording {
				err = enc.Encode(scr)
				handle(err)
			}
		}
	}()
}

func stopRecording() {
	recording = false
	err := enc.Cleanup()
	handle(err)
}

func getDisplays() int {
	return screenshot.NumActiveDisplays()
}
