# A demo of golang call rust

Compile the Rust library into cdylib using cbindgen, and then call it in Go using cgo

## dependencies

### Rust

- [cbindgen](https://github.com/eqrion/cbindgen)

## Notes

### 1. Add the `#[no_mangle]` macro to the function that Rust is exporting


### 2. Add `pub extern "C"` before the Rust function to be exported

### 3. Generating `.h` files using `cbindgen`

### 4. Golang needs to add these comments above `import "C"` and cannot have blank lines

```go
/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L../rust/target/debug -lrust_demo

#include "../rust/rust_demo.h"
*/
import "C"
```
