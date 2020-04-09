package callbacks

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
	"github.com/valeyard1/pineapple-pen/pkg/utils"
	"log"
)

const (
	Title = "Pineapple Pen"
)

var x = 0.0
var y = 0.0

func Activate(application *gtk.Application) bool {
	Width, Height := 640, 480

	window, err := gtk.ApplicationWindowNew(application)
	utils.ErrorCheck("Failed to create window: ", err)

	frame, err := gtk.FrameNew("")
	utils.ErrorCheck("Failed to create frame: ", err)

	// Set window configuration
	window.SetTitle(Title)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetDefaultSize(Width, Height)

	// Drawing Area
	da, _ := gtk.DrawingAreaNew()

	// Frame settings
	frame.SetShadowType(gtk.SHADOW_IN)

	// Signal Handlers
	window.Connect("destroy", gtk.MainQuit)
	da.Connect("draw", draw)

	window.Add(da)
	//window.Add(frame)
	window.ShowAll()

	/*
	TODO:
	1. Create the frames
	 */

	return true
}

func Startup(application *gtk.Application) bool {
	log.Println("Application started")

	return true
}

func Shutdown(application *gtk.Application) bool {
	log.Println("Exited")

	return true
}

func draw(da *gtk.DrawingArea, cr *cairo.Context) bool {
	cr.SetSourceRGB(0, 0, 0)
	cr.Rectangle(x, y, 0, 0)
	cr.Fill()

	return true
}
