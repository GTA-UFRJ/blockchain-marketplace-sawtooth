# Packages
import statistics
import sys

# Import global configuration parameters
if(sys.argv[1]=="one-org"):
	import one-org-config
	ROUNDS = one-org-config.rounds
	ENTITY_LIST = one-org-config.entityList
	ENTITY_TYPE = one-org-config.entityType
	TRANSACTIONS = one-org-config.transactions

if(sys.argv[1]=="poet-scalab"):
	import poet-scalab-config
	ROUNDS = poet-scalab.rounds
	ENTITY_LIST = poet-scalab.entityList
	ENTITY_TYPE = poet-scalab.entityType
	TRANSACTIONS = poet-scalab.transactions

if(sys.argv[1]=="pbtt-scalab"):
	import pbft-scalab-config
	ROUNDS = pbft-scalab.rounds
	ENTITY_LIST = pbft-scalab.entityList
	ENTITY_TYPE = pbft-scalab.entityType
	TRANSACTIONS = pbft-scalab.transactions

def CalculateRoundThrowput (roundCount, entity):

	# Read time log file content
	initialTimeFileName = "initial-time-"+ENTITY_TYPE+"-"entity+"-transaction-"+TRANSACTIONS+"-round-"+roundCount
	f=open(initialTimeFileName,"r")
	initialTimeFileLines = f.readlines()
	f.close()
	
	finalTimeFileName = "final-time-"+ENTITY_TYPE+"-"entity+"-transaction-"+TRANSACTIONS+"-round-"+roundCount
	f=open(finalTimeFileName,"r")
	finalTimeFileLines = f.readlines()
	f.close()
	
	# Verify if the number of transactions read is diferent
	if (len(finalTimeFileLines)==2):
		realTransaction = finalTimeFileLines[0].split(":")[1]
		initialTime = initialTimeFileLines[0].split(" ")[1]
		finalTime = finalTimeFileLines[1].split(" ")[1] - 10
	else:
		realTransaction = TRANSACTIONS
		initialTime = initialTimeFileLines[0].split(" ")[1]
		finalTime = finalTimeFileLines[0].split(" ")[1]

	# Calculate throwput
	throwput = realTransaction / (finalTime - initialTime)
	return round(throwput,2)

def LoopThrowRounds (entity):

	# Loop throw ROUNDS
	throwputsCalculations = []
	for roundCount in range ROUNDS:
		throwputsCalculations.append(CalculateRoundThrouwput(roundCount, entity))

	# Calculate mean and standard deviation of throwput for certein number of entities
	meanThrowput = statistics.mean(throwputsCalculations)
	stdevThrowput = statistics.pstdev(throwputsCalculations)
	return meanThrowput, stdevThrowput
	
def main ():

	# File to append
	f = open("results.log","a+")
	f.write("-------------------------------")
	
	# Loop throw entities quantity (number of clients or organizations)
	resultsLineList = [none] * len(ENTITY_LIST)
	for entity in ENTITY_LIST:
		entityMeanThrowput, entityStdevThrowput = LoopThrowRounds(entity)
		resultsLineList[entity] = "Entity "+entity+" = "+entityMeanThrowput+" pm "+entityStdevThrowput+" txps"
		f.write(resultsLineList[entity])
	f.write("-------------------------------")
	f.close()

if __name__ == "__main__":
	main()
