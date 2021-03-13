#!/bin/bash

rounds=10

#faz backup dos resultados atuais antes de pegar os proximos
if [ "$(ls -A ./scripts/results/)" ]; then
	cp -rf ./scripts/results/* ./backups/
	rm -rf ./scripts/results/*
fi
bash ./poet-scalab-results_3.sh $rounds
bash ./pbft-scalab-results_3.sh $rounds
bash ./one-org-results.sh $rounds
#bash ./data-processing/compute-results.sh
#cp ./data-processing/results.log ..
#mv ./results.log ./final-results-$(date '+%F-%H-%M-%S')
