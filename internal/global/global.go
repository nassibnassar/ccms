package global

import "path/filepath"

// Version is defined at build time via -ldflags.
var Version = ""

const DefaultPort = "8504"

const ServerProgram = "ccd"

const ClientProgram = "ccc"

func ServerConfigFileName(datadir string) string {
	return filepath.Join(datadir, "ccd.conf")
}
