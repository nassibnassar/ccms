package server

import (
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

	// _, _ = testclient.Send("drop project p cascade")

	m.Run()
}

func setupTest() {
	srv := option.Server{
		Debug:     false,
		Trace:     false,
		Datadir:   "/home/nrn/work/cclp/data/localtest",
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

func teststatus(resp *ccms.Response, resultIndex int) string {
	i := 0
	for result := range resp.Results() {
		if i == resultIndex {
			return result.Status()
		}
		i++
	}
	return ""
}

func testresult(resp *ccms.Response, resultIndex, row, column int) string {
	// testprintresult(resp)
	i := 0
	for result := range resp.Results() {
		j := 0
		for data := range result.Data() {
			values := data.Values()
			for k := range values {
				if i == resultIndex && j == row && k == column {
					str, ok := values[k].(string)
					if ok {
						return str
					} else {
						return ""
					}
				}
			}
			j++
		}
		i++
	}
	return ""
}

/*
func testprintresult(resp *ccms.Response) {
	i := 0
	for result := range resp.Results() {
		if i > 0 {
			fmt.Printf("\n")
		}
		j := 0
		for data := range result.Data() {
			values := data.Values()
			for k := range values {
				if k > 0 {
					fmt.Printf("\t")
				}
				fmt.Printf("%v", values[k])
			}
			fmt.Printf("\n")
			j++
		}
		i++
	}
}
*/
