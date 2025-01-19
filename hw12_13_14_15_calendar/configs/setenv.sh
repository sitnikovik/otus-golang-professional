#!/bin/bash

if [ $1 == "" ]; then
    echo "ERROR: Please provide the env file path"
    exit 1
fi

echo "Using env config file: $1"
. $1

# parse all env vars from the env file and export them
export $(grep -v '^#' $1 | xargs)