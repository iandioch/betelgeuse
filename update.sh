#!/bin/bash

cd ~/betelgeuse/
git fetch
git pull
go run generate.go
