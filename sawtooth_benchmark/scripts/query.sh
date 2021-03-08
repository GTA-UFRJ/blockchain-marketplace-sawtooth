#! /bin/sh

VALIDATED_INTKEY_TRANSACTIONS=0
while [ $VALIDATED_INTKEY_TRANSACTIONS -le $1 ]
do
	VALIDATED_INTKEY_TRANSACTIONS=$(( $(sawtooth transaction list --url http://sawtooth-rest-api-default-0:8008 | wc -l) - 6 ))
	sleep 1
done
date >> /tmp/sawtooth-timestamp
