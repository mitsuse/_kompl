#!/bin/bash

go-bindata -o="predictor/bindata_test.go" -pkg="predictor" "test/"
go-bindata -o="bindata_test.go" -pkg="kompl" "test/"
