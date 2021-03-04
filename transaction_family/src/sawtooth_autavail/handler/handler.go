package handler

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/hyperledger/sawtooth-sdk-go/logging"
	"github.com/hyperledger/sawtooth-sdk-go/processor"
	"protobuf/processor_pb2"
	"strings"
)

const (
	FAMILY_NAME   = "autavail"
	ADVERT_TYPE   = "advert"
	BUY_TYPE      = "buy"
	REGISTER_TYPE = "register"
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
	payloadSplit := strings.Split(payload,":")
	txType := payloadSplit[0]
	txID := payloadSplit[1]
	advertisementTxID := payloadSplit[2]
	price := payloadSplit[3]
	ipAddress := payloadSplit[4]
	orgID := payloadSplit[5]
	title := payloadSplit[6]
	description := payloadSplit[7]
	dataType := payloadSplit[8]

	// Verify transaction type
	switch txType {
		case ADVERT_TYPE:
			// Get state address
			// Validate advertisement transaction
			// Apply advertisement transaction
		case BUY_TYPE:
			// Get state address
			// Validate advertisement transaction
			// Apply advertisement transaction
		case REGISTER_TYPE:
			// Get state address
			// Validate advertisement transaction
			// Apply advertisement transaction
		default:
			return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Invalid transaction type: %v", txType)}
	}
}

// Compute the SHA512 hash value in hexadecimal string format
func Hexdigest(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}
