use log::debug;
use gtk::prelude::*;

// struct PineappleData {
    // // width:  i32,            // monitor width
    // // height: i32,            // monitor height
    // cr:     cairo::Context, // cairo context should be 
// }

// Default signals
pub fn on_activate(application: &gtk::Application) {
    let window = gtk::ApplicationWindow::new(application);
    let drawing_area = gtk::DrawingArea::new();

    let width = 640;
    let height = 480;

    // Set window configuration
    window.set_title("Pineapple Pen");
    window.set_default_size(width, height);
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

// Triggered when I move the mouse with the button pressed
fn motion_notify_event(application: &gtk::ApplicationWindow,
    ev: &gdk::EventMotion) -> glib::signal::Inhibit {

    // debug!("Got motion event on coord: ({}, {})",
        // ev.get_coords().unwrap().0,
        // ev.get_coords().unwrap().1);

    let surface = cairo::ImageSurface::create(cairo::Format::ARgb32, 640, 480)
        .expect("Couldn't create a surface!");

    let cr = cairo::Context::new(&surface);

    cr.set_source_rgb(255.0, 0.0, 0.0);
    cr.set_line_width(7.0);
    cr.set_line_cap(cairo::LineCap::Round);
    cr.set_line_join(cairo::LineJoin::Round);

    cr.move_to(ev.get_coords().unwrap().0, ev.get_coords().unwrap().1);
    cr.line_to(ev.get_coords().unwrap().0, ev.get_coords().unwrap().1);
    cr.stroke();

    gtk::Inhibit(false)
}

// Triggered when I press the button
fn button_press_event(application: &gtk::ApplicationWindow,
    ev: &gdk::EventButton) -> glib::signal::Inhibit {

    debug!("Got button event on coord: ({}, {})",
        ev.get_coords().unwrap().0,
        ev.get_coords().unwrap().1);

    let surface = cairo::ImageSurface::create(cairo::Format::ARgb32, 640, 480)
        .expect("Couldn't create a surface!");

    let cr = cairo::Context::new(&surface);

    cr.set_source_rgb(255.0, 0.0, 0.0);
    cr.set_line_width(7.0);
    cr.set_line_cap(cairo::LineCap::Round);
    cr.set_line_join(cairo::LineJoin::Round);

    cr.move_to(ev.get_coords().unwrap().0, ev.get_coords().unwrap().1);
    cr.line_to(0.0, 0.0);
    cr.stroke();
    cr.paint();
    cr.fill();

    gtk::Inhibit(false)
}

// draw_line() will draw the points continuously which
// will draw the whole line
fn draw_line(ctx: &cairo::Context, x: f64, y: f64) -> bool {
    debug!("Drawing line from: ({}, {}) to ({}, {})", x, y, x, y);

    ctx.set_source_rgb(255.0, 0.0, 0.0);
    ctx.set_line_width(2.0);
    ctx.set_line_cap(cairo::LineCap::Round);
    ctx.set_line_join(cairo::LineJoin::Round);

    ctx.move_to(x, y);
    ctx.line_to(x, y);
    ctx.paint();
    ctx.fill();
    ctx.stroke();

    true
}
