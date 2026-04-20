package catalog

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"github.com/indexdata/ccms/internal/crypto"
	"github.com/indexdata/ccms/internal/global"
)

type User struct {
	UserName  string
	Superuser bool
	Login     bool
}

func (c *Catalog) initAuth() error {
	sql := "select name, superuser, login, password, salt from ccms.auth"
	rows, err := c.dp.Query(context.TODO(), sql)
	if err != nil {
		return fmt.Errorf("selecting users: %v", err)
	}
	defer rows.Close()
	users := make(map[string]auth)
	for rows.Next() {
		var username, password, salt string
		var superuser, login bool
		if err := rows.Scan(&username, &superuser, &login, &password, &salt); err != nil {
			return fmt.Errorf("reading users: %v", err)
		}
		s, _ := crypto.DecodeFromHexString(salt)
		users[username] = auth{
			superuser: superuser,
			login:     login,
			password:  password,
			salt:      s,
		}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("reading users: %v", err)
	}
	c.users = users
	return nil
}

func (c *Catalog) Authenticate(username, password string) bool {
	return c.users[username].login &&
		c.users[username].password == crypto.HashPassword(password, c.users[username].salt, c.secretKey)
}

func (c *Catalog) UserExists(userName string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.users[userName]
	return ok
}

func (c *Catalog) CreateUser(userName, password string, superuser, login bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("creating user %q: opening transaction: %v", userName, err)
	}
	defer tx.Rollback(context.TODO())

	salt := crypto.RandomKey()
	hash := crypto.HashPassword(password, salt, c.secretKey)
	sql := "insert into ccms.auth (name, superuser, login, password, salt) values ($1, $2, $3, $4, $5)"
	if _, err := tx.Exec(context.TODO(), sql, userName, superuser, login, hash, crypto.EncodeToHexString(salt)); err != nil {
		return fmt.Errorf("registering user %q: %v", userName, global.PGErr(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("creating user %q: committing changes: %v", userName, err)
	}

	c.users[userName] = auth{
		superuser: superuser,
		login:     login,
		password:  hash,
		salt:      salt,
	}
	return nil
}

func (c *Catalog) AllUsers() []User {
	c.mu.Lock()
	defer c.mu.Unlock()
	users := make([]User, len(c.users))
	i := 0
	for k, v := range c.users {
		users[i] = User{
			UserName:  k,
			Superuser: v.superuser,
			Login:     v.login,
		}
		i++
	}
	sortUserNames(users)
	return users
}

func sortUserNames(users []User) {
	slices.SortFunc(users, func(x, y User) int {
		return cmp.Compare(x.UserName, y.UserName)
	})
}
