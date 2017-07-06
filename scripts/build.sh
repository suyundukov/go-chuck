#!/usr/bin/env bash
set -e

# Building main stuff (npm)
npm run build:ci

# Minifying html files
minify -r -o ./public/static/css/ --match=\.css ./public/static/css/
minify -r -o ./public/static/js/ --match=\.js ./public/static/js/
# Minifying Chuck Norris' facts database files
minify -o ./data/db.json ./data/

# Removing unnecessary files to upload
rm -rf package.json README* LICENSE .travis* node_modules/
