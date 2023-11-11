#!/bin/bash

GOOS=js GOARCH=wasm go build -o public/orbital.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js public/