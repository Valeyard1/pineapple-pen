package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/valeyard1/pineapple-pen/pkg/callbacks"
	"github.com/valeyard1/pineapple-pen/pkg/utils"
	"os"
)

const (
	appID = "com.github.valeyard1.pineapple-pen"
)

func main() {
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	utils.ErrorCheck("Failed to create application: ", err)

	signals := map[string]interface{}{
		"startup": callbacks.Startup,
		"activate": callbacks.Activate,
		"shutdown": callbacks.Shutdown,
	}

	application.Connect("startup", signals["startup"])
	application.Connect("activate", signals["activate"])
	application.Connect("shutdown", signals["shutdown"])

	os.Exit(application.Run(os.Args))
}
