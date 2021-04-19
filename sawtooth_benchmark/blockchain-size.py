# Packages
import statistics
import sys
from pathlib import Path

def main():
    f = open("blockchain-size","r")
    lines = f.readlines()
    for line in lines:
        if (line[0:4] == 'next'):
            print ("Next!")

if __name__ == "__main__":
        main()
