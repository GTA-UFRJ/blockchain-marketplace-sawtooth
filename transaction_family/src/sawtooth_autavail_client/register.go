package main

import (
	"github.com/jessevdk/go-flags"
)

// autavail register --url http://rest-api:8008 --keyfile ./foo.priv 119131225 /
type Register struct {
	Args struct {
		OrgId        string `positional-arg-name:"orgid" required:"true" description:"Identification number of organization"`
	} `positional-args:"true"`
	Url     string `long:"url" description:"Specify URL of REST API"`
	Keyfile string `long:"keyfile" description:"Identify file containing user's private key"`
}

func (args *Register) Name() string {
	return "register"
}

func (args *Register) KeyfilePassed() string {
	return args.Keyfile
}

func (args *Register) UrlPassed() string {
	return args.Url
}

// Apply AddCommand method to parser object
func (args *Register) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(args.Name(), "Register organization", "Sends an autavail register transaction", args)
	if err != nil {
		return err
	}
	return nil
}

func (args *Register) Run() error {
	// Collect data to construct payload
	txtype := args.Name()
	orgid := args.Args.OrgId

	// Construct client using URL to submit and keyfile to sign transactions
	autavailClient, err := GetClient(args, true)
	if err != nil {
		return err
	}

	// Call function in the program to submitt transactions (autavail_client.go)
	_, err = autavailClient.Register(txtype, orgid)
	return err
}
