package doc

import (
	_ "embed"
	"strings"
)

/*
//go:embed info.txt
var infoTxt string

func Info() string {
	return strings.TrimSpace(infoTxt)
}
*/

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
