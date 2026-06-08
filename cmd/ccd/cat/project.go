package cat

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode"

	"github.com/indexdata/ccms/internal/dbx"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/jackc/pgx/v5"
)

// type Project struct {
// 	ProjectName string
// 	//ProjectTitle string
// 	//Action       string
// 	//MOULink      string
// }

func ProjectExists(d *dbx.DB, projectOrFullSetName string) (bool, error) {
	s := strings.Split(projectOrFullSetName, ".")
	if len(s) < 1 || len(s) > 2 {
		return false, nil
	}
	q := "select 1 from ccms.project where name=$1"
	var n int32
	err := d.Q.QueryRow(d.C, q, s[0]).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func IsValidTargetProject(project string) bool {
	if strings.ContainsRune(project, '.') {
		return false
	}
	return true
}

func AllProjects(d *dbx.DB) ([]string, error) {
	rows, _ := d.Q.Query(d.C, "select name from ccms.project")
	projects, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil {
		return nil, pgerr.Error(err)
	}
	sortProjectNames(projects)
	return projects, nil
}

func sortProjectNames(projects []string) {
	slices.SortFunc(projects, func(x, y string) int {
		return cmp.Compare(x, y)
	})
}

func DropProject(d *dbx.DB, project string) error {
	sets, err := SetsInProject(d, project)
	if err != nil {
		return err
	}
	if len(sets) != 0 {
		return errors.New("project \"" + project + "\" contains one or more sets")
	}

	sql := "drop table " + project + ".object"
	if _, err := d.Q.Exec(d.C, sql); err != nil {
		return pgerr.Error(err)
	}
	sql = "drop schema " + project
	if _, err := d.Q.Exec(d.C, sql); err != nil {
		return pgerr.Error(err)
	}
	sql = "delete from ccms.project where name=$1"
	if _, err := d.Q.Exec(d.C, sql, project); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func CreateProject(d *dbx.DB, project string) error {
	sql := "create schema " + project
	if _, err := d.Q.Exec(d.C, sql); err != nil {
		return pgerr.Error(err)
	}
	sql = "insert into ccms.project (name, title) values ($1, $2)"
	if _, err := d.Q.Exec(d.C, sql, project, makeTitle(project)); err != nil {
		return pgerr.Error(err)
	}
	sql = "create table " + project + ".object (" +
		"id bigint primary key," +
		"fund_id integer references ccms.fund (id))"
	if _, err := d.Q.Exec(d.C, sql); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func AlterProjectAddToProperty(d *dbx.DB, project, property, value string, stringLiteral bool) error {
	// look up project id
	projectID, err := selectProjectID(d, project)
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
		if err := alterProjectAddFund(d, project, value, projectID); err != nil {
			return err
		}
	case "locations":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := alterProjectAddLocation(d, project, value, projectID); err != nil {
			return err
		}
	case "origins":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := alterProjectAddOrigin(d, project, value, projectID); err != nil {
			return err
		}
	case "destinations":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := alterProjectAddDestination(d, project, value, projectID); err != nil {
			return err
		}
	case "tracks":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := alterProjectAddTrack(d, project, value, projectID); err != nil {
			return err
		}
	case "title", "action", "mou_link":
		return errors.New("property \"" + property + "\" is not composite")
	default:
		return errors.New("unknown property \"" + property + "\"")
	}
	return nil
}

func AlterProjectDropFromProperty(d *dbx.DB, project, property, value string, stringLiteral bool) error {
	// look up project id
	projectID, err := selectProjectID(d, project)
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
		if err := alterProjectDropFund(d, project, value, projectID); err != nil {
			return err
		}
	case "locations":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := alterProjectDropLocation(d, project, value, projectID); err != nil {
			return err
		}
	case "origins":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := alterProjectDropOrigin(d, project, value, projectID); err != nil {
			return err
		}
	case "destinations":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := alterProjectDropDestination(d, project, value, projectID); err != nil {
			return err
		}
	case "tracks":
		if stringLiteral {
			return invalidValueError(property, value)
		}
		if err := alterProjectDropTrack(d, project, value, projectID); err != nil {
			return err
		}
	case "title", "action", "mou_link":
		return errors.New("property \"" + property + "\" is not composite")
	default:
		return errors.New("unknown property \"" + property + "\"")
	}
	return nil
}

func alterProjectAddFund(d *dbx.DB, project, fund string, projectID int64) error {
	// look up fund id
	fundID, err := SelectFundID(d, fund)
	if err != nil {
		return err
	}
	if fundID == -1 {
		return errors.New("fund \"" + fund + "\" does not exist")
	}
	// check if project fund exists
	projectFundExists, err := projectFundExists(d, projectID, fundID)
	if err != nil {
		return err
	}
	if projectFundExists {
		return nil
	}
	// add project fund
	sql := "insert into ccms.project_fund (project_id, fund_id) values ($1, $2)"
	if _, err := d.Q.Exec(d.C, sql, projectID, fundID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func alterProjectDropFund(d *dbx.DB, project, fund string, projectID int64) error {
	if fund == "*" {
		sql := "delete from ccms.project_fund where project_id=$1"
		if _, err := d.Q.Exec(d.C, sql, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up fund id
	fundID, err := SelectFundID(d, fund)
	if err != nil {
		return err
	}
	if fundID == -1 {
		return errors.New("fund \"" + fund + "\" does not exist")
	}
	// check if project fund exists
	projectFundExists, err := projectFundExists(d, projectID, fundID)
	if err != nil {
		return err
	}
	if !projectFundExists {
		return errors.New("project \"" + project + "\" does not have fund \"" + fund + "\"")
	}
	// drop project fund
	sql := "delete from ccms.project_fund where project_id=$1 and fund_id=$2"
	if _, err := d.Q.Exec(d.C, sql, projectID, fundID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func alterProjectAddLocation(d *dbx.DB, project, location string, projectID int64) error {
	// look up location id
	locationID, err := selectLocationID(d, location)
	if err != nil {
		return err
	}
	if locationID == -1 {
		return errors.New("location \"" + location + "\" does not exist")
	}
	// check if project location exists
	projectLocationExists, err := projectLocationExists(d, projectID, locationID)
	if err != nil {
		return err
	}
	if projectLocationExists {
		return nil
	}
	// add project location
	sql := "insert into ccms.project_location (project_id, location_id) values ($1, $2)"
	if _, err := d.Q.Exec(d.C, sql, projectID, locationID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func alterProjectDropLocation(d *dbx.DB, project, location string, projectID int64) error {
	if location == "*" {
		sql := "delete from ccms.project_location where project_id=$1"
		if _, err := d.Q.Exec(d.C, sql, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up location id
	locationID, err := selectLocationID(d, location)
	if err != nil {
		return err
	}
	if locationID == -1 {
		return errors.New("location \"" + location + "\" does not exist")
	}
	// check if project location exists
	projectLocationExists, err := projectLocationExists(d, projectID, locationID)
	if err != nil {
		return err
	}
	if !projectLocationExists {
		return errors.New("project \"" + project + "\" does not have location \"" + location + "\"")
	}
	// drop project location
	sql := "delete from ccms.project_location where project_id=$1 and location_id=$2"
	if _, err := d.Q.Exec(d.C, sql, projectID, locationID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func alterProjectAddOrigin(d *dbx.DB, project, origin string, projectID int64) error {
	// look up origin id
	originID, err := selectOriginID(d, origin)
	if err != nil {
		return err
	}
	if originID == -1 {
		return errors.New("origin \"" + origin + "\" does not exist")
	}
	// check if project origin exists
	projectOriginExists, err := projectOriginExists(d, projectID, originID)
	if err != nil {
		return err
	}
	if projectOriginExists {
		return nil
	}
	// add project origin
	sql := "insert into ccms.project_origin (project_id, origin_id) values ($1, $2)"
	if _, err := d.Q.Exec(d.C, sql, projectID, originID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func alterProjectDropOrigin(d *dbx.DB, project, origin string, projectID int64) error {
	if origin == "*" {
		sql := "delete from ccms.project_origin where project_id=$1"
		if _, err := d.Q.Exec(d.C, sql, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up origin id
	originID, err := selectOriginID(d, origin)
	if err != nil {
		return err
	}
	if originID == -1 {
		return errors.New("origin \"" + origin + "\" does not exist")
	}
	// check if project origin exists
	projectOriginExists, err := projectOriginExists(d, projectID, originID)
	if err != nil {
		return err
	}
	if !projectOriginExists {
		return errors.New("project \"" + project + "\" does not have origin \"" + origin + "\"")
	}
	// drop project origin
	sql := "delete from ccms.project_origin where project_id=$1 and origin_id=$2"
	if _, err := d.Q.Exec(d.C, sql, projectID, originID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func alterProjectAddDestination(d *dbx.DB, project, destination string, projectID int64) error {
	// look up destination id
	destinationID, err := selectDestinationID(d, destination)
	if err != nil {
		return err
	}
	if destinationID == -1 {
		return errors.New("destination \"" + destination + "\" does not exist")
	}
	// check if project destination exists
	projectDestinationExists, err := projectDestinationExists(d, projectID, destinationID)
	if err != nil {
		return err
	}
	if projectDestinationExists {
		return nil
	}
	// add project destination
	sql := "insert into ccms.project_destination (project_id, destination_id) values ($1, $2)"
	if _, err := d.Q.Exec(d.C, sql, projectID, destinationID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func alterProjectDropDestination(d *dbx.DB, project, destination string, projectID int64) error {
	if destination == "*" {
		sql := "delete from ccms.project_destination where project_id=$1"
		if _, err := d.Q.Exec(d.C, sql, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up destination id
	destinationID, err := selectDestinationID(d, destination)
	if err != nil {
		return err
	}
	if destinationID == -1 {
		return errors.New("destination \"" + destination + "\" does not exist")
	}
	// check if project destination exists
	projectDestinationExists, err := projectDestinationExists(d, projectID, destinationID)
	if err != nil {
		return err
	}
	if !projectDestinationExists {
		return errors.New("project \"" + project + "\" does not have destination \"" + destination + "\"")
	}
	// drop project destination
	sql := "delete from ccms.project_destination where project_id=$1 and destination_id=$2"
	if _, err := d.Q.Exec(d.C, sql, projectID, destinationID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func alterProjectAddTrack(d *dbx.DB, project, track string, projectID int64) error {
	// look up track id
	trackID, err := selectTrackID(d, track)
	if err != nil {
		return err
	}
	if trackID == -1 {
		return errors.New("track \"" + track + "\" does not exist")
	}
	// check if project track exists
	projectTrackExists, err := projectTrackExists(d, projectID, trackID)
	if err != nil {
		return err
	}
	if projectTrackExists {
		return nil
	}
	// add project track
	sql := "insert into ccms.project_track (project_id, track_id) values ($1, $2)"
	if _, err := d.Q.Exec(d.C, sql, projectID, trackID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

func alterProjectDropTrack(d *dbx.DB, project, track string, projectID int64) error {
	if track == "*" {
		sql := "delete from ccms.project_track where project_id=$1"
		if _, err := d.Q.Exec(d.C, sql, projectID); err != nil {
			return pgerr.Error(err)
		}
		return nil
	}
	// look up track id
	trackID, err := selectTrackID(d, track)
	if err != nil {
		return err
	}
	if trackID == -1 {
		return errors.New("track \"" + track + "\" does not exist")
	}
	// check if project track exists
	projectTrackExists, err := projectTrackExists(d, projectID, trackID)
	if err != nil {
		return err
	}
	if !projectTrackExists {
		return errors.New("project \"" + project + "\" does not have track \"" + track + "\"")
	}
	// drop project track
	sql := "delete from ccms.project_track where project_id=$1 and track_id=$2"
	if _, err := d.Q.Exec(d.C, sql, projectID, trackID); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

// returns project id, or -1 if project does not exist
func selectProjectID(d *dbx.DB, project string) (int64, error) {
	var q = "select id from ccms.project where name=$1"
	var id int64
	err := d.Q.QueryRow(d.C, q, project).Scan(&id)
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
func selectLocationID(d *dbx.DB, location string) (int64, error) {
	sql := "select id from ccms.location where name=$1"
	var id int64
	err := d.Q.QueryRow(d.C, sql, location).Scan(&id)
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
func selectOriginID(d *dbx.DB, origin string) (int64, error) {
	var sql = "select id from ccms.origin where name=$1"
	var id int64
	err := d.Q.QueryRow(d.C, sql, origin).Scan(&id)
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
func selectDestinationID(d *dbx.DB, destination string) (int64, error) {
	sql := "select id from ccms.destination where name=$1"
	var id int64
	err := d.Q.QueryRow(d.C, sql, destination).Scan(&id)
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
func selectTrackID(d *dbx.DB, track string) (int64, error) {
	sql := "select id from ccms.track where name=$1"
	var id int64
	err := d.Q.QueryRow(d.C, sql, track).Scan(&id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return -1, nil
	case err != nil:
		return 0, pgerr.Error(err)
	default:
		return id, nil
	}
}

func projectFundExists(d *dbx.DB, projectID, fundID int64) (bool, error) {
	var q = "select 1 from ccms.project_fund where project_id=$1 and fund_id=$2"
	var n int32
	err := d.Q.QueryRow(d.C, q, projectID, fundID).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func projectLocationExists(d *dbx.DB, projectID, locationID int64) (bool, error) {
	var q = "select 1 from ccms.project_location where project_id=$1 and location_id=$2"
	var n int32
	err := d.Q.QueryRow(d.C, q, projectID, locationID).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func projectOriginExists(d *dbx.DB, projectID, originID int64) (bool, error) {
	var sql = "select 1 from ccms.project_origin where project_id=$1 and origin_id=$2"
	var n int32
	err := d.Q.QueryRow(d.C, sql, projectID, originID).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func projectDestinationExists(d *dbx.DB, projectID, destinationID int64) (bool, error) {
	var sql = "select 1 from ccms.project_destination where project_id=$1 and destination_id=$2"
	var n int32
	err := d.Q.QueryRow(d.C, sql, projectID, destinationID).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func projectTrackExists(d *dbx.DB, projectID, trackID int64) (bool, error) {
	sql := "select 1 from ccms.project_track where project_id=$1 and track_id=$2"
	var n int32
	err := d.Q.QueryRow(d.C, sql, projectID, trackID).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func AlterProjectSetProperty(d *dbx.DB, projectName, property, value string, stringLiteral bool) error {
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
	if _, err := d.Q.Exec(d.C, sql, value, projectName); err != nil {
		return errors.New("updating project: " + pgerr.String(err))
	}
	return nil
}

func ProjectProperties(d *dbx.DB, projectName string) ([][2]string, error) {
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
	err := d.Q.QueryRow(d.C, q, projectName).Scan(&title, &action, &mouLink, &funds, &locations, &origins, &destinations, &tracks)
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
