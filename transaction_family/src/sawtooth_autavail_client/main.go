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

// Methods implemented by each subcommand
type Command interface {
	Register(*flags.Command) error
	Name() string
	KeyfilePassed() string
	UrlPassed() string
	Run() error
}

type Opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Enable more verbose output"`
}

// Use the SDK logger tool
var logger *logging.Logger = logging.Get()

func main() {

	// Creates an object to parse the arguments after the "autavail" comand
	var opts Opts
	parser := flags.NewParser(&opts, flags.Default)
	parser.Command.Name = "autavail"

	// Create sub-commands using pointers to structs, i.e. autavail advert ...
	// The structs: Advert, Buy and List, and it's methods are defined in the corresponding files (.go)
	commands := []Command{
		&Advert{},
		&Buy{},
		&List{},
		&Register{},
	}

	// Apply the Register method to all parser sub-commands, i.e. Advert.Register (advert.go)
	for _, cmd := range commands {
		err := cmd.Register(parser.Command)
		if err != nil {
			logger.Errorf("Couldn't register command %v: %v", cmd.Name(), err)
			os.Exit(1)
		}
	}

	// Parse the remaining arguments. Show help (-h / --help) 
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

	// Get passed sub-command name
	if parser.Command.Active == nil {
		os.Exit(2)
	}
	name := parser.Command.Active.Name

	// Apply Run method to passed sub-command, i.e. Advert.Run (advert.go)
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

// Get hexadecimal lowercase hash string
func Sha512HashValue(value string) string {
	hashHandler := sha512.New()
	hashHandler.Write([]byte(value))
	return strings.ToLower(hex.EncodeToString(hashHandler.Sum(nil)))
}

// Read URL from sub-command option. Default: http://localhost:8008
// Get keyfile path from sub-command option. Default: ~/.sawtooth/keys/$USER.priv
// Return client struct (autavail_client.go)
func GetClient(args Command, readFile bool) (AutavailClient, error) {
	url := args.UrlPassed()
	if url == "" {
		url = DEFAULT_URL
	}

	keyfile := ""
	if readFile {
		var err error
		_ = err

		// Get default keyfile path if keyfile arguement was not passed
		if args.KeyfilePassed() == "" {
			username, err := user.Current()
			if err != nil {
				return NewAutavailClient("","")
			}
			keyfile = path.Join(username.HomeDir, ".sawtooth", "keys", username.Username+".priv")
		} else {
			keyfile = args.KeyfilePassed()
		}
	}
	return NewAutavailClient (url, keyfile)
}
