#!/bin/bash

rounds=$1
file=$2

#path=/scripts/results
transaction=$3

cmd=()

cmd[0]="/binary/v3/autavail-go-v3 register 000001 --url="http://sawtooth-rest-api-default-0:8008""
cmd[1]="/binary/v3/autavail-go-v3 register 000002 --url="http://sawtooth-rest-api-default-0:8008""
cmd[2]="/binary/v3/autavail-go-v3 register 000003 --url="http://sawtooth-rest-api-default-0:8008""
cmd[3]="/binary/v3/autavail-go-v3 register 000004 --url="http://sawtooth-rest-api-default-0:8008""
cmd[4]="/binary/v3/autavail-go-v3 register 000005 --url="http://sawtooth-rest-api-default-0:8008""
cmd[5]="/binary/v3/autavail-go-v3 register 000006 --url="http://sawtooth-rest-api-default-0:8008""
cmd[6]="/binary/v3/autavail-go-v3 register 000007 --url="http://sawtooth-rest-api-default-0:8008""
cmd[7]="/binary/v3/autavail-go-v3 register 000008 --url="http://sawtooth-rest-api-default-0:8008""

cmd[8]="/binary/v3/autavail-go-v3 advert 2 000001 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[9]="/binary/v3/autavail-go-v3 advert 2 000002 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[10]="/binary/v3/autavail-go-v3 advert 2 000003 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[11]="/binary/v3/autavail-go-v3 advert 2 000004 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[12]="/binary/v3/autavail-go-v3 advert 2 000005 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[13]="/binary/v3/autavail-go-v3 advert 2 000006 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[14]="/binary/v3/autavail-go-v3 advert 2 000007 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[15]="/binary/v3/autavail-go-v3 advert 2 000008 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""

cmd[16]="/binary/v3/autavail-go-v3 list --url="http://sawtooth-rest-api-default-0:8008""

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

for groups in 1 2 4 8;
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

	for index in $(seq 0 7);
	do
		docker exec sawtooth-shell-default-0 ${cmd[$index]}
        	sleep 5
	done
	
	#adverttxid=("0" "0" "0" "0" "0" "0" "0" "0")
	#for index in $(seq 8 15);
	#do	
	#	adverttxid[$(($index-8))]=$(docker exec -t sawtooth-shell-default-0 ${cmd[$index]})
	#	echo ${adverttxid[@]}
	#	sleep 5
	#done
	#docker exec -t sawtooth-shell-default-0 ${cmd[16]}
	
	adverttxid1=$(docker exec -t sawtooth-shell-default-0 ${cmd[8]})
	sleep 2
	adverttxid2=$(docker exec -t sawtooth-shell-default-0 ${cmd[9]})
	sleep 2
	adverttxid3=$(docker exec -t sawtooth-shell-default-0 ${cmd[10]})
	sleep 2
	adverttxid4=$(docker exec -t sawtooth-shell-default-0 ${cmd[11]})
	sleep 2
	adverttxid5=$(docker exec -t sawtooth-shell-default-0 ${cmd[12]})
	sleep 2
	adverttxid6=$(docker exec -t sawtooth-shell-default-0 ${cmd[13]})
	sleep 2
	adverttxid7=$(docker exec -t sawtooth-shell-default-0 ${cmd[14]})
	sleep 2
	adverttxid8=$(docker exec -t sawtooth-shell-default-0 ${cmd[15]})
	sleep 2



	if [ $groups -gt 0 ]; then
		docker exec -t sawtooth-shell-default-0 /scripts/generate-workload-5.sh $(($transaction/$groups)) ${adverttxid1:0:16} 000001
	fi
	if [ $groups -gt 1 ]; then
		docker exec -t sawtooth-shell-default-0 /scripts/generate-workload-5.sh $(($transaction/$groups)) ${adverttxid2:0:16} 000002
	fi
	if [ $groups -gt 3 ]; then
		docker exec -t sawtooth-shell-default-0 /scripts/generate-workload-5.sh $(($transaction/$groups)) ${adverttxid3:0:16} 000003
		docker exec -t sawtooth-shell-default-0 /scripts/generate-workload-5.sh $(($transaction/$groups)) ${adverttxid4:0:16} 000004
	fi
	if [ $groups -gt 5 ]; then
		docker exec -t sawtooth-shell-default-0 /scripts/generate-workload-5.sh $(($transaction/$groups)) ${adverttxid5:0:16} 000005
		docker exec -t sawtooth-shell-default-0 /scripts/generate-workload-5.sh $(($transaction/$groups)) ${adverttxid6:0:16} 000006
	fi
	if [ $groups -gt 7 ]; then
		docker exec -t sawtooth-shell-default-0 /scripts/generate-workload-5.sh $(($transaction/$groups)) ${adverttxid7:0:16} 000007
		docker exec -t sawtooth-shell-default-0 /scripts/generate-workload-5.sh $(($transaction/$groups)) ${adverttxid8:0:16} 000008
	fi

    	#gera workload
	#for index in $(seq 0 $(($groups-1)));
	#do
	#	docker exec sawtooth-shell-default-0 /scripts/generate-workload-5.sh $(( $transaction / $(($index+1)) )) ${adverttxid[$index]} "00000$(($index+1))"
	#done
	#docker exec sawtooth-shell-default-0 mv autavail.workload autavail.workload-1
	#docker logs sawtooth-validator-default-0

    	# marcar o tempo
	date '+%M %s %N' >> .$path/initial-time-txperbatch-$groups-transaction-$transaction-round-$round

    	#  envia transacoes
	./scripts/query-5.sh $path $transaction $groups $round &
   	docker exec sawtooth-shell-default-0 sawtooth batch submit -f autavail.workload --url http://sawtooth-rest-api-default-0:8008 &

	#printf "\n ta na hora do query"
	# consultar transações
	#docker exec sawtooth-shell-default-0 /scripts/query-3.sh $path $transaction 1 $round &
	
	#printf "\n mimiu again"
	# dormir esperando o resultado
	sleep 60
	docker logs sawtooth-validator-default-0

	docker-compose -f $file down -v --remove-orphans >> /dev/null 2>&1 &
	#docker-compose -f docker-poet-1.yaml down -v --remove-orphans
	sleep 30
	done
done
