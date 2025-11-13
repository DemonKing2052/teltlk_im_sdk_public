#!/bin/bash

# swag_init.sh - Generate Swagger documentation for media-user-Server API

# Find all relevant directories and format them for swag
Model_PATH="./models"
API_DIRS=$(find ./handlers ./proto -type d | paste -sd ',' -)
DIR_STR="${API_DIRS},${Model_PATH}"

echo "Generating Swagger docs with directories: ${DIR_STR}"

# Run swag init
swag init -g ./main.go -d "./,${DIR_STR}"

# Check if swag init succeeded
if [ $? -eq 0 ]; then
    echo "Swagger documentation generated successfully"
else
    echo "Failed to generate Swagger documentation"
    exit 1
fi
