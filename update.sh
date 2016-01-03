#!/bin/bash

cd /var/www/betelgeuse/
git fetch
git pull
go run generate.go
cp styles/* site/styles/
