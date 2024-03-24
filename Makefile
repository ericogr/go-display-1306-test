#sudo apt install inotify-tools

BIN := display-test
BIN_PATH := bin
REMOTE_HOST := erico@192.168.0.35
REMOTE_DIR := /tmp
WATCH_DIR := .
DISPLAY_BUS := /dev/i2c-3
DISPLAY_WIDTH := 128
DISPLAY_HEIGHT := 32
DISPLAY_SEQUENTIAL := true

#--
PARAMETERS="-bus=$(DISPLAY_BUS) -with=$(DISPLAY_WIDTH) -height=$(DISPLAY_HEIGHT) -sequential=$(DISPLAY_SEQUENTIAL)"

.PHONY: all clean build copy run

all: build copy run

build:
	@echo "Building..."
	@env GOOS=linux GOARCH=arm64 GOARM=5 go build -o $(BIN_PATH)/$(BIN) ./cmd/display-test

copy: build
	@echo "Copying..."
	@rsync -avz $(BIN_PATH)/$(BIN) $(REMOTE_HOST):$(REMOTE_DIR)

run: copy stop
	@echo "Running..."
	@echo "Parameters:" $(PARAMETERS)
	@echo ssh $(REMOTE_HOST) $(REMOTE_DIR)/$(BIN) $(PARAMETERS)&
	@ssh $(REMOTE_HOST) $(REMOTE_DIR)/$(BIN) $(PARAMETERS)&

stop:
	@echo "Stopping..."
	@ssh $(REMOTE_HOST) pkill $(BIN) || true

clean:
	@echo "Cleaning..."
	@rm -f $(BIN_PATH)/$(BIN)

watch: all
	@echo "Monitoring... $(WATCH_DIR)..."
	@while inotifywait -r -e modify,create,delete --include '\.go' $(WATCH_DIR); do \
		$(MAKE) watch; \
	done