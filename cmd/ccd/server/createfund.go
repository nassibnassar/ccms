package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func createFundStmt(s *svr, db *dbx.DB, rqid int64, cmd *ast.CreateFundStmt) *ccms.Result {
	fundID, err := cat.FundID(db, cmd.Fund)
	if err != nil {
		return cmderr(err.Error())
	}
	if fundID != 0 {
		return cmderr("fund \"" + cmd.Fund + "\" already exists")
	}

	if !cat.IsValidFundName(cmd.Fund) {
		return cmderr("invalid fund name \"" + cmd.Fund + "\"")
	}

	if err := cat.CreateFund(db, cmd.Fund); err != nil {
		return cmderr("creating fund: " + err.Error())
	}

	return ccms.NewResult("create fund")
}
