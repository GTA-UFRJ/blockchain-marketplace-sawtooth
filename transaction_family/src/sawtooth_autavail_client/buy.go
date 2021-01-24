package main

import (
	"github.com/jessevdk/go-flags"
	"strconv"
	"net"
	"time"
	"crypto/rand"
  "crypto/sha512"
  "encoding/hex"
)

type Buy struct {
	Args struct {
		OrgId				 string `positional-arg-name:"orgid" required:"true" description:"Identification number of organization"`
		AdvertTxId	 string `positional-arg-name:"adverttxid" required:"true" description:"Identification number of advertisement transaction"`
		IpAddr       string `positional-arg-name:"ip" required:"true" description:"Client IP address"`
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

func (args *Buy) TxId () string {
	nonce := make([]byte, 10)
  _, err := rand.Read(nonce)
  if err != nil {
    return err
  }
  txid := sha512.Sum512(nonce)
  return hex.EncodeToString(txid[0:])
}

func (args *Buy) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(args.Name(), "Buy data transactions value", "Sends an autavail buy transaction", args)
	if err != nil {
		return err
	}
	return nil
}

func (args *Buy) Run() error {
	txtype := args.Name()
	txid := args.TxId()
	adverttxid = args.Args.AdvertTxId
	ipaddr := args.Args.IpAddr
	orgid := args.Args.OrgId

	autavailClient, err := GetClient(args, true)
	if err != nil {
		return err
	}
	_, err = autavailClient.Buy(txtype, txid, adverttxid, ipaddr, orgid)
	return err
}
