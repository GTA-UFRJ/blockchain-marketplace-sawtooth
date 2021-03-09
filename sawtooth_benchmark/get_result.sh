#!/bin/bash

path=/scripts/results
transaction=50
cmd="/binary/autavail-go register 123456 --url="http://sawtooth-rest-api-default-0:8008""

docker-compose -f docker-poet-1.yaml down --remove-orphans -v >> /dev/null 2>&1

for round in $(seq 1 10); 
do
	for client in 1 2 4 8 16 32; 
	do #4 8 16 32 64 do

		#printf "\n round $k start for $i clis start"	
		# levantar a rede
		docker-compose -f docker-poet-$client.yaml up -d >> /dev/null 2>&1
		
		#printf "\n mimiu"
		# dormir esperando a rede levantar
		sleep 10
		

		docker exec sawtooth-shell-default-0 $cmd &
		sleep 5
	
		# marcar o tempo 
		date '+%M %s %N' >> ./scripts/results/initial-time-client-$client-transaction-$transaction-round-$round
		
		#printf "\n executando transacoes"
		# executar o script
		for i in $(seq 0 $(($client-1)));
		do
			docker exec sawtooth-shell-default-$i ./scripts/send_transactions.sh $path $transaction $client & 
	
		done
		
		#printf "\n ta na hora do query"

		# consultar transações
		docker exec sawtooth-shell-default-0 ./scripts/query.sh $path $transaction $client $round  
	
		#printf "\n mimiu again"
		# dormir esperando o resultado
		sleep 10
	
		docker-compose -f docker-poet-$client.yaml down -v --remove-orphans >> /dev/null 2>&1
	done;
done
