package handler

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/hyperledger/sawtooth-sdk-go/logging"
	"github.com/hyperledger/sawtooth-sdk-go/processor"
	"protobuf/processor_pb2"
	"strings"
	"strconv"
)

const (
	FAMILY_NAME         = "autavail"
	ADVERT_TYPE         = "advert"
	BUY_TYPE            = "buy"
	REGISTER_TYPE       = "register"
	TXID_LENGTH         = 16
	VERB_ADDRESS_LENGTH = 64
	INITIAL_ORG_BALANCE = 2000.0
)

var logger *logging.Logger = logging.Get()

type AutAvailHandler struct {
	namespace string
}

func (self *AutAvailHandler) FamilyName() string {
	return FAMILY_NAME
}

func (self *AutAvailHandler) FamilyVersions() []string {
	return []string{"1.0"}
}

func (self *AutAvailHandler) Namespaces() []string {
	return []string{self.namespace}
}

func NewAutAvailHandler(namespace string) *AutAvailHandler {
	return &AutAvailHandler{
		namespace: namespace,
	}
}

// Apply method gets called when a command (request) needs to be executed in order to change the state (context)
func (self *AutAvailHandler) Apply(request *processor_p0b2.TpProcessRequest, context *processor.Context) error {

	// Get transaction payload, transaction header and header signer public key
	payloadData := request.GetPayload()
	header := request.GetHeader()
  sender := header.GetSignerPublicKey()
	if payloadData == nil {
		return &processor.InvalidTransactionError{Msg: "Must contain payload"}
	}

	//As we define the payload as a ":"-separated string, we split the payload to process the transaction
	payload := string(payloadData)
	payloadSplit := strings.Split(payload,":")
	if len(payloadSplit) != 9 {
		return &processor.InvalidTransactionError{Msg: "Malformed payload: %s", payload}
	}
	txType := payloadSplit[0]
	txID := payloadSplit[1]
	advertisementTxID := payloadSplit[2]
	advertisementOrgID := payloadSplit[3]
	price := payloadSplit[4]
	ipAddress := payloadSplit[5]
	orgID := payloadSplit[6]
	title := payloadSplit[7]
	description := payloadSplit[8]
	dataType := payloadSplit[9]

	// Verify transaction type
	switch txType {

		// ------------------------- ADVERTISEMENT -------------------------

		case ADVERT_TYPE:

		// Validate advertisement transaction
		if (len(txID) != TXID_LENGTH)     ||
		   (len(advertisementTxID) != 0)  ||
		   (len(advertisementOrgID) != 0) ||
			 (len(price) == 0)              ||
			 (len(ipAddress) == 0)          ||
			 (len(orgID) == 0)              ||
			 (len(title) == 0)              ||
			 (len(description) == 0)        ||
			 (len(dataType) == 0) {
			return &processor.InvalidTransactionError{Msg: "Incorrect fields: %s", payload}
		}
	  if floatPrice, err := strconv.ParseFloat(price,32); err != nil {
			return &processor.InvalidTransactionError{Msg: "Price is not a valid float: %s. Error converting: %v", price, err}
		}

		// Get state addresses
		hashedTxID := Hexdigest(txID)
		advertAddress := self.namespace + hashedTxID[len(hashedTxID)-VERB_ADDRESS_LENGTH:]
		hashedOrgID := Hexdigest(orgID)
		orgAddress := self.namespace + hashedOrgID[len(hashedOrgID)-VERB_ADDRESS_LENGTH:]

		// Get state information
		stateQuery, err := context.GetState([]string{advertAddress, orgAdress})
		if err != nil {
			return err
		}
		if len(string(stateQuery[advertAddress])) > 0 {
			return &processor.InvalidTransactionError{Msg: "Transaction %s alredy exists", txID}
		}
		orgData := string(stateQuery[orgAddress])
		if len(orgData) == 0 {
			return &processor.InvalidTransactionError{Msg: "Organization %s not registred", orgID}
		}

		// Construct state advert string
		advertStateData := fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s",
		txID,
		price,
		ipAddress,
		orgID,
		title,
		description,
		dataType)

		// Apply advertisement transaction
		adresses, err := context.SetState(map[string][]byte{
			advertAddress: []byte(advertStateData),
		})
		if err != nil {
			return err
		}
		if len(addresses) == 0 {
			return &processor.InternalError{Msg: "No addresses in set response"}
		}
		break;

		// ------------------------- BUY -------------------------

		case BUY_TYPE:

		// Validate buy transaction
		if (len(txID) != TXID_LENGTH)              ||
		   (len(advertisementTxID) != TXID_LENGTH) ||
		   (len(advertisementOrgID) == 0)          ||
			 (len(price) != 0)                       ||
			 (len(ipAddress) == 0)                   ||
			 (len(orgID) == 0)                       ||
			 (len(title) != 0)                       ||
			 (len(description) != 0)                 ||
			 (len(dataType) != 0) {
			return &processor.InvalidTransactionError{Msg: "Incorrect fields: %s", payload}
		}

		// Get state addresses
		hashedTxID := Hexdigest(txID)
		buyAddress := self.namespace + hashedTxID[len(hashedTxID)-VERB_ADDRESS_LENGTH:]
		hashedAdvertTxId := Hexdigest(advertisementTxID)
		advertAddress := self.namespace + hashedAdvertTxID[len(hashedAdvertTxID)-VERB_ADDRESS_LENGTH:]
		hashedOrgID := Hexdigest(orgID)
		buyOrgAddress := self.namespace + hashedOrgID[len(hashedOrgID)-VERB_ADDRESS_LENGTH:]
		hashedAdvertOrgID := Hexdigest(advertisementOrgID)
		advertOrgAddress := self.namespace + hashedAdvertOrgID[len(hashedAdvertOrgID)-VERB_ADDRESS_LENGTH:]

		// Get state information
		stateQuery, err := context.GetState([]string{buyAddress, buyOrgAddress, advertAddress, advertOrgAddress})
		if err != nil {
			return err
		}
		if len(string(stateQuery[buyAddress])) > 0 {
			return &processor.InvalidTransactionError{Msg: "Transaction %s alredy exists", txID}
		}
		advertTxData := string(stateQuery[advertAddress])
		if len(advertTxData) == 0 {
			return &processor.InvalidTransactionError{Msg: "Advertisement transaction %s does not exist", advertisementTxID}
		}
		buyOrgData := string(stateQuery[buyOrgAddress])
		if len(buyOrgData) == 0 {
			return &processor.InvalidTransactionError{Msg: "Organization %s not registred", orgID}
		}
		advertOrgData := string(stateQuery[advertOrgAddress])
		if len(advertOrgData) == 0 {
			return &processor.InvalidTransactionError{Msg: "Organization %s not registred", advertisementOrgID}
		}

		// Obtain informations about advertisement transaction
		advertTxDataSplit := strings.Split(advertTxData)
		if assetPrice, err := strconv.ParseFloat(advertTxDataSplit[1],32), err != nil {
			return &processor.InternalError{Msg: "Error reading price: %s", advertTxDataSplit[1]}
		}
		if advertTxDataSplit[3] != advertisementOrgId {
			return &processor.InvalidTransactionError{Msg: "Advertisement transaction was not sent from this organization: %s", advertTxData}
		}

		// Obtain balances
		buyOrgDataSplit := strings.Split(buyOrgData,":")
		if buyerBalance, err := strconv.ParseFloat(buyOrgDataSplit[1],32), err != nil {
			return &processor.InternalError{Msg: "Error reading balance: %s", buyOrgDataSplit[1]}
		}
		advertOrgDataSplit := strings.Split(advertOrgData,":")
		if sellerBalance, err := strconv.ParseFloat(advertOrgDataSplit[1],32), err != nil {
			return &processor.InternalError{Msg: "Error reading balance: %s", advertOrgDataSplit[1]}
		}

		// Compute balances if payment is allowed
		if buyerBalance >= assetPrice {
			buyerBalance -= assetPrice
			sellerBalance += assetPrice
		}

		// Construct state buy string
		buyStateData := fmt.Sprintf("%s:%s:%s:%s:%s",
		txID,
		advertisementTxID,
		advertisementOrgID,
		ipAddress,
		orgID)

		// Construct state organizations string
		buyOrgStateData := fmt.Sprintf("%s:%f", orgID, buyerBalance)
		advertOrgStateData := fmt.Sprintf("%s:%f", advertisementOrgID, sellerBalance)

		// Apply buy transaction
		adresses, err := context.SetState(map[string][]byte{
			buyAddress: []byte(buyStateData),
			buyOrgAddress: []byte(buyOrgStateData),
			buyAdvertAddress: []byte(buyAdvertStateData),
		})
		if err != nil {
			return err
		}
		if len(addresses) == 0 {
			return &processor.InternalError{Msg: "No addresses in set response"}
		}
		break;

		// ------------------------- REGISTER -------------------------

		case REGISTER_TYPE:

		// Validate register transaction
    if (len(txID) != 0)               ||
       (len(advertisementTxID) != 0)  ||
       (len(advertisementOrgID) != 0) ||
       (len(price) != 0)              ||
       (len(ipAddress) != 0)          ||
       (len(orgID) == 0)              ||
       (len(title) != 0)              ||
       (len(description) != 0)        ||
       (len(dataType) != 0) {
      return &processor.InvalidTransactionError{Msg: "Incorrect fields: %s", payload}
    }

		// Get state address
		hashedOrgID := Hexdigest(orgID)
    orgAddress := self.namespace + hashedOrgID[len(hashedOrgID)-VERB_ADDRESS_LENGTH:]

		// Get state information
    stateQuery, err := context.GetState([]string{orgAddress})
    if err != nil {
      return err
    }
    if len(string(stateQuery[orgAddress])) > 0 {
      return &processor.InvalidTransactionError{Msg: "Organization %s alredy registred", orgID}
    }

		// Construct state register string
    registerStateData := fmt.Sprintf("%s:%f", orgID, INITIAL_ORG_BALANCE)

		// Apply register transaction
		adresses, err := context.SetState(map[string][]byte{
      orgAddress: []byte(registerStateData),
    })
    if err != nil {
      return err
    }
    if len(addresses) == 0 {
      return &processor.InternalError{Msg: "No addresses in set response"}
    }
		break;

		default:
			return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Invalid transaction type: %v", txType)}
	}
	return nil
}

// Compute the SHA512 hash value in hexadecimal string format
func Hexdigest(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}
