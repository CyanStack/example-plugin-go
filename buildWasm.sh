#!/bin/bash
# Test script for building and running WASM using our new wasmFeature
# Ensure you have wasmer installed

echo "Building Go to WebAssembly with WASI target..."
cd src
export GOOS=wasip2
export GOARCH=wasm

echo Go does not support wasip2 yet, so we use tinygo instead

tinygo build -o ../wasm/main.wasm ./...


if [ ! -f "../wasm/main.wasm" ]; then
    echo "WASM compilation failed!"
    exit 1
fi

