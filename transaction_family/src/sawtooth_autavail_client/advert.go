package main

import (
	"github.com/jessevdk/go-flags"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

// autavail advert --url http://rest-api:8008 --keyfile ./foo.priv 500 119131225 /
// "Biomedical Electronic Instrumentation Dataset" "Heartbeat monitoring collected from wearable IoT devices" 10.14.291.123 Medical
type Advert struct {
	Args struct {
		Price        string `positional-arg-name:"price" required:"true" description:"Price of the asset"`
		OrgId        string `positional-arg-name:"orgid" required:"true" description:"Identification number of organization"`
		Title        string `positional-arg-name:"title" required:"true" description:"Title of the anouncement"`
		Description  string `positional-arg-name:"description" required:"true" description:"Description of the anouncement"`
		IpAddr       string `positional-arg-name:"ip" required:"true" description:"Client IP address"`
		DataType     string `positional-arg-name:"datatype" required:"true" description:"Data type that are being anounced"`
	} `positional-args:"true"`
	Url     string `long:"url" description:"Specify URL of REST API"`
	Keyfile string `long:"keyfile" description:"Identify file containing user's private key"`
}

func (args *Advert) Name() string {
	return "advert"
}

func (args *Advert) KeyfilePassed() string {
	return args.Keyfile
}

func (args *Advert) UrlPassed() string {
	return args.Url
}

// Generate a 16 character hexadecimal lowercase random string to be the advertisement identifier number
func (args *Advert) TxId () (string, error) {
	nonce := make([]byte, 10)
	_, err := rand.Read(nonce)
	if err != nil {
		return "", err
	}
	txid := sha512.Sum512(nonce)
	return hex.EncodeToString(txid[0:TXID_LENGTH]), nil
}

// Apply AddCommand method to parser object
func (args *Advert) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(args.Name(), "Advertise data transactions value", "Sends an autavail advertisement transaction", args)
	if err != nil {
		return err
	}
	return nil
}

func (args *Advert) Run() error {
	// Collect data to construct payload
	txtype := args.Name()
	txid, err := args.TxId()
	price := args.Args.Price
	ipaddr := args.Args.IpAddr
	orgid := args.Args.OrgId
	title := args.Args.Title
	description := args.Args.Description
	datatype := args.Args.DataType
	if err != nil {
		return err
	}

	// Construct client using URL to submit and keyfile to sign transactions
	autavailClient, err := GetClient(args, true)
	if err != nil {
		return err
	}

	// Call function in the program to submitt transactions (autavail_client.go)
	_, err = autavailClient.Advert(txtype, txid, price, ipaddr, orgid, title, description, datatype)
	return err
}
