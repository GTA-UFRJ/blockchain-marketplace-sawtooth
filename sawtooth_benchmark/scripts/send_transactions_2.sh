#!/bin/bash
#
#  Created by Gustavo Camilo
#
#
# Usage ./issueTimeIntructionTransaction NumberOfTransactions 


path=$1
transaction=$2
clients=$3
batch=$4
rate=$5

txperbatch=$(($transaction/$batch))
delay=$(( 1000 / $rate))

function issueTransaction() {
    { for batchNumber in $(seq 1 $batch);
    do
		initialTime=$(date +%s%N | cut -b1-13)
		while [ $(($current-$initialTime)) -lt $delay ]
		do
			current=$(date +%s%N | cut -b1-13)
		done
		#echo "ola"
        /binary/v2/autavail-go batch $txperbatch --url="http://sawtooth-rest-api-default-0:8008" &
    done ; }  
}

issueTransaction 
printf "\n\n"
