package main

import (
	"fmt"
	"os"

	"github.com/andlabs/ui"
)

var win *ui.Window

func handle(err error) {
	if err != nil {
		ui.MsgBoxError(win, "Error!", err.Error())
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func setupUI() {
	win = ui.NewWindow("Screenlapse", 640, 200, true)
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
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	outputGroup := ui.NewGroup("Output")
	outputGroup.SetMargined(true)
	outputForm := ui.NewForm()
	outputForm.SetPadded(true)

	speedInput := ui.NewSpinbox(1, 100)
	speedInput.SetValue(2)

	filename := ""
	hbox := ui.NewHorizontalBox()

	fileShower := ui.NewEntry()
	saveBtn := ui.NewButton("Select")
	saveBtn.OnClicked(func(*ui.Button) {
		filename = ui.SaveFile(win)
		fileShower.SetText(filename)
	})

	hbox.Append(fileShower, true)
	hbox.Append(saveBtn, false)

	encoderBox := ui.NewCombobox()
	encoderNames := make([]string, len(encoders))
	i := 0
	for k := range encoders {
		encoderBox.Append(k)
		encoderNames[i] = k
		i++
	}
	encoderBox.SetSelected(0)

	outputForm.Append("Speed Up", speedInput, false)
	outputForm.Append("Output File", hbox, false)
	outputForm.Append("Encoder", encoderBox, false)
	outputGroup.SetChild(outputForm)

	recordBtn := ui.NewButton("Record!")
	recordBtn.OnClicked(func(*ui.Button) {
		ui.MsgBox(win, "Info", fmt.Sprintf("Encoder: %s\nSpeed Up: %dx\nOutput File: %s", encoderNames[encoderBox.Selected()], speedInput.Value(), filename))
	})

	vbox.Append(outputGroup, false)
	vbox.Append(recordBtn, false)
	win.SetChild(vbox)

	win.Show()
}
