# Ozone

A declarative and language-agnostic CLI framework

# Thoughts on Wasm

For first version only allow a single Wasm module per command. Furthermore, we will only support Wasi. We will not support custom libraries like Suborbital or capabilities like WasmCloud

# TODO

- Temporarily just pass in arguments to the WASM via WASI arguments
- Validate recursive struct YAML parsing
- Add more sophistication to arguments
- Handle single and multi command CLIS
