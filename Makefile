BIN=display-test
BIN_PATH=bin
REMOTE_HOST=erico@192.168.0.35
REMOTE_DIR=/tmp

.PHONY: all clean build copy run

all: build copy run

build:
	@echo "Building..."
	@env GOOS=linux GOARCH=arm64 GOARM=5 go build -o $(BIN_PATH)/$(BIN) ./cmd/display-test

copy: build
	@echo "Copying..."
	@rsync -avz $(BIN_PATH)/$(BIN) $(REMOTE_HOST):$(REMOTE_DIR)

run: copy
	@echo "Running..."
	@ssh $(REMOTE_HOST) $(REMOTE_DIR)/$(BIN)

clean:
	@echo "Cleaning..."
	@rm -f $(BIN_PATH)/$(BIN)
