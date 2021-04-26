#!/bin/bash

rounds=$1
file=$2

#path=/scripts/results
transaction=200
cmd1="/binary/v2/autavail-go-v2 register 123456 --url="http://sawtooth-rest-api-default-0:8008""

docker-compose -f $file down --remove-orphans -v >> /dev/null 2>&1
#docker-compose -f docker-poet-1.yaml down --remove-orphans -v
sleep 10

#faz backup dos resultados atuais antes de pegar os proximos
#if [ "$(ls -A .$path)" ]; then
#	backup="./backups/poet-scalab-results-$(date '+%F-%H-%M-%S')"
#	mkdir $backup
#	mv .$path/* $backup
#fi

path="/scripts/results/serial-$(date '+%F-%H-%M-%S')"
mkdir .$path

for round in $(seq 1 $rounds); 
do
	echo "Round $round"
    #printf "\n round $round start for $i clis start"
    # levantar a rede
    docker-compose -f $file up -d >> /dev/null 2>&1 &
    #docker-compose -f docker-poet-1.yaml up -d

	#printf "\n mimiu"
	# dormir esperando a rede levantar
	sleep 20
		
	docker exec sawtooth-shell-default-0 $cmd1
	#sleep 5
	#docker exec sawtooth-shell-default-0 $cmd2
	#sleep 5
	#output=$(docker exec -t sawtooth-shell-default-0 $cmd3 | head -n 1)
	#adverttxid=${output:7:16}

    #gera workload
   	docker exec sawtooth-shell-default-0 /scripts/generate-workload.sh $transaction
	docker exec sawtooth-shell-default-0 mv autavail.workload autavail.workload-1
        docker exec sawtooth-shell-default-0 /scripts/generate-workload.sh $transaction
        docker exec sawtooth-shell-default-0 mv autavail.workload autavail.workload-2
        docker exec sawtooth-shell-default-0 /scripts/generate-workload.sh $transaction
        docker exec sawtooth-shell-default-0 mv autavail.workload autavail.workload-3
        docker exec sawtooth-shell-default-0 /scripts/generate-workload.sh $transaction
        docker exec sawtooth-shell-default-0 mv autavail.workload autavail.workload-4

	#docker logs sawtooth-validator-default-0

    # marcar o tempo
    date '+%M %s %N' >> .$path/initial-time-txperbatch-1-transaction-$(($transaction*4))-round-$round

    #  envia transacoes
    	./scripts/query-4.sh $path $(($transaction*4)) 1 $round &
   	docker exec sawtooth-shell-default-0 sawtooth batch submit -f autavail.workload-1 --url http://sawtooth-rest-api-default-0:8008
	sleep 1
        docker exec sawtooth-shell-default-0 sawtooth batch submit -f autavail.workload-2 --url http://sawtooth-rest-api-default-0:8008
        sleep 1
        docker exec sawtooth-shell-default-0 sawtooth batch submit -f autavail.workload-3 --url http://sawtooth-rest-api-default-0:8008
        sleep 1
        docker exec sawtooth-shell-default-0 sawtooth batch submit -f autavail.workload-4 --url http://sawtooth-rest-api-default-0:8008 &

	#printf "\n ta na hora do query"
	# consultar transações
	#docker exec sawtooth-shell-default-0 /scripts/query-3.sh $path $transaction 1 $round &
	
	#printf "\n mimiu again"
	# dormir esperando o resultado
	sleep 80
	docker logs sawtooth-validator-default-0
	
	docker-compose -f $file down -v --remove-orphans >> /dev/null 2>&1 &
	#docker-compose -f docker-poet-1.yaml down -v --remove-orphans
	sleep 30
done
