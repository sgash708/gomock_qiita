#!/bin/sh

go test `go list ./... | grep -v server/mock` -coverprofile=../test_result/cover.out ./...  > ../test_result/coverage.txt
go tool cover -html=../test_result/cover.out -o ../test_result/cover.html
