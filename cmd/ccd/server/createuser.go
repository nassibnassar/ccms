package server

import (
	"strings"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func createUserStmt(s *svr, rqid int64, cmd *ast.CreateUserStmt) *ccms.Result {
	userExists, err := cat.UserExists(s.d, cmd.UserName)
	if err != nil {
		return cmderrint("checking if user exists", err)
	}
	if userExists {
		return cmderr("user \"" + cmd.UserName + "\" already exists")
	}

	if strings.TrimSpace(cmd.EncryptedPassword) == "" {
		return cmderr("password is required")
	}

	if err := cat.CreateUser(s.conf.Security.SecretKey, s.d, cmd.UserName, cmd.EncryptedPassword, false, true); err != nil {
		return cmderrint("writing user", err)
	}

	return ccms.NewResult("create user")
}
