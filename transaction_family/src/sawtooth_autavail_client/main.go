package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/hyperledger/sawtooth-sdk-go/logging"
	flags "github.com/jessevdk/go-flags"
	"os"
	"os/user"
	"path"
	"strings"
)

type Command interface {
	Register(*flags.Command) error
	Name() string
	UrlPassed() string
	Run() error
}

type Opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Enable more verbose output"`
}

var DISTRIBUTION_VERSION string

var logger *logging.Logger = logging.Get()

func init () {
	DISTRIBUTION_VERSION = "Unknown"
}

func main() {

	var opts Opts
	parser := flags.NewParser(&opts, flags.Default)
	parser.Command.Name = "autavail"

	commands := []Command{
		&Advert{},
		&Buy{},
		&List{},
	}

	for _, cmd := range commands {
		err := cmd.Register(parser.Command)
		if err != nil {
			logger.Errorf("Couldn't register command %v: %v", cmd.Name(), err)
			os.Exit(1)
		}
	}

	remaining, err := parser.Parse()
	if e, ok := err.(*flags.Error); ok {
		if e.Type == flags.ErrHelp {
			return
		} else {
			os.Exit(1)
		}
	}

	if len(remaining) > 0 {
		fmt.Println("Error: Unrecognized arguments passed: ", remaining)
		os.Exit(2)
	}

	switch len(opts.Verbose) {
		case 2:
			logger.SetLevel(logging.DEBUG)
		case 1:
			logger.SetLevel(logging.INFO)
		default:
			logger.SetLevel(logging.WARN)
	}

	if parser.Command.Active == nil {
		os.Exit(2)
	}

	name := parser.Command.Active.Name
	for _, cmd := range commands {
		if cmd.Name() == name {
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Println("Error: Command not found: ", name)
}

func Sha512HashValue(value string) string {
	hashHandler := sha512.New()
	hashHandler.Write([]byte(value))
	return strings.ToLower(hex.EncodeToString(hashHandler.Sum(nil)))
}

func GetClient(args Command, readFile bool) (AutavailClient, error) {
	url := args.UrlPassed()
	if url == "" {
		url = DEFAULT_URL
	}
	keyfile := ""
	if readFile {
		var err error
		keyfile, err = GetKeyfile()
		if err != nil {
			return AutavailClient{}, err
		}
	}
	return NewAutavailClient (url, keyfile)
}

func GetKeyfile() (string, error) {
	username, err := user.Current()
	if err != nil {
		return "", err
	}
	return path.Join(username.HomeDir, ".sawtooth", "keys", username.Username+".priv"), nil
}
