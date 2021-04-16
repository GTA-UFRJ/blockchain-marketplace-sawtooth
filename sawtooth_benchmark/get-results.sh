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
#echo "Starting PoET scalability test"
#bash ./poet-scalab-results.sh $rounds
#echo "Starting PBFT scalability test"
#bash ./pbft-scalab-results.sh $rounds
echo "Starting serial scheduler test"
bash ./scheduler-results-serial.sh $rounds
echo "Starting parallel scheduler test"
bash ./scheduler-results-parallel.sh $rounds
echo "Starting serial scheduler test ad"
bash ./scheduler-results-serial-ad.sh $rounds
echo "Starting parallel scheduler test ad"
bash ./scheduler-results-parallel-ad.sh $rounds


#bash ./data-processing/compute-results.sh
#cp ./data-processing/results.log ..
#mv ./results.log ./final-results-$(date '+%F-%H-%M-%S')
