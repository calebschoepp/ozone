all:
	@$(MAKE) clean
	@$(MAKE) setup-example
	@go build
	@./ozone build example/
	@$(MAKE) compile

compile:
	@cd .ozone; go get ./...; go build

setup-example:
	@cd example/echo; cargo +nightly build --target=wasm32-wasi --release;
	@cp example/echo/target/wasm32-wasi/release/echo.wasm example/echo.wasm
	@cd example/cat; cargo +nightly build --target=wasm32-wasi --release;
	@cp example/cat/target/wasm32-wasi/release/cat.wasm example/cat.wasm

clean:
	@rm -rf .ozone