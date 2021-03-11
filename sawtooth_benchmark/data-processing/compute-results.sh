#!bin/bash

cp ../scripts/results/one-org-results-*/* .
python3 ./compute-results.py one-org
rm ./*-time-*

cp ../scripts/results/poet-scalab-results-*/* .
python3 ./compute-results.py poet-scalab
rm ./*-time-*

cp ../scripts/results/pbft-scalab-results-*/* .
python3 ./compute-results.py pbtt-scalab
rm ./*-time-*
