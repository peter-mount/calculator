#!/bin/sh
#
# Runs all tests
#
for bin in test
do
  echo "Testing ${bin}"
  go test -v github.com/peter-mount/calculator/${bin}
done
