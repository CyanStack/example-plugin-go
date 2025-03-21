@echo off
REM Auto-generated WASM wrapper script
set PORT=4002
"C:\Users\Raket\.wasmer\bin\wasmer.exe" run --net --env PORT=4002 ./main.wasm %*
