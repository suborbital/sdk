# Download and build a compatible version wit-bindgen
# Originally from:
# https://github.com/ospencer/witx-bindgen/tree/oscar/js-buffer-size
FROM rust:slim-bullseye AS tooling

RUN apt update && apt install -y git 

WORKDIR /tmp

RUN mkdir witx-bindgen && \
    git clone https://github.com/bytecodealliance/wit-bindgen.git --single-branch wit-bindgen && \
    cd wit-bindgen && \
    git checkout 48f364d1d4e42275b219d0862af8aa8b17c49d51 && \
    cargo build --release

# Executable in /tmp/wit-bindgen/target/release/wit-bindgen


FROM rust:slim-bullseye AS sdk

#RUN mkdir wit-bindgen && \
#    git clone https://github.com/suborbital/sdk.git --single-branch --branch v0.16.3 && \
#    cd wit-bindgen && \

WORKDIR /tmp

COPY --from=tooling /tmp/wit-bindgen/target/release/wit-bindgen /usr/bin/
COPY env.wit .

# Generate JS & TS bindings
RUN wit-bindgen js --import env.wit

# Generate Rust bindings
RUN wit-bindgen rust-wasm --import env.wit

CMD wit-bindgen --help
