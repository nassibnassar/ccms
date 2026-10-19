package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func createTrackStmt(s *svr, db *dbx.DB, rqid int64, cmd *ast.CreateStmt) *ccms.Result {
	trackID, err := cat.TrackID(db, cmd.Name)
	if err != nil {
		return cmderr(err.Error())
	}
	if trackID != 0 {
		return cmderr("track \"" + cmd.Name + "\" already exists")
	}

	if !cat.IsValidTrackName(cmd.Name) {
		return cmderr("invalid track name \"" + cmd.Name + "\"")
	}

	if err := cat.CreateTrack(db, cmd.Name); err != nil {
		return cmderr("creating track: " + err.Error())
	}

	return ccms.NewResult("create track")
}
