package server

import (
	"strings"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/dbx"
)

func showStmt(s *svr, cmd *ast.ShowStmt) *ccms.Result {
	result := ccms.NewResult("show")
	switch cmd.Type {
	case "filters":
		result.AddField("filter_name", "text")
	case "funds":
		result.AddField("fund_name", "text")
		result.AddField("fund_title", "text")
		if err := addShowFundsData(s.d, result); err != nil {
			return cmderrint("retrieving funds", err)
		}
	//case "roles":
	//        result.AddField("role_name", "text")
	//        result.AddField("user_names", "text")
	//        addShowRolesData(s.cat, result)
	case "projects":
		result.AddField("project_name", "text")
		err := addShowProjectsData(s.d, result)
		if err != nil {
			return cmderrint("retrieving projects", err)
		}
	case "project":
		result.AddField("property", "text")
		result.AddField("value", "text")
		if err := addShowProjectData(s.d, result, cmd.Name); err != nil {
			return cmderrint("retrieving project data", err)
		}
	case "sets":
		if cmd.In != "" {
			projectExists, err := cat.ProjectExists(s.d, cmd.In)
			if err != nil {
				return cmderrint("checking if project exists", err)
			}
			if !projectExists {
				return cmderr("project \"" + cmd.In + "\" does not exist")
			}
		}
		result.AddField("set_name", "text")
		if err := addShowSetsData(s.d, result, cmd.In); err != nil {
			return cmderrint("retrieving sets", err)
		}
	case "tags":
		result.AddField("tag_name", "text")
	case "users":
		result.AddField("user_name", "text")
		result.AddField("superuser", "boolean")
		result.AddField("login", "boolean")
		if err := addShowUsersData(s.d, result); err != nil {
			return cmderrint("retrieving users", err)
		}
	default:
		return cmderr("unknown variable \"" + cmd.Type + "\"")
	}
	return result
}

func addShowFundsData(d *dbx.DB, result *ccms.Result) error {
	funds, err := cat.AllFunds(d)
	if err != nil {
		return err
	}
	for i := range funds {
		result.AddData([]any{funds[i].Name, funds[i].Title})
	}
	return nil
}

func addShowRolesData(d *dbx.DB, result *ccms.Result) error {
	roles, err := cat.AllRoles(d)
	if err != nil {
		return err
	}
	for i := range roles {
		users := strings.Join(roles[i].UserNames, ", ")
		result.AddData([]any{roles[i].RoleName, users})
	}
	return nil
}

func addShowProjectsData(d *dbx.DB, result *ccms.Result) error {
	projects, err := cat.AllProjects(d)
	if err != nil {
		return err
	}
	for i := range projects {
		result.AddData([]any{projects[i]})
	}
	return nil
}

func addShowProjectData(d *dbx.DB, result *ccms.Result, projectName string) error {
	prop, err := cat.ProjectProperties(d, projectName)
	if err != nil {
		return err
	}
	for i := range prop {
		result.AddData([]any{prop[i][0], prop[i][1]})
	}
	return nil
}

func addShowSetsData(d *dbx.DB, result *ccms.Result, in string) error {
	var sets []string
	var err error
	if in == "" {
		sets, err = cat.AllSets(d)
	} else {
		sets, err = cat.SetsInProject(d, in)
	}
	if err != nil {
		return err
	}
	for i := range sets {
		result.AddData([]any{sets[i]})
	}
	return nil
}

func addShowUsersData(d *dbx.DB, result *ccms.Result) error {
	users, err := cat.AllUsers(d)
	if err != nil {
		return err
	}
	for i := range users {
		result.AddData([]any{users[i].UserName, users[i].Superuser, users[i].Login})
	}
	return nil
}
