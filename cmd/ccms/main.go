package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccms/log"
	"github.com/indexdata/ccms/cmd/ccms/option"
	"github.com/indexdata/ccms/cmd/ccms/osutil"
	"github.com/indexdata/ccms/cmd/ccms/server"
	"github.com/indexdata/ccms/cmd/ccms/stop"
	"github.com/indexdata/ccms/internal/eout"
	"github.com/indexdata/ccms/internal/eout/color"
	"github.com/spf13/cobra"
)

const defaultListen = "127.0.0.1"

func main() {
	initColor(os.Getenv("CCMS_COLOR"))
	if err := run(); err != nil {
		// fmt.Fprintf(os.Stderr, "%s: %s\n", ccms.ServerProgram, err)
		eout.Error("%s", err)
		os.Exit(1)
	}
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
			serverOpt.Listen = defaultListen
			if err = server.Start(&serverOpt); err != nil {
				return fatal(err, logf, csvlogf)
			}
			return nil
		},
	}
	cmdStart.SetHelpFunc(help)
	_ = dirFlag(cmdStart, &serverOpt.Datadir)
	_ = logFlag(cmdStart, &logfile)
	//_ = csvlogFlag(cmdStart, &csvlogfile)
	//_ = listenFlag(cmdStart, &serverOpt.Listen)
	_ = portFlag(cmdStart, &serverOpt.Port)
	//_ = certFlag(cmdStart, &serverOpt.TLSCert)
	//_ = keyFlag(cmdStart, &serverOpt.TLSKey)
	// _ = debugFlag(cmdStart, &serverOpt.Debug)
	// _ = traceLogFlag(cmdStart, &serverOpt.Trace)

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
			fmt.Printf("%s version %s\n", ccms.ServerProgram, ccms.Version)
			return nil
		},
	}
	cmdVersion.SetHelpFunc(help)

	var rootCmd = &cobra.Command{
		Use:                ccms.ServerProgram,
		SilenceErrors:      true,
		SilenceUsage:       true,
		DisableSuggestions: true,
		CompletionOptions:  cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	rootCmd.SetHelpFunc(help)
	// Redefine help flag without -h; so we can use it for something else.
	var helpFlag bool
	rootCmd.PersistentFlags().BoolVarP(&helpFlag, "help", "", false, "Help for "+ccms.ServerProgram)
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
var helpVersion = "Print " + ccms.ServerProgram + " version\n"

func help(cmd *cobra.Command, commandLine []string) {
	_ = commandLine
	switch cmd.Use {
	case ccms.ServerProgram:
		fmt.Print("" +
			ccms.ServerProgram + " server\n" +
			"\n" +
			"Usage:  " + ccms.ServerProgram + " <command> <arguments>\n" +
			"\n" +
			"Commands:\n" +
			"  start                       - " + helpStart +
			"  stop                        - " + helpStop +
			"  init                        - " + helpInit +
			"  version                     - " + helpVersion +
			"\n" +
			"Use \"" + ccms.ServerProgram + " help <command>\" for more information about a command.\n")
	case "start":
		fmt.Print("" +
			helpStart +
			"\n" +
			"Usage:  " + ccms.ServerProgram + " start <options>\n" +
			"\n" +
			"Options:\n" +
			dirFlag(nil, nil) +
			logFlag(nil, nil) +
			//csvlogFlag(nil, nil) +
			//listenFlag(nil, nil) +
			//portFlag(nil, nil) +
			//certFlag(nil, nil) +
			//keyFlag(nil, nil) +
			//debugFlag(nil, nil) +
			//noTLSFlag(nil, nil) +
			//traceLogFlag(nil, nil) +
			//noKafkaCommitFlag(nil, nil) +
			//logSourceFlag(nil, nil) +
			//memoryLimitFlag(nil, nil) +
			"")
	case "stop":
		fmt.Print("" +
			helpStop +
			"\n" +
			"Usage:  " + ccms.ServerProgram + " stop <options>\n" +
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
			"Usage:  " + ccms.ServerProgram + " init <options>\n" +
			"\n" +
			"Options:\n" +
			dirFlag(nil, nil) +
			"")

	case "version":
		fmt.Print("" +
			helpVersion +
			"\n" +
			"Usage:  " + ccms.ServerProgram + " version\n")
	default:
	}
}

func portFlag(cmd *cobra.Command, adminPort *string) string {
	if cmd != nil {
		cmd.Flags().StringVarP(adminPort, "port", "p", ccms.DefaultPort, "")
	}
	return "" +
		"  -p, --port <p>              - Port to listen on (default: " + ccms.DefaultPort + ")\n"
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
	var s = "[main]\n" +
		"host = \n" +
		"port = 5432\n" +
		"database = \n" +
		"user = ccms\n" +
		"password = \n" +
		"sslmode = require\n"
	_, err = f.WriteString(s)
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
	eout.Init(ccms.ServerProgram)
	switch colorMode {
	case "always":
		color.AlwaysColor()
	case "auto":
		color.AutoColor()
	default:
		color.NeverColor()
	}
}
