OUTPUT_DIR=./bin

# List of supported OS and architecture combinations
OS_ARCH_LIST = \
    darwin/amd64 \
    darwin/arm64 \
    linux/386 \
    linux/amd64 \
    linux/arm64 \
    windows/386 \
    windows/amd64 \
    windows/arm64

all: $(OS_ARCH_LIST)

# Rule to build a specific target binary
$(OS_ARCH_LIST):
	@echo "Building $@"
	@mkdir -p $(dir $(OUTPUT_DIR)/$(subst /,_,$@))
	cd cmd && GOOS=$(word 1,$(subst /, ,$@)) GOARCH=$(word 2,$(subst /, ,$@)) go build -o ../$(OUTPUT_DIR)/$(subst /,_,$@)/drawbox

clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: all clean
