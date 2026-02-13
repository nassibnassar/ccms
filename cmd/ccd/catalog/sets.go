package catalog

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

func (c *Catalog) initSets() error {
	sql := "select setname from ccms.sets"
	rows, err := c.dp.Query(context.TODO(), sql)
	if err != nil {
		return fmt.Errorf("selecting sets: %v", err)
	}
	defer rows.Close()
	sets := make(map[string]struct{})
	for rows.Next() {
		var set string
		if err := rows.Scan(&set); err != nil {
			return fmt.Errorf("reading sets: %v", err)
		}
		sets[set] = struct{}{}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("reading sets: %v", err)
	}
	c.sets = sets
	return nil
}

func (c *Catalog) SetExists(setName string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if setName == "reserve" {
		return true
	}
	_, ok := c.sets[setName]
	return ok
}

func IsValidTargetSet(setName string) bool {
	if setName == "reserve" {
		return false
	}
	if !strings.ContainsRune(setName, '.') {
		return false
	}
	if strings.HasPrefix(setName, ".") || strings.HasSuffix(setName, ".") {
		return false
	}
	return ProjectExists(setName)
}

// return table containing set
func SetTable(setName string) string {
	if setName == "reserve" {
		return "ccms.reserve"
	}
	return setName
}

func (c *Catalog) AllSets() []string {
	sets := make([]string, len(c.sets)+1)
	i := 0
	sets[i] = "reserve"
	i++
	for k := range c.sets {
		sets[i] = k
		i++
	}
	return sets
}

func (c *Catalog) CreateSet(setName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("creating set %q: opening transaction: %v", setName, err)
	}
	defer tx.Rollback(context.TODO())

	sql := "create table " + setName + "(" +
		"id bigint primary key)"
	if _, err := tx.Exec(context.TODO(), sql); err != nil {
		return fmt.Errorf("creating set %q: %v", setName, err)
	}
	sql = "insert into ccms.sets (setname) values ($1)"
	if _, err := tx.Exec(context.TODO(), sql, setName); err != nil {
		return fmt.Errorf("registering set %q: %v", setName, PGErr(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("creating set %q: committing changes: %v", setName, err)
	}

	c.sets[setName] = struct{}{}
	return nil
}

func (c *Catalog) DropSet(setName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tx, err := c.dp.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("dropping set %q: opening transaction: %v", setName, err)
	}
	defer tx.Rollback(context.TODO())

	sql := "drop table " + setName
	if _, err := tx.Exec(context.TODO(), sql); err != nil {
		return fmt.Errorf("dropping set %q: %v", setName, err)
	}
	sql = "delete from ccms.sets where setname=$1"
	if _, err := tx.Exec(context.TODO(), sql, setName); err != nil {
		return fmt.Errorf("deregistering set %q: %v", setName, PGErr(err))
	}

	if err := tx.Commit(context.TODO()); err != nil {
		return fmt.Errorf("dropping set %q: committing changes: %v", setName, err)
	}

	delete(c.sets, setName)
	return nil
}

func PGErr(err error) error {
	e := err.(*pgconn.PgError)
	var b strings.Builder
	b.WriteString(e.Message)
	if e.Detail != "" {
		b.WriteString(": ")
		b.WriteString(e.Detail)
	}
	if e.Hint != "" {
		b.WriteString(" (")
		b.WriteString(e.Hint)
		b.WriteRune(')')
	}
	return errors.New(b.String())
}
