package cat

import (
	"cmp"
	"maps"
	"slices"

	"github.com/indexdata/ccms/cmd/ccd/dberr"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func CreateRole(d *dbx.DB, role string) error {
	sql := "insert into ccms.role (rolename) values ($1)"
	if _, err := d.Q.Exec(d.C, sql, role); err != nil {
		return dberr.Error(err)
	}
	return nil
}

type Role struct {
	RoleName  string
	UserNames []string
}

func SortRoles(roles []Role) {
	slices.SortFunc(roles, func(x, y Role) int {
		return cmp.Compare(x.RoleName, y.RoleName)
	})
}

func Roles(d *dbx.DB) ([]Role, error) {
	q := "select r.name, u.name from ccms.role r left join ccms.role_user ru on r.id=ru.role_id left join ccms.auth u on ru.user_id=u.id"
	rows, err := d.Q.Query(d.C, q)
	if err != nil {
		return nil, dberr.Error(err)
	}
	defer rows.Close()
	roles := make(map[string]roleUsers)
	for rows.Next() {
		var rolename, username string
		if err := rows.Scan(&rolename, &username); err != nil {
			return nil, dberr.Error(err)
		}
		users := roles[rolename]
		if users.users == nil {
			users.users = make(map[string]struct{})
			roles[rolename] = users
		}
		if username != "" {
			users.users[username] = struct{}{}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, dberr.Error(err)
	}

	roles1 := make([]Role, len(roles))
	i := 0
	for k, v := range roles {
		roles1[i] = Role{
			RoleName:  k,
			UserNames: slices.Sorted(maps.Keys(v.users)),
		}
		i++
	}
	return roles1, nil
}
