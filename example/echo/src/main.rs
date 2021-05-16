use std::env;

fn main() {
    println!("Arguments:");
    let args: Vec<String> = env::args().collect();
    println!("{:?}", args);
}
