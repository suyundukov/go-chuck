#!/usr/bin/env bash
set -e

# Building main stuff (npm)
npm run build

# Minifying html files
minify -r -o ./public/static/css/ --match=\.css ./public/static/css/
minify -r -o ./public/static/js/ --match=\.js ./public/static/js/

# Removing unnecessary files to upload
rm -rf package.json README* LICENSE .travis* node_modules/
