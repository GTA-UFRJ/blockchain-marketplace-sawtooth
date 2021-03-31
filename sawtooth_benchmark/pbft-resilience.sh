#!/bin/bash

leader=$1

#path=/scripts/results
transaction=200
#cmd="/binary/autavail-go register 123456 --url="http://sawtooth-rest-api-default-0:8008""

path="/scripts/results/pbft-resilience-$(date '+%F-%H-%M-%S')"
mkdir .$path
			
#docker exec sawtooth-shell-default-0 $cmd
#sleep 1

# gera workload
docker exec sawtooth-shell-default-0 /scripts/generate-workload.sh $transaction

# derruba lider
docker stop sawtooth-validator-default-$leader --time 1
sleep 1

# marcar o tempo
date '+%M %s %N' >> .$path/initial-time-org-4-transaction-$transaction-round-1

# envia transacoes
docker exec sawtooth-shell-default-0 sawtooth batch submit -f autavail.workload --url http://sawtooth-rest-api-default-0:8008 &

# consultar transações
docker exec sawtooth-shell-default-0 /scripts/query_2.sh $path $transaction 4 1 &

# dormir esperando o resultado
sleep 10
