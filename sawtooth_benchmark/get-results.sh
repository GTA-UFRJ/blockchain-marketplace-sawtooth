#!/bin/bash

#faz backup dos resultados atuais antes de pegar os proximos
if [ "$(ls -A ./scripts/results/)" ]; then
	cp -rf ./scripts/results/* ./backups/
	rm -rf ./scripts/results/*
fi
bash ./one-org-results.sh
bash ./poet-scalab-results.sh
bash ./pbft-scalab-results.sh
bash ./data-processing/compute-results.sh
cp ./data-processing/results.log .
mv ./results.log ./final-results-$(date '+%F-%H-%M-%S')
