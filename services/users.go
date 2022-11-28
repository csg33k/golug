package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/safaci2000/golug/dbmodels"
	"github.com/safaci2000/golug/models"
	"github.com/sirupsen/logrus"
)

func (s *MagicService) ListUsers() []models.LinuxUser {
	ctx := context.Background()

	users, err := s.Query.GetLinuxUsers(ctx)
	if err != nil {
		logrus.Panic("Cannot get contact forms data")
	}
	result := make([]models.LinuxUser, 0)
	for _, val := range users {
		c := models.NewLinusUser(val)
		result = append(result, *c)
	}
	return result
}

func (s *MagicService) GetUser(id int64) (*models.LinuxUser, error) {
	ctx := context.Background()
	user, err := s.Query.GetLinuxUser(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return models.NewLinusUser(user), nil

}

func (s *MagicService) UpdateUser(user models.LinuxUser) (*models.LinuxUser, error) {
	_, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	request := dbmodels.UpdateLinuxUserParams{
		LinuxUserID:   user.Id,
		LinuxUserName: sql.NullString{String: user.Name, Valid: true},
		LinuxDistro:   sql.NullString{String: user.Distro, Valid: true},
	}
	err = s.Query.UpdateLinuxUser(ctx, request)
	return s.GetUser(user.Id)
}

func (s *MagicService) DeleteUser(id int64) error {
	ctx := context.Background()
	_, err := s.Query.GetLinuxUser(ctx, id)
	if err != nil {
		return err
	}
	err = s.Query.DeleteLinuxUser(ctx, id)
	return err
}

func (s *MagicService) CreateUser(user models.LinuxUser) (*models.LinuxUser, error) {
	ctx := context.Background()

	userRequest := dbmodels.CreateLinuxUserParams{
		LinuxUserName: sql.NullString{String: user.Name, Valid: true},
		LinuxDistro:   sql.NullString{String: user.Distro, Valid: true},
	}
	newUser, err := s.Query.CreateLinuxUser(ctx, userRequest)
	if err != nil {
		return nil, fmt.Errorf("could not create user %s", user.Name)
	}
	return s.GetUser(newUser.LinuxUserID)

}
