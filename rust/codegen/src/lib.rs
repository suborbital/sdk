use proc_macro::TokenStream;
use quote::quote;
use syn::parse_macro_input;
use syn::DeriveInput;

#[proc_macro_derive(Plugin)]
pub fn derive_plugin(token_stream: TokenStream) -> TokenStream {
	let input = parse_macro_input!(token_stream as DeriveInput);

	let plugin_name = input.ident;

	let expanded = quote! {
		static PLUGIN: &#plugin_name = &#plugin_name{};

		#[no_mangle]
		pub extern fn init() {
			suborbital::plugin::use_plugin(PLUGIN);
		}
	};

	TokenStream::from(expanded)
}
