#!/bin/bash

clients=$3

VALIDATED_TRANSACTIONS=0
while [ $VALIDATED_TRANSACTIONS -lt $(($2*$clients)) ]
do
	VALIDATED_TRANSACTIONS=$(( $(sawtooth transaction list --url http://sawtooth-rest-api-default-0:8008 | wc -l) - 12 ))
	sleep 1
done
date '+%M %s %N' >> $1/sawtooth-timestamp-transactions-$2-client-$3-round-$4
