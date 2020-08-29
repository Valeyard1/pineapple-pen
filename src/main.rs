use std::error::Error;
use std::env;

use gio::prelude::*;

mod callbacks;
pub use crate::callbacks::*;


const APPID: &str = "com.github.valeyard1.pineapple-pen";

fn main() -> Result<(), Box<dyn Error>> {
    env::set_var("RUST_LOG" , "debug");
    env_logger::init();

    let app = gtk::Application::new(Some(APPID), Default::default())
        .expect("Failed to create application");

    app.connect_startup(|app| on_startup(app));
    app.connect_activate(|app| on_activate(app));
    app.connect_shutdown(|app| on_shutdown(app));

    app.run(&std::env::args().collect::<Vec<_>>());

    Ok(())
}
