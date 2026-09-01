module github.com/indexdata/ccms

go 1.27.0

require (
	github.com/chzyer/readline v1.5.1
	github.com/essentialkaos/ek/v13 v13.38.3
	github.com/jackc/pgx/v5 v5.9.2
	github.com/nassibnassar/goharvest v0.0.0-20160726165741-cbaf6f70f07d
	github.com/spf13/cobra v1.8.1
	golang.org/x/crypto v0.47.0
	golang.org/x/term v0.40.0
	gopkg.in/ini.v1 v1.67.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/kisielk/errcheck v1.20.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.39.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/telemetry v0.0.0-20260811182544-a038080d80e5 // indirect
	golang.org/x/text v0.33.0 // indirect
	golang.org/x/tools v0.49.0 // indirect
	golang.org/x/tools/go/packages/packagestest v0.1.1-deprecated // indirect
	golang.org/x/vuln v1.7.0 // indirect
)

tool (
	github.com/kisielk/errcheck
	golang.org/x/tools/cmd/goyacc
	golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
	golang.org/x/vuln/cmd/govulncheck
)
