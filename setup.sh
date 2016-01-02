#!/bin/bash

sudo apt-get install golang
go get gopkg.in/yaml.v1 # go get throws an error with v2 :(
go get github.com/robertkrimen/otto
go get github.com/russross/blackfriday