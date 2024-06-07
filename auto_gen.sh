#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <nama feature>"
    exit 1
fi

feature_name=$1

mkdir -p "./feature/$feature_name/dto"
mkdir -p "./feature/$feature_name/handler"
mkdir -p "./feature/$feature_name/repository"
mkdir -p "./feature/$feature_name/service"

# Membuat file-filenya
touch "./feature/$feature_name/dto/req.go"
touch "./feature/$feature_name/dto/res.go"
touch "./feature/$feature_name/handler/handler.go"
touch "./feature/$feature_name/interface.go"
touch "./feature/$feature_name/repository/repository.go"
touch "./feature/$feature_name/service/service.go"

echo "feature '$feature_name' berhasil dibuat!"