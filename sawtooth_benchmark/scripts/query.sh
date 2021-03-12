#!/bin/bash

clients=$3

VALIDATED_TRANSACTIONS=0
TIME_CONTROL=0
PREVIOUS_TRANSACTION_NUMBER=0
while [ $VALIDATED_TRANSACTIONS -lt $(($2*$clients)) ]
do
	VALIDATED_TRANSACTIONS=$(( $(sawtooth transaction list --url http://sawtooth-rest-api-default-0:8008 | wc -l) - 12 ))
	FLAG_WAIT_TIME=0
	if [ $VALIDATED_TRANSACTIONS -eq $PREVIOUS_TRANSACTION_NUMBER ]; then 
		TIME_CONTROL=$((TIME_CONTROL+1))
		FLAG_WAIT_TIME=1
		if [ $TIME_CONTROL -eq 10 ]; then
			if [ $2 -gt 200 ]; then
				echo "Number of transactions read:"$PREVIOUS_TRANSACTION_NUMBER >> $1/final-time-org-$3-transaction-$2-round-$4
			else
				echo "Number of transactions read:"$PREVIOUS_TRANSACTION_NUMBER >> date '+%M %s %N' >> $1/final-time-client-$3-transaction-$2-round-$4
			fi
			break
		fi
	fi
	PREVIOUS_TRANSACTION_NUMBER=$VALIDATED_TRANSACTIONS
	TIME_CONTROL=$((TIME_CONTROL*FLAG_WAIT_TIME))
	sleep 1
done
if [ $2 -gt 180 ]; then
	date '+%M %s %N' >> $1/final-time-org-$3-transaction-$2-round-$4
else
	date '+%M %s %N' >> $1/final-time-client-$3-transaction-$2-round-$4
fi
