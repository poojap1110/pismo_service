#!/bin/sh

echo "Running Test.."
go test -coverprofile=cover.out -short -v

if [ -e "cover.out" ]; then
  echo "Running Cover.."
  go tool cover -func=cover.out

  rm -rf cover.out
fi

echo "Running Vet.."
go tool vet -all -v *.go
