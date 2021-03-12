#!/bin/bash
#
#  Created by Gustavo Camilo
#
#
# Usage ./issueTimeIntructionTransaction NumberOfTransactions 


path=$1
transaction=$2
clients=$3
rate=$4

delay=$(( 1000 / $rate))

function issueTransaction() {
    { for transactionNumber in $(seq 1 $transaction);
    do
		initialTime=$(date +%s%N | cut -b1-13)
		while [ $(($current-$initialTime)) -lt $delay ]
		do
			current=$(date +%s%N | cut -b1-13)
		done
		#echo "ola"
        /binary/autavail-go advert 200 123456 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008" &
    done ; }  
}

issueTransaction 
printf "\n\n"
