#!bin/bash

if [ -f "./results.log" ]; then
	mv ./results.log ../final-results-$(date '+%F-%H-%M-%S')
fi

echo "one org:"
cp ../test/one-org-*/* .
python3 ./compute-results.py one-org
rm ./*-time-*

echo "poet scalab:"
cp ../test/poet-scalab-*/* .
python3 ./compute-results.py poet-scalab
rm ./*-time-*

echo "pbft scalab:"
cp ../test/pbft-scalab-*/* .
python3 ./compute-results.py pbft-scalab
rm ./*-time-*

cp ./results.log ../test
mv ../test/results.log ../test/final-results-$(date '+%F-%H-%M-%S')
rm ./results.log
