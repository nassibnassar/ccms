package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
)

func createFundStmt(s *svr, rqid int64, cmd *ast.CreateFundStmt) *ccms.Result {
	exists, err := s.cat.FundExists(cmd.FundName)
	if err != nil {
		return cmderr(err.Error())
	}
	if exists {
		return cmderr("fund \"" + cmd.FundName + "\" already exists")
	}

	if err := s.cat.CreateFund(cmd.FundName); err != nil {
		return cmderr(err.Error())
	}

	return ccms.NewResult("create fund")
}
