package catalog

import (
	"cmp"
	"context"
	"fmt"
	"maps"
	"slices"

	"github.com/indexdata/ccms/internal/global"
)

func (c *Catalog) initRoles() error {
	sql := "select r.name, u.name from ccms.role r left join ccms.role_user ru on r.id=ru.role_id left join ccms.auth u on ru.user_id=u.id"
	rows, err := c.dp.Query(context.TODO(), sql)
	if err != nil {
		return fmt.Errorf("selecting roles: %v", err)
	}
	defer rows.Close()
	roles := make(map[string]roleUsers)
	for rows.Next() {
		var rolename, username string
		if err := rows.Scan(&rolename, &username); err != nil {
			return fmt.Errorf("reading roles: %v", err)
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
		return fmt.Errorf("reading roles: %v", err)
	}
	c.roles = roles
	return nil
}

func (c *Catalog) RoleExists(roleName string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.roles[roleName]
	return ok
}

func (c *Catalog) CreateRole(roleName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("creating role %q: opening transaction: %v", roleName, err)
	}
	defer tx.Rollback(context.TODO())

	sql := "insert into ccms.role (rolename) values ($1)"
	if _, err := tx.Exec(context.TODO(), sql, roleName); err != nil {
		return fmt.Errorf("registering role %q: %v", roleName, global.PGErr(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("creating role %q: committing changes: %v", roleName, err)
	}

	c.roles[roleName] = roleUsers{users: make(map[string]struct{})}
	return nil
}

type Role struct {
	RoleName  string
	UserNames []string
}

func (c *Catalog) AllRoles() []Role {
	c.mu.Lock()
	defer c.mu.Unlock()
	roles := make([]Role, len(c.roles))
	i := 0
	for k, v := range c.roles {
		roles[i] = Role{
			RoleName:  k,
			UserNames: slices.Sorted(maps.Keys(v.users)),
		}
		i++
	}
	slices.SortFunc(roles, func(x, y Role) int {
		return cmp.Compare(x.RoleName, y.RoleName)
	})
	return roles
}
