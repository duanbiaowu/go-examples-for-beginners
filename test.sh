#!/usr/bin/env bash

echo "Test coverage running ....."
echo "" > coverage.txt

for d in $(go list ./...); do
  go test -coverprofile=profile.out -covermode=atomic "$d"
  if [ -f profile.out ]; then
      cat profile.out >> coverage.txt
      rm profile.out
  fi
done

# delete last profile.out
if [ -f profile.out ]; then
    rm profile.out
fi

echo "Test coverage finished."
