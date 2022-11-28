package services

import (
	"context"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/safaci2000/golug/dbmodels"
	"github.com/safaci2000/golug/models"
	log "github.com/sirupsen/logrus"

	"sync"
)

var doOnce sync.Once

type ServiceContract interface {
	ListUsers() []models.LinuxUser
	GetUser(id int64) (*models.LinuxUser, error)
	UpdateUser(user models.LinuxUser) (*models.LinuxUser, error)
	DeleteUser(id int64) error
	CreateUser(user models.LinuxUser) (*models.LinuxUser, error)
}

type MagicService struct {
	DB    *pgx.Conn
	Query *dbmodels.Queries
}

var instance *MagicService

func GetServices() ServiceContract {
	return instance
}

func InitializeServices(dbURI string) ServiceContract {
	doOnce.Do(func() {
		instance = &MagicService{}
		var err error
		ctx := context.Background()
		instance.DB, err = pgx.Connect(ctx, dbURI)
		if err != nil {
			log.Fatalln(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		instance.Query = dbmodels.New(instance.DB)
	})

	return instance
}
