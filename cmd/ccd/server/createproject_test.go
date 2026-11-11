package server

import (
	"testing"
)

func TestCreateProject2(t *testing.T) {
	TestCreateProject(t)
}

func TestCreateProject(t *testing.T) {
	// client := &ccms.Client{
	// 	Host:          "localhost",
	// 	Port:          "8503",
	// 	User:          "nrn",
	// 	Password:      "testpass49",
	// 	NoTLS:         true,
	// 	TLSSkipVerify: true,
	// }

	resp, err := testclient.Send("create project p")
	if err != nil {
		t.Fatal(err)
		return
	}
	got := teststatus(resp)
	want := "create project"
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
		return
	}
	resp, err = testclient.Send("show project p")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if teststatus(resp) != "show" {
		t.Errorf("created project not found")
	}

	resp, err = testclient.Send("drop project p cascade")
	if err != nil {
		t.Fatal(err)
		return
	}
	got = teststatus(resp)
	want = "drop project"
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
		return
	}

	// for result := range resp.Results() {
	// 	for data := range result.Data() {
	// 		values := data.Values()
	// 		for k := range values {
	// 			fmt.Printf("\t%v", values[k])
	// 		}
	// 		fmt.Printf("\n")
	// 	}
	// }
	// if testresult(resp, 0, 0, 0) != "title" {
	// 	t.Errorf("created project not found")
	// }

	// log.Printf("%v\n", resp)

	// log.Println("in test")
	// t.Errorf("my error")
}
