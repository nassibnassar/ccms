package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/dbx"
)

func createFundStmt(s *svr, d *dbx.DB, rqid int64, cmd *ast.CreateFundStmt) *ccms.Result {
	fundID, err := cat.FundID(d, cmd.Fund)
	if err != nil {
		return cmderr(err.Error())
	}
	if fundID != 0 {
		return cmderr("fund \"" + cmd.Fund + "\" already exists")
	}

	if !cat.IsValidFundName(cmd.Fund) {
		return cmderr("invalid fund name \"" + cmd.Fund + "\"")
	}

	if err := cat.CreateFund(d, cmd.Fund); err != nil {
		return cmderr("creating fund: " + err.Error())
	}

	return ccms.NewResult("create fund")
}
