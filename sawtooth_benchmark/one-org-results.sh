#!/bin/bash

#path=/scripts/results
transaction=50
cmd="/binary/autavail-go register 123456 --url="http://sawtooth-rest-api-default-0:8008""

docker-compose -f docker-poet-64.yaml down --remove-orphans -v >> /dev/null 2>&1
sleep 20

#faz backup dos resultados atuais antes de pegar os proximos
#if [ "$(ls -A .$path)" ]; then
#	backup="./backups/backup-$(date '+%F-%H-%M-%S')"
#	mkdir $backup
#	mv .$path/* $backup
#fi

path="/scripts/results/one-org-results-$(date '+%F-%H-%M-%S')"
mkdir .$path

for round in $(seq 1 15); 
do
	for client in 1 2 4 8 16 32; 
	do #4 8 16 32 64 do

		#printf "\n round $k start for $i clis start"	
		# levantar a rede
		docker-compose -f docker-poet-$client.yaml up -d >> /dev/null 2>&1
		
		#printf "\n mimiu"
		# dormir esperando a rede levantar
		sleep 10
		
		docker exec sawtooth-shell-default-0 $cmd
		sleep 1
	
		# marcar o tempo 
		date '+%M %s %N' >> .$path/initial-time-client-$client-transaction-$transaction-round-$round
		
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
		sleep $((25 + $client))
	
		docker-compose -f docker-poet-$client.yaml down -v --remove-orphans >> /dev/null 2>&1
		sleep $((10 + $client))
	done;
done
