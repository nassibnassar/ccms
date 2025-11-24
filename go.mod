module github.com/indexdata/ccms

go 1.25.4

require (
	github.com/chzyer/readline v1.5.1
	github.com/fatih/color v1.18.0
	github.com/mattn/go-isatty v0.0.20
	github.com/nassibnassar/goharvest v0.0.0-20160726165741-cbaf6f70f07d
	github.com/spf13/cobra v1.8.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kisielk/errcheck v1.8.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/telemetry v0.0.0-20240522233618-39ace7a40ae7 // indirect
	golang.org/x/tools v0.30.0 // indirect
	golang.org/x/vuln v1.1.4 // indirect
)

tool (
	github.com/kisielk/errcheck
	golang.org/x/tools/cmd/goyacc
	golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
	golang.org/x/vuln/cmd/govulncheck
)
