package server

import (
	"strings"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
	"github.com/indexdata/ccms/internal/util"
)

func createFilterStmt(s *svr, db *dbx.DB, rqid int64, cmd *ast.CreateStmt) *ccms.Result {
	project, filter, err := util.ParsePair(cmd.Name)
	if err != nil {
		return cmderr(err.Error())
	}

	// is target filter valid?
	if project == "" || filter == "" {
		return cmderr("invalid filter name \"" + cmd.Name + "\"")
	}
	projectID, err := cat.ProjectID(db, project)
	if err != nil {
		return cmderr(err.Error())
	}
	if projectID == 0 {
		return cmderr("invalid filter name \"" + cmd.Name +
			"\" (project \"" + project + "\" does not exist)")
	}

	filterExists, err := cat.FilterExists(db, projectID, filter)
	if err != nil {
		return cmderr(err.Error())
	}
	if filterExists {
		return cmderr("filter \"" + cmd.Name + "\" already exists")
	}

	if !cmd.Where.(*ast.WhereClause).Valid {
		return cmderr("required \"where\" clause is missing")
	}

	var cmdsql strings.Builder
	cmdsql.WriteString("create filter ")
	cmdsql.WriteString(cmd.Name)
	cmdsql.WriteString(" where ")

	sql, err := cmd.SQL(db, &cmdsql)
	if err != nil {
		return cmderr(err.Error())
	}

	if err := cat.CreateFilter(db, projectID, filter, cmdsql.String(), sql); err != nil {
		return cmderr(err.Error())
	}
	return ccms.NewResult("create filter")
}
