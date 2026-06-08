package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func createFundStmt(s *svr, rqid int64, cmd *ast.CreateFundStmt) *ccms.Result {
	exists, err := cat.FundExists(s.d, cmd.FundName)
	if err != nil {
		return cmderr(err.Error())
	}
	if exists {
		return cmderr("fund \"" + cmd.FundName + "\" already exists")
	}

	if err := cat.CreateFund(s.d, cmd.FundName); err != nil {
		return cmderrint("creating fund", err)
	}

	return ccms.NewResult("create fund")
}
