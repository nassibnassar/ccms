package process

import (
	"fmt"
	"os"
	"sync/atomic"
	"syscall"

	"github.com/indexdata/ccms/cmd/ccd/osutil"
)

var stop int32

func SetStop() {
	atomic.StoreInt32(&stop, 1)
}

func Stop() bool {
	return atomic.LoadInt32(&stop) == 1
}

func ReadPIDFile(datadir string) (int, error) {
	var err error
	var f *os.File
	if f, err = os.Open(osutil.SystemPIDFileName(datadir)); err != nil {
		return 0, err
	}
	var pid int
	var n int
	if n, err = fmt.Fscanf(f, "%d", &pid); err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, fmt.Errorf("unable to read data from file: %s", osutil.SystemPIDFileName(datadir))
	}
	if pid <= 0 {
		return 0, fmt.Errorf("invalid data in file: %s", osutil.SystemPIDFileName(datadir))
	}
	return pid, nil
}

func WritePIDFile(datadir string) error {
	var err error
	var f *os.File
	if f, err = os.OpenFile(osutil.SystemPIDFileName(datadir), os.O_RDWR|os.O_CREATE, osutil.ModePermRW); err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if _, err = f.WriteString(fmt.Sprintf("%d\n", os.Getpid())); err != nil {
		return err
	}
	return nil
}

func RemovePIDFile(datadir string) {
	_ = os.Remove(osutil.SystemPIDFileName(datadir))
}

func IsServerRunning(datadir string) (bool, int, error) {
	// check for lock file
	lockfile := osutil.SystemPIDFileName(datadir)
	fexists, err := osutil.FileExists(lockfile)
	if err != nil {
		return false, 0, fmt.Errorf("reading lock file %q: %s", lockfile, err)
	}
	if !fexists {
		return false, 0, nil
	}
	// read pid
	pid, err := ReadPIDFile(datadir)
	if err != nil {
		return false, 0, fmt.Errorf("reading lock file %q: %s", lockfile, err)
	}
	// check for running process
	p, err := os.FindProcess(pid)
	if err != nil {
		return false, 0, nil
	}
	err = p.Signal(syscall.Signal(0))
	if err != nil {
		errno, ok := err.(syscall.Errno)
		if !ok {
			return false, 0, nil
		}
		if errno == syscall.EPERM {
			return true, pid, nil
		} else {
			return false, 0, nil
		}
	}
	return true, pid, nil
}
