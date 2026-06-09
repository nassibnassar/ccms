package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func createFundStmt(s *svr, rqid int64, cmd *ast.CreateFundStmt) *ccms.Result {
	fundID, err := cat.FundID(s.d, cmd.Fund)
	if err != nil {
		return cmderr(err.Error())
	}
	if fundID != 0 {
		return cmderr("fund \"" + cmd.Fund + "\" already exists")
	}

	if err := cat.CreateFund(s.d, cmd.Fund); err != nil {
		return cmderrint("creating fund", err)
	}

	return ccms.NewResult("create fund")
}
