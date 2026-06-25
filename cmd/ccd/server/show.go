package server

import (
	"slices"
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
		result.AddField("definition", "text")
		if err := addShowFiltersData(s.d, result); err != nil {
			return cmderr("retrieving filters: " + err.Error())
		}
	case "funds":
		result.AddField("fund_name", "text")
		result.AddField("fund_title", "text")
		if err := addShowFundsData(s.d, result); err != nil {
			return cmderr("retrieving funds: " + err.Error())
		}
	//case "roles":
	//        result.AddField("role_name", "text")
	//        result.AddField("user_names", "text")
	//        addShowRolesData(s.cat, result)
	case "projects":
		result.AddField("project_name", "text")
		err := addShowProjectsData(s.d, result, cmd.Archived)
		if err != nil {
			return cmderr("retrieving projects: " + err.Error())
		}
	case "project":
		result.AddField("property", "text")
		result.AddField("value", "text")
		if err := addShowProjectData(s.d, result, cmd.Name); err != nil {
			return cmderr("retrieving project data: " + err.Error())
		}
	case "sets":
		if cmd.In != "" {
			projectID, err := cat.ProjectID(s.d, cmd.In)
			if err != nil {
				return cmderr("checking if project exists: " + err.Error())
			}
			if projectID == 0 {
				return cmderr("project \"" + cmd.In + "\" does not exist")
			}
		}
		result.AddField("set_name", "text")
		if err := addShowSetsData(s.d, result, cmd.In); err != nil {
			return cmderr("retrieving sets: " + err.Error())
		}
	case "tags":
		result.AddField("tag_name", "text")
	case "users":
		result.AddField("user_name", "text")
		result.AddField("superuser", "boolean")
		result.AddField("login", "boolean")
		if err := addShowUsersData(s.d, result); err != nil {
			return cmderr("retrieving users: " + err.Error())
		}
	default:
		return cmderr("unknown variable \"" + cmd.Type + "\"")
	}
	return result
}

func addShowFundsData(d *dbx.DB, result *ccms.Result) error {
	funds, err := cat.Funds(d)
	if err != nil {
		return err
	}
	funds.Sort()
	for i := range funds {
		result.AddData([]any{funds[i].Name, funds[i].Title})
	}
	return nil
}

func addShowRolesData(d *dbx.DB, result *ccms.Result) error {
	roles, err := cat.Roles(d)
	if err != nil {
		return err
	}
	cat.SortRoles(roles)
	for i := range roles {
		users := strings.Join(roles[i].UserNames, ", ")
		result.AddData([]any{roles[i].RoleName, users})
	}
	return nil
}

func addShowProjectsData(d *dbx.DB, result *ccms.Result, archived bool) error {
	projects, err := cat.Projects(d, archived)
	if err != nil {
		return err
	}
	slices.Sort(projects)
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

func addShowFiltersData(d *dbx.DB, result *ccms.Result) error {
	filters, err := cat.Filters(d)
	if err != nil {
		return err
	}
	cat.SortFilters(filters)
	for i := range filters {
		result.AddData([]any{filters[i].Name, filters[i].Definition})
	}
	return nil
}

func addShowSetsData(d *dbx.DB, result *ccms.Result, in string) error {
	var sets []string
	var err error
	if in == "" {
		sets, err = cat.Sets(d)
	} else {
		sets, err = cat.SetsInProject(d, in)
	}
	if err != nil {
		return err
	}
	slices.Sort(sets)
	for i := range sets {
		result.AddData([]any{sets[i]})
	}
	return nil
}

func addShowUsersData(d *dbx.DB, result *ccms.Result) error {
	users, err := cat.Users(d)
	if err != nil {
		return err
	}
	cat.SortUsers(users)
	for i := range users {
		result.AddData([]any{users[i].UserName, users[i].Superuser, users[i].Login})
	}
	return nil
}
