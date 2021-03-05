#!/bin/bash

docker run $1 cat /go/src/github.com/hyperledger/sawtooth-sdk-go/examples/autavail_go/bin/autavail-go > ./autavail-go
chmod +x ./autavail-go
mv ./autavail-go ../sawtooth_benchmark/binary
