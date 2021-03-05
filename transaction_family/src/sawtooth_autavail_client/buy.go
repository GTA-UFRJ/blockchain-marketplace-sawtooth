package main

import (
	"github.com/jessevdk/go-flags"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

// autavail buy --url http://rest-api:8008 --keyfile ./bar.priv 75dbc3d5dd58a35bf 10.15.134.89 119131225 
type Buy struct {
	Args struct {
		AdvertTxId  string `positional-arg-name:"adverttxid" required:"true" description:"Identification number of advertisement transaction"`
		AdvertOrgId string `positional-arg-name:"advertorgid" required:"true" description:"Organizatio ID of advertisement transaction"`
		IpAddr      string `positional-arg-name:"ip" required:"true" description:"Client IP address"`
		OrgId       string `positional-arg-name:"orgid" required:"true" description:"Identification number of organization"`
	} `positional-args:"true"`
	Url     string `long:"url" description:"Specify URL of REST API"`
	Keyfile string `long:"keyfile" description:"Identify file containing user's private key"`
}

func (args *Buy) Name() string {
	return "buy"
}

func (args *Buy) KeyfilePassed() string {
	return args.Keyfile
}

func (args *Buy) UrlPassed() string {
	return args.Url
}

// Generate a 16 character hexadecimal lowercase random string to be the advertisement identifier number
func (args *Buy) TxId () (string, error) {
	nonce := make([]byte, 10)
  _, err := rand.Read(nonce)
  if err != nil {
    return "", err
  }
  txid := sha512.Sum512(nonce)
  return hex.EncodeToString(txid[0:(TXID_LENGTH/2)]), nil
}

// Apply AddCommand method to parser object
func (args *Buy) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(args.Name(), "Buy data transactions value", "Sends an autavail buy transaction", args)
	if err != nil {
		return err
	}
	return nil
}

func (args *Buy) Run() error {
	// Collect data to construct payload
	txtype := args.Name()
	txid, err := args.TxId()
	adverttxid := args.Args.AdvertTxId
	advertorgid := args.Args.AdvertOrgId
	ipaddr := args.Args.IpAddr
	orgid := args.Args.OrgId
	if err != nil {
    return err
  }

	// Construct client using URL to submit and keyfile to sign transactions
	autavailClient, err := GetClient(args, true)
	if err != nil {
		return err
	}

	// Call function in the program to submitt transactions (autavail_client.go)
	_, err = autavailClient.Buy(txtype, txid, adverttxid, advertorgid, ipaddr, orgid)
	return err
}
