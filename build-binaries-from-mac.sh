#!/bin/bash

set -e

# Change to the project root
cd "cmd/godo" || exit

# Remove dist directory if it exists
if [ -d "dist" ]; then
  rm -r dist
fi

RELEASE_VERSION=0.0.2-alpha

platforms=("linux/amd64" "linux/arm64" "windows/amd64" "windows/arm64" "darwin/amd64" "darwin/arm64")

# Create output folders
mkdir -p dist/linux
mkdir -p dist/windows
mkdir -p dist/darwin
mkdir -p dist/release



for platform in "${platforms[@]}"; do
  IFS='/' read -ra target <<< "$platform"
  output_name="${target[0]}_${target[1]}"

  # Create ZIP for Windows
  if [ "${target[0]}" == "windows" ]; then
    GOOS=${target[0]} GOARCH=${target[1]} go build -o dist/"${target[0]}"/godo.exe
    zip -j dist/release/godo-v${RELEASE_VERSION}-"$output_name".zip dist/"${target[0]}"/godo.exe
    cd dist/release || exit
    shasum -a 256 godo-v${RELEASE_VERSION}-"$output_name".zip | tr -d '\n' > godo-v${RELEASE_VERSION}-"$output_name".sha256sum
    cd ../../
  fi

  # Create TAR.GZ for Linux and macOS
  if [ "${target[0]}" == "linux" ] || [ "${target[0]}" == "darwin" ]; then
    GOOS=${target[0]} GOARCH=${target[1]} go build -o dist/"${target[0]}"/godo
    tar -czf dist/release/godo-v${RELEASE_VERSION}-"$output_name".tar.gz -C dist/"${target[0]}" godo
    cd dist/release || exit
    shasum -a 256 godo-v${RELEASE_VERSION}-"$output_name".tar.gz | tr -d '\n' > godo-v${RELEASE_VERSION}-"$output_name".sha256sum
    cd ../../
  fi
done


