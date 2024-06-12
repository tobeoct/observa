# Variables (with default values)
SOURCE_DIR = projects/endpoint-monitor
DEST_DIR = build
EXECUTABLE_NAME ?= observa.endpoint-monitor
FOLDER_PATH ?= projects/endpoint-monitor

# Determine OS
ifeq ($(OS),Windows_NT)
	RM = del /s /q
	MKDIR = mkdir
	COPY = cmd /c copy
else
	RM = rm -rf
	MKDIR = mkdir -p
	COPY = cp -r
endif

# Default target
all: clean copy-json build-go

copy-json:
	$(MKDIR) $(DEST_DIR)
	for file in $(SOURCE_DIR)/*.json; do \
		$(COPY_COMMAND) $$file $(DEST_DIR)/; \
	done
	@echo "All JSON files copied successfully to $(DEST_DIR)"

build-go:
	go build -o $(DEST_DIR)/$(EXECUTABLE_NAME).exe $(FOLDER_PATH)/main.go
	@echo "$(EXECUTABLE_NAME) project was built successfully as $(EXECUTABLE_NAME).exe."

# Clean target
clean:
	$(RM) $(DEST_DIR)
	@echo "Cleaned up build directory."