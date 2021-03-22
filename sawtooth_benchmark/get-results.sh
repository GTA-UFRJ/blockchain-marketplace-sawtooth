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
echo "Starting clients scalability test - 1 tx/batch"
bash ./one-org-results-2.sh $rounds 1
echo "Starting clients scalability test - 2 tx/batch"
bash ./one-org-results-2.sh $rounds 2
echo "Starting clients scalability test - 5 tx/batch"
bash ./one-org-results-2.sh $rounds 5
echo "Starting clients scalability test - 10 tx/batch"
bash ./one-org-results-2.sh $rounds 10

#bash ./data-processing/compute-results.sh
#cp ./data-processing/results.log ..
#mv ./results.log ./final-results-$(date '+%F-%H-%M-%S')
