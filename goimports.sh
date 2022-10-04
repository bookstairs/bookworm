#!/bin/bash

# remove all blank lines in go 'imports' statements,
# then sort with goimports

# shellcheck disable=SC2032
imports() {
  gsed -i '
    /^import/,/)/ {
      /^$/ d
    }
  ' "$1"
  goimports -w -local github.com/bookstairs/bookworm "$1"
}

find . -name \*.go -not -path . > gofiles.txt

while IFS= read -r line
do
  imports "$line"
done < <(grep -v '^ *#' < gofiles.txt)

rm -rf gofiles.txt
