#!/bin/bash

transaction=$1
adverttxid=$2
orgid=$3
echo "$1 $2 $3"
for transaction in $(seq 1 $transaction); do
	/binary/v3/autavail-go-v3 buy $adverttxid $orgid "slow" $orgid --url="file"
done
