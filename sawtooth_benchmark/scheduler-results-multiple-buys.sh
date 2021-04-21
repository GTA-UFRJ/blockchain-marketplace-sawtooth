#!/bin/bash

rounds=$1
file=$2

#path=/scripts/results
transaction=400
cmd1="/binary/v2/autavail-go-v2 register 000001 --url="http://sawtooth-rest-api-default-0:8008""
cmd2="/binary/v2/autavail-go-v2 register 000002 --url="http://sawtooth-rest-api-default-0:8008""
cmd3="/binary/v2/autavail-go-v2 register 000003 --url="http://sawtooth-rest-api-default-0:8008""
cmd4="/binary/v2/autavail-go-v2 register 000004 --url="http://sawtooth-rest-api-default-0:8008""
cmd5="/binary/v2/autavail-go-v2 register 000005 --url="http://sawtooth-rest-api-default-0:8008""
cmd6="/binary/v2/autavail-go-v2 register 000006 --url="http://sawtooth-rest-api-default-0:8008""
cmd7="/binary/v2/autavail-go-v2 register 000007 --url="http://sawtooth-rest-api-default-0:8008""
cmd8="/binary/v2/autavail-go-v2 register 000008 --url="http://sawtooth-rest-api-default-0:8008""

cmd9="/binary/v2/autavail-go-v2 advert 100 000001 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd10="/binary/v2/autavail-go-v2 advert 100 000002 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd11="/binary/v2/autavail-go-v2 advert 100 000003 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd12="/binary/v2/autavail-go-v2 advert 100 000004 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd13="/binary/v2/autavail-go-v2 advert 100 000005 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd14="/binary/v2/autavail-go-v2 advert 100 000006 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd15="/binary/v2/autavail-go-v2 advert 100 000007 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd16="/binary/v2/autavail-go-v2 advert 100 000008 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""

cmd17="/binary/v2/autavail-go-v2 list --url="http://sawtooth-rest-api-default-0:8008""

docker-compose -f $file down --remove-orphans -v >> /dev/null 2>&1
#docker-compose -f docker-poet-1-serial.yaml down --remove-orphans -v >> /dev/null 2>&1
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

for groups in 1 2 4 6 8;
do
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
	for index in $(seq 1 8);
	do	
		docker exec sawtooth-shell-default-0 $cmd$index
		sleep 5
	done
	adverttxid=()
	for index in $(seq 9 16);
        do
                docker exec sawtooth-shell-default-0 $cmd$index
                sleep 5
		output=$(docker exec -t sawtooth-shell-default-0 $cmd17 | head -n 1)
		adverttxid[$(($index-9))]=${output:7:16}
        done
	
	output=$(docker exec -t sawtooth-shell-default-0 $cmd3 | head -n 1)
	adverttxid=${output:7:16}

    #gera workload
   	docker exec sawtooth-shell-default-0 /scripts/generate-workload-5.sh $transaction $adverttxid $pac
	docker exec sawtooth-shell-default-0 mv autavail.workload autavail.workload-1
	docker exec sawtooth-shell-default-0 /scripts/generate-workload-5.sh $transaction $adverttxid 
        docker exec sawtooth-shell-default-0 mv autavail.workload autavail.workload-2
	docker exec sawtooth-shell-default-0 /scripts/generate-workload-5.sh $transaction $adverttxid
        docker exec sawtooth-shell-default-0 mv autavail.workload autavail.workload-3
	docker exec sawtooth-shell-default-0 /scripts/generate-workload-5.sh $transaction $adverttxid
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
done
