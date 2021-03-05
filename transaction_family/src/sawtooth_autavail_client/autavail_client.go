package main

import (
  bytes2 "bytes"
  "encoding/base64"
  "encoding/hex"
  "errors"
  "fmt"
  // cbor "github.com/brianolson/cbor_go"
  "github.com/golang/protobuf/proto"
  "github.com/hyperledger/sawtooth-sdk-go/signing"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "math/rand"
  "net/http"
  "protobuf/batch_pb2"
  "protobuf/transaction_pb2"
  "strconv"
  "strings"
	"os"
)

type AutavailClient struct {
  url    string
  signer *signing.Signer
}

/*
  Run (advert, buy, list) calls GetClient (main)
  GetClient (main) calls this function bellow
  Creates a signer private key
  Puts the signer and the URL in a struct
*/
func NewAutavailClient(url string, keyfile string) (IntkeyClient, error) {

  var privateKey signing.PrivateKey

  // keyfile = ~/.sawtooth/keys/$USER.priv by default
  if keyfile != "" {
    // Read private key file
    privateKeyStr, err := ioutil.ReadFile(keyfile)
    if err != nil {
      return AutavailClient{},
        errors.New(fmt.Sprintf("Failed to read private key: %v", err))
    }
    // Get private key object
    privateKey = signing.NewSecp256k1PrivateKey(privateKeyStr)
  } else {
    privateKey = signing.NewSecp256k1Context().NewRandomPrivateKey()
  }
  /*
    https://sawtooth.hyperledger.org/docs/core/releases/latest/_autogen/sdk_submit_tutorial_go.html#creating-a-private-key-and-signer

    context -> cryptoFactory
                            \
                             signer
                            /
       keyfile -> privatekey
  */
  context := signing.NewSecp256k1Context()
  cryptoFactory := signing.NewCryptoFactory(context)
  signer := cryptoFactory.NewSigner(privateKey)

	// url = http://localhost:8008 by default
  return AutavailClient{url, signer}, nil
}

/*
  main (main) calls Run (advert, buy, list)
  Run (advert, buy, list) calls Advert, Buy, List down bellow
  Calls the method to send the transaction passing the informations to build the payload
  Tx type, Tx ID, advert. Tx ID, price, IP address, organization ID, title, description, data type
*/
func (autavailClient AutavailClient) Advert(
  txtype string, txid string, price string, ipaddr string, orgid string, title string, description string, datatype string) (string, error) {
  adverttxid := ""
	advertorgid := ""
  return autavailClient.sendTransaction(txtype, txid, adverttxid, advertorgid, price, ipaddr, orgid, title, description, datatype)
}

func (autavailClient AutavailClient) Buy(
  txtype string, txid string, adverttxid string, advertorgid string, ipaddr string, orgid string) (string, error) {
  price := ""
  title := ""
  description := ""
  datatype := ""
  return autavailClient.sendTransaction(txtype, txid, adverttxid, advertorgid, price, ipaddr, orgid, title, description, datatype)
}

/*
	curl http://sawtooth-rest-api-default-0:8008/state?address=1cf126
  {
    "data": [
      {
        "address": "1cf12660cf60810ceb9d4e4b49e2777c69a3ee188a8ebfa539b3ed7ea29494abc60773",
        "data": "oWVvdXRyYRhQ"
      },
      {
        "address": "1cf126a8d4e16cceab74612bfd24ac64c61e4043cf242dacb0aaffaaa94ef3837644b2",
        "data": "oWVjaGF2ZRha"
      }
    ],
    "head": "9308ee23e96e598103993c34e61e396 ...",
    "link": "http://sawtooth-rest-api-default-0:8008/state?head=9308ee23e96e598103993c34e61e396 ...",
    "paging": {
      "limit": null,
      "start": null
    }
  }

  Autavail list implementation

  1- Makes a GET request to REST API on the specific TF address
  2- Reads the body of the response in YAML string format
  3- Converts (unmarshal) the YAML bytes into a map 
  4- Access the data field of YAML response which contains an entry list (value of data key)
  5- Access the data field of each entry list elements which contains a string (value of data key)
  6- Decodes bytes in base 64 standard encoding
  7- Append strings to a list (the only differnce)

  Returns a slice of states entries got from a HTTP GET request from validator throw REST API
*/
func (autavailClient AutavailClient) List() ([]string, error) {

  // API to call (state?address=d7ad2c)
  apiSuffix := fmt.Sprintf("%s?address=%s",
    STATE_API, autavailClient.getPrefix())

  // HTTP request using GET method
  response, err := autavailClient.sendRequest(apiSuffix, []byte{}, "", "")
  if err != nil {
    return []string{}, err
  }

	var toReturn []string
  responseMap := make(map[interface{}]interface{})

  // Decodes YAML response to map
  err = yaml.Unmarshal([]byte(response), &responseMap)
  if err != nil {
    return []string{},
      errors.New(fmt.Sprintf("Error reading response: %v", err))
  }

  // List and access each state entry (in date field)
  encodedEntries := responseMap["data"].([]interface{})
	for _, entry := range encodedEntries {
    entryData, ok := entry.(map[interface{}]interface{})
    if !ok {
      return []string{},
        errors.New("Error reading entry data")
    }

    // Access the content of the state entry
    stringData, ok := entryData["data"].(string)
    if !ok {
      return []string{},
        errors.New("Error reading string data")
    }

    // Decodes the base64 encoded state string
    decodedBytes, err := base64.StdEncoding.DecodeString(stringData)
    if err != nil {
      return []string{},
        errors.New(fmt.Sprint("Error decoding: %v", err))
    }

    // Append the state information to the slice
    toReturn = append(toReturn, string(decodedBytes[:]))
  }
  return toReturn, nil
}

func (autavailClient AutavailClient) Register(
  txtype string, orgid string) (string, error) {
	txid := ""
	adverttxid := ""
	advertorgid := ""
  price := ""
	ipaddr := ""
  title := ""
  description := ""
  datatype := ""
  return autavailClient.sendTransaction(txtype, txid, adverttxid, advertorgid, price, ipaddr, orgid, title, description, datatype)
}

// Makes a HTTP request to the validator throw the REST API and returns response body
func (autavailClient AutavailClient) sendRequest(
  apiSuffix string,   // state?address=d7ad2c
  data []byte,        // Data to POST. If empty, GET method will be applied
  contentType string, // application/octet-stream in case of transactions 
  txid string) (string, error) {

  // Construct URL (http://localhost:8008/state?address=d7ad2c)
  var url string
  if strings.HasPrefix(autavailClient.url, "http://") {
    url = fmt.Sprintf("%s/%s", autavailClient.url, apiSuffix)
  } else {
    url = fmt.Sprintf("http://%s/%s", autavailClient.url, apiSuffix)
  }

	// Send request to validator URL
  var response *http.Response
  var err error
  if len(data) > 0 {
    response, err = http.Post(url, contentType, bytes2.NewBuffer(data))
  } else {
    response, err = http.Get(url)
  }
  if err != nil {
    return "", errors.New(
      fmt.Sprintf("Failed to connect to REST API: %v", err))
  }
  if response.StatusCode == 404 {
    logger.Debug(fmt.Sprintf("%v", response))
    return "", errors.New(fmt.Sprintf("No such transaction: %s", txid))
  } else if response.StatusCode >= 400 {
    return "", errors.New(
      fmt.Sprintf("Error %d: %s", response.StatusCode, response.Status))
  }
  defer response.Body.Close()
  reponseBody, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return "", errors.New(fmt.Sprintf("Error reading response: %v", err))
  }
  return string(reponseBody), nil
}

// Encode payload, build transaction, wrap transaction in bacth, send batch, return HTTP response and error
func (autavailClient AutavailClient) sendTransaction(
	txtype string,
	txid string,
	adverttxid string,
	advertorgid string,
	price string,
	ipaddr string,
	orgid string,
	title string,
	description string,
	datatype string) (string, error) {

	// Payload: <txtype>:<txid>:<adverttxid>:<price>:<ipaddr>:<orgid>:<title>:<description>:<datatype>
	payload := fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s:%s:%s",
	txtype,
	txid,
	adverttxid,
	advertorgid,
	price,
	ipaddr,
	orgid,
	title,
	description,
	datatype)

	// Construct the 70 characters hex encoded transaction address string
	inputAdresses, outputAdresses := autavailClient.getAddress(txtype, txid, adverttxid, orgid, advertorgid)

	// Construct the TransactionHeader
	rawTransactionheader := transaction_pb2.TransactionHeader{
				SignerPublicKey:   autavailClient.signer.GetPublicKey().AsHex(),
				FamilyName:        FAMILY_NAME,
				FamilyVersion:     FAMILY_VERSION,
				Dependencies:      []string{},
				Nonce:             strconv.Itoa(rand.Int()),
				BatcherPublicKey:  autavailClient.signer.GetPublicKey().AsHex(),
				Inputs:            inputAdresses,
				Outputs:           outputAdresses,
				PayloadSha512:     Sha512HashValue(string(payload)),
	}

	// Serialize (marshall) in protocol buffer (Google Protobuf)
	transactionHeader, err := proto.Marshal(&rawTransactionHeader)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to serialize transaction header: %v", err))
	}

	// Signature of transaction header
	transactionHeaderSignature := hex.EncodeToString(
					autavailClient.signer.Sign(transactionHeader))

	// Construct transaction
	transaction  := transaction_pb2.Transaction{
		Header:           transactionHeader,
		HeaderSignature:  transactionHeaderSignature,
		Payload:          []byte(payload),
	}

	// Get BatchList
	rawBatchList, err := autavailClient.createBatchList(
					[]*transaction_pb2.Transaction{&transaction})
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to construct batch list: %v", err))
	}

	// Serialize batch list
	batchId := rawBatchList.Batches[0].HeaderSignature
	batchList, err := proto.Marshal(&rawBatchList)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to serialize batch list: %v", err))
	}

	// If url is "file", append batch list in file named "autavail.workload"
	if autavailClient.url == "file" {
		fileDescriptor, err := os.OpenFile("autavail.workload", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Unable to open file: %v", err))
    }
		if _, err := f.Write(batchList); err != nil {
			return "", errors.New(fmt.Sprintf("Unable to write into file: %v", err))
    }
		if err := f.Close(); err != nil {
			return "", errors.New(fmt.Sprintf("Unable to close file: %v", err))
    }
		return "Write into autavail.workload file", nil
	}

	// Submmit batch list throw HTTP
	return autavailClient.sendRequest(BATCH_SUBMIT_API, batchList, CONTENT_TYPE_OCTET_STREAM, txid)
}

// Prefix = 6 first characters of hex encoding of SHA512 of "autavail" = d7ad2c
func (autavailClient AutavailClient) getPrefix() string {
  return Sha512HashValue(FAMILY_NAME)[:FAMILY_NAMESPACE_ADDRESS_LENGTH]
}

/*
 Addresses determine where TP (transaction processor) are allowed to read from (inputs) or write to (outputs) the state

    address = prefix + 64 characters of SHA512 of something
       |        |                       |
     state  TF specific     Depends on the transaction
*/
func (autavailClient AutavailClient) getAddress(
	txtype string,
	txid string,
	adverttxid string,
	orgid string,
	advertorgid string) ([]string, []string) {

	prefix := autavailClient.getPrefix()
	switch (txtype) {
		case "advert":
		// Address where advert transaction will be stored
		hashedTxId := Sha512HashValue(txid)
		advertAddress := prefix + hashedTxId[len(hashedTxId)-FAMILY_VERB_ADDRESS_LENGTH:]

		// Address where seller organization balance is stored
		hashedOrdId := Sha512HashValue(orgid)
		advertOrgAddress := prefix + hashedOrdId[len(hashedOrdId)-FAMILY_VERB_ADDRESS_LENGTH:]

		// TP will read to validate if the transacation and the organization alredy exists
		inputAddresses := []string{advertAddress, advertOrgAddress}

		// TP will write the transaction
		outputAddresses := []string{advertAddress}
		break;

		case "buy":
		// Address where buy transaction will be stored 
		hashedTxId := Sha512HashValue(txid)
		buyAddress := prefix + hashedTxId[len(hashedTxId)-FAMILY_VERB_ADDRESS_LENGTH:]

		// Address where correspondent advert transaction is stored
		hashedAdvertTxId := Sha512HashValue(adverttxid)
		advertAddress := prefix + hashedAdvertTxId[len(hashedAdvertTxId)-FAMILY_VERB_ADDRESS_LENGTH:]

		// Address where buyer organization balance is stored
		hashedOrgId := Sha512HashValue(orgid)
		buyOrgAddress := prefix + hashedOrgId[len(hashedOrgId)-FAMILY_VERB_ADDRESS_LENGTH:]

		// Address where seller organization balance is stored
		hashedAdvertOrgId := Sha512HashValue(advertorgid)
		advertOrgAddress := prefix + hashedAdvertOrgId[len(hashedAdvertOrgId)-FAMILY_VERB_ADDRESS_LENGTH:]

		// TP will read to validate the buyer and seller balances and the buy and advert transactions existance
		inputAddresses := []string{buyAddress, buyOrgAddress, advertAddress, advertOrgAddress}

		// TP will write to the balances and apply the buy transactions
		inputAddresses := []string{buyAddress, buyOrgAddress, advertOrgAddress}
		break;

		case "register":
		// Address where organization balance will be stored
		hashedOrgId := Sha512HashValue(orgid)
		orgAddress := prefix + hashedOrgId[len(hashedOrgId)-FAMILY_VERB_ADDRESS_LENGTH:]

		// TP will read to validate if the organization is alredy registred
		inputAddresses := []string{orgAddress}

		// TP will write to register the organization
		inputAddresses := []string{orgAddress}
		break;
	}
	return inputAddresses, outputAddresses
}

/*
	Get a list of transactions and wrapps than into a batch
	Collect the batch into a batch list
*/
func (autavailClient AutavailClient) createBatchList(
				transactions []*transaction_pb2.Transaction) (batch_pb2.BatchList, error) {

	// Get list of TransactionHeader signatures
	transactionSignatures := []string{}
	for _, transaction := range transactions {
		transactionSignatures = append(transactionSignatures, transaction.HeaderSignature)
	}

	// Construct BatchHeader
	rawBatchHeader := batch_pb2.BatchHeader{
		SignerPublicKey: autavailClient.signer.GetPublicKey().AsHex(),
		TransactionIds:  transactionSignatures,
	}
	batchHeader, err := proto.Marshal(rawBatchHeader)
	if err != nil {
		return batch_pb2.BatchList{}, errors.New(fmt.Sprintf("Unable to serialize batch header: %v", err))
	}

	// Signature of batch header
	batchHeaderSignature := hex.EncodeToString(autavailClient.signer.Sign(batchHeader))

	//Construct batch
	batch := batch_pb2.Batch{
		Header:          batchHeader,
		Transactions:    transactions,
		HeaderSignature: batchHeaderSignatures,
	}

	// Construct BatchList
	return batch_pb2.BatchList{
		Batches: []*batch_pb2.Batch{&batch}
	}, nil
}
