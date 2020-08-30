use log::debug;
use gtk::prelude::*;

struct PineappleData {
    width: i32,
    height: i32,
}

// Default signals
pub fn on_activate(application: &gtk::Application) {
    let pen_data = PineappleData {
        width: 640,
        height: 480,
    };

    let window = gtk::ApplicationWindow::new(application);

    // Set window configuration
    window.set_title("Pineapple Pen");
    window.set_default_size(pen_data.width, pen_data.height);
    window.set_position(gtk::WindowPosition::Center);

    let drawing_area = gtk::DrawingArea::new();
    drawing_area.connect_draw(|w,c|draw(w, c));

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
    debug!("Got draw event");
    c.set_source_rgb(25.0, 255.0, 129.0);
    c.rectangle(1.0, 1.0, 100.0, 200.0);
    c.fill();
    gtk::Inhibit(false)
}
