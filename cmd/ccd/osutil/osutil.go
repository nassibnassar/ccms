package osutil

import (
	"net"
	"os"
	"path/filepath"

	"github.com/indexdata/ccms"
)

// ModePermRW is the umask "-rw-------".
const ModePermRW = 0600

// ModePermRWX is the umask "-rwx------".
const ModePermRWX = 0700

// FileExists returns true if f is an existing file or directory.
func FileExists(f string) (bool, error) {
	_, err := os.Stat(f)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SystemPIDFileName(datadir string) string {
	return filepath.Join(datadir, ccms.ServerProgram+".pid")
}

func ConfigFileName(datadir string) string {
	return filepath.Join(datadir, ccms.ServerProgram+".conf")
}

func AddrHost(addr string) string {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return ""
	}
	return host
}
