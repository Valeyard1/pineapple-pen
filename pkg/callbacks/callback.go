package callbacks

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
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

	// Set window configuration
	window.SetTitle(Title)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetDefaultSize(Width, Height)

	// Drawing Area
	da, _ := gtk.DrawingAreaNew()

	// Signal Handlers
	da.Connect("draw", draw)
	window.Connect("motion_notify_event", motionNotifyEvent)

	window.Add(da)
	window.ShowAll()

	/*
		TODO:
		1. Create the frames
	*/

	return true
}

func Startup() bool {
	log.Println("Application started")

	return true
}

func Shutdown() bool {
	log.Println("Exited")

	return true
}

func draw(_ *gtk.DrawingArea, cr *cairo.Context) bool {
	cr.SetSourceRGB(0, 0, 0)
	cr.Rectangle(x, y, 0, 0)
	cr.Fill()

	return false
}

func motionNotifyEvent(_ *gtk.ApplicationWindow, evm *gdk.Event) bool {
	evMotion := gdk.EventMotionNewFromEvent(evm)
	x, y := evMotion.MotionVal()
	log.Printf("Got motion event on coord: (%f, %f)", x, y)

	return false
}
