package server

import (
	"strings"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
)

func showStmt(s *svr, cmd *ast.ShowStmt) *ccms.Result {
	result := ccms.NewResult("show")
	switch cmd.Name {
	case "filters":
		result.AddField("filter_name", "text")
	//case "roles":
	//        result.AddField("role_name", "text")
	//        result.AddField("user_names", "text")
	//        addShowRolesData(s.cat, result)
	case "sets":
		result.AddField("set_name", "text")
		addShowSetsData(s.cat, result)
	case "users":
		result.AddField("user_name", "text")
		result.AddField("superuser", "boolean")
		result.AddField("login", "boolean")
		addShowUsersData(s.cat, result)
	default:
		return cmderr("unknown variable \"" + cmd.Name + "\"")
	}
	return result
}

func addShowRolesData(cat *catalog.Catalog, result *ccms.Result) {
	roles := cat.AllRoles()
	for i := range roles {
		users := strings.Join(roles[i].UserNames, ", ")
		result.AddData([]any{roles[i].RoleName, users})
	}
}

func addShowSetsData(cat *catalog.Catalog, result *ccms.Result) {
	sets := cat.AllSets()
	for i := range sets {
		result.AddData([]any{sets[i]})
	}
}

func addShowUsersData(cat *catalog.Catalog, result *ccms.Result) {
	users := cat.AllUsers()
	for i := range users {
		result.AddData([]any{users[i].UserName, users[i].Superuser, users[i].Login})
	}
}
