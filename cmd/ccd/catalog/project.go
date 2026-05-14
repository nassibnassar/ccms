package catalog

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/indexdata/ccms/internal/pgerr"
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
		return errors.New("registering project \"" + projectName + "\": " + pgerr.String(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("creating project %q: committing changes: %v", projectName, err)
	}

	c.projects[projectName] = struct{}{}
	return nil
}

// AlterProjectAddToProperty does not do synchronization and must not access the catalog cache
func (c *Catalog) AlterProjectAddToProperty(project, property, value string) error {
	// look up project id
	projectID, err := c.selectProjectID(project)
	if err != nil {
		return err
	}
	if projectID == -1 {
		return errors.New("project \"" + project + "\" does not exist")
	}

	switch property {
	case "funds":
		if err := c.alterProjectAddFund(project, value, projectID); err != nil {
			return err
		}
	case "locations":
		if err := c.alterProjectAddLocation(project, value, projectID); err != nil {
			return err
		}
	case "tracks":
		if err := c.alterProjectAddTrack(project, value, projectID); err != nil {
			return err
		}
	case "title", "action", "mou_link":
		return errors.New("property \"" + property + "\" is not composite")
	default:
		return errors.New("unknown property \"" + property + "\"")
	}
	return nil
}

// AlterProjectDropFromProperty does not do synchronization and must not access the catalog cache
func (c *Catalog) AlterProjectDropFromProperty(project, property, value string) error {
	// look up project id
	projectID, err := c.selectProjectID(project)
	if err != nil {
		return err
	}
	if projectID == -1 {
		return errors.New("project \"" + project + "\" does not exist")
	}

	switch property {
	case "funds":
		if err := c.alterProjectDropFund(project, value, projectID); err != nil {
			return err
		}
	case "locations":
		if err := c.alterProjectDropLocation(project, value, projectID); err != nil {
			return err
		}
	case "tracks":
		if err := c.alterProjectDropTrack(project, value, projectID); err != nil {
			return err
		}
	case "title", "action", "mou_link":
		return errors.New("property \"" + property + "\" is not composite")
	default:
		return errors.New("unknown property \"" + property + "\"")
	}
	return nil
}

func (c *Catalog) alterProjectAddFund(project, fund string, projectID int64) error {
	// look up fund id
	fundID, err := c.selectFundID(fund)
	if err != nil {
		return err
	}
	if fundID == -1 {
		return errors.New("fund \"" + fund + "\" does not exist")
	}
	// check if project fund exists
	projectFundExists, err := c.projectFundExists(projectID, fundID)
	if err != nil {
		return err
	}
	if projectFundExists {
		return errors.New("project \"" + project + "\" already has fund \"" + fund + "\"")
	}
	// add project fund
	q := "insert into ccms.project_fund (project_id, fund_id) values ($1, $2)"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, fundID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func (c *Catalog) alterProjectDropFund(project, fund string, projectID int64) error {
	if fund == "*" {
		q := "delete from ccms.project_fund where project_id=$1"
		if _, err := c.dp.Exec(context.TODO(), q, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up fund id
	fundID, err := c.selectFundID(fund)
	if err != nil {
		return err
	}
	if fundID == -1 {
		return errors.New("fund \"" + fund + "\" does not exist")
	}
	// check if project fund exists
	projectFundExists, err := c.projectFundExists(projectID, fundID)
	if err != nil {
		return err
	}
	if !projectFundExists {
		return errors.New("project \"" + project + "\" does not have fund \"" + fund + "\"")
	}
	// drop project fund
	q := "delete from ccms.project_fund where project_id=$1 and fund_id=$2"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, fundID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func (c *Catalog) alterProjectAddLocation(project, location string, projectID int64) error {
	// look up location id
	locationID, err := c.selectLocationID(location)
	if err != nil {
		return err
	}
	if locationID == -1 {
		return errors.New("location \"" + location + "\" does not exist")
	}
	// check if project location exists
	projectLocationExists, err := c.projectLocationExists(projectID, locationID)
	if err != nil {
		return err
	}
	if projectLocationExists {
		return errors.New("project \"" + project + "\" already has location \"" + location + "\"")
	}
	// add project location
	q := "insert into ccms.project_location (project_id, location_id) values ($1, $2)"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, locationID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func (c *Catalog) alterProjectDropLocation(project, location string, projectID int64) error {
	if location == "*" {
		q := "delete from ccms.project_location where project_id=$1"
		if _, err := c.dp.Exec(context.TODO(), q, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up location id
	locationID, err := c.selectLocationID(location)
	if err != nil {
		return err
	}
	if locationID == -1 {
		return errors.New("location \"" + location + "\" does not exist")
	}
	// check if project location exists
	projectLocationExists, err := c.projectLocationExists(projectID, locationID)
	if err != nil {
		return err
	}
	if !projectLocationExists {
		return errors.New("project \"" + project + "\" does not have location \"" + location + "\"")
	}
	// drop project location
	q := "delete from ccms.project_location where project_id=$1 and location_id=$2"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, locationID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func (c *Catalog) alterProjectAddTrack(project, track string, projectID int64) error {
	// look up track id
	trackID, err := c.selectTrackID(track)
	if err != nil {
		return err
	}
	if trackID == -1 {
		return errors.New("track \"" + track + "\" does not exist")
	}
	// check if project track exists
	projectTrackExists, err := c.projectTrackExists(projectID, trackID)
	if err != nil {
		return err
	}
	if projectTrackExists {
		return errors.New("project \"" + project + "\" already has track \"" + track + "\"")
	}
	// add project track
	q := "insert into ccms.project_track (project_id, track_id) values ($1, $2)"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, trackID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func (c *Catalog) alterProjectDropTrack(project, track string, projectID int64) error {
	if track == "*" {
		q := "delete from ccms.project_track where project_id=$1"
		if _, err := c.dp.Exec(context.TODO(), q, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up track id
	trackID, err := c.selectTrackID(track)
	if err != nil {
		return err
	}
	if trackID == -1 {
		return errors.New("track \"" + track + "\" does not exist")
	}
	// check if project track exists
	projectTrackExists, err := c.projectTrackExists(projectID, trackID)
	if err != nil {
		return err
	}
	if !projectTrackExists {
		return errors.New("project \"" + project + "\" does not have track \"" + track + "\"")
	}
	// drop project track
	q := "delete from ccms.project_track where project_id=$1 and track_id=$2"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, trackID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

// returns project id, or -1 if project does not exist
func (c *Catalog) selectProjectID(project string) (int64, error) {
	var q = "select id from ccms.project where name=$1"
	var id int64
	err := c.dp.QueryRow(context.TODO(), q, project).Scan(&id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return -1, nil
	case err != nil:
		return 0, pgerr.Error(err)
	default:
		return id, nil
	}
}

// returns fund id, or -1 if fund does not exist
func (c *Catalog) selectFundID(fund string) (int64, error) {
	var q = "select id from ccms.fund where name=$1"
	var id int64
	err := c.dp.QueryRow(context.TODO(), q, fund).Scan(&id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return -1, nil
	case err != nil:
		return 0, pgerr.Error(err)
	default:
		return id, nil
	}
}

// returns location id, or -1 if location does not exist
func (c *Catalog) selectLocationID(location string) (int64, error) {
	var q = "select id from ccms.location where name=$1"
	var id int64
	err := c.dp.QueryRow(context.TODO(), q, location).Scan(&id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return -1, nil
	case err != nil:
		return 0, pgerr.Error(err)
	default:
		return id, nil
	}
}

// returns track id, or -1 if track does not exist
func (c *Catalog) selectTrackID(track string) (int64, error) {
	var q = "select id from ccms.track where name=$1"
	var id int64
	err := c.dp.QueryRow(context.TODO(), q, track).Scan(&id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return -1, nil
	case err != nil:
		return 0, pgerr.Error(err)
	default:
		return id, nil
	}
}

func (c *Catalog) projectFundExists(projectID, fundID int64) (bool, error) {
	var q = "select 1 from ccms.project_fund where project_id=$1 and fund_id=$2"
	var n int32
	err := c.dp.QueryRow(context.TODO(), q, projectID, fundID).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func (c *Catalog) projectLocationExists(projectID, locationID int64) (bool, error) {
	var q = "select 1 from ccms.project_location where project_id=$1 and location_id=$2"
	var n int32
	err := c.dp.QueryRow(context.TODO(), q, projectID, locationID).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func (c *Catalog) projectTrackExists(projectID, trackID int64) (bool, error) {
	var q = "select 1 from ccms.project_track where project_id=$1 and track_id=$2"
	var n int32
	err := c.dp.QueryRow(context.TODO(), q, projectID, trackID).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

// AlterProjectSetProperty does not do synchronization and must not access the catalog cache
func (c *Catalog) AlterProjectSetProperty(projectName, property, value string) error {
	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return errors.New("opening transaction: " + pgerr.String(err))
	}
	defer tx.Rollback(context.TODO())

	switch property {
	case "funds", "locations", "tracks":
		return errors.New("property \"" + property + "\" is composite")
	}

	sql := "update ccms.project set \"" + property + "\"=nullif($1, '') where name=$2"
	if _, err := tx.Exec(context.TODO(), sql, value, projectName); err != nil {
		return errors.New("updating project: " + pgerr.String(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return errors.New("committing changes: " + pgerr.String(err))
	}
	return nil
}

func (c *Catalog) ProjectProperties(projectName string) ([][2]string, error) {
	var title, action, mouLink, funds, locations, tracks string
	q := `with fnd as (
    select p.id project_id,
           coalesce(string_agg(f.name||':'||f.title, '|' order by f.name), '') funds
        from ccms.project p
            join ccms.project_fund pf on p.id=pf.project_id
            join ccms.fund f on pf.fund_id=f.id
        group by p.id
),
loc as (
    select p.id project_id,
           coalesce(string_agg(l.name||':'||l.title, '|' order by l.name), '') locations
        from ccms.project p
            join ccms.project_location pl on p.id=pl.project_id
            join ccms.location l on pl.location_id=l.id
        group by p.id
),
trk as (
    select p.id project_id,
           coalesce(string_agg(t.name||':'||t.title, '|' order by t.name), '') tracks
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
