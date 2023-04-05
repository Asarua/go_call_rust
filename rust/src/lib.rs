#[no_mangle]
pub extern "C" fn rust_add(a: u32, b: u64) -> u64 {
  a as u64 + b
}