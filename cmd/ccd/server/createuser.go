package server

import (
	"strings"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func createUserStmt(s *svr, d *dbx.DB, rqid int64, cmd *ast.CreateUserStmt) *ccms.Result {
	userExists, err := cat.UserExists(d, cmd.User)
	if err != nil {
		return cmderr("checking if user exists: " + err.Error())
	}
	if userExists {
		return cmderr("user \"" + cmd.User + "\" already exists")
	}

	if strings.TrimSpace(cmd.EncryptedPassword) == "" {
		return cmderr("password is required")
	}

	if err := cat.CreateUser(s.conf.Security.SecretKey, d, cmd.User, cmd.EncryptedPassword, false, true); err != nil {
		return cmderr("writing user: " + err.Error())
	}

	return ccms.NewResult("create user")
}
