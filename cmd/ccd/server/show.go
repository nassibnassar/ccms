package server

import (
	"cmp"
	"slices"
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
	"github.com/indexdata/ccms/internal/protocol"
)

func showStmt(s *svr, cmd *ast.ShowStmt) *protocol.CommandResponse {
	switch cmd.Name {
	case "filters":
		return &protocol.CommandResponse{
			Status: "show",
			Fields: []protocol.FieldDescription{
				{
					Name: "filter",
				},
			},
			Data: []protocol.DataRow{},
		}
	case "sets":
		return &protocol.CommandResponse{
			Status: "show",
			Fields: []protocol.FieldDescription{
				{
					Name: "set",
				},
			},
			Data: data(s.cat),
		}
	default:
		return cmderr("unknown variable \"" + cmd.Name + "\"")
	}
}

func data(cat *catalog.Catalog) []protocol.DataRow {
	rows := make([]protocol.DataRow, 0)
	sets := cat.AllSets()
	sortSetNames(sets)
	for i := range sets {
		rows = append(rows, protocol.DataRow{Values: []any{sets[i]}})
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
