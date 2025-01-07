package pg

import (
	"github.com/jackc/pgx"
)

// NewConnPool creates and returns a new postgresql connection pool.
func NewConnPool(dbname, user, pwd, host string, port int) (*pgx.ConnPool, error) {
	p, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			User:     user,
			Password: pwd,
			Host:     host,
			Port:     uint16(port),
			Database: dbname,
		},
	})
	if err != nil {
		return nil, err
	}

	return p, nil
}
