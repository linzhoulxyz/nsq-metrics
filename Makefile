BUILD_DIR = build

GO       = go
GOX      = gox
GOX_ARGS = -output="$(BUILD_DIR)/{{.Dir}}_{{.OS}}_{{.Arch}}" -osarch="linux/amd64 linux/arm linux/arm64 darwin/amd64 darwin/arm64 windows/amd64"

.PHONY: build
build:
	$(GO) build -o $(BUILD_DIR)/nsq-metrics .

.PHONY: clean
clean:
	rm -R $(BUILD_DIR)/* || true

.PHONY: test
test:
	$(GO) test ./...

.PHONY: release-build
release-build:
	@go install github.com/mitchellh/gox@latest
	@$(GOX) $(GOX_ARGS) github.com/leonkunert/nsq-metrics
