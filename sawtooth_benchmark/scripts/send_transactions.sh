#!/bin/bash
#
#  Created by Gustavo Camilo
#
#
# Usage ./issueTimeIntructionTransaction NumberOfTransactions 


path=$1
transaction=$2
clients=$3


function issueTransaction() {
    { for transactionNumber in $(seq 1 $transaction);
    do
        /binary/autavail-go advert 200 123456 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008"
		sleep 0.5
    done ; }  
}

issueTransaction 
printf "\n\n"
