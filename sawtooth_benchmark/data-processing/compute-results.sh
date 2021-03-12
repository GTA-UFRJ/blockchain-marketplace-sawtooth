#!bin/bash

echo "one org:"
cp ./scripts/results/one-org-*/* ./data-processing
python3 ./data-processing/compute-results.py one-org
rm ./data-processing/*-time-*

echo "poet scalab:"
cp ./scripts/results/poet-scalab-*/* ./data-processing
python3 ./data-processing/compute-results.py poet-scalab
rm ./data-processing/*-time-*

echo "pbft scalab:"
cp ./scripts/results/pbft-scalab-*/* ./data-processing
python3 ./data-processing/compute-results.py pbft-scalab
rm ./data-processing/*-time-*
