#!bin/bash

echo "one org:"
cp ../scripts/results/one-org-*/* .
python3 ./compute-results.py one-org
rm ./*-time-*

echo "poet scalab:"
cp ../scripts/results/poet-scalab-*/* .
python3 ./compute-results.py poet-scalab
rm ./*-time-*

echo "pbft scalab:"
cp ../scripts/results/pbft-scalab-*/* .
python3 ./compute-results.py pbft-scalab
rm ./*-time-*
