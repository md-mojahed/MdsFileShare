#!/bin/bash

# Build script for MdsFileShare
# Creates cross-platform binaries in ./builds directory

set -e

APP_NAME="MdsFileShare"
BUILD_DIR="./builds"
VERSION=$(date +%Y%m%d-%H%M%S)

echo "🚀 Building $APP_NAME..."
echo "📦 Version: $VERSION"
echo ""

# Create builds directory
mkdir -p "$BUILD_DIR"

# Build for different platforms
echo "🔨 Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o "$BUILD_DIR/${APP_NAME}-windows-amd64.exe"

echo "🔨 Building for Windows (386)..."
GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o "$BUILD_DIR/${APP_NAME}-windows-386.exe"

echo "🔨 Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "$BUILD_DIR/${APP_NAME}-linux-amd64"

echo "🔨 Building for Linux (arm64)..."
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o "$BUILD_DIR/${APP_NAME}-linux-arm64"

echo "🔨 Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o "$BUILD_DIR/${APP_NAME}-macos-amd64"

echo "🔨 Building for macOS (arm64/M1)..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o "$BUILD_DIR/${APP_NAME}-macos-arm64"

echo ""
echo "✅ Build complete!"
echo ""
echo "📁 Binaries created in $BUILD_DIR:"
ls -lh "$BUILD_DIR"
echo ""
echo "📊 Total size:"
du -sh "$BUILD_DIR"
