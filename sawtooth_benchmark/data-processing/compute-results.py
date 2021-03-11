# Packages
import statistics
import sys

# Import global configuration parameters
if(sys.argv[1]=="one-org"):
	ROUNDS = 15
	ENTITY_LIST = [1, 2, 4, 8, 16, 32]
	ENTITY_TYPE = "client"
	TRANSACTIONS = 50

if(sys.argv[1]=="poet-scalab"):
	ROUNDS = 15
	ENTITY_LIST = [2, 4, 6, 8]
	ENTITY_TYPE = "org"
	TRANSACTIONS = 250

if(sys.argv[1]=="pbft-scalab"):
	ROUNDS = 15
	ENTITY_LIST = [2, 4, 6, 8]
	ENTITY_TYPE = "org"
	TRANSACTIONS = 250

def CalculateRoundThrowput (roundCount, entity):
	
	# Read time log file content
	initialTimeFileName = "initial-time-"+ENTITY_TYPE+"-"+str(entity)+"-transaction-"+str(TRANSACTIONS)+"-round-"+str(roundCount)
	f=open(initialTimeFileName,"r")
	initialTimeFileLines = f.readlines()
	f.close()
	
	finalTimeFileName = "final-time-"+ENTITY_TYPE+"-"+str(entity)+"-transaction-"+str(TRANSACTIONS)+"-round-"+str(roundCount)
	f=open(finalTimeFileName,"r")
	finalTimeFileLines = f.readlines()
	f.close()

	#initialTimeFileLines = ["23 123456 123456"]
	#finalTimeFileLines = ["23 123454 123456"]
	
	#initialTimeFileLines = ["23 123454 123456"]
	#finalTimeFileLines = ["Number of transactions read:40", "23 123456 123456"]

	# Verify if the number of transactions read is diferent
	if (len(finalTimeFileLines)==2):
		realTransaction = int(finalTimeFileLines[0].split(":")[1])
		initialTime = int(initialTimeFileLines[0].split(" ")[1])
		finalTime = int(finalTimeFileLines[1].split(" ")[1]) - 10
	else:
		realTransaction = TRANSACTIONS
		initialTime = int(initialTimeFileLines[0].split(" ")[1])
		finalTime = int(finalTimeFileLines[0].split(" ")[1])

	#realTransaction = 50
	#finalTime = 12345678
	#initialTime = 12345675
	
	# Calculate throwput
	throwput = realTransaction / (finalTime - initialTime)
	return round(throwput,2)

def LoopThrowRounds (entity):

	# Loop throw ROUNDS
	throwputsCalculations = []
	for roundCount in range (ROUNDS):
		throwputsCalculations.append(CalculateRoundThrowput(roundCount, entity))

	# Calculate mean and standard deviation of throwput for certein number of entities
	meanThrowput = statistics.mean(throwputsCalculations)
	stdevThrowput = statistics.pstdev(throwputsCalculations)
	return meanThrowput, stdevThrowput
	
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
		resultsLine = "Entity "+str(entity)+" = "+str(entityMeanThrowput)+" pm "+str(entityStdevThrowput)+" txps\n"
		f.write(resultsLine)
	f.write("\n")
	f.close()

if __name__ == "__main__":
	main()
