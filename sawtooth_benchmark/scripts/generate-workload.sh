#!/bin/bash

for transaction in $(seq 1 $1); do
	/binary/autavail-go advert 200 123456 title description 10.0.0.1 datatype --url="file"
done
