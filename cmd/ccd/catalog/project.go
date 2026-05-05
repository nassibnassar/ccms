package catalog

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/indexdata/ccms/internal/global"
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

func (c *Catalog) IsValidTargetProject(projectName string) bool {
	if projectName == "reserve" {
		return false
	}
	if strings.ContainsRune(projectName, '.') {
		return false
	}
	return true
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

func (c *Catalog) CreateProject(projectName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("creating project %q: opening transaction: %v", projectName, err)
	}
	defer tx.Rollback(context.TODO())

	sql := "create schema " + projectName
	if _, err := tx.Exec(context.TODO(), sql); err != nil {
		return fmt.Errorf("creating project %q: %v", projectName, err)
	}
	sql = "insert into ccms.project (name) values ($1)"
	if _, err := tx.Exec(context.TODO(), sql, projectName); err != nil {
		return fmt.Errorf("registering project %q: %v", projectName, global.PGErr(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("creating project %q: committing changes: %v", projectName, err)
	}

	c.projects[projectName] = struct{}{}
	return nil
}

func (c *Catalog) ProjectProperties(projectName string) ([][2]string, error) {
	var title, action, mouLink, funds, locations, tracks string
	q := `with fnd as (
    select p.id project_id,
           coalesce(string_agg(f.name||':'||f.title, '|'), '') funds
        from ccms.project p
            join ccms.project_fund pf on p.id=pf.project_id
            join ccms.fund f on pf.fund_id=f.id
        group by p.id
),
loc as (
    select p.id project_id,
           coalesce(string_agg(l.name||':'||l.title, '|'), '') locations
        from ccms.project p
            join ccms.project_location pl on p.id=pl.project_id
            join ccms.location l on pl.location_id=l.id
        group by p.id
),
trk as (
    select p.id project_id,
           coalesce(string_agg(t.name||':'||t.title, '|'), '') tracks
        from ccms.project p
            join ccms.project_track pl on p.id=pl.project_id
            join ccms.track t on pl.track_id=t.id
        group by p.id
)
select coalesce(p.title, '') title,
       coalesce(p.action, '') action,
       coalesce(p.mou_link, '') mou_link,
       coalesce(fnd.funds, '') funds,
       coalesce(loc.locations, '') locations,
       coalesce(trk.tracks, '') tracks
       from ccms.project p
           left join fnd on p.id=fnd.project_id
           left join loc on p.id=loc.project_id
           left join trk on p.id=trk.project_id
       where p.name=$1`
	err := c.dp.QueryRow(context.TODO(), q, projectName).Scan(&title, &action, &mouLink, &funds, &locations, &tracks)
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
		{"locations", locations},
		{"tracks", tracks},
	}
	return prop, nil
}
