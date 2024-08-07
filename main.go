package main

import (
	"fmt"
	"os"

	"github.com/Nv7-Github/ui"
	_ "github.com/Nv7-Github/ui/winmanifest"
)

func handle(err error) {
	if err != nil {
		ui.QueueMain(func() {
			ui.MsgBoxError(win, "Error!", err.Error())
			fmt.Printf("Error: %s\n", err.Error())
			win.Destroy()
			ui.Quit()
			os.Exit(1)
		})
	}
}

func main() {
	ui.Main(setupUI)
}
