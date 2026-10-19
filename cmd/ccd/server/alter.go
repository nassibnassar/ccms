package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func alterStmt(s *svr, db *dbx.DB, rqid int64, cmd *ast.AlterStmt) *ccms.Result {
	switch cmd.Type {
	case "fund":
		return alterFundStmt(s, db, rqid, cmd)
	case "project":
		return alterProjectStmt(s, db, rqid, cmd)
	case "set":
		return alterSetStmt(s, db, rqid, cmd)
	case "track":
		return alterTrackStmt(s, db, rqid, cmd)
	default:
		return cmderr("syntax error")
	}
}
