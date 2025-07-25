TARGET = waveshare-rp2040-zero

# Build for TinyGo (uses real machine package)
.PHONY: build
build:
	tinygo build -target $(TARGET) -o main.uf2 .

# Flash to device
.PHONY: flash
flash:
	tinygo flash -target $(TARGET) main.go

# Run tests (uses mock machine package)
.PHONY: test
test:
	go test -v

# Run with mock (useful for development)
.PHONY: run-mock
run-mock:
	go run .

# This prints out the appropriate GOROOT for your target. Useful if your editor/IDE
# does not have native support for tinygo. You would use it like this:
#
#     GOROOT=$(make goroot) edit main.go
#
.PHONY: goroot
goroot:
	@tinygo info -json $(TARGET) | jq -r .goroot
