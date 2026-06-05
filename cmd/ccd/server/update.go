package server

import (
	"context"
	"errors"
	"strconv"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func updateStmt(s *svr, rqid int64, cmd *ast.UpdateStmt) *ccms.Result {

	schema, table := catalog.SplitSchemaTable(cmd.SetName)
	if table != "object" {
		return cmderr("set \"" + cmd.SetName + "\" is not valid for update")
	}
	if !s.cat.ProjectExists(schema) {
		return cmderr("project \"" + schema + "\" does not exist")
	}

	if cmd.IDAttr != "id" {
		return cmderr("attribute \"" + cmd.IDAttr + "\" is not valid with update")
	}

	idInt, _ := strconv.ParseInt(cmd.IDValue.Value, 10, 64)
	idExists, err := objectIDExists(s.dp, idInt)
	if err != nil {
		return cmderr(err.Error())
	}
	if !idExists {
		return cmderr("object id = " + cmd.IDValue.Value + " does not exist")
	}

	switch cmd.Attr {
	case "fund":
		cmd.Attr = "fund_id"
		if cmd.Value != "" {
			var fundID int64
			fundID, err = s.cat.SelectFundID(cmd.Value)
			if err != nil {
				return cmderr(err.Error())
			}
			if fundID == -1 {
				return cmderr("fund \"" + cmd.Value + "\" does not exist")
			}
			cmd.Value = strconv.FormatInt(fundID, 10)
		}
	default:
		return cmderr("unknown attribute \"" + cmd.Attr + "\"")
	}

	sql, err := cmd.SQL()
	if err != nil {
		return cmderr(err.Error())
	}
	log.Info("[%d] %s", rqid, sql)
	if _, err := s.dp.Exec(context.TODO(), sql); err != nil {
		return cmderr(pgerr.String(err))
	}

	return ccms.NewResult("update")
}

func objectIDExists(dp *pgxpool.Pool, id int64) (bool, error) {
	var q = "select 1 from ccms.reserve where id=$1"
	var n int32
	err := dp.QueryRow(context.TODO(), q, id).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}
