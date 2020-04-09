package main

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

var unitSize = 5.0
var x = 0.0
var y = 0.0

func drawWindow(da *gtk.DrawingArea, cr *cairo.Context) {
	cr.SetSourceRGB(0, 0, 0)
	cr.Rectangle(x, y, unitSize, unitSize)
	cr.Fill()
}

func clicked(win *gtk.Window, ev *gdk.Event) bool {
	evMotion := gdk.EventMotionNewFromEvent(ev)
	xFromEvent, yFromEvent := evMotion.MotionVal()
	x, y = xFromEvent, yFromEvent
	win.QueueDraw()

	return false
}

func setupWindow(title string) *gtk.Window {
	Width, Height := 640, 480

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Pineapple Pen")
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(Width, Height)

	drawArea, _ := gtk.DrawingAreaNew()

	win.Connect("destroy", gtk.MainQuit)

	drawArea.Connect("draw", drawWindow)
	win.Connect("button-press-event", clicked)

	win.Add(drawArea)
	win.ShowAll()
	return win
}



func main() {
	gtk.Init(nil)

	win := setupWindow("Go Example Drawing area mouse event")

	win.ShowAll()

	gtk.Main()
}
