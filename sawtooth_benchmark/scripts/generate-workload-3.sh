#!/bin/bash

transaction=$1
txperbatch=$2
batches=$((transaction / txperbatch))
for batch in $(seq 1 $batches); do
	/binary/v2/autavail-go-v2 batch 123456 $2 --url="file"
done
