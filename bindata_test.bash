#!/bin/bash

go-bindata -o="predictor/bindata_test.go" -pkg="predictor" "test/"
go-bindata -o="server/bindata_test.go" -pkg="server" "test/"
