package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func dropTrackStmt(s *svr, db *dbx.DB, rqid int64, cmd *ast.DropTrackStmt) *ccms.Result {
	if err := cat.DropTrack(db, cmd.Track); err != nil {
		return cmderr(err.Error())
	}
	return ccms.NewResult("drop track")
}
