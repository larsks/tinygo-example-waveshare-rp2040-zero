TARGET = waveshare-rp2040-zero

.PHONY: flash
flash:
	tinygo flash -target $(TARGET) main.go

# This prints out the appropriate GOROOT for your target. Useful if your editor/IDE
# does not have native support for tinygo. You would use it like this:
#
#     GOROOT=$(make goroot) edit main.go
#
.PHONY: goroot
goroot:
	@tinygo info -json $(TARGET) | jq -r .goroot
