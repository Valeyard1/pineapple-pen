use gtk::prelude::*;
use log::*;

// Default signals
pub fn on_activate(application: &gtk::Application) {
    // … create a new window …
    let window = gtk::ApplicationWindow::new(application);
    // … with a button in it …
    let button = gtk::Button::with_label("Hello World!");
    // … which closes the window when clicked
    button.connect_clicked(glib::clone!(@weak window => move |_| window.close()));
    window.add(&button);
    window.show_all();
}

pub fn on_startup(_: &gtk::Application) {
    debug!("Application started");
}

pub fn on_shutdown(_: &gtk::Application) {
    debug!("Application exited");
}


// Custom callbacks
