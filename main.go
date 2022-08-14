package main

import (
	"github.com/go-system-tasks/helpers"
	"github.com/go-system-tasks/tui"
)

func main() {

	//Refresh lists of services and process
	svcs, procs := helpers.Refresh()

	//Create TUI
	tui.Start(svcs, procs)
}
