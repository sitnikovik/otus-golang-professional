package depinjection

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v5"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	memorystorage "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/storage/sql"
)

// DIContainer describes dependency injection container to work with app dependencies
type DIContainer interface {
	// Redis returns the redis storage instance
	Redis() memorystorage.Storage
	// Postgres returns the postgres storage instance
	Postgres() sqlstorage.Storage
}

// dicontainer describes the DI container instance
type dicontainer struct {
	redis    memorystorage.Storage
	postgres sqlstorage.Storage

	conf config.Config
}

// NewDIContainer creates and returns the DI container instance
func NewDIContainer(conf config.Config) DIContainer {
	di := &dicontainer{
		conf: conf,
	}

	return di
}

// Redis returns the redis storage instance
func (d *dicontainer) Redis() memorystorage.Storage {
	if d.redis == nil {
		cl := redis.NewClient(&redis.Options{
			Network: d.conf.Redis.Network,
			Addr:    d.conf.Redis.Addr,
		})
		if cl == nil {
			log.Fatal("redis client is nil on creation")
			return nil
		}
		d.redis = memorystorage.NewRedis(cl)
	}

	return d.redis
}

// Postgres returns the postgres storage instance
func (d *dicontainer) Postgres() sqlstorage.Storage {
	if d.postgres == nil {
		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			d.conf.PG.User,
			d.conf.PG.Password,
			d.conf.PG.Host,
			d.conf.PG.Port,
			d.conf.PG.Database,
		)
		db, err := pgx.Connect(context.Background(), dsn)
		if err != nil {
			log.Fatal("failed to connect to postgres: " + err.Error())
			return nil
		}
		d.postgres = sqlstorage.New(db)
	}

	return d.postgres
}
