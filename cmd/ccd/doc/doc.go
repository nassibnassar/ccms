package doc

import (
	_ "embed"
	"strings"
)

//go:embed createset.txt
var createSetTxt string

func CreateSet() string {
	return strings.TrimSpace(createSetTxt)
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
