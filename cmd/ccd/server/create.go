package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func createStmt(s *svr, db *dbx.DB, rqid int64, cmd *ast.CreateStmt) *ccms.Result {
	switch cmd.Type {
	case "filter":
		return createFilterStmt(s, db, rqid, cmd)
	case "fund":
		return createFundStmt(s, db, rqid, cmd)
	case "project":
		return createProjectStmt(s, db, rqid, cmd)
	case "set":
		return createSetStmt(s, db, rqid, cmd)
	case "track":
		return createTrackStmt(s, db, rqid, cmd)
	case "user":
		return createUserStmt(s, db, rqid, cmd)
	default:
		return cmderr("syntax error")
	}
}
