# Packages
import statistics
import sys
from pathlib import Path

# Import global configuration parameters
if(sys.argv[1]=="one-org"):
	ROUNDS = 10
	ENTITY_LIST = [1, 2, 4, 8, 16, 32]
	ENTITY_TYPE = "client"
	TRANSACTIONS = 50

if(sys.argv[1]=="poet-scalab"):
	ROUNDS = 10
	ENTITY_LIST = [2, 4, 6, 8, 10]
	ENTITY_TYPE = "org"
	TRANSACTIONS = 200

if(sys.argv[1]=="pbft-scalab"):
	ROUNDS = 10
	ENTITY_LIST = [2, 4, 6, 8, 10]
	ENTITY_TYPE = "org"
	TRANSACTIONS = 200

if(sys.argv[1]=="batch"):
	ROUNDS = 10
	ENTITY_LIST = [1, 2, 5, 10, 20, 50, 100, 200]
	ENTITY_TYPE = "txperbatch"
	TRANSACTIONS = 200

if(sys.argv[1]=="serial"):
    ROUNDS = 10
    ENTITY_LIST = [1]
    ENTITY_TYPE = "txperbatch"
    TRANSACTIONS = 300

if(sys.argv[1]=="parallel"):
    ROUNDS = 10
    ENTITY_LIST = [1]
    ENTITY_TYPE = "txperbatch"
    TRANSACTIONS = 300

def CalculateRoundThrowput (roundCount, entity):
	
	# Read time log file content
	initialTimeFileName = "initial-time-"+ENTITY_TYPE+"-"+str(entity)+"-transaction-"+str(TRANSACTIONS)+"-round-"+str(roundCount)
	if Path(initialTimeFileName).is_file():
		f=open(initialTimeFileName,"r")
		initialTimeFileLines = f.readlines()
		f.close()
	else:
		return -1
	
	finalTimeFileName = "final-time-"+ENTITY_TYPE+"-"+str(entity)+"-transaction-"+str(TRANSACTIONS)+"-round-"+str(roundCount)
	if Path(finalTimeFileName).is_file():
		f=open(finalTimeFileName,"r")
		finalTimeFileLines = f.readlines()
		f.close()
	else:
		return -1

	#initialTimeFileLines = ["23 123456 123456"]
	#finalTimeFileLines = ["23 123454 123456"]
	
	#initialTimeFileLines = ["23 123454 123456"]
	#finalTimeFileLines = ["Number of transactions read:40", "23 123456 123456"]

	# Verify if the number of transactions read is diferent
	if (len(finalTimeFileLines)==2):
		realTransaction = int(finalTimeFileLines[0].split(":")[1].split(" ")[0])
		initialTime = int(initialTimeFileLines[0].split(" ")[1])
		finalTime = int(finalTimeFileLines[1].split(" ")[1]) - 10
	else:
		if (ENTITY_TYPE=="org"):
			realTransaction = TRANSACTIONS
		elif (ENTITY_TYPE=="txperbatch"):
			realTransaction = TRANSACTIONS
		else:
			realTransaction = TRANSACTIONS * entity
		initialTime = int(initialTimeFileLines[0].split(" ")[1])
		finalTime = int(finalTimeFileLines[0].split(" ")[1])

	#realTransaction = 50
	#finalTime = 12345678
	#initialTime = 12345675
	
	# Calculate throwput
	# (TRANSACTIONS/5)
	# ((TRANSACTIONS * entity)/5)
	if (ENTITY_TYPE=="org"):
		if (realTransaction < (TRANSACTIONS/5)):
			return -1
	elif (ENTITY_TYPE=="txperbatch"):
		if (realTransaction < (TRANSACTIONS/5)):
			return -1
	else:
		if (realTransaction < ((TRANSACTIONS * entity)/5)):
			return -1
	return (realTransaction-1) / (finalTime - initialTime)
	#return (realTransaction) / (finalTime - initialTime)

def LoopThrowRounds (entity):

	# Loop throw ROUNDS
	throwputsCalculationsWithNegative = []
	for roundCount in range (ROUNDS):
		throwputsCalculationsWithNegative.append(CalculateRoundThrowput(roundCount+1, entity))

	# Eliminate negative values
	throwputsCalculations = throwputsCalculationsWithNegative[:]
	for item in throwputsCalculationsWithNegative:
		if item < 0:
			throwputsCalculations.remove(item)

	# Verify if there are no values to calculate
	if (len(throwputsCalculations) == 0):
		return 0, 0

	# Calculate mean and standard deviation of throwput for certein number of entities
	meanThrowput = statistics.mean(throwputsCalculations)
	stdevThrowput = statistics.pstdev(throwputsCalculations)
	return round(meanThrowput,2), round(stdevThrowput,2)
	
def main ():
	
	#print(ROUNDS)
	#print(ENTITY_LIST)
	#print(ENTITY_TYPE)
	#print(TRANSACTIONS)

	# File to append
	f = open("results.log","a+")
	header = "\n----------"+sys.argv[1]+"----------\n"
	f.write(header)
	
	# Loop throw entities quantity (number of clients or organizations)
	for entity in ENTITY_LIST:
		#entityMeanThrowput, entityStdevThrowput = 15.0, 3.1
		entityMeanThrowput, entityStdevThrowput = LoopThrowRounds(entity)
		resultsLine = "Entity "+str(entity)+" = ("+str(entityMeanThrowput)+" pm "+str(entityStdevThrowput)+") txns/sec\n"
		f.write(resultsLine)
	f.write("\n")
	f.close()

if __name__ == "__main__":
	main()
