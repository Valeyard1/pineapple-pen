package callbacks

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/valeyard1/pineapple-pen/pkg/utils"
	"log"
)

const (
	Title  = "Pineapple Pen"
)

type pineapplePenData struct{
	surface  *cairo.Surface            // context surface
    cr       *cairo.Context            // main cairo context
	debug    bool                      // true if you want debug messages
	width    int                       // width of window
	height   int                       // height of window
	x        float64                   // idk yet
	y        float64                   // idk yet
}

var penData = pineapplePenData{
    debug: true,
	width: 640,
	height: 480,
	x: 0.0,
	y: 0.0,
}

// Startup signal
func Startup() bool {
	log.Println("Application started")

	return true
}

// Shutdown signal
func Shutdown() bool {
	log.Println("Exited")

	return true
}


// Main application  functions

func Activate(application *gtk.Application) bool {

	Width, Height := penData.width, penData.height

	window, err := gtk.ApplicationWindowNew(application)
	utils.ErrorCheck("Failed to create window: ", err)

	// Set window configuration
	window.SetTitle(Title)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetDefaultSize(Width, Height)

    penData.surface = cairo.CreateImageSurface(cairo.FORMAT_ARGB32, Width, Height)
    penData.cr = cairo.Create(penData.surface)

	// Drawing Area
	da, _ := gtk.DrawingAreaNew()

	// Signal Handlers
	da.Connect("draw", draw)
    window.Connect("motion_notify_event", motionNotifyEvent)
    window.Connect("button_press_event", buttonPressEvent)

	window.Add(da)
	window.ShowAll()


	return true
}

//
// Signal callbacks functions
//

func draw(_ *gtk.DrawingArea, cr *cairo.Context) bool {
    if penData.debug {
        log.Println("DEBUG: Got draw event")
    }

    cr.Save()
    cr.SetSourceSurface(penData.surface, 0, 0)
    cr.SetOperator(cairo.OPERATOR_SOURCE)
    cr.Paint()

	return true
}

func motionNotifyEvent(app *gtk.ApplicationWindow, ev *gdk.Event) bool {
	x, y := gdk.EventMotionNewFromEvent(ev).MotionVal()

    if penData.debug {
        log.Printf("DEBUG: Got motion event on coord: (%f, %f)", x, y)
    }

    drawLine(x, y)
    app.QueueDraw()

	return true
}

func buttonPressEvent(app *gtk.ApplicationWindow, ev *gdk.Event) bool {
	x, y := gdk.EventMotionNewFromEvent(ev).MotionVal()

    if penData.debug {
        log.Printf("DEBUG: Got button Press event on coord: (%f, %f)", x, y)
    }

    drawLine(x, y)
    app.QueueDraw()

    return true
}


// helpers
func drawLine(x float64, y float64) {
    if penData.debug {
        log.Printf("DEBUG: Drawing line from: (%f, %f) to (%f, %f)", x, y, x, y)
    }

    penData.cr.SetSourceRGB(255, 0, 0)
    penData.cr.SetLineWidth(7)
    penData.cr.SetLineCap(cairo.LINE_CAP_ROUND)
    penData.cr.SetLineJoin(cairo.LINE_JOIN_ROUND)

    penData.cr.MoveTo(x, y)
    penData.cr.LineTo(x, y)
    penData.cr.Stroke()
}

