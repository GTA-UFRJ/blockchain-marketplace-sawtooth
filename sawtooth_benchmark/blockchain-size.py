# Packages
import statistics
import sys
from pathlib import Path
import re

def MaxSizeSearch (resultsMatrix):
    sizeCache = 0
    for roundArray in resultsMatrix:
        roundSize = len (roundArray)
        if (roundSize > sizeCache):
            sizeCache = roundSize
    return sizeCache

def MatrixFill (resultsMatrix):
    maxRoundSize = MaxSizeSearch (resultsMatrix)
    for roundIndex in range(len(resultsMatrix)):
        roundArray = resultsMatrix[roundIndex]
        lastValue = roundArray[-1]

        # Array fill
        for index in range(len(roundArray),maxRoundSize):
            resultsMatrix[roundIndex].append(lastValue)
    
    return resultsMatrix

def CalculateSize (resultsMatrix):
    roundsAmount = len (resultsMatrix)
    preprocessesdMatrix = MatrixFill (resultsMatrix)
    
    # MISSING CODE
    
    print(preprocessesdMatrix)
    return 0, 0


def main():
    experimentCount = 0
    roundCount = 0
    resultsMatrix = []
    resultsMatrix.append([])
    f = open("blockchain-size-1","r")
    lines = f.readlines()
    for line in lines:
        if (line[0:8] == 'Starting'):
            if (experimentCount != 0):
                resultsMatrix.pop()
                mean, deviation = CalculateSize (resultsMatrix)
                print("Experiment ", experimentCount, ": ", str(round(mean,2)), " +- ", str(round(deviation,2)))
                # print(resultsMatrix)
                resultsMatrix = []
                resultsMatrix.append([])
            experimentCount += 1
            roundCount = 1
        elif (line[0:4] == 'next'):
            roundCount += 1
            resultsMatrix.append([])
        else:
            number = re.split(r'\t+', line.rstrip('\t'))[0]
            resultsMatrix[roundCount-1].append(int(number))

if __name__ == "__main__":
        main()
