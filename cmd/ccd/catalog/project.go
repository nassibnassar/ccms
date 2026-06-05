package catalog

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode"

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
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.projectExists(projectOrFullSetName)
}

func (c *Catalog) projectExists(projectOrFullSetName string) bool {
	s := strings.Split(projectOrFullSetName, ".")
	if len(s) < 1 || len(s) > 2 {
		return false
	}
	_, ok := c.projects[s[0]]
	return ok
}

func (c *Catalog) IsValidTargetProject(project string) bool {
	if strings.ContainsRune(project, '.') {
		return false
	}
	return true
}

func (c *Catalog) AllProjects() []Project {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.allProjects()
}

func (c *Catalog) allProjects() []Project {
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

func (c *Catalog) DropProject(project string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	sets := c.setsInProject(project)
	if len(sets) != 0 {
		return errors.New("project \"" + project + "\" contains one or more sets")
	}

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("dropping project %q: opening transaction: %v", project, err)
	}
	defer tx.Rollback(context.TODO())

	sql := "drop table " + project + ".object"
	if _, err := tx.Exec(context.TODO(), sql); err != nil {
		return pgerr.Error(err)
	}
	sql = "drop schema " + project
	if _, err := tx.Exec(context.TODO(), sql); err != nil {
		return pgerr.Error(err)
	}
	sql = "delete from ccms.project where name=$1"
	if _, err := tx.Exec(context.TODO(), sql, project); err != nil {
		return pgerr.Error(err)
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return pgerr.Error(err)
	}

	delete(c.projects, project)
	return nil
}

func (c *Catalog) CreateProject(project string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return pgerr.Error(err)
	}
	defer tx.Rollback(context.TODO())

	sql := "create schema " + project
	if _, err := tx.Exec(context.TODO(), sql); err != nil {
		return pgerr.Error(err)
	}
	sql = "insert into ccms.project (name) values ($1)"
	if _, err := tx.Exec(context.TODO(), sql, project); err != nil {
		return pgerr.Error(err)
	}

	q := "create table " + project + ".object (" +
		"id bigint primary key," +
		"fund_id integer references ccms.fund (id))"
	if _, err := tx.Exec(context.TODO(), q); err != nil {
		return pgerr.Error(err)
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return pgerr.Error(err)
	}

	c.projects[project] = struct{}{}
	return nil
}

// AlterProjectAddToProperty does not do synchronization and must not access the catalog cache
func (c *Catalog) AlterProjectAddToProperty(project, property, value string, stringLiteral bool) error {
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
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := c.alterProjectAddFund(project, value, projectID); err != nil {
			return err
		}
	case "locations":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := c.alterProjectAddLocation(project, value, projectID); err != nil {
			return err
		}
	case "origins":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := c.alterProjectAddOrigin(project, value, projectID); err != nil {
			return err
		}
	case "destinations":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := c.alterProjectAddDestination(project, value, projectID); err != nil {
			return err
		}
	case "tracks":
		if stringLiteral {
			return invalidValueError(property, value)
		}
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
func (c *Catalog) AlterProjectDropFromProperty(project, property, value string, stringLiteral bool) error {
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
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := c.alterProjectDropFund(project, value, projectID); err != nil {
			return err
		}
	case "locations":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := c.alterProjectDropLocation(project, value, projectID); err != nil {
			return err
		}
	case "origins":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := c.alterProjectDropOrigin(project, value, projectID); err != nil {
			return err
		}
	case "destinations":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := c.alterProjectDropDestination(project, value, projectID); err != nil {
			return err
		}
	case "tracks":
		if stringLiteral {
			return invalidValueError(property, value)
		}
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
		return nil
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
		return nil
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

func (c *Catalog) alterProjectAddOrigin(project, origin string, projectID int64) error {
	// look up origin id
	originID, err := c.selectOriginID(origin)
	if err != nil {
		return err
	}
	if originID == -1 {
		return errors.New("origin \"" + origin + "\" does not exist")
	}
	// check if project origin exists
	projectOriginExists, err := c.projectOriginExists(projectID, originID)
	if err != nil {
		return err
	}
	if projectOriginExists {
		return nil
	}
	// add project origin
	q := "insert into ccms.project_origin (project_id, origin_id) values ($1, $2)"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, originID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func (c *Catalog) alterProjectDropOrigin(project, origin string, projectID int64) error {
	if origin == "*" {
		q := "delete from ccms.project_origin where project_id=$1"
		if _, err := c.dp.Exec(context.TODO(), q, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up origin id
	originID, err := c.selectOriginID(origin)
	if err != nil {
		return err
	}
	if originID == -1 {
		return errors.New("origin \"" + origin + "\" does not exist")
	}
	// check if project origin exists
	projectOriginExists, err := c.projectOriginExists(projectID, originID)
	if err != nil {
		return err
	}
	if !projectOriginExists {
		return errors.New("project \"" + project + "\" does not have origin \"" + origin + "\"")
	}
	// drop project origin
	q := "delete from ccms.project_origin where project_id=$1 and origin_id=$2"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, originID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func (c *Catalog) alterProjectAddDestination(project, destination string, projectID int64) error {
	// look up destination id
	destinationID, err := c.selectDestinationID(destination)
	if err != nil {
		return err
	}
	if destinationID == -1 {
		return errors.New("destination \"" + destination + "\" does not exist")
	}
	// check if project destination exists
	projectDestinationExists, err := c.projectDestinationExists(projectID, destinationID)
	if err != nil {
		return err
	}
	if projectDestinationExists {
		return nil
	}
	// add project destination
	q := "insert into ccms.project_destination (project_id, destination_id) values ($1, $2)"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, destinationID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func (c *Catalog) alterProjectDropDestination(project, destination string, projectID int64) error {
	if destination == "*" {
		q := "delete from ccms.project_destination where project_id=$1"
		if _, err := c.dp.Exec(context.TODO(), q, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up destination id
	destinationID, err := c.selectDestinationID(destination)
	if err != nil {
		return err
	}
	if destinationID == -1 {
		return errors.New("destination \"" + destination + "\" does not exist")
	}
	// check if project destination exists
	projectDestinationExists, err := c.projectDestinationExists(projectID, destinationID)
	if err != nil {
		return err
	}
	if !projectDestinationExists {
		return errors.New("project \"" + project + "\" does not have destination \"" + destination + "\"")
	}
	// drop project destination
	q := "delete from ccms.project_destination where project_id=$1 and destination_id=$2"
	if _, err := c.dp.Exec(context.TODO(), q, projectID, destinationID); err != nil {
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
		return nil
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

// returns origin id, or -1 if origin does not exist
func (c *Catalog) selectOriginID(origin string) (int64, error) {
	var q = "select id from ccms.origin where name=$1"
	var id int64
	err := c.dp.QueryRow(context.TODO(), q, origin).Scan(&id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return -1, nil
	case err != nil:
		return 0, pgerr.Error(err)
	default:
		return id, nil
	}
}

// returns destination id, or -1 if destination does not exist
func (c *Catalog) selectDestinationID(destination string) (int64, error) {
	var q = "select id from ccms.destination where name=$1"
	var id int64
	err := c.dp.QueryRow(context.TODO(), q, destination).Scan(&id)
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

func (c *Catalog) projectOriginExists(projectID, originID int64) (bool, error) {
	var q = "select 1 from ccms.project_origin where project_id=$1 and origin_id=$2"
	var n int32
	err := c.dp.QueryRow(context.TODO(), q, projectID, originID).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func (c *Catalog) projectDestinationExists(projectID, destinationID int64) (bool, error) {
	var q = "select 1 from ccms.project_destination where project_id=$1 and destination_id=$2"
	var n int32
	err := c.dp.QueryRow(context.TODO(), q, projectID, destinationID).Scan(&n)
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
func (c *Catalog) AlterProjectSetProperty(projectName, property, value string, stringLiteral bool) error {
	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return errors.New("opening transaction: " + pgerr.String(err))
	}
	defer tx.Rollback(context.TODO())

	switch property {
	case "funds", "locations", "origins", "destinations", "tracks":
		return errors.New("property \"" + property + "\" is composite")
	case "title", "mou_link":
		if !stringLiteral {
			return invalidValueError(property, value)
		}
	case "action":
		if stringLiteral {
			return invalidValueError(property, value)
		}
	default:
		return errors.New("property \"" + property + "\" does not exist")
	}

	if property == "action" {
		switch value {
		case "", "acquire", "retire", "digitize", "move", "other":
			// NOP
		default:
			return errors.New("action \"" + value + "\" does not exist")
		}
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
	var title, action, mouLink, funds, locations, origins, destinations, tracks string
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
org as (
    select p.id project_id,
           coalesce(string_agg(l.name||':'||l.title, '|' order by l.name), '') origins
        from ccms.project p
            join ccms.project_origin pl on p.id=pl.project_id
            join ccms.origin l on pl.origin_id=l.id
        group by p.id
),
dst as (
    select p.id project_id,
           coalesce(string_agg(l.name||':'||l.title, '|' order by l.name), '') destinations
        from ccms.project p
            join ccms.project_destination pl on p.id=pl.project_id
            join ccms.destination l on pl.destination_id=l.id
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
       coalesce(org.origins, '') origins,
       coalesce(dst.destinations, '') destinations,
       coalesce(trk.tracks, '') tracks
       from ccms.project p
           left join fnd on p.id=fnd.project_id
           left join loc on p.id=loc.project_id
           left join org on p.id=org.project_id
           left join dst on p.id=dst.project_id
           left join trk on p.id=trk.project_id
       where p.name=$1`
	err := c.dp.QueryRow(context.TODO(), q, projectName).Scan(&title, &action, &mouLink, &funds, &locations, &origins, &destinations, &tracks)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, fmt.Errorf("project %q does not exist", projectName)
	case err != nil:
		return nil, err
	default:
	}
	if action != "" {
		action = action + ":" + string(unicode.ToUpper(rune(action[0]))) + action[1:]
	}
	prop := [][2]string{
		{"title", title},
		{"action", action},
		{"mou_link", mouLink},
		{"funds", funds},
		{"locations", locations},
		{"origins", origins},
		{"destinations", destinations},
		{"tracks", tracks},
	}
	return prop, nil
}

func invalidValueError(property, value string) error {
	return errors.New("invalid value for property \"" + property + "\": \"" + value + "\"")
}
