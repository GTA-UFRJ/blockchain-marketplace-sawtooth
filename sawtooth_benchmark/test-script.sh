#!/bin/bash

cmd=()

cmd[0]="/binary/v2/autavail-go-v2 register 000001 --url="http://sawtooth-rest-api-default-0:8008""
cmd[1]="/binary/v2/autavail-go-v2 register 000002 --url="http://sawtooth-rest-api-default-0:8008""
cmd[2]="/binary/v2/autavail-go-v2 register 000003 --url="http://sawtooth-rest-api-default-0:8008""
cmd[3]="/binary/v2/autavail-go-v2 register 000004 --url="http://sawtooth-rest-api-default-0:8008""
cmd[4]="/binary/v2/autavail-go-v2 register 000005 --url="http://sawtooth-rest-api-default-0:8008""
cmd[5]="/binary/v2/autavail-go-v2 register 000006 --url="http://sawtooth-rest-api-default-0:8008""
cmd[6]="/binary/v2/autavail-go-v2 register 000007 --url="http://sawtooth-rest-api-default-0:8008""
cmd[7]="/binary/v2/autavail-go-v2 register 000008 --url="http://sawtooth-rest-api-default-0:8008""

cmd[8]="/binary/v2/autavail-go-v2 advert 100 000001 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[9]="/binary/v2/autavail-go-v2 advert 100 000002 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[10]="/binary/v2/autavail-go-v2 advert 100 000003 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[11]="/binary/v2/autavail-go-v2 advert 100 000004 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[12]="/binary/v2/autavail-go-v2 advert 100 000005 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[13]="/binary/v2/autavail-go-v2 advert 100 000006 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[14]="/binary/v2/autavail-go-v2 advert 100 000007 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""
cmd[15]="/binary/v2/autavail-go-v2 advert 100 000008 title description 10.0.0.1 datatype --url="http://sawtooth-rest-api-default-0:8008""

cmd[16]="/binary/v2/autavail-go-v2 list --url="http://sawtooth-rest-api-default-0:8008""


for index in $(seq 0 7);
do
	docker exec sawtooth-shell-default-0 ${cmd[$index]}
        sleep 5
done
adverttxid=()
for index in $(seq 8 15);
do
        docker exec sawtooth-shell-default-0 ${cmd[$index]}
        sleep 5
        output=$(docker exec -t sawtooth-shell-default-0 ${cmd[16]} | head -n 1)
        adverttxid[$(($index-8))]=${output:7:16}
done
echo $adverttxid
