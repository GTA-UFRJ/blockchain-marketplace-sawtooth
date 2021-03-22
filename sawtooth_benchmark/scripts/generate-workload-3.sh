#!/bin/bash

batches=$(($2/$1))
for transaction in $(seq 1 $batches); do
	/binary/v2/autavail-go batch $2 --url="file"
done
