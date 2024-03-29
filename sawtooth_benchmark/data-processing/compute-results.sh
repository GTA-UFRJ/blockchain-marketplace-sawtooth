#!bin/bash

if [ -f "./results.log" ]; then
	mv ./results.log ../final-results-$(date '+%F-%H-%M-%S')
fi

#echo "one org:"
#cp ../scripts/results/one-org-*/* .
#python3 ./compute-results.py one-org
#rm ./*-time-*

echo "batch:"
cp ../scripts/results/batch-*/* .
python3 ./compute-results.py batch
rm ./*-time-*

echo "poet scalab:"
cp ../scripts/results/poet-scalab-*/* .
python3 ./compute-results.py poet-scalab
rm ./*-time-*

echo "pbft scalab:"
cp ../scripts/results/pbft-scalab-*/* .
python3 ./compute-results.py pbft-scalab
rm ./*-time-*

echo "serial:"
cp ../scripts/results/serial-*/* .
python3 ./compute-results.py serial
rm ./*-time-*

echo "parallel:"
cp ../scripts/results/parallel-*/* .
python3 ./compute-results.py parallel
rm ./*-time-*

cp ./results.log ..
mv ../results.log ../final-results-$(date '+%F-%H-%M-%S')
rm ./results.log
