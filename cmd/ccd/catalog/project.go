package catalog

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Project struct {
	ProjectName string
	//ProjectTitle string
	//Action       string
	//MOULink      string
}

func (c *Catalog) initProjects() error {
	sql := "select name from ccms.project"
	rows, err := c.dp.Query(context.TODO(), sql)
	if err != nil {
		return fmt.Errorf("selecting projects: %v", err)
	}
	defer rows.Close()
	projects := make(map[string]struct{})
	for rows.Next() {
		var project string
		if err := rows.Scan(&project); err != nil {
			return fmt.Errorf("reading projects: %v", err)
		}
		projects[project] = struct{}{}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("reading projects: %v", err)
	}
	c.projects = projects
	return nil
}

// given a project name or fully qualified set name (project.set), return true if the project exists
func (c *Catalog) ProjectExists(projectOrFullSetName string) bool {
	var p string // project name
	sp := strings.Split(projectOrFullSetName, ".")
	switch len(sp) {
	case 1: // project
		fallthrough
	case 2: // full set name
		p = sp[0]
	}
	_, ok := c.projects[p]
	return ok
}

func (c *Catalog) AllProjects() []Project {
	c.mu.Lock()
	defer c.mu.Unlock()
	projects := make([]Project, len(c.projects))
	i := 0
	for k, v := range c.projects {
		_ = v
		projects[i] = Project{
			ProjectName: k,
		}
		i++
	}
	sortProjectNames(projects)
	return projects
}

func sortProjectNames(projects []Project) {
	slices.SortFunc(projects, func(x, y Project) int {
		return cmp.Compare(x.ProjectName, y.ProjectName)
	})
}

func (c *Catalog) ProjectProperties(projectName string) ([][2]string, error) {
	var title, action, mouLink, funds string
	q := "select coalesce(p.title, ''), coalesce(p.action, ''), coalesce(p.mou_link, ''), " +
		"coalesce(string_agg(f.name||':'||f.title, '|'), '') funds " +
		"from ccms.project p " +
		"left join ccms.project_fund pf on p.id=pf.project_id " +
		"left join ccms.fund f on pf.fund_id=f.id " +
		"where p.name=$1 " +
		"group by p.title, p.action, p.mou_link"
	err := c.dp.QueryRow(context.TODO(), q, projectName).Scan(&title, &action, &mouLink, &funds)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, fmt.Errorf("project %q does not exist", projectName)
	case err != nil:
		return nil, err
	default:
	}
	prop := [][2]string{
		{"title", title},
		{"action", action},
		{"mou_link", mouLink},
		{"funds", funds},
	}
	return prop, nil
}
