package server

import (
	"context"
	"fmt"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/internal/global"
)

func deleteStmt(s *svr, rqid int64, cmd *ast.DeleteStmt) *ccms.Result {
	if !s.cat.IsValidTargetSet(cmd.From) {
		return cmderr("invalid target set \"" + cmd.From + "\"")
	}

	if !s.cat.SetExists(cmd.From) {
		return cmderr("set \"" + cmd.From + "\" does not exist")
	}

	sql, err := cmd.SQL()
	if err != nil {
		return cmderr(err.Error())
	}
	//log.Info("[%d] %s", rqid, sql)
	if _, err := s.dp.Exec(context.TODO(), sql); err != nil {
		return cmderr(fmt.Sprintf("deleting from %q: %v", cmd.From, global.PGErr(err)))
	}

	return ccms.NewResult("delete")
}
