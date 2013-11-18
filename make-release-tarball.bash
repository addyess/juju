#!/bin/bash

# if someone invokes this with bash
set -e

unset GOPATH

# build release tarball from a bzr branch
DEFAULT_JUJU_CORE="lp:juju-core"


usage() {
    echo "usage: $0 REVNO [JUJU_CORE_BRANCH]"
    echo "  REVNO: The juju-core revno to build"
    echo "  JUJU_CORE_BRANCH: The juju-core branch; defaults to ${DEFAULT_JUJU_CORE}"
    exit 1
}


test $# -ge 1 ||  usage
REVNO=$1
JUJU_CORE_BRANCH=${2:-$DEFAULT_JUJU_CORE}
TMP_DIR=$(mktemp -d --tmpdir=$(pwd))
mkdir $TMP_DIR/RELEASE
WORK=$TMP_DIR/RELEASE

echo "Getting juju-core and all its dependencies."
GOPATH=$WORK go get -v -d launchpad.net/juju-core/...

echo "Setting juju-core tree to $JUJU_CORE_BRANCH $REVNO."
(cd "${WORK}/src/launchpad.net/juju-core/" &&
 bzr pull --no-aliases --remember --overwrite -r $REVNO $JUJU_CORE_BRANCH)

echo "Updating juju-core dependencies to the required versions."
GOPATH=$WORK go get -v launchpad.net/godeps
GOPATH=$WORK go install -v launchpad.net/godeps
GODEPS="$WORK/bin/godeps"
if [[ ! -f $GODEPS ]]; then
    echo "! Could not install godeps."
    ls $WORK
    ls $work/src/launchpad.net/godeps/
    ls $WORK/bin/
    exit 1
di
GOPATH=$WORK $GODEPS -u "${WORK}/src/launchpad.net/juju-core/dependencies.tsv"
# Remove godeps.
rm -r $WORK/bin

# Smoke test
GOPATH=$WORK go build -v launchpad.net/juju-core/...

# Change the generic release to the proper juju-core version.
VERSION=$(sed -n 's/^const version = "\(.*\)"/\1/p' \
    $WORK/src/launchpad.net/juju-core/version/version.go)
mv $WORK $TMP_DIR/juju-core_${VERSION}/

# Tar it up.
TARFILE=$(pwd)/juju-core_${VERSION}.tar.gz
cd $TMP_DIR
tar cfz $TARFILE --exclude .hg --exclude .git --exclude .bzr juju-core_${VERSION}

echo "release tarball: ${TARFILE}"
