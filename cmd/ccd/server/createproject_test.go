package server

import (
	"testing"
)

func TestCreateProject(t *testing.T) {
	resp, _ := testclient.Send(`
create project p;
show project p;
create project p;
drop project p cascade;
`)
	if teststatus(resp, 0) != "create project" {
		t.Errorf("unexpected status")
	}
	if teststatus(resp, 1) != "show" {
		t.Errorf("created project not found")
	}
	if teststatus(resp, 2) != "error" {
		t.Errorf("unexpected status")
	}
}
