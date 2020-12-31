PROG_NAME := "salt256pem"
VERSION := "0.1"


build:
	GO111MODULE=off CGO_ENABLED=0 go build -o ${PROG_NAME} main.go
