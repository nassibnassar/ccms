package doc

import (
	_ "embed"
	"strings"
)

//go:embed alterproject.txt
var alterProjectTxt string

func AlterProject() string {
	return strings.TrimSpace(alterProjectTxt)
}

//go:embed createproject.txt
var createProjectTxt string

func CreateProject() string {
	return strings.TrimSpace(createProjectTxt)
}

//go:embed createset.txt
var createSetTxt string

func CreateSet() string {
	return strings.TrimSpace(createSetTxt)
}

//go:embed createuser.txt
var createUserTxt string

func CreateUser() string {
	return strings.TrimSpace(createUserTxt)
}

//go:embed delete.txt
var deleteTxt string

func Delete() string {
	return strings.TrimSpace(deleteTxt)
}

//go:embed dropset.txt
var dropSetTxt string

func DropSet() string {
	return strings.TrimSpace(dropSetTxt)
}

/*
//go:embed info.txt
var infoTxt string

func Info() string {
	return strings.TrimSpace(infoTxt)
}
*/

//go:embed insert.txt
var insertTxt string

func Insert() string {
	return strings.TrimSpace(insertTxt)
}

//go:embed select.txt
var selectTxt string

func Select() string {
	return strings.TrimSpace(selectTxt)
}

//go:embed show.txt
var showTxt string

func Show() string {
	return strings.TrimSpace(showTxt)
}
