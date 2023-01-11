pub use suborbital_macro::*;

pub mod ffi;
pub mod graphql;
pub mod http;
pub mod log;
pub mod plugin;
pub mod req;
pub mod resp;
pub mod util;

/// This file represents the Rust "API" for E2Core Wasm plugins. The functions defined herein are used to exchange
/// data between the host (E2Core, written in Go) and the Plugin (a Wasm module, in this case written in Rust).

/// State struct to hold our dynamic Plugin
struct State<'a> {
	ident: i32,
	plugin: Option<&'a dyn plugin::Plugin>,
}

/// The state that holds the user-provided Plugin and the current ident
static mut STATE: State = State {
	ident: 0,
	plugin: None,
};
