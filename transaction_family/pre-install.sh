#!/bin/bash

if [ -d "$GOPATH/src/github.com/hyperledger/sawtooth-sdk-go/examples/autavail_go" ]
then
	rm -rf $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go/examples/autavail_go
fi
mkdir $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go/examples/autavail_go
cp -R * $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go/examples/autavail_go
cd $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go && docker build -f ./examples/autavail_go/Dockerfile-installed-bionic -t autavail_img .
