package server

import (
	"cmp"
	"slices"
	"strings"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
)

func showStmt(s *svr, cmd *ast.ShowStmt) *ccms.Result {
	switch cmd.Name {
	case "filters":
		return &ccms.Result{
			Status: "show",
			Fields: []*ccms.FieldDescription{
				&ccms.FieldDescription{
					Name: "filter",
					Type: "text",
				},
			},
			Data: []*ccms.DataRow{},
		}
	case "sets":
		return &ccms.Result{
			Status: "show",
			Fields: []*ccms.FieldDescription{
				&ccms.FieldDescription{
					Name: "set",
					Type: "text",
				},
			},
			Data: data(s.cat),
		}
	default:
		return cmderr("unknown variable \"" + cmd.Name + "\"")
	}
}

func data(cat *catalog.Catalog) []*ccms.DataRow {
	rows := make([]*ccms.DataRow, 0)
	sets := cat.AllSets()
	sortSetNames(sets)
	for i := range sets {
		rows = append(rows, &ccms.DataRow{Values: []any{sets[i]}})
	}
	return rows
}

func sortSetNames(sets []string) {
	slices.SortFunc(sets, func(x, y string) int {
		a := !strings.ContainsRune(x, '.')
		b := !strings.ContainsRune(y, '.')
		if a && !b {
			return -1
		}
		if !a && b {
			return 1
		}
		return cmp.Compare(x, y)
	})
}
