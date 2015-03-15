#!/bin/bash

base_package=github.com/mitsuse/kompl
base_path=${GOPATH}/src/${base_package}

package_list=(
    ${base_package}/binary
    ${base_package}/ngram
    ${base_package}/predictor
    ${base_package}/predictor/data
    ${base_package}/server
    ${base_package}/sentencizer
    ${base_package}/tokenizer
    ${base_package}/trie
)

if [ ! -d ${base_path}/coverprofile ]
then 
    mkdir ${base_path}/coverprofile
else
    rm ${base_path}/coverprofile/*.coverprofile
fi

for package in ${package_list[@]}
do
    cover_name=$(echo ${package} | sed -e "s/\//__/g").coverprofile
    cover_path=${base_path}/coverprofile/${cover_name}
    go test -covermode=count -coverprofile ${cover_path} ${package}
done

cd ${base_path}/coverprofile && gover
