package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/csg33k/golug/dbmodels"
)

func (s *MagicService) ListUsers() ([]dbmodels.LinuxUser, error) {
	ctx := context.Background()
	return s.query.GetLinuxUsers(ctx)
}

func (s *MagicService) GetUser(id int64) (*dbmodels.LinuxUser, error) {
	ctx := context.Background()
	res, err := s.query.GetLinuxUser(ctx, id)
	return &res, err
}

func (s *MagicService) UpdateUser(user dbmodels.LinuxUser) (*dbmodels.LinuxUser, error) {
	_, err := s.GetUser(user.LinuxUserID)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	valid := s.validateDistro(user.LinuxDistro)
	if !valid {
		return nil, errors.New("invalid Linux distribution selected")
	}
	request := dbmodels.UpdateLinuxUserParams{
		LinuxUserID:   user.LinuxUserID,
		LinuxUserName: user.LinuxUserName,
		LinuxDistro:   user.LinuxDistro,
	}
	err = s.query.UpdateLinuxUser(ctx, request)
	return s.GetUser(user.LinuxUserID)
}

func (s *MagicService) DeleteUser(id int64) error {
	ctx := context.Background()
	_, err := s.query.GetLinuxUser(ctx, id)
	if err != nil {
		return err
	}
	err = s.query.DeleteLinuxUser(ctx, id)
	return err
}

func (s *MagicService) validateDistro(distro string) bool {
	q, err := s.query.GetLinuxDistro(context.Background(), distro)
	if err != nil || q == "" {
		return false
	}
	return true
}

func (s *MagicService) CreateUser(user dbmodels.LinuxUser) (*dbmodels.LinuxUser, error) {
	ctx := context.Background()
	valid := s.validateDistro(user.LinuxDistro)
	if !valid {
		return nil, errors.New("invalid Linux distribution selected")
	}

	userRequest := dbmodels.CreateLinuxUserParams{
		LinuxUserName: user.LinuxUserName,
		LinuxDistro:   user.LinuxDistro,
	}
	newUser, err := s.query.CreateLinuxUser(ctx, userRequest)
	if err != nil {
		return nil, fmt.Errorf("could not create user %s", user.LinuxUserName)
	}
	return s.GetUser(newUser.LinuxUserID)

}
