#!/bin/bash

# abort on errors
set -e

# build
./tasks/build-wasm.sh

# navigate into the build output directory
cd public

git init
git checkout -b main
git add -A
git commit -m 'deploy'

git push -f git@github.com:lucaspiller/orbital-go.git main:gh-pages

cd -