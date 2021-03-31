#!/bin/bash

transaction=$1
adverttxid=$2
for transaction in $(seq 1 $transaction); do
	/binary/v2/autavail-go-v2 buy $adverttxid 123456 10.0.0.1 123456 --url="file"
done
