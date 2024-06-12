#!/bin/bash

# Check if folder name is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <project_name>"
    exit 1
fi

project_name=$1
folder_path=$(realpath "projects/$project_name")

echo "Folder Path: $folder_path"
# Look for config.sh in the folder path
config_file="$folder_path/config/config.sh"
if [ -f "$config_file" ]; then
    source "$config_file"
else
    EXECUTABLE_NAME="$project_name"
fi

# Define the source directory where endpoints.json is located
SOURCE_DIR="projects\endpoint-monitor"

# Define the destination directory where endpoints.json should be copied
DEST_DIR="build"

# Create the destination directory if it doesn't exist
mkdir -p "$DEST_DIR"

# Copy all JSON files from source directory and its subdirectories to destination directory
find "$SOURCE_DIR" -type f -name "*.json" -exec cp {} "$DEST_DIR" \;
echo "All JSON files copied successfully to $DEST_DIR"

# Copy all TOML files from source directory and its subdirectories to destination directory
find "$SOURCE_DIR" -type f -name "*.toml" -exec cp {} "$DEST_DIR" \;
echo "All TOML files copied successfully to $DEST_DIR"

# Build the specified Go application with the desired executable name
go build -o "build/$EXECUTABLE_NAME.exe" "$folder_path/main.go"

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "$project_name project was built successfully as $EXECUTABLE_NAME.exe."
else
    echo "Error: $project_name project build failed."
fi
