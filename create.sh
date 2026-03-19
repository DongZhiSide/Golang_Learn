#!/bin/sh

# Create a new project
# Usage: ./create.sh <project_name>

# Check if the project name is provided
if [ -z "$1" ]; then
  echo "Please provide a project name."
  exit 1
fi

# Create the project directory
mkdir "$1"

# Create the main.go file
cd "$1"
touch "main.go"
go mod init "golanglearn/$1"