#!/bin/sh

cd ../db/migrations
goose mysql "mock-user:password@tcp(mock-mysql:3306)/mock-db?charset=utf8mb4&parseTime=true" up
