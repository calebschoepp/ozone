compile:
	@cd .ozone; go get ./...; go build

clean:
	@rm -rf .ozone