#!/bin/sh

figlet Gomock Backend Service
reflex -r '(\.go|go\.mod)' -s go run main.go
