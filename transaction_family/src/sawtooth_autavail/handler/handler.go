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
		case ADVERT_TYPE:

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

		// Validate advertisement transaction
		if (len(txID) != TXID_LENGTH)     ||
		   (len(advertisementTxID) != 0)  ||
		   (len(advertisementOrgID) != 0) ||
		   (len(advertisementTxID) != 0)  ||
			 (len(price) == 0)              ||
			 (len(ipAddress) == 0)          ||
			 (len(orgID) == 0)              ||
			 (len(title) == 0)              ||
			 (len(description) == 0)        ||
			 (len(dataType) == 0) {
			return &processor.InvalidTransactionError{Msg: "Missing fields: %s", payload}
		}
	  if floatPrice, err := strconv.ParseFloat(price,32); err != nil {
			return &processor.InvalidTransactionError{Msg: "Price is not a valid float: %s. Error converting: %v", price, err}
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
		address := advertAddress
		data := advertStateData,
		adresses, err := context.SetState(map[string][]byte{
			address: data,
		})
		if err != nil {
			return err
		}
		if len(addresses) == 0 {
			return &processor.InternalError{Msg: "No addresses in set response"}
		}
		break;

		case BUY_TYPE:
		// Get state address
		// Validate advertisement transaction
		// Apply advertisement transaction
		break;

		case REGISTER_TYPE:
		// Get state address
		// Validate advertisement transaction
		// Apply advertisement transaction
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
