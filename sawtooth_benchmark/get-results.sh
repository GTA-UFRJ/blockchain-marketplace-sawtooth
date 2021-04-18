#!/bin/bash

exec &> logfile.txt

rounds=10
echo "$rounds rounds defined"

#faz backup dos resultados atuais antes de pegar os proximos
if [ "$(ls -A ./scripts/results/)" ]; then
	cp -rf ./scripts/results/* ./backups/
	rm -rf ./scripts/results/*
	echo "Backup of last results"
else
	echo "No last results found"
fi

echo "Starting serial scheduler test PBFT"
bash ./scheduler-results.sh $rounds docker-pbft-org5-serial.yaml
echo "Starting parallel scheduler test PBFT"
bash ./scheduler-results.sh $rounds docker-pbft-org5-parallel.yaml
echo "Starting serial scheduler test ad PBFT"
bash ./scheduler-results-ad.sh $rounds docker-pbft-org5-serial.yaml
echo "Starting parallel scheduler test ad PBFT"
bash ./scheduler-results-ad.sh $rounds docker-pbft-org5-parallel.yaml

#echo "Starting serial scheduler test PoET simulator"
#bash ./scheduler-results.sh $rounds docker-poet-org5-serial.yaml
#echo "Starting parallel scheduler test PoET simulator"
#bash ./scheduler-results.sh $rounds docker-poet-org5-parallel.yaml
#echo "Starting serial scheduler test ad PoET simulator"
#bash ./scheduler-results-ad.sh $rounds docker-poet-org5-serial.yaml
#echo "Starting parallel scheduler test ad PoET simulator"
#bash ./scheduler-results-ad.sh $rounds docker-poet-org5-parallel.yaml

#echo "Starting serial scheduler test PoET SGX"
#bash ./scheduler-results.sh $rounds docker-sgx-org5-serial.yaml
#echo "Starting parallel scheduler test PoET SGX"
#bash ./scheduler-results.sh $rounds docker-sgx-org5-parallel.yaml
#echo "Starting serial scheduler test ad PoET SGX"
#bash ./scheduler-results-ad.sh $rounds docker-sgx-org5-serial.yaml
#echo "Starting parallel scheduler test ad PoET SGX"
#bash ./scheduler-results-ad.sh $rounds docker-sgx-org5-parallel.yaml

#bash ./data-processing/compute-results.sh
#cp ./data-processing/results.log ..
#mv ./results.log ./final-results-$(date '+%F-%H-%M-%S')
