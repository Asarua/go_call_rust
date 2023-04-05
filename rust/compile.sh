#! /bin/bash

cd rust && cbindgen --config cbindgen.toml --crate rust_demo --output rust_demo.h --lang c
