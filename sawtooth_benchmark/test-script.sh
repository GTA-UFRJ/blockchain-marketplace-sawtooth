#!/bin/bash

docker exec -t sawtooth-shell-default-0 /binary/v3/autavail-go-v3 register 000001 --url="http://sawtooth-rest-api-default-0:8008"
sleep 5
docker exec -t sawtooth-shell-default-0 /binary/v3/autavail-go-v3 list --url="http://sawtooth-rest-api-default-0:8008"

#for i in "${adverttxid[@]}":
#do
#	i=$(docker exec -t sawtooth-shell-default-0 /binary/v3/autavail-go-v3 advert 100 000001 title description 10.0.0.1 tipo --url="http://sawtooth-rest-api-default-0:8008")
#	echo ${adverttxid[@]}
#	sleep 5
#done

adverttxid=("0000000000000000" "0000000000000000" "0000000000000000" "0000000000000000" "0000000000000000" "0000000000000000" "0000000000000000" "0000000000000000")
#adverttxid=()
for index in $(seq 8 15);
do
	id=$(docker exec -t sawtooth-shell-default-0 /binary/v3/autavail-go-v3 advert 100 000001 title description 10.0.0.1 tipo --url="http://sawtooth-rest-api-default-0:8008")
	adverttxid[$(($index-8))]=$id
	echo ${adverttxid[@]}
	sleep 5
done
docker exec -t sawtooth-shell-default-0 /binary/v3/autavail-go-v3 list --url="http://sawtooth-rest-api-default-0:8008"
