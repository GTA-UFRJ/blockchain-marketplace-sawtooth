#!/bin/bash

VALIDATED_TRANSACTIONS=0
TIME_CONTROL=0
PREVIOUS_TRANSACTION_NUMBER=0
while [ $VALIDATED_TRANSACTIONS -lt $2 ]
do
	VALIDATED_TRANSACTIONS=$(( $(docker exec -t sawtooth-shell-default-0 sawtooth transaction list --url http://sawtooth-rest-api-default-0:8008 | wc -l) - 8 ))
	FLAG_WAIT_TIME=0
	if [ $VALIDATED_TRANSACTIONS -eq $PREVIOUS_TRANSACTION_NUMBER ]; then 
		TIME_CONTROL=$((TIME_CONTROL+1))
		FLAG_WAIT_TIME=1
		if [ $TIME_CONTROL -eq 10 ]; then
			echo "Number of transactions read:"$PREVIOUS_TRANSACTION_NUMBER >> date '+%M %s %N' >> .$1/final-time-txperbatch-$3-transaction-$2-round-$4
			break
		fi
	fi
	PREVIOUS_TRANSACTION_NUMBER=$VALIDATED_TRANSACTIONS
	TIME_CONTROL=$((TIME_CONTROL*FLAG_WAIT_TIME))
	#docker exec -t sawtooth-validator-default-0 du /var/lib/sawtooth >> ./blockchain-size
	sleep 1
done
date '+%M %s %N' >> .$1/final-time-txperbatch-$3-transaction-$2-round-$4
#sleep 3
#docker exec -t sawtooth-validator-default-0 du /var/lib/sawtooth >> ./blockchain-size
#docker exec -t sawtooth-validator-default-0 cat bytes.log >> ./bytes.log 
#sleep 1
#echo "next" >> ./bytes.log
