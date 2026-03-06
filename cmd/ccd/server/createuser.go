package server

import (
	"strings"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
)

func createUserStmt(s *svr, rqid int64, cmd *ast.CreateUserStmt) *ccms.Result {
	if s.cat.UserExists(cmd.UserName) {
		return cmderr("user \"" + cmd.UserName + "\" already exists")
	}

	if strings.TrimSpace(cmd.EncryptedPassword) == "" {
		return cmderr("password is required")
	}

	if err := s.cat.CreateUser(cmd.UserName, cmd.EncryptedPassword); err != nil {
		return cmderr("error writing user \"" + cmd.UserName + "\"")
	}

	return ccms.NewResult("create user")
}
