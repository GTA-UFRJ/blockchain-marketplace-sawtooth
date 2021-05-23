# Packages
import statistics
import sys
from pathlib import Path
import re

f = open("bytes-10.log","r")
lines = f.readlines()
for index in range(len(lines)):
    if ((index % 2) != 1):
        size = (len(lines[index]) - len(lines[index+1]))
        print(str(size))
f.close()
