use std::env;
use std::fs;
use std::path::Path;

fn main() {
    // Get the directory that the output of the build will be placed in.
    let out_dir = env::var("OUT_DIR").unwrap();
    let debug_dir = Path::new(&out_dir).parent().unwrap().parent().unwrap().parent().unwrap();

    // Copy the file to the output directory.
    fs::copy("src/input.txt", debug_dir.join("input.txt")).unwrap();
    fs::copy("src/testinput.txt", debug_dir.join("testinput.txt")).unwrap();
}