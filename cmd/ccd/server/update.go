package server

import (
	"errors"
	"strconv"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/dbx"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/jackc/pgx/v5"
)

func updateStmt(s *svr, rqid int64, cmd *ast.UpdateStmt) *ccms.Result {

	schema, table := cat.SplitSchemaTable(cmd.SetName)
	if table != "object" {
		return cmderr("set \"" + cmd.SetName + "\" is not valid for update")
	}
	projectExists, err := cat.ProjectExists(s.d, schema)
	if err != nil {
		return cmderrint("checking if project exists", err)
	}
	if !projectExists {
		return cmderr("project \"" + schema + "\" does not exist")
	}

	if cmd.IDAttr != "id" {
		return cmderr("attribute \"" + cmd.IDAttr + "\" is not valid with update")
	}

	idInt, _ := strconv.ParseInt(cmd.IDValue.Value, 10, 64)
	idExists, err := objectIDExists(s.d, idInt)
	if err != nil {
		return cmderrint("checking if object ID exists", err)
	}
	if !idExists {
		return cmderr("object id = " + cmd.IDValue.Value + " does not exist")
	}

	switch cmd.Attr {
	case "fund":
		cmd.Attr = "fund_id"
		if cmd.Value != "" {
			var fundID int64
			fundID, err = cat.SelectFundID(s.d, cmd.Value)
			if err != nil {
				return cmderrint("looking up fund", err)
			}
			if fundID == -1 {
				return cmderr("fund \"" + cmd.Value + "\" does not exist")
			}
			cmd.Value = strconv.FormatInt(fundID, 10)
		}
	default:
		return cmderr("attribute \"" + cmd.Attr + "\" is not valid for update")
	}

	sql, err := cmd.SQL()
	if err != nil {
		return cmderr(err.Error())
	}
	if _, err := s.d.Q.Exec(s.d.C, sql); err != nil {
		return cmderrint("executing update", err)
	}

	return ccms.NewResult("update")
}

func objectIDExists(d *dbx.DB, id int64) (bool, error) {
	var q = "select 1 from ccms.reserve where id=$1"
	var n int32
	err := d.Q.QueryRow(d.C, q, id).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}
