package server

import (
	"strings"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
)

func showStmt(s *svr, cmd *ast.ShowStmt) *ccms.Result {
	result := ccms.NewResult("show")
	switch cmd.Type {
	case "filters":
		result.AddField("filter_name", "text")
	//case "roles":
	//        result.AddField("role_name", "text")
	//        result.AddField("user_names", "text")
	//        addShowRolesData(s.cat, result)
	case "projects":
		result.AddField("project_name", "text")
		addShowProjectsData(s.cat, result)
	case "project":
		result.AddField("property", "text")
		result.AddField("value", "text")
		if err := addShowProjectData(s.cat, result, cmd.Name); err != nil {
			return cmderr(err.Error())
		}
	case "sets":
		result.AddField("set_name", "text")
		addShowSetsData(s.cat, result)
	case "users":
		result.AddField("user_name", "text")
		result.AddField("superuser", "boolean")
		result.AddField("login", "boolean")
		addShowUsersData(s.cat, result)
	default:
		return cmderr("unknown variable \"" + cmd.Type + "\"")
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

func addShowProjectsData(cat *catalog.Catalog, result *ccms.Result) {
	projects := cat.AllProjects()
	for i := range projects {
		result.AddData([]any{projects[i].ProjectName})
	}
}

func addShowProjectData(cat *catalog.Catalog, result *ccms.Result, projectName string) error {
	prop, err := cat.ProjectProperties(projectName)
	if err != nil {
		return err
	}
	for i := range prop {
		result.AddData([]any{prop[i][0], prop[i][1]})
	}
	return nil
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
