package server

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/indexdata/ccms"
	clog "github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/cmd/ccd/option"
)

func TestMain(m *testing.M) {
	go setupTest()
	time.Sleep(1 * time.Second)

	client := &ccms.Client{
		Host:          "localhost",
		Port:          "8503",
		User:          "nrn",
		Password:      "testpass49",
		NoTLS:         true,
		TLSSkipVerify: true,
	}
	_, _ = client.Send("drop project p cascade")

	m.Run()
}

func setupTest() {
	srv := option.Server{
		Debug:     false,
		Trace:     false,
		Datadir:   "/home/nrn/work/cclp/data/test",
		Listen:    "localhost",
		Port:      "8503",
		TLSCert:   "",
		TLSKey:    "",
		NoTLS:     true,
		NoHarvest: true,
		Program:   "ccd_test",
		Test:      true,
	}
	clog.Init(os.Stderr, false, false)
	if err := Start(&srv); err != nil {
		panic(err)
	}
}

var testclient = &ccms.Client{
	Host:          "localhost",
	Port:          "8503",
	User:          "nrn",
	Password:      "testpass49",
	NoTLS:         true,
	TLSSkipVerify: true,
}

func teststatus(resp *ccms.Response) string {
	for result := range resp.Results() {
		return result.Status()
	}
	return ""
}

func testresult(resp *ccms.Response, resultIndex, row, column int) string {
	i := 0
	for result := range resp.Results() {
		j := 0
		for data := range result.Data() {
			values := data.Values()
			for k := range values {
				// fmt.Printf("\t%v", values[k])
				if i == resultIndex && j == row && k == column {
					str, ok := values[k].(string)
					if ok {
						return str
					} else {
						return ""
					}
				}
			}
			fmt.Printf("\n")
			j++
		}
		i++
	}
	return ""
}
