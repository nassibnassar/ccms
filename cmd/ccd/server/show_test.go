package server

import "testing"

func TestShowFilters(t *testing.T) {
	resp, _ := testclient.Send(`
create project p;
create filter p.f1 where id > 100;
create filter p.f2 where id > 200;
show filters;
drop filter p.f2;
drop filter p.f1;
drop project p cascade;
`)
	if testresult(resp, 3, 0, 0) != "p" {
		t.Errorf("unexpected value at 0, 0")
	}
	if testresult(resp, 3, 0, 1) != "f1" {
		t.Errorf("unexpected value at 0, 1")
	}
	if testresult(resp, 3, 0, 2) != "create filter p.f1 where id > 100" {
		t.Errorf("unexpected value at 0, 2")
	}
	if testresult(resp, 3, 1, 0) != "p" {
		t.Errorf("unexpected value at 1, 0")
	}
	if testresult(resp, 3, 1, 1) != "f2" {
		t.Errorf("unexpected value at 1, 1")
	}
	if testresult(resp, 3, 1, 2) != "create filter p.f2 where id > 200" {
		t.Errorf("unexpected value at 1, 2")
	}
}

func TestShowFiltersInProject(t *testing.T) {
	resp, _ := testclient.Send(`
create project p1;
create project p2;
create filter p1.f1 where id > 100;
create filter p1.f2 where id > 200;
create filter p2.f3 where id > 300;
show filters in project p1;
show filters in project p2;
drop filter p2.f3;
drop filter p1.f2;
drop filter p1.f1;
drop project p2 cascade;
drop project p1 cascade;
`)
	if testresult(resp, 5, 0, 0) != "p1" {
		t.Errorf("unexpected value at 0")
	}
	if testresult(resp, 5, 0, 1) != "f1" {
		t.Errorf("unexpected value at 1")
	}
	if testresult(resp, 5, 0, 2) != "create filter p1.f1 where id > 100" {
		t.Errorf("unexpected value at 2")
	}
	if testresult(resp, 5, 1, 0) != "p1" {
		t.Errorf("unexpected value at 3")
	}
	if testresult(resp, 5, 1, 1) != "f2" {
		t.Errorf("unexpected value at 4")
	}
	if testresult(resp, 5, 1, 2) != "create filter p1.f2 where id > 200" {
		t.Errorf("unexpected value at 5")
	}
	if testresult(resp, 5, 2, 0) != "" {
		t.Errorf("unexpected value at 6")
	}

	if testresult(resp, 6, 0, 0) != "p2" {
		t.Errorf("unexpected value at 7")
	}
	if testresult(resp, 6, 0, 1) != "f3" {
		t.Errorf("unexpected value at 8")
	}
	if testresult(resp, 6, 0, 2) != "create filter p2.f3 where id > 300" {
		t.Errorf("unexpected value at 9")
	}
	if testresult(resp, 6, 1, 0) != "" {
		t.Errorf("unexpected value at 10")
	}
}

func TestShowFund(t *testing.T) {
	resp, _ := testclient.Send(`
create fund f;
show fund f;
drop fund f;
`)
	if teststatus(resp, 1) != "show" {
		t.Errorf("created fund not found")
	}
	if testresult(resp, 1, 0, 0) != "name" {
		t.Errorf("unexpected property")
	}
	if testresult(resp, 1, 0, 1) == "" {
		t.Errorf("unexpected value")
	}
}

func TestShowFunds(t *testing.T) {
	resp, _ := testclient.Send(`
create fund f1;
create fund f2;
show funds;
drop fund f2;
drop fund f1;
`)
	if testresult(resp, 2, 0, 0) != "f1" {
		t.Errorf("unexpected value at 0, 0")
	}
	if testresult(resp, 2, 0, 1) == "" {
		t.Errorf("unexpected value at 0, 1")
	}
	if testresult(resp, 2, 1, 0) != "f2" {
		t.Errorf("unexpected value at 1, 0")
	}
	if testresult(resp, 2, 1, 1) == "" {
		t.Errorf("unexpected value at 1, 1")
	}
}

func TestShowProject(t *testing.T) {
	resp, _ := testclient.Send(`
create project p;
show project p;
drop project p cascade;
`)
	if teststatus(resp, 1) != "show" {
		t.Errorf("created project not found")
	}
	if testresult(resp, 1, 0, 0) != "title" {
		t.Errorf("unexpected property")
	}
	if testresult(resp, 1, 0, 1) == "" {
		t.Errorf("unexpected value")
	}
}

func TestShowProjects(t *testing.T) {
	resp, _ := testclient.Send(`
create project p1;
create project p2;
show projects;
drop project p2 cascade;
drop project p1 cascade;
`)
	if testresult(resp, 2, 0, 0) != "p1" {
		t.Errorf("unexpected value at 0, 0")
	}
	if testresult(resp, 2, 0, 1) == "" {
		t.Errorf("unexpected value at 0, 1")
	}
	if testresult(resp, 2, 1, 0) != "p2" {
		t.Errorf("unexpected value at 1, 0")
	}
	if testresult(resp, 2, 1, 1) == "" {
		t.Errorf("unexpected value at 1, 1")
	}
}
