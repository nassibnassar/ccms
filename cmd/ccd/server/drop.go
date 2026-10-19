package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func dropStmt(s *svr, db *dbx.DB, rqid int64, cmd *ast.DropStmt) *ccms.Result {
	switch cmd.Type {
	case "filter":
		return dropFilterStmt(s, db, rqid, cmd)
	case "fund":
		return dropFundStmt(s, db, rqid, cmd)
	case "project":
		return dropProjectStmt(s, db, rqid, cmd)
	case "set":
		return dropSetStmt(s, db, rqid, cmd)
	case "track":
		return dropTrackStmt(s, db, rqid, cmd)
	default:
		return cmderr("syntax error")
	}
}
