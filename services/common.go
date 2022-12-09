package services

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"

	"sync"

	"github.com/csg33k/golug/dbmodels"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var doOnce sync.Once

type ServiceContract interface {
	ListDistributions() ([]string, error)
	ListUsers() ([]dbmodels.LinuxUser, error)
	GetUser(id int64) (*dbmodels.LinuxUser, error)
	UpdateUser(user dbmodels.LinuxUser) (*dbmodels.LinuxUser, error)
	DeleteUser(id int64) error
	CreateUser(user dbmodels.LinuxUser) (*dbmodels.LinuxUser, error)
	LinuxDistroCount() ([]dbmodels.GetLinuxDistroCountRow, error)
}

type MagicService struct {
	dbPool         *pgxpool.Pool
	query          *dbmodels.Queries //sqlc generated code.
	queryBuilder   *sqlx.DB
	poolConnection *pgxpool.Conn
}

var instance *MagicService

func GetServices() ServiceContract {
	return instance
}

// CleaUp: Handle connection cleanup
func (s *MagicService) CleaUp() {
	s.poolConnection.Release()
	s.dbPool.Close()

}

// InitializeServices: initialize service and setups connection pool
func InitializeServices(dbURI string) ServiceContract {
	doOnce.Do(func() {
		instance = &MagicService{}
		ctx := context.Background()
		var err error
		instance.queryBuilder, err = sqlx.Connect("pgx", dbURI)
		if err != nil {
			log.Fatalln(err)
		}
		instance.dbPool, err = pgxpool.Connect(ctx, dbURI)
		instance.poolConnection, err = instance.dbPool.Acquire(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		instance.query = dbmodels.New(instance.poolConnection)
		if err != nil {
			log.Fatal(err)
		}
	})

	return instance
}
