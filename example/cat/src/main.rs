use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();
    let filename = args.get(1).expect("Expected at least one argument");
    let contents = fs::read_to_string(filename).expect("Something went wrong reading the file");
    print!("{}", contents);
}
