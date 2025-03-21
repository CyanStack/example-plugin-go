@echo off
REM Test script for building and running WASM using WASI target
REM Ensure you have wasmer installed

echo Building Go to WebAssembly with WASI target...
cd src
set GOOS=wasip2
set GOARCH=wasm

echo Go does not support wasip2 yet, so we use tinygo instead

tinygo build -o ..\wasm\main.wasm .\...


if not exist "..\wasm\main.wasm" (
    echo WASM compilation failed!
    exit /b 1
)
