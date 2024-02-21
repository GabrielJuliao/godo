#!/bin/bash

set -euo pipefail

# Define log levels
readonly LOG_LEVEL_TRACE=0
readonly LOG_LEVEL_DEBUG=1
readonly LOG_LEVEL_INFO=2
readonly LOG_LEVEL_WARN=3
readonly LOG_LEVEL_ERROR=4

# Default log level
LOG_LEVEL=$LOG_LEVEL_INFO

# Function to set log level
set_log_level() {
    case $1 in
        "TRACE") LOG_LEVEL=$LOG_LEVEL_TRACE ;;
        "DEBUG") LOG_LEVEL=$LOG_LEVEL_DEBUG ;;
        "INFO") LOG_LEVEL=$LOG_LEVEL_INFO ;;
        "WARN") LOG_LEVEL=$LOG_LEVEL_WARN ;;
        "ERROR") LOG_LEVEL=$LOG_LEVEL_ERROR ;;
        *) echo "Invalid log level: $1" >&2 ;;
    esac
}

# Function to log messages
log() {
    local level=$1
    shift
    local message="$*"
    local timestamp=$(date +"%Y-%m-%d %H:%M:%S")

    # Check log level
    if [ $LOG_LEVEL -le $level ]; then
        case $level in
            $LOG_LEVEL_TRACE) printf "\e[34m[%s] [TRACE]\e[0m %s\n" "$timestamp" "$message" ;;
            $LOG_LEVEL_DEBUG) printf "\e[32m[%s] [DEBUG]\e[0m %s\n" "$timestamp" "$message" ;;
            $LOG_LEVEL_INFO) printf "\e[36m[%s] [INFO]\e[0m %s\n" "$timestamp" "$message" ;;
            $LOG_LEVEL_WARN) printf "\e[33m[%s] [WARN]\e[0m %s\n" "$timestamp" "$message" ;;
            $LOG_LEVEL_ERROR) printf "\e[31m[%s] [ERROR]\e[0m %s\n" "$timestamp" "$message" >&2 ;;
        esac
    fi
}

buildReleaseLink() {
    local os="$1"
    local arch="$2"
    local version="$3"
    local bin_filename="$4"
    local sha="$5"
    local sha_filename="$6"

    case $os in
        darwin)
            os="MacOS"
            ;;
        linux)
            os="Linux"
            ;;
        windows)
            os="Windows"
            ;;
    esac

    local bin_url="[${os} ${arch}](https://github.com/GabrielJuliao/godo/releases/download/${version}/${bin_filename})"
    local sha_url="[checksum](https://github.com/GabrielJuliao/godo/releases/download/${version}/${sha_filename})"
    local link=" - ${bin_url} (${sha_url} / ${sha})"
    echo "$link" >>RELEASE.md
}

cleanup() {
    log $LOG_LEVEL_INFO "Cleaning up..."
    if [ -d "dist" ]; then
        rm -rf dist
        log $LOG_LEVEL_DEBUG "Removed directory 'dist'"
    else
        log $LOG_LEVEL_INFO "Directory 'dist' does not exist, skipping..."
    fi

    if [ -f "RELEASE.md" ]; then
        rm RELEASE.md
        log $LOG_LEVEL_DEBUG "Removed file RELEASE.md"
    else
        log $LOG_LEVEL_INFO "File RELEASE.md does not exist, skipping..."
    fi
    log $LOG_LEVEL_INFO "Cleanup completed!"
}

main() {
    local version
    version=$(<VERSION)
    log $LOG_LEVEL_INFO "Starting the build process for version $version"

    platforms=("linux/amd64" "linux/arm64" "windows/amd64" "windows/arm64" "darwin/amd64" "darwin/arm64")
    platforms=($(printf "%s\n" "${platforms[@]}" | sort -t'/' -k1.1))

    log $LOG_LEVEL_DEBUG "Platforms to be built: ${platforms[*]}"

    echo "GODO ${version} add title here" >RELEASE.md
    mkdir -p dist/release

    for platform in "${platforms[@]}"; do
        local os
        local arch
        IFS='/' read -r os arch <<<"$platform"
        local output_name="${os}_${arch}"

        log $LOG_LEVEL_INFO "Building: ${output_name}"

        # Build and package for Windows
        if [ "$os" == "windows" ]; then
            GOOS="$os" GOARCH="$arch" go build -o "dist/${os}/${arch}/godo.exe"
            (
                cd dist/release || exit
                zip -j "godo-${version}-${output_name}.zip" "../${os}/${arch}/godo.exe"
                shasum -a 256 "godo-${version}-${output_name}.zip" >"godo-${version}-${output_name}.sha256sum"
            )
            buildReleaseLink "$os" "$arch" "$version" "godo-${version}-${output_name}.zip" "$(cat dist/release/godo-${version}-${output_name}.sha256sum | awk '{print $1}')" "godo-${version}-${output_name}.sha256sum"
        fi

        # Build and package for Linux and macOS
        if [ "$os" == "linux" ] || [ "$os" == "darwin" ]; then
            GOOS="$os" GOARCH="$arch" go build -o "dist/${os}/${arch}/godo"
            (
                cd dist/release || exit
                tar -czf "godo-${version}-${output_name}.tar.gz" -C "../${os}/${arch}" godo
                shasum -a 256 "godo-${version}-${output_name}.tar.gz" >"godo-${version}-${output_name}.sha256sum"
            )
            buildReleaseLink "$os" "$arch" "$version" "godo-${version}-${output_name}.tar.gz" "$(cat dist/release/godo-${version}-${output_name}.sha256sum | awk '{print $1}')" "godo-${version}-${output_name}.sha256sum"
        fi
    done

    log $LOG_LEVEL_INFO "Build process completed!"
}

# Main script execution
set_log_level "DEBUG"
if [[ $# -gt 0 && "$1" == "clean" ]]; then
    cleanup
else
    cleanup
    main
fi