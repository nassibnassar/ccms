package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/internal/eout"
	"github.com/indexdata/ccms/internal/global"
	"github.com/spf13/cobra"
)

var colorMode string
var devMode bool

var colorInitialized bool

var option struct {
	Host          string
	Port          string
	NoTLS         bool
	TLSSkipVerify bool
	Version       bool
}

func main() {
	colorMode = os.Getenv("CCC_COLOR")
	cccMain()
}

func cccMain() {
	eout.Init(global.ClientProgram)
	run()
}

func errorExit(err error) {
	if !colorInitialized {
		eout.NeverColor()
	}
	eout.Error("%s", err)
	os.Exit(1)
}

func run() {
	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			printHelp()
		}
		if os.Args[1][0] != '-' {
			fmt.Fprintf(os.Stderr, "%s: unknown argument: %s\n", global.ClientProgram, os.Args[1])
			os.Exit(1)
		}
	}
	var rootCmd = &cobra.Command{
		Use:                global.ClientProgram,
		SilenceErrors:      true,
		SilenceUsage:       true,
		DisableSuggestions: true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runClient(); err != nil {
				errorExit(err)
			}
		},
	}
	rootCmd.SetHelpFunc(help)
	// redefine help without -h, so we can use it for --port
	var helpFlag bool
	rootCmd.PersistentFlags().BoolVarP(&helpFlag, "help", "", false, "Help for ccc")
	_ = hostFlag(rootCmd, &option.Host)
	_ = portFlag(rootCmd, &option.Port)
	_ = noTLSFlag(rootCmd, &option.NoTLS)
	_ = skipVerifyFlag(rootCmd, &option.TLSSkipVerify)
	_ = versionFlag(rootCmd, &option.Version)
	if err := rootCmd.Execute(); err != nil {
		errorExit(err)
	}
}

func runClient() error {
	if option.Version {
		fmt.Printf("%s %s\n", global.ClientProgram, global.Version)
		return nil
	}
	if option.NoTLS && option.Host != "127.0.0.1" {
		eout.Warning("disabling TLS (insecure)")
	}

	rl, err := readline.New("=> ")
	if err != nil {
		return err
	}
	defer rl.Close()

	client := &ccms.Client{
		Host:          option.Host,
		Port:          option.Port,
		User:          "nemo",
		Password:      "testpass",
		NoTLS:         option.NoTLS,
		TLSSkipVerify: option.TLSSkipVerify,
	}
	if _, err = client.Send("ping"); err != nil {
		errorExit(err)
	}
	eout.Interactive()
	fmt.Printf("%s %s: type \"help\" for help, \"quit\" to quit\n",
		global.ClientProgram, global.Version)
	for {
		rline, err := rl.Readline()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			continue
		}
		line := strings.TrimSpace(rline)
		if line == "" {
			continue
		}
		if line == "quit" {
			break
		}
		resp, err := client.Send(line)
		if err != nil {
			eout.Error("%v", err)
		}
		if resp.Status == "error" {
			eout.Error("%s", resp.Message)
			continue
		}
		if resp.Status == "help" {
			fmt.Println(resp.Message)
			continue
		}
		if resp.Status == "ping" {
			continue
		}
		for i := range resp.Fields {
			if i != 0 {
				fmt.Print("\t")
			}
			fmt.Print(resp.Fields[i].Name)
		}
		fmt.Print("\n")
		for i := range resp.Data {
			for j := range resp.Data[i].Values {
				if j != 0 {
					fmt.Print("\t")
				}
				fmt.Print(resp.Data[i].Values[j])
			}
			fmt.Print("\n")
		}
	}

	return nil
}

func help(cmd *cobra.Command, commandLine []string) {
	printHelp()
}

func printHelp() {
	fmt.Print("" +
		global.ClientProgram + " is the CCMS client\n" +
		"\n" +
		"Usage:  " + global.ClientProgram + " [options]\n" +
		"\n" +
		"Options:\n" +
		hostFlag(nil, nil) +
		portFlag(nil, nil) +
		noTLSFlag(nil, nil) +
		skipVerifyFlag(nil, nil) +
		versionFlag(nil, nil) +
		"")
	os.Exit(0)
}

func versionFlag(cmd *cobra.Command, version *bool) string {
	if cmd != nil {
		cmd.Flags().BoolVar(version, "version", false, "")
	}
	return "" +
		"      --version               print ccc version\n"
}

func verboseFlag(cmd *cobra.Command, verbose *bool) string {
	if cmd != nil {
		cmd.Flags().BoolVarP(verbose, "verbose", "v", false, "")
	}
	return "" +
		"  -v, --verbose               enable verbose output\n"
}

func hostFlag(cmd *cobra.Command, host *string) string {
	if cmd != nil {
		cmd.Flags().StringVarP(host, "host", "h", "127.0.0.1", "")
	}
	return "" +
		"  -h, --host <h>              server host (default: 127.0.0.1)\n"
}

func portFlag(cmd *cobra.Command, port *string) string {
	if cmd != nil {
		cmd.Flags().StringVarP(port, "port", "p", global.DefaultPort, "")
	}
	return "" +
		"  -p, --port <p>              server port (default: " + global.DefaultPort + ")\n"
}

func noTLSFlag(cmd *cobra.Command, noTLS *bool) string {
	if cmd != nil {
		cmd.Flags().BoolVar(noTLS, "notls", false, "")
	}
	return "" +
		"      --notls                 disable TLS in connection to server (insecure)\n"
}

func skipVerifyFlag(cmd *cobra.Command, skipVerify *bool) string {
	if cmd != nil {
		cmd.Flags().BoolVar(skipVerify, "skipverify", false, "")
	}
	return "" +
		"      --skipverify            do not verify server certificate chain and host\n" +
		"                              name (insecure)\n"
}

func initColor() error {
	switch colorMode {
	case "always":
		eout.AlwaysColor()
	case "auto":
		eout.AutoColor()
	default:
		eout.NeverColor()
	}
	colorInitialized = true
	return nil
}

func ValueFromFile(filename string) (string, error) {
	var err error
	var f *os.File
	if f, err = os.Open(filename); err != nil {
		return "", err
	}
	var scanner = bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	var ok bool
	if ok = scanner.Scan(); !ok {
		if scanner.Err() != nil {
			return "", err
		}
	}
	return scanner.Text(), nil
}
