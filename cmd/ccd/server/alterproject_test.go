package server

import (
	"log"
	"testing"

	"github.com/indexdata/ccms"
)

func TestCreateProject(t *testing.T) {
	cmd := "create project pppp"
	client := &ccms.Client{
		Host:          "localhost",
		Port:          "8503",
		User:          "nrn",
		Password:      "testpass49",
		NoTLS:         true,
		TLSSkipVerify: true,
	}
	resp, err := client.Send(cmd)
	if err != nil {
		t.Errorf("%v", err)
	}
	_ = resp
	log.Printf("%v\n", resp)

	// log.Println("in test")
	t.Errorf("my error")
}

func TestSomething(t *testing.T) {
	// log.Println("in test")
	// t.Errorf("my error")
}

func TestSomethingElse(t *testing.T) {
	// log.Println("in test 2")
	// t.Errorf("my error")
}
