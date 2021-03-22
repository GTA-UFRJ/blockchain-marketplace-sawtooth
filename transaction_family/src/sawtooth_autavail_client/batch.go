package main

import (
	"github.com/jessevdk/go-flags"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"strconv"
)

// autavail batch --url http://rest-api:8008 --keyfile ./foo.priv 119131225 5
type Batch struct {
	Args struct {
		OrgId        string `positional-arg-name:"orgid" required:"true" description:"Identification number of organization"`
		TxCount      string `positional-arg-name:"txcount" required:"true" description:"Number of random advertisement transactions"`
	} `positional-args:"true"`
	Url     string `long:"url" description:"Specify URL of REST API"`
	Keyfile string `long:"keyfile" description:"Identify file containing user's private key"`
}

func (args *Batch) Name() string {
	return "batch"
}

func (args *Batch) KeyfilePassed() string {
	return args.Keyfile
}

func (args *Batch) UrlPassed() string {
	return args.Url
}

// Generate a 16 character hexadecimal lowercase random string to be the advertisement identifier number
func (args *Batch) TxId () (string, error) {
	nonce := make([]byte, 10)
	_, err := rand.Read(nonce)
	if err != nil {
		return "", err
	}
	txid := sha512.Sum512(nonce)
	return hex.EncodeToString(txid[0:(TXID_LENGTH/2)]), nil
}

// Generate a list of advert transactions IDs
func (args *Batch) IdList () ([]string, error) {
	var txidList []string

	txCount, err := strconv.Atoi(args.Args.TxCount)
	if err != nil {
    return []string{}, err
  }

	for txIndex := 0; txIndex < txCount; txIndex++ {
		txid, err := args.TxId()
		if err != nil {
			return []string{}, err
		}
		txidList = append(txidList, txid)
	}
	return txidList, nil
}

// Apply AddCommand method to parser object
func (args *Batch) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(args.Name(), "Benchmark performance", "Generates a batch of random advertisement transactions", args)
	if err != nil {
		return err
	}
	return nil
}

func (args *Batch) Run() error {
	// Construct client using URL to submit and keyfile to sign transactions
	autavailClient, err := GetClient(args, true)
	if err != nil {
		return err
	}

	txidList, err := args.IdList()
	if err != nil {
    return err
  }

	// Call function in the program to submitt transactions (autavail_client.go)
	_, err = autavailClient.Batch(txidList, args.Args.OrgId)
	return err
}
