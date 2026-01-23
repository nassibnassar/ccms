package doc

import (
	_ "embed"
	"strings"
)

//go:embed select.txt
var selectTxt string

func Select() string {
	return strings.TrimSpace(selectTxt)
}

//go:embed showfilters.txt
var showfiltersTxt string

func ShowFilters() string {
	return strings.TrimSpace(showfiltersTxt)
}

//go:embed showsets.txt
var showsetsTxt string

func ShowSets() string {
	return strings.TrimSpace(showsetsTxt)
}
