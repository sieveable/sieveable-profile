#!/bin/bash
mode=count
profile="coverage.cov"
# Generate coverage profile files for each package
for pkg in $(go list ./...)
do
  f="profile_${pkg##*/}.cov"
  go test -covermode="$mode" -coverprofile="$f" "$pkg"
  rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
done
echo "mode: $mode" >"$profile"
grep -h -v "^mode:" profile_*.cov >>"$profile"
rm profile_*.cov
