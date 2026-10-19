package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func createFundStmt(s *svr, db *dbx.DB, rqid int64, cmd *ast.CreateStmt) *ccms.Result {
	fundID, err := cat.FundID(db, cmd.Name)
	if err != nil {
		return cmderr(err.Error())
	}
	if fundID != 0 {
		return cmderr("fund \"" + cmd.Name + "\" already exists")
	}

	if !cat.IsValidFundName(cmd.Name) {
		return cmderr("invalid fund name \"" + cmd.Name + "\"")
	}

	if err := cat.CreateFund(db, cmd.Name); err != nil {
		return cmderr("creating fund: " + err.Error())
	}

	return ccms.NewResult("create fund")
}
