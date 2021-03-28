rm bundle.wasm
set GOROOT=C:\Users\user\Desktop\portprog\installs\go
cp %GOROOT%/misc/wasm/wasm_exec.js .
set GOOS=js
set GOARCH=wasm
go build -o bundle.wasm bundle.go
set GOOS=
set GOARCH=
go build -o wasm-rotating-cube.exe server/main.go
REM wasm-rotating-cube.exe
