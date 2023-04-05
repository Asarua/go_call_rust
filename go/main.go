package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L../rust/target/debug -lrust_demo

#include "../rust/rust_demo.h"
*/
import "C"
import "fmt"

func main() {
	a := 1
	b := 2

	clib_a := C.uint32_t(a)
	clib_b := C.uint64_t(b)

	clib_sum := C.rust_add(clib_a, clib_b)

	fmt.Printf("clib_sum is %d", clib_sum)
}
