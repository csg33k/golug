package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/safaci2000/golug/dbmodels"
	"github.com/safaci2000/golug/models"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *MagicService) ListUsers() []models.LinuxUser {
	ctx := context.Background()

	users, err := dbmodels.LinuxUsers().All(ctx, s.DB)
	if err != nil {
		logrus.Panic("Cannot get contact forms data")
	}
	result := make([]models.LinuxUser, 0)
	for _, val := range users {
		c := models.NewLinusUser(*val)
		result = append(result, *c)
	}
	return result
}

func (s *MagicService) GetUser(id int64) (*models.LinuxUser, error) {
	ctx := context.Background()
	user, err := dbmodels.FindLinuxUser(ctx, s.DB, id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return models.NewLinusUser(*user), nil

}

func (s *MagicService) UpdateUser(user models.LinuxUser) (*models.LinuxUser, error) {
	_, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	var newUser dbmodels.LinuxUser
	newUser.LinuxUserName.SetValid(user.Name)
	newUser.LinuxDistro.SetValid(user.Distro)
	newUser.LinuxUserID = user.Id
	_, err = newUser.Update(ctx, s.DB, boil.Infer())

	return models.NewLinusUser(newUser), nil
}

func (s *MagicService) DeleteUser(id int64) error {
	ctx := context.Background()
	user, err := dbmodels.FindLinuxUser(ctx, s.DB, id)
	if err != nil {
		return err
	}
	_, err = user.Delete(ctx, s.DB)
	return err
}

func (s *MagicService) CreateUser(user models.LinuxUser) (*models.LinuxUser, error) {
	ctx := context.Background()
	var newUser dbmodels.LinuxUser
	newUser.LinuxUserName.SetValid(user.Name)
	newUser.LinuxDistro.SetValid(user.Distro)
	err := newUser.Insert(ctx, s.DB, boil.Infer()) // Insert the first pilot with name "Larry"
	if err != nil {
		return nil, fmt.Errorf("could not create user %s", user.Name)
	}
	return s.GetUser(newUser.LinuxUserID)

}
