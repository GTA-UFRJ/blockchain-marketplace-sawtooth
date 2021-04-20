# Packages
import statistics
import sys
from pathlib import Path
import re

def main():
    experimentCount = 1
    experimentList = []
    
    f = open("blockchain-size","r")
    lines = f.readlines()
    for line in lines:
        if (line[0:4] == 'next'):
            experimentCount += 1
        else:
            number = re.split('r\t+', line.rstrip('\t'))
            numberList.append(number[0])


if __name__ == "__main__":
        main()
