# Ozone

A declarative and language-agnostic CLI framework

# Thoughts on Wasm

For first version only allow a single Wasm module per command. Furthermore, we will only support Wasi. We will not support custom libraries like Suborbital or capabilities like WasmCloud

# TODO

- Validate recursive struct YAML parsing
- Add more sophistication to arguments
- Handle single and multi command CLIS
- TODO build correct usage string based on provided args
- Temporarily just pass in flags as env variables to the WASM
