/**
 * Copyright 2017 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ------------------------------------------------------------------------------
 */

package handler

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	cbor "github.com/brianolson/cbor_go"
	"github.com/hyperledger/sawtooth-sdk-go/logging"
	"github.com/hyperledger/sawtooth-sdk-go/processor"
	"protobuf/processor_pb2"
	"strings"
)

var logger *logging.Logger = logging.Get()

type AutAvailPayload struct {
	CompletePayload	string
}

type AutAvailHandler struct {
	namespace string
}

func NewIntkeyHandler(namespace string) *IntkeyHandler {
	return &IntkeyHandler{
		namespace: namespace,
	}
}

const (
	FAMILY_NAME     = "autavail"
)

func (self *AutAvailHandler) FamilyName() string {
	return FAMILY_NAME
}

func (self *AutAvailHandler) FamilyVersions() []string {
	return []string{"1.0"}
}

func (self *AutAvailHandler) Namespaces() []string {
	return []string{self.namespace}
}

func (self *AutAvailHandler) Apply(request *processor_pb2.TpProcessRequest, context *processor.Context) error {
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

	

	//As the different transactions have different states, we analyze the transaction type to proceed with the processing
	if !(txType == "advertisement" || txType == "purchase" ) {
		return &processor.InvalidTransactionError{Msg: fmt.Sprintf("Invalid transaction type: %v", txType)}
	}

	autAvailState := autAvail_state.NewAutAvailState(context)

	switch txType {
	case "advertisement":
		err := validateAdvertisement (autAvailState, price)
		if err != nil {
			return err
		}
		stateAdString := txType + ":" + price + ":" + ipAddress + ":" + orgID + ":" + title + ":" + description + ":" + dataType
		ad := &autAvail_state.Ad{
			CompleteState:   stateAdString,
		}
		return autavailState.SetAd(txID, ad)
	case "purchase":
		newValue = storedValue - value
	}

	if newValue < MIN_VALUE {
		return &processor.InvalidTransactionError{
			Msg: fmt.Sprintf("New Value must be >= ", MIN_VALUE),
		}
	}
	if newValue > MAX_VALUE {
		return &processor.InvalidTransactionError{
			Msg: fmt.Sprintf("New Value must be <= ", MAX_VALUE),
		}
	}
	}

	if verb == "set" {
		if exists {
			return &processor.InvalidTransactionError{Msg: "Cannot set existing value"}
		}
		newValue = value
	}

	collisionMap[name] = newValue
	data, err = EncodeCBOR(collisionMap)
	if err != nil {
		return &processor.InternalError{
			Msg: fmt.Sprint("Failed to encode new map:", err),
		}
	}

	addresses, err := context.SetState(map[string][]byte{
		address: data,
	})
	if err != nil {
		return err
	}
	if len(addresses) == 0 {
		return &processor.InternalError{Msg: "No addresses in set response"}
	}

	return nil
}

func EncodeCBOR(value interface{}) ([]byte, error) {
	data, err := cbor.Dumps(value)
	return data, err
}

func DecodeCBOR(data []byte, pointer interface{}) error {
	defer func() error {
		if recover() != nil {
			return &processor.InvalidTransactionError{Msg: "Failed to decode payload"}
		}
		return nil
	}()
	err := cbor.Loads(data, pointer)
	if err != nil {
		return err
	}
	return nil
}

func Hexdigest(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}