use log::debug;
use gtk::prelude::*;

struct PineappleData {
    width:  i32,            // monitor width
    height: i32,            // monitor height
    // cr:     cairo::Context, // cairo context should be 
}

// Default signals
pub fn on_activate(application: &gtk::Application) {
    let pen_data = PineappleData {
        width: 640,
        height: 480,
    };

    let window = gtk::ApplicationWindow::new(application);
    let drawing_area = gtk::DrawingArea::new();

    // Set window configuration
    window.set_title("Pineapple Pen");
    window.set_default_size(pen_data.width, pen_data.height);
    window.set_position(gtk::WindowPosition::Center);

    // Connect signals
    drawing_area.connect_draw(|da, ctx|draw(da, ctx));
    window.connect_motion_notify_event(|app, event|motion_notify_event(app, event));
    window.connect_button_press_event(|app, event|button_press_event(app, event));

    window.add(&drawing_area);
    window.show_all();
}

pub fn on_startup(_: &gtk::Application) {
    debug!("Application started");
}

pub fn on_shutdown(_: &gtk::Application) {
    debug!("Application exited");
}


// Custom callbacks
fn draw(_: &gtk::DrawingArea, c: &cairo::Context) -> glib::signal::Inhibit {
    debug!("Got a draw event");

    c.save();
    c.set_operator(cairo::Operator::Source);
    c.paint();

    gtk::Inhibit(false)
}

fn motion_notify_event(application: &gtk::ApplicationWindow,
    ev: &gdk::EventMotion) -> glib::signal::Inhibit {

    debug!("Got motion event on coord: ({}, {})",
        ev.get_coords().unwrap().0,
        ev.get_coords().unwrap().1);

    application.queue_draw();

    gtk::Inhibit(false)
}

fn button_press_event(application: &gtk::ApplicationWindow,
    ev: &gdk::EventButton) -> glib::signal::Inhibit {

    debug!("Got button event on coord: ({}, {})",
        ev.get_coords().unwrap().0,
        ev.get_coords().unwrap().1);

    // draw_line(ctx: &cairo::Context, x: f64, y: f64);
    application.queue_draw();

    gtk::Inhibit(false)
}

// draw_line() will draw the points continuously which
// will draw the whole line
fn draw_line(ctx: &cairo::Context, x: f64, y: f64) -> bool {
    debug!("Drawing line from: ({}, {}) to ({}, {})", x, y, x, y);

    // I wish I could do PineappleData.cr.line_to(x, y), but
    // I can't make it because it would be a mutable global structure,
    // so I need to create a cairo context here
    // let ctx = cairo::Context::new();

    ctx.set_source_rgb(255.0, 0.0, 0.0);
    ctx.set_line_width(7.0);
    ctx.set_line_cap(cairo::LineCap::Round);
    ctx.set_line_join(cairo::LineJoin::Round);

    ctx.move_to(x, y);
    ctx.line_to(x, y);
    ctx.stroke();

    true
}
