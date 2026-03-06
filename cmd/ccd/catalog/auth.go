package catalog

import (
	"context"
	"fmt"

	"github.com/indexdata/ccms/internal/crypto"
	"github.com/indexdata/ccms/internal/global"
)

type auth struct {
	password string
	salt     []byte
}

func (c *Catalog) initAuth() error {
	sql := "select username, password, salt from ccms.auth"
	rows, err := c.dp.Query(context.TODO(), sql)
	if err != nil {
		return fmt.Errorf("selecting users: %v", err)
	}
	defer rows.Close()
	users := make(map[string]auth)
	for rows.Next() {
		var username, password, salt string
		if err := rows.Scan(&username, &password, &salt); err != nil {
			return fmt.Errorf("reading users: %v", err)
		}
		s, _ := crypto.DecodeFromHexString(salt)
		users[username] = auth{password: password, salt: s}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("reading users: %v", err)
	}
	c.users = users
	return nil
}

func (c *Catalog) Authenticate(username, password string) bool {
	return c.users[username].password == crypto.HashPassword(password, c.users[username].salt, c.secretKey)
}

func (c *Catalog) UserExists(userName string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.users[userName]
	return ok
}

func (c *Catalog) CreateUser(userName, password string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("creating user %q: opening transaction: %v", userName, err)
	}
	defer tx.Rollback(context.TODO())

	salt := crypto.RandomKey()
	hash := crypto.HashPassword(password, salt, c.secretKey)
	sql := "insert into ccms.auth (username, password, salt) values ($1, $2, $3)"
	if _, err := tx.Exec(context.TODO(), sql, userName, hash, crypto.EncodeToHexString(salt)); err != nil {
		return fmt.Errorf("registering user %q: %v", userName, global.PGErr(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("creating user %q: committing changes: %v", userName, err)
	}

	c.users[userName] = auth{password: hash, salt: salt}
	return nil
}

func (c *Catalog) AllUsers() []string {
	users := make([]string, len(c.users))
	i := 0
	for k := range c.users {
		users[i] = k
		i++
	}
	return users
}
