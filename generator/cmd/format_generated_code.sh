#!/bin/bash

go run generator.go >sample.go
go fmt sample.go
cat sample.go
