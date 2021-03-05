package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

// autavail list --url http://rest-api:8008
type List struct {
	Url string `long:"url" description:"Specify URL of REST API"`
}

func (args *List) Name() string {
	return "list"
}

// This is a query transaction that don't get signed and registred in blockchain
func (args *List) KeyfilePassed() string {
	return ""
}

func (args *List) UrlPassed() string {
	return args.Url
}

// Apply AddCommand method to parser object
func (args *List) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(args.Name(), "Display all data anounced", "Show all advertisement transactions information", args)
	if err != nil {
		return err
	}
	return nil
}

func (args *List) Run() error {
	// Construct client using URL to submit
	autavailClient, err := GetClient(args, false)
	if err != nil {
		return err
	}

	// Get list of transactions from state and print than
	txSlice, err := autavailClient.List()
	if err != nil {
		return err
	}
	for index, tx := range txSlice {
		fmt.Println(index, "--> ", tx)
	}
	return nil
}
