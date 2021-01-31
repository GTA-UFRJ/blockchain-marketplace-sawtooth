package main

import (
  bytes2 "bytes"
  "encoding/base64"
  "encoding/hex"
  "errors"
  "fmt"
  cbor "github.com/brianolson/cbor_go"
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
  "time"
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

  // keyfile = ~/.sawtooth/keys/$USER.priv by default (--key-file option is disabled)
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
  contex := signing.NewSecp256k1Context()
  cryptoFactory := signing.NewCryptoFactory(context)
  signer := cryptoFactory.NewSigner(privateKey)

	// url = http://sawtooth-rest-api-default-0:8008
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
  return autavailClient.sendTransaction(txtype, txid, adverttxid, price, ipaddr, orgid, title, description, datatype)
}

func (autavailClient AutavailClient) Buy(
  txtype string, txid string, adverttxid string, ipaddr string, orgid string) (string, error) {
  price := ""
  title := ""
  description := ""
  datatype := ""
  return autavailClient.sendTransaction(txtype, txid, adverttxid, price, ipaddr, orgid, title, description, datatype)
}

