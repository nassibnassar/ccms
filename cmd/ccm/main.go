package main

import (
	"bufio"
	"os"

	"github.com/chzyer/readline"
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccm/client"
	"github.com/indexdata/ccms/internal/eout"
	"github.com/indexdata/ccms/internal/eout/color"
	"github.com/spf13/cobra"
)

var colorMode string
var devMode bool

var colorInitialized bool

func main() {
	ccmMain()
}

func ccmMain() {
	// Initialize error output
	eout.Init(ccms.ClientProgram)
	// Run
	var err error
	if err = run(); err != nil {
		if !colorInitialized {
			color.NeverColor()
		}
		eout.Error("%s", err)
		os.Exit(1)
	}
}

func run() error {

	var rootCmd = &cobra.Command{
		Use:                "ccm",
		SilenceErrors:      true,
		SilenceUsage:       true,
		DisableSuggestions: true,
	}
	rootCmd.SetHelpFunc(help)
	// Redefine help flag without -h; so we can use it for something else.
	var helpFlag bool
	rootCmd.PersistentFlags().BoolVarP(&helpFlag, "help", "", false, "Help for ccm")
	// Add commands.
	// rootCmd.AddCommand(cmdConfig, cmdUser, cmdEnable, cmdDisable, cmdStatus, cmdVersion, cmdCompletion)
	// var err error
	// if err = rootCmd.Execute(); err != nil {
	// 	return err
	// }

	rl, err := readline.New("> ")
	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil || line == "quit" || line == "exit" {
			break
		}
		client.Send(line)
	}

	return nil
}

var helpConfig = "Configure or show server settings\n"

func help(cmd *cobra.Command, commandLine []string) {
	_ = commandLine
}

func verboseFlag(cmd *cobra.Command, verbose *bool) string {
	if cmd != nil {
		cmd.Flags().BoolVarP(verbose, "verbose", "v", false, "")
	}
	return "" +
		"  -v, --verbose               - Enable verbose output\n"
}

func traceFlag(cmd *cobra.Command, trace *bool) string {
	if devMode {
		if cmd != nil {
			cmd.Flags().BoolVar(trace, "xtrace", false, "")
		}
		return "" +
			"      --xtrace                - Enable extremely verbose output\n"
	}
	return ""
}

func skipVerifyFlag(cmd *cobra.Command, skipVerify *bool) string {
	if cmd != nil {
		cmd.Flags().BoolVar(skipVerify, "skipverify", false, "")
	}
	return "" +
		"      --skipverify            - Do not verify server certificate chain and host\n" +
		"                                name [insecure, use for testing only]\n"
}

func initColor() error {
	switch colorMode {
	case "always":
		color.AlwaysColor()
	case "auto":
		color.AutoColor()
	default:
		color.NeverColor()
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
