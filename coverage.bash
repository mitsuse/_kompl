#!/bin/bash

base_package=github.com/mitsuse/kompl
base_path=${GOPATH}/src/${base_package}

package_list=(
    ${base_package}
    ${base_package}/binary
    ${base_package}/cmd/kompl
    ${base_package}/ngram
    ${base_package}/predictor
    ${base_package}/trie
)

if [ ! -d ${base_path}/coverprofile ]
then 
    mkdir ${base_path}/coverprofile
fi

for package in ${package_list[@]}
do
    cover_path=${base_path}/coverprofile/$(basename ${package}).coverprofile
    go test -coverprofile ${cover_path} ${package}
done


cd ${base_path}/coverprofile && gover
