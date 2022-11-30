package services

import (
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/safaci2000/golug/dbmodels"
	log "github.com/sirupsen/logrus"
	"sync"
)

var doOnce sync.Once

type ServiceContract interface {
	ListDistributions() ([]string, error)
	ListUsers() ([]dbmodels.LinuxUser, error)
	GetUser(id int64) (*dbmodels.LinuxUser, error)
	UpdateUser(user dbmodels.LinuxUser) (*dbmodels.LinuxUser, error)
	DeleteUser(id int64) error
	CreateUser(user dbmodels.LinuxUser) (*dbmodels.LinuxUser, error)
}

type MagicService struct {
	DbPool       *pgxpool.Pool
	query        *dbmodels.Queries //sqlc generated code.
	queryBuilder goqu.DialectWrapper
	dbExtended   *sqlx.DB
}

var instance *MagicService

func GetServices() ServiceContract {
	return instance
}

func InitializeServices(dbURI string) ServiceContract {
	doOnce.Do(func() {
		instance = &MagicService{}
		instance.queryBuilder = goqu.Dialect("postgres")
		instance.queryBuilder.Select("foobar")
		var err error
		ctx := context.Background()
		instance.DbPool, err = pgxpool.Connect(ctx, dbURI)
		instance.dbExtended, err = sqlx.Connect("postgres", dbURI)
		connection, err := instance.DbPool.Acquire(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		instance.query = dbmodels.New(connection)

		if err != nil {
			log.Fatal(err)
		}
	})

	return instance
}
