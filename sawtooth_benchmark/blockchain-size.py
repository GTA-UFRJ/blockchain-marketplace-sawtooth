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
    roundArraySize = len(preprocessesdMatrix[0])
    meanArray = []
    deviationArray = []
    cacheArray = []
    for index in range(roundArraySize):
        for roundArray in preprocessesdMatrix:
            cacheArray.append(roundArray[index])
        mean = statistics.mean(cacheArray)
        deviation = 0.62 * statistics.pstdev(cacheArray)
        meanArray.append(round(mean,2))
        deviationArray.append(round(deviation,2))
    #print(preprocessesdMatrix)
    return meanArray, deviationArray

def main():
    experimentCount = 0
    roundCount = 0
    resultsMatrix = []
    resultsMatrix.append([])
    f = open("blockchain-size","r")
    lines = f.readlines()
    index = 0
    for line in lines:
        if ((line[0:8] == 'Starting') or ((len(lines)-1) == index)):
            if (experimentCount != 0):
                #resultsMatrix.pop()
                mean, deviation = CalculateSize (resultsMatrix)
                print("Experiment ", experimentCount, ": ")
                print(mean)
                print(deviation)
                #print(resultsMatrix)
                resultsMatrix = []
                resultsMatrix.append([])
            experimentCount += 1
            roundCount = 1
        elif (line[0:4] == 'next'):
            roundCount += 1
            if ((len(lines)-1) != index):
                if (lines[index+1][0:8] != "Starting") :
                    resultsMatrix.append([])
        else:
            number = re.split(r'\t+', line.rstrip('\t'))[0]
            resultsMatrix[roundCount-1].append(int(number))
        #print(str(experimentCount), " ", str(roundCount))
        index += 1

if __name__ == "__main__":
    main()
