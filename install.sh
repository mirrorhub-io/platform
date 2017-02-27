#!/bin/bash

TMP_BUILD=/tmp/mirrorhub/pkg

DIST="freebsd openbsd windows linux"
ARCH="amd64 386 arm64"
ARM="5 6 7"

binary() {
  for os in $DIST; do
    for arch in $ARCH; do
	    mkdir -p ${TMP_BUILD}/${os}/${arch}
      echo "Build OS: ${os} Arch: ${arch}"
	    GOOS=$os GOARCH=$arch go build -o ${TMP_BUILD}/${os}/${arch}/mirrorhub .	
    done
    for av in $ARM; do
	    mkdir -p ${TMP_BUILD}/${os}/${arch}
      echo "Build OS: ${os} Arch: arm${av}"
	    GOOS=$os GOARCH=arm GOARM=$av \
        go build -o ${TMP_BUILD}/${os}/arm${av}/mirrorhub .	
    done
  done
}

binary
