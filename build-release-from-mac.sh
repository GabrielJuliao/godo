#!/bin/bash

set -e

# Remove dist directory if it exists
if [ -d "dist" ]; then
  rm -r dist
fi

if [ "$1" != "clean" ]; then

VERSION=$(cat VERSION | tr -d '[:space:]')

platforms=("linux/amd64" "linux/arm64" "windows/amd64" "windows/arm64" "darwin/amd64" "darwin/arm64")

# Create output folders
mkdir -p dist/release

for platform in "${platforms[@]}"; do

  IFS='/' read -ra target <<< "$platform"
  output_name="${target[0]}_${target[1]}"

  echo -e builing: "${output_name}"

    # Create ZIP for Windows
    if [ "${target[0]}" == "windows" ]; then
      GOOS=${target[0]} GOARCH=${target[1]} go build -o dist/"${target[0]}"/"${target[1]}"/godo.exe
      zip -j dist/release/godo-"${VERSION}"-"$output_name".zip dist/"${target[0]}"/"${target[1]}"/godo.exe
      cd dist/release || exit
      shasum -a 256 godo-"${VERSION}"-"$output_name".zip | tr -d '\n' > godo-"${VERSION}"-"$output_name".sha256sum
      cd ../../
    fi

  # Create TAR.GZ for Linux and macOS
  if [ "${target[0]}" == "linux" ] || [ "${target[0]}" == "darwin" ]; then
    GOOS=${target[0]} GOARCH=${target[1]} go build -o dist/"${target[0]}"/"${target[1]}"/godo
    tar -czf dist/release/godo-"${VERSION}"-"$output_name".tar.gz -C dist/"${target[0]}"/"${target[1]}" godo
    cd dist/release || exit
    shasum -a 256 godo-"${VERSION}"-"$output_name".tar.gz | tr -d '\n' > godo-"${VERSION}"-"$output_name".sha256sum
    cd ../../
  fi

done

fi