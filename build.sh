#!/bin/bash

go mod download
docker build -q --tag go-go-go:latest .
docker run -q --rm -it go-go-go:latest