package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func createTrackStmt(s *svr, db *dbx.DB, rqid int64, cmd *ast.CreateTrackStmt) *ccms.Result {
	trackID, err := cat.TrackID(db, cmd.Track)
	if err != nil {
		return cmderr(err.Error())
	}
	if trackID != 0 {
		return cmderr("track \"" + cmd.Track + "\" already exists")
	}

	if !cat.IsValidTrackName(cmd.Track) {
		return cmderr("invalid track name \"" + cmd.Track + "\"")
	}

	if err := cat.CreateTrack(db, cmd.Track); err != nil {
		return cmderr("creating track: " + err.Error())
	}

	return ccms.NewResult("create track")
}
