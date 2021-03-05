#!/bin/bash

if [ -f "$GOPATH/src/github.com/hyperledger/sawtooth-sdk-go/examples/autavail_go" ]
then
	rm -rf $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go/examples/autavail_go
fi
mkdir -p $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go /go/src/github.com/hyperledger/sawtooth-sdk-go/examples/autavail_go
mv * $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go /go/src/github.com/hyperledger/sawtooth-sdk-go/examples/autavail_go
cd $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go /go/src/github.com/hyperledger/sawtooth-sdk-go
