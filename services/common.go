package services

import (
	"database/sql"
	_ "github.com/lib/pq"
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
	DB *sql.DB
}

var instance *MagicService

func GetServices() ServiceContract {
	return instance
}

func InitializeServices(dbURI string) ServiceContract {
	doOnce.Do(func() {
		instance = &MagicService{}
		var err error
		instance.DB, err = sql.Open("postgres", dbURI)
		if err != nil {
			log.Fatal(err)
		}
	})

	return instance
}
