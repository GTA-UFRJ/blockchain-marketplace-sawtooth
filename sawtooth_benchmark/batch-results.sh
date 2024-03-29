#!/bin/bash

rounds=$1

#path=/scripts/results
transaction=200
cmd="/binary/v2/autavail-go-v2 register 123456 --url="http://sawtooth-rest-api-default-0:8008""

docker-compose -f docker-poet-1.yaml down --remove-orphans -v >> /dev/null 2>&1
#docker-compose -f docker-poet-1.yaml down --remove-orphans -v
sleep 10

#faz backup dos resultados atuais antes de pegar os proximos
#if [ "$(ls -A .$path)" ]; then
#	backup="./backups/poet-scalab-results-$(date '+%F-%H-%M-%S')"
#	mkdir $backup
#	mv .$path/* $backup
#fi

path="/scripts/results/batch-$(date '+%F-%H-%M-%S')"
mkdir .$path

for round in $(seq 1 $rounds); 
do
	echo "Round $round"
	for txperbatch in 1 2 5 10 20 50 100 200;
	do	
		echo "$txperbatch txns per batch"
		#printf "\n round $round start for $i clis start"	
		# levantar a rede
		docker-compose -f docker-poet-1.yaml up -d >> /dev/null 2>&1 &
		#docker-compose -f docker-poet-1.yaml up -d
		
		#printf "\n mimiu"
		# dormir esperando a rede levantar
		sleep 20
		
		docker exec sawtooth-shell-default-0 $cmd
		sleep 5

	    #gera workload
    	docker exec sawtooth-shell-default-0 /scripts/generate-workload-3.sh $transaction $txperbatch

	    # marcar o tempo
    	date '+%M %s %N' >> .$path/initial-time-txperbatch-$txperbatch-transaction-$transaction-round-$round

	    #  envia transacoes
    	docker exec sawtooth-shell-default-0 sawtooth batch submit -f autavail.workload --url http://sawtooth-rest-api-default-0:8008 &

		#printf "\n ta na hora do query"
		# consultar transações
		docker exec sawtooth-shell-default-0 /scripts/query-3.sh $path $transaction $txperbatch $round &
	
		#printf "\n mimiu again"
		# dormir esperando o resultado
		sleep 25
	
		docker-compose -f docker-poet-1.yaml down -v --remove-orphans >> /dev/null 2>&1 &
		#docker-compose -f docker-poet-1.yaml down -v --remove-orphans
		sleep 30
	done;
done
