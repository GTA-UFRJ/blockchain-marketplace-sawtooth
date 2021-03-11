#!/bin/bash

#faz backup dos resultados atuais antes de pegar os proximos
if [ "$(ls -A ./scripts/results/)" ]; then
   mv ./scripts/results/* ./backups/
fi
bash ./one-org-results.sh
bash ./poet-scalab-results.sh
bash ./pbft-scalab-results.sh
bash ./data-processing/compute-results.sh
mv ./data-processing/results.log ./final-results-$(date '+%F-%H-%M-%S')
