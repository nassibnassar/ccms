package cat

import (
	"cmp"
	"errors"
	"fmt"
	"slices"

	"github.com/indexdata/ccms/internal/crypto"
	"github.com/indexdata/ccms/internal/dbx"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/jackc/pgx/v5"
)

type User struct {
	UserName  string
	Superuser bool
	Login     bool
}

func Authenticate(secretKey []byte, d *dbx.DB, user, password string) (bool, error) {
	sql := "select login, password, salt from ccms.auth where name=$1"
	var login bool
	var passwd, salt string
	err := d.Q.QueryRow(d.C, sql, user).Scan(&login, &passwd, &salt)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		s, _ := crypto.DecodeFromHexString(salt)
		return login && passwd == crypto.HashPassword(password, s, secretKey), nil
	}
}

func UserExists(d *dbx.DB, user string) (bool, error) {
	var q = "select 1 from ccms.auth where name=$1"
	var n int32
	err := d.Q.QueryRow(d.C, q, user).Scan(&n)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return false, nil
	case err != nil:
		return false, pgerr.Error(err)
	default:
		return true, nil
	}
}

func CreateUser(secretKey []byte, d *dbx.DB, userName, password string, superuser, login bool) error {
	salt := crypto.RandomKey()
	hash := crypto.HashPassword(password, salt, secretKey)
	q := "insert into ccms.auth (name, superuser, login, password, salt) values ($1, $2, $3, $4, $5)"
	if _, err := d.Q.Exec(d.C, q, userName, superuser, login, hash, crypto.EncodeToHexString(salt)); err != nil {
		return fmt.Errorf("registering user %q: %v", userName, pgerr.Error(err))
	}
	return nil
}

func AllUsers(d *dbx.DB) ([]User, error) {
	sql := "select name, superuser, login, password, salt from ccms.auth"
	rows, err := d.Q.Query(d.C, sql)
	if err != nil {
		return nil, pgerr.Error(err)
	}
	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		var username, password, salt string
		var superuser, login bool
		if err := rows.Scan(&username, &superuser, &login, &password, &salt); err != nil {
			return nil, pgerr.Error(err)
		}
		users = append(users, User{
			UserName:  username,
			Superuser: superuser,
			Login:     login,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, pgerr.Error(err)
	}
	sortUserNames(users)
	return users, nil
}

func sortUserNames(users []User) {
	slices.SortFunc(users, func(x, y User) int {
		return cmp.Compare(x.UserName, y.UserName)
	})
}
