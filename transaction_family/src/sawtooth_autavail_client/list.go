package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

type List struct {
	Url     string `long:"url" description:"Specify URL of REST API"`
}

func (args *List) Name() string {
	return "list"
}

func (args *List) KeyfilePassed() string {
	return ""
}

func (args *List) UrlPassed() string {
	return args.Url
}

func (args *List) Register(parent *flags.Command) error {
	_, err := parent.AddCommand(args.Name(), "Display all data anounced", "Show all advertisement transactions information", args)
	if err != nil {
		return err
	}
	return nil
}

func (args *List) Run() error {
	autavailClient, err := GetClient(args, true)
	if err != nil {
		return err
	}
	advertSlice, err = autavailClient.List()
	if err != nil {
		return err
	}
	for index, advertTx := range advertSlice {
		fmt.Println(index, ": ", advertTx)
	}
	return nil
}
