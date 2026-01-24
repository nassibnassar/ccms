package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/indexdata/ccms/cmd/ccd/config"
	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/cmd/ccd/option"
	"github.com/indexdata/ccms/cmd/ccd/osutil"
	"github.com/indexdata/ccms/cmd/ccd/server"
	"github.com/indexdata/ccms/cmd/ccd/stop"
	"github.com/indexdata/ccms/internal/eout"
	"github.com/indexdata/ccms/internal/global"
	"github.com/nassibnassar/goharvest/oai"
	"github.com/spf13/cobra"
)

func main() {
	//testOAI()
	initColor(os.Getenv("CCMS_COLOR"))
	if err := run(); err != nil {
		// fmt.Fprintf(os.Stderr, "%s: %s\n", global.ServerProgram, err)
		eout.Error("%s", err)
		os.Exit(1)
	}
}

func testOAI() {
	/*
		(&oai.Request{
			BaseURL: "http://services.kb.nl/mdo/oai", Set: "DTS", MetadataPrefix: "dcx",
			From: "2012-09-06T014:00:00.000Z",
		}).HarvestRecords(func(record *oai.Record) {
			fmt.Printf("%s\n\n", record.Metadata.Body[0:500])
		})
	*/

	//(&oai.Request{
	//        BaseURL: "https://cclp-okapi.reshare-dev.indexdata.com/_/invoke/tenant/cclp/reservoir/oai",

	//        //Verb:  "ListRecords",
	//        //From:  "2025-10-14T22:52:14Z",
	//        //Until: "2025-10-14T22:52:16Z",

	//        Verb:           "GetRecord",
	//        Identifier:     "b6c6160c-6bbb-41cc-9e07-690049d7d537",
	//        MetadataPrefix: "marcxml",
	//}).HarvestRecords(func(record *oai.Record) {
	//        //fmt.Printf("identifier: %s\n", record.Header.Identifier)
	//        //fmt.Printf("datestamp: %s\n", record.Header.DateStamp)
	//        //fmt.Printf("setspec: %v\n", record.Header.SetSpec)
	//        //fmt.Printf("status: %s\n", record.Header.Status)
	//        //fmt.Printf("about: %s\n", record.About.Body)
	//        //fmt.Printf("metadata: %s\n", record.Metadata.Body)

	//        fmt.Printf("%s %s\n", record.Header.Identifier, record.Header.DateStamp)
	//})

	// The following OAI-PMH verbs are supported by the Reservoir: ListIdentifiers, ListRecords,
	// GetRecord, Identify.
	rq := &oai.Request{
		BaseURL: "https://cclp-okapi.reshare-dev.indexdata.com/_/invoke/tenant/cclp/reservoir/oai",

		Verb:  "ListRecords",
		From:  "2025-10-14T22:52:14Z",
		Until: "2025-10-14T22:52:16Z",

		//Verb:           "GetRecord",
		//Identifier:     "b6c6160c-6bbb-41cc-9e07-690049d7d537",
		//MetadataPrefix: "marcxml",
	}
	fmt.Printf("%#v\n\n", rq)
	//rq.HarvestRecords(func(record *oai.Record) {
	rq.Harvest(func(rs *oai.Response) {
		//fmt.Printf("identifier: %s\n", record.Header.Identifier)
		//fmt.Printf("datestamp: %s\n", record.Header.DateStamp)
		//fmt.Printf("setspec: %v\n", record.Header.SetSpec)
		//fmt.Printf("status: %s\n", record.Header.Status)
		//fmt.Printf("about: %s\n", record.About.Body)
		//fmt.Printf("metadata: %s\n", record.Metadata.Body)

		//fmt.Printf("%s %s\n", record.Header.Identifier, record.Header.DateStamp)

		fmt.Printf("%#v\n", rs)
	})
	// TODO add to ccd.conf:
	// [oai]
	// base_url = https://cclp-okapi.reshare-dev.indexdata.com/_/invoke/tenant/cclp/reservoir/oai

	os.Exit(0)
}

func run() error {
	var globalOpt = option.Global{}
	var initOpt = option.Init{}
	var serverOpt = option.Server{}
	var stopOpt = option.Stop{}
	var logfile string

	var cmdStart = &cobra.Command{
		Use: "start",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			serverOpt.Global = globalOpt
			var logf, csvlogf *os.File
			if logf, csvlogf, err = setupLog(logfile, serverOpt.Debug, serverOpt.Trace); err != nil {
				return err
			}
			if err = server.Start(&serverOpt); err != nil {
				return fatal(err, logf, csvlogf)
			}
			return nil
		},
	}
	cmdStart.SetHelpFunc(help)
	_ = dirFlag(cmdStart, &serverOpt.Datadir)
	_ = logFlag(cmdStart, &logfile)
	_ = listenFlag(cmdStart, &serverOpt.Listen)
	_ = portFlag(cmdStart, &serverOpt.Port)
	_ = certFlag(cmdStart, &serverOpt.TLSCert)
	_ = keyFlag(cmdStart, &serverOpt.TLSKey)
	// _ = debugFlag(cmdStart, &serverOpt.Debug)
	// _ = traceLogFlag(cmdStart, &serverOpt.Trace)
	_ = noTLSFlag(cmdStart, &serverOpt.NoTLS)
	_ = noHarvestFlag(cmdStart, &serverOpt.NoHarvest)

	var cmdStop = &cobra.Command{
		Use: "stop",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			stopOpt.Global = globalOpt
			if err = stop.Stop(&stopOpt); err != nil {
				return err
			}
			return nil
		},
	}
	cmdStop.SetHelpFunc(help)
	_ = dirFlag(cmdStop, &stopOpt.Datadir)

	var cmdInit = &cobra.Command{
		Use: "init",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			initOpt.Global = globalOpt
			if err = initSystem(&initOpt); err != nil {
				return err
			}
			return nil
		},
	}
	cmdInit.SetHelpFunc(help)
	_ = dirFlag(cmdInit, &initOpt.Datadir)

	var cmdVersion = &cobra.Command{
		Use: "version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%s version %s\n", global.ServerProgram, global.Version)
			return nil
		},
	}
	cmdVersion.SetHelpFunc(help)

	var rootCmd = &cobra.Command{
		Use:                global.ServerProgram,
		SilenceErrors:      true,
		SilenceUsage:       true,
		DisableSuggestions: true,
		CompletionOptions:  cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	rootCmd.SetHelpFunc(help)
	// Redefine help flag without -h; so we can use it for something else.
	var helpFlag bool
	rootCmd.PersistentFlags().BoolVarP(&helpFlag, "help", "", false, "Help for "+global.ServerProgram)
	// Add commands.
	rootCmd.AddCommand(cmdStart, cmdStop, cmdInit, cmdVersion)
	var err error
	if err = rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

var helpStart = "Start server\n"
var helpStop = "Shutdown server\n"
var helpInit = "Initialize new ccms instance\n"
var helpVersion = "Print " + global.ServerProgram + " version\n"

func help(cmd *cobra.Command, commandLine []string) {
	_ = commandLine
	switch cmd.Use {
	case global.ServerProgram:
		fmt.Print("" +
			global.ServerProgram + " is the CCMS server\n" +
			"\n" +
			"Usage:  " + global.ServerProgram + " <command> [options]\n" +
			"\n" +
			"Commands:\n" +
			"  start                       - " + helpStart +
			"  stop                        - " + helpStop +
			"  init                        - " + helpInit +
			"  version                     - " + helpVersion +
			"\n" +
			"Use \"" + global.ServerProgram + " help <command>\" for more information about a command.\n")
	case "start":
		fmt.Print("" +
			helpStart +
			"\n" +
			"Usage:  " + global.ServerProgram + " start [options]\n" +
			"\n" +
			"Options:\n" +
			dirFlag(nil, nil) +
			logFlag(nil, nil) +
			listenFlag(nil, nil) +
			portFlag(nil, nil) +
			certFlag(nil, nil) +
			keyFlag(nil, nil) +
			//debugFlag(nil, nil) +
			noTLSFlag(nil, nil) +
			noHarvestFlag(nil, nil) +
			//traceLogFlag(nil, nil) +
			//memoryLimitFlag(nil, nil) +
			"")
	case "stop":
		fmt.Print("" +
			helpStop +
			"\n" +
			"Usage:  " + global.ServerProgram + " stop [options]\n" +
			"\n" +
			"Options:\n" +
			dirFlag(nil, nil) +
			//verboseFlag(nil, nil) +
			//traceFlag(nil, nil) +
			"")
	case "init":
		fmt.Print("" +
			helpInit +
			"\n" +
			"Usage:  " + global.ServerProgram + " init [options]\n" +
			"\n" +
			"Options:\n" +
			dirFlag(nil, nil) +
			"")

	case "version":
		fmt.Print("" +
			helpVersion +
			"\n" +
			"Usage:  " + global.ServerProgram + " version\n")
	default:
	}
}

func portFlag(cmd *cobra.Command, adminPort *string) string {
	if cmd != nil {
		cmd.Flags().StringVarP(adminPort, "port", "p", global.DefaultPort, "")
	}
	return "" +
		"  -p, --port <p>              - Port to listen on (default: " + global.DefaultPort + ")\n"
}

func logFlag(cmd *cobra.Command, logfile *string) string {
	if cmd != nil {
		cmd.Flags().StringVarP(logfile, "log", "l", "", "")
	}
	return "" +
		"  -l, --log <f>               - File name for server log output\n"
}

func dirFlag(cmd *cobra.Command, datadir *string) string {
	if cmd != nil {
		cmd.Flags().StringVarP(datadir, "dir", "D", "", "")
	}
	return "" +
		"  -D, --dir <d>               - Data directory\n"
}

func listenFlag(cmd *cobra.Command, listen *string) string {
	if cmd != nil {
		cmd.Flags().StringVar(listen, "listen", "", "")
	}
	return "" +
		"      --listen <a>            - Address to listen on (default: 127.0.0.1)\n"
}

func certFlag(cmd *cobra.Command, cert *string) string {
	if cmd != nil {
		cmd.Flags().StringVar(cert, "cert", "", "")
	}
	return "" +
		"      --cert <f>              - File name of server certificate, including the\n" +
		"                                CA's certificate and intermediates\n"
}

func keyFlag(cmd *cobra.Command, key *string) string {
	if cmd != nil {
		cmd.Flags().StringVar(key, "key", "", "")
	}
	return "" +
		"      --key <f>               - File name of server private key\n"
}

func noTLSFlag(cmd *cobra.Command, noTLS *bool) string {
	if cmd != nil {
		cmd.Flags().BoolVar(noTLS, "notls", false, "")
	}
	return "" +
		"      --notls                 - Disable TLS in client connections (insecure)\n"
}

func noHarvestFlag(cmd *cobra.Command, noHarvest *bool) string {
	if cmd != nil {
		cmd.Flags().BoolVar(noHarvest, "noharvest", false, "")
	}
	return "" +
		"      --noharvest             - Do not harvest records\n"
}

func setupLog(logfile string, debug bool, trace bool) (*os.File, *os.File, error) {
	var err error
	var logf, csvlogf *os.File
	if logfile != "" {
		if logfile != "" {
			if logf, err = log.OpenLogFile(logfile); err != nil {
				return nil, nil, err
			}
		}
		log.Init(logf, debug, trace)
		return logf, csvlogf, nil
	}
	log.Init(os.Stderr, debug, trace)
	return nil, nil, nil
}

func fatal(err error, logf, csvlogf *os.File) error {
	if logf != nil {
		_ = logf.Close()
	}
	if csvlogf != nil {
		_ = csvlogf.Close()
	}
	return err
}

func initSystem(opt *option.Init) error {
	// Check for required options.
	if opt.Datadir == "" {
		return fmt.Errorf("data directory not specified")
	}
	dd, err := filepath.Abs(opt.Datadir)
	if err != nil {
		return fmt.Errorf("absolute path: %w", err)
	}
	// Require that the data directory not already exist.
	exists, err := osutil.FileExists(dd)
	if err != nil {
		return fmt.Errorf("checking if path exists: %w", err)
	}
	if exists {
		return fmt.Errorf("%s already exists", dd)
	}
	// Create the data directory.
	eout.Verbose("creating data directory")
	eout.Trace("mkdir: %s", dd)
	err = os.MkdirAll(dd, osutil.ModePermRWX)
	if err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}
	f, err := os.Create(osutil.ConfigFileName(dd))
	if err != nil {
		return fmt.Errorf("creating configuration file: %w", err)
	}
	_, err = f.WriteString(config.InitStub())
	if err != nil {
		return fmt.Errorf("writing configuration file: %w", err)
	}
	err = f.Close()
	if err != nil {
		return fmt.Errorf("closing configuration file: %w", err)
	}
	eout.Info("initialized new data directory in %s", dd)
	return nil
}

func initColor(colorMode string) {
	eout.Init(global.ServerProgram)
	switch colorMode {
	case "always":
		eout.AlwaysColor()
	case "auto":
		eout.AutoColor()
	default:
		eout.NeverColor()
	}
}
