#!/bin/sh
#
# Embed static files into library as go packages.
#
# See: github.com/rakyll/statik for details
#

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

echo "create package 'statik' under internal/names containing the static file data from /data/names folder"
statik -src ${DIR}/data/names -dest internal/names -f

echo "done"
