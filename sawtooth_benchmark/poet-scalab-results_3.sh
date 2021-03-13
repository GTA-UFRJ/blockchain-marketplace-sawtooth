#!/bin/bash

rounds=$1

#path=/scripts/results
transaction=50
cmd="/binary/autavail-go register 123456 --url="http://sawtooth-rest-api-default-0:8008""

docker-compose -f docker-pbft-org10.yaml down --remove-orphans -v >> /dev/null 2>&1
#docker-compose -f docker-pbft-org8.yaml down --remove-orphans -v
sleep 10

#faz backup dos resultados atuais antes de pegar os proximos
#if [ "$(ls -A .$path)" ]; then
#	backup="./backups/poet-scalab-results-$(date '+%F-%H-%M-%S')"
#	mkdir $backup
#	mv .$path/* $backup
#fi

path="/scripts/results/poet-scalab-results-$(date '+%F-%H-%M-%S')"
mkdir .$path

for round in $(seq 1 $rounds); 
do
	for org in 2 4 6 8 10; 
	do # 2 4 6 8 10 12 do

		#printf "\n round $round start for $i clis start"	
		# levantar a rede
		docker-compose -f docker-pbft-org$org.yaml up -d >> /dev/null 2>&1 &
		#docker-compose -f docker-pbft-org$org.yaml up -d
		
		#printf "\n mimiu"
		# dormir esperando a rede levantar
		sleep 20
		
		docker exec sawtooth-shell-default-0 $cmd
		sleep 1

        # marcar o tempo
        date '+%M %s %N' >> .$path/initial-time-org-$org-transaction-$transaction-round-$round

		#printf "\n executando transacoes"
        # executar o script
        for i in $(seq 0 15);
        do
            sleep 0.1 ; docker exec sawtooth-shell-default-$i ./scripts/send_transactions.sh $path $transaction 0 2 &
        done

		#printf "\n ta na hora do query"
		# consultar transações
		docker exec sawtooth-shell-default-0 /scripts/query_2.sh $path $(($transaction*16)) $org $round &

		finish=$(ps aux | grep send_transactions.sh | wc -l)
        while [ $finish -gt 1 ]; do
            finish=$(ps aux | grep send_transactions.sh | wc -l)
        done
        echo "Finish: $(date '+%M %s %N')" >> .$path/initial-time-org-$org-transaction-$transaction-round-$round

		#printf "\n mimiu again"
		# dormir esperando o resultado
		sleep $((25+$org))
	
		docker-compose -f docker-pbft-org$org.yaml down -v --remove-orphans >> /dev/null 2>&1 &
		#docker-compose -f docker-pbft-org$org.yaml down -v --remove-orphans >> /dev/null 2>&1
		sleep $((10+$org))
	done;
done
