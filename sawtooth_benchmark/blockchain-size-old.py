# Packages
import statistics
import sys
from pathlib import Path
import re

def IncrementExperiment(experimentList):
    if (len(experimentList) == 0):
        expeirmentList.append(0)
        expeirmentList.append(0)
        expeirmentList.append(0)
        expeirmentList.append(0)
    else:
        experimentList[3] = experimentList[0] + 1

    if (experimentList[3] == 10):
        experimentList[2] = experimentList[2] + 1
        experimentList[3] = 0
    if (experimentList[2] == 2):
        experimentList[1] = experimentList[1] + 1
        experimentList[2] = 0
    if (experimentList[1] == 2):
        experimentList[0] = experimentList[0] + 1
        experimentList[1] = 0
    if (experimentList[0] == 2):
        return experimentList, "end"
    return experimentList, "continue"

def main():
    experimentList = []
    newExperiment = 0
    numberList[]
    f = open("blockchain-size","r")
    lines = f.readlines()
    for line in lines:
        if (line[0:4] == 'next'):
            experimentList, continueFlag = IncrementExperiment(experimentList)
            newExperiment = 1
        else:
            if (newExperiment == 1):
                numberList = []
                newExperiment = 0
            number = re.split('r\t+', line.rstrip('\t'))
            numberList.append(number[0])


if __name__ == "__main__":
        main()
