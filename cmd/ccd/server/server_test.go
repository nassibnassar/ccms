package server

import (
	"os"
	"testing"
	"time"

	clog "github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/cmd/ccd/option"
)

func TestMain(m *testing.M) {
	go setupTest()
	time.Sleep(3 * time.Second)
	m.Run()
}

func setupTest() {
	srv := option.Server{
		Debug:     false,
		Trace:     false,
		Datadir:   "/home/nrn/work/cclp/data/ccmslocal",
		Listen:    "localhost",
		Port:      "8503",
		TLSCert:   "",
		TLSKey:    "",
		NoTLS:     true,
		NoHarvest: true,
		Program:   "ccd_test",
	}
	clog.Init(os.Stderr, false, false)
	if err := Start(&srv); err != nil {
		panic(err)
	}
}
