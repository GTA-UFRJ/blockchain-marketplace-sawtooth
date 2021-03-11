#!bin/bash

cp ../scripts/results/one-org-*/* .
python3 ./compute-results.py one-org
rm ./*-time-*

cp ../scripts/results/poet-scalab-*/* .
python3 ./compute-results.py poet-scalab
rm ./*-time-*

cp ../scripts/results/pbft-scalab-*/* .
python3 ./compute-results.py pbtt-scalab
rm ./*-time-*
