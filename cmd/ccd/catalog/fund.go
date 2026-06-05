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
	"github.com/indexdata/ccms/prop"
	"github.com/jackc/pgx/v5"
)

// CreateFund does not do synchronization and must not access the catalog cache
func (c *Catalog) CreateFund(fund string) error {
	q := "insert into ccms.fund (name, title) values ($1, $2)"
	if _, err := c.dp.Exec(context.TODO(), q, fund, makeTitle(fund)); err != nil {
		return pgerr.Error(err)
	}
	return nil
}

// FundExists does not do synchronization and must not access the catalog cache
func (c *Catalog) FundExists(fund string) (bool, error) {
	id, err := c.selectFundID(fund)
	if err != nil {
		return false, err
	}
	return id != -1, nil
}

// returns fund id, or -1 if fund does not exist
// SelectFundID does not do synchronization and must not access the catalog cache
func (c *Catalog) SelectFundID(fund string) (int64, error) {
	return c.selectFundID(fund)
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

func (c *Catalog) AllFunds() ([]prop.Prop, error) {
	q := "select name, title from ccms.fund"
	rows, err := c.dp.Query(context.TODO(), q)
	if err != nil {
		return nil, fmt.Errorf("selecting funds: %v", err)
	}
	defer rows.Close()
	funds := make([]prop.Prop, 0)
	for rows.Next() {
		var f prop.Prop
		if err := rows.Scan(&f.Name, &f.Title); err != nil {
			return nil, fmt.Errorf("reading funds: %v", err)
		}
		funds = append(funds, f)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("reading funds: %v", err)
	}
	sortFundNames(funds)
	return funds, nil
}

func sortFundNames(funds []prop.Prop) {
	slices.SortFunc(funds, func(x, y prop.Prop) int {
		return cmp.Compare(x.Name, y.Name)
	})
}

func makeTitle(name string) string {
	var b strings.Builder
	q := '_'
	for _, r := range name {
		if r == '_' {
			if q != '_' {
				b.WriteRune(' ')
			}
			q = r
			continue
		}
		if q == '_' {
			b.WriteRune(unicode.ToUpper(r))
			q = r
			continue
		}
		b.WriteRune(r)
		q = r
	}
	return b.String()
}
