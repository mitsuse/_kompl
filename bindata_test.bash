#!/bin/bash

go-bindata -o="predictor/bindata_test.go" -pkg="predictor" "predictor/test/" 
