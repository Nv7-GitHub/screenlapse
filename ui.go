package main

import "github.com/andlabs/ui"

func setupUI() {
	win := ui.NewWindow("Screenlapse", 640, 480, true)
	win.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		win.Destroy()
		return true
	})
	win.SetMargined(true)

	// Setup UI

	win.Show()
}
