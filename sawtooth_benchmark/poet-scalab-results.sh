#!/bin/bash

#path=/scripts/results
transaction=250
cmd="/binary/autavail-go register 123456 --url="http://sawtooth-rest-api-default-0:8008""

docker-compose -f docker-poet-org8.yaml down --remove-orphans -v >> /dev/null 2>&1
#docker-compose -f docker-poet-org8.yaml down --remove-orphans -v
sleep 10

#faz backup dos resultados atuais antes de pegar os proximos
#if [ "$(ls -A .$path)" ]; then
#	backup="./backups/backups-$(date '+%F-%H-%M-%S')"
#	mkdir $backup
#	mv .$path/* $backup
#fi

path="/scripts/results/poet-scalab-results-$(date '+%F-%H-%M-%S')"
mkdir .$path

for round in $(seq 1 15); 
do
	for org in 2 4 6 8 10; 
	do # 2 4 6 8 10 12 do

		#printf "\n round $round start for $i clis start"	
		# levantar a rede
		docker-compose -f docker-poet-org$org.yaml up -d >> /dev/null 2>&1
		#docker-compose -f docker-poet-org$org.yaml up -d
		
		#printf "\n mimiu"
		# dormir esperando a rede levantar
		sleep 10
		
		docker exec sawtooth-shell-default-0 $cmd
		sleep 1
	
		# gera workload
		if [ -f "./scripts/autavail.workload" ]; then
			rm ./scripts/autavail.workload
		fi
		if [ -f "./binary/autavail.workload" ]; then
            rm ./binary/autavail.workload
        fi
		docker exec sawtooth-shell-default-0 ./scripts/generate-workload.sh $transaction
		sleep 20
	
		# marcar o tempo 
		date '+%M %s %N' >> .$path/initial-time-org-$org-transaction-$transaction-round-$round
				
		# enviar transacoes
		docker exec sawtooth-shell-default-0 sawtooth batch submit -f autavail.workload --url http://sawtooth-rest-api-default-0:8008 &

		#printf "\n ta na hora do query"
		# consultar transações
		docker exec sawtooth-shell-default-0 ./scripts/query_2.sh $path $transaction $org $round
	
		#printf "\n mimiu again"
		# dormir esperando o resultado
		sleep $((25+$org))
	
		docker-compose -f docker-poet-org$org.yaml down -v --remove-orphans >> /dev/null 2>&1
		#docker-compose -f docker-poet-org$org.yaml down -v --remove-orphans >> /dev/null 2>&1
		sleep $((10+$org))
	done;
done
