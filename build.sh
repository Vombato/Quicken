#!/usr/bin/env bash

NAMESPACE=github.com/matteojoliveau/quicken
MODULES="cmd/ modules/ recipe/ utils/ resources/"
WORKDIR=build/go/src
sep="-----------------------------------"
echo ${sep}
echo ${sep}
echo "Building Quicken!"
echo ${sep}
echo ${sep}

printf "Creating Workdir...\n\n"

mkdir -p ${WORKDIR}/${NAMESPACE}
GOPATH=$(pwd)/build/go
printf "Using GOPATH: ${GOPATH}\n"
printf "Copying project modules...\n\n"
cp -r ${MODULES} ${WORKDIR}/${NAMESPACE}
if [ "$1" != "-d" ]; then #debug mode
    printf "Getting dependencies...\n\n"
    env GOPATH=${GOPATH} go get ./...
else
    printf "DEBUG MODE: Skipping dependencies download\n"
fi
wait $!
printf "Building project...\n\n"
env GOPATH=${GOPATH} go build -o build/qk
wait $!
printf "Done!\n"
chmod +x build/
