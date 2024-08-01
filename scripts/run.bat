@echo off
REM Download Go module dependencies
go mod download

REM Run Go program
go run cmd/MultiSender/main.go
