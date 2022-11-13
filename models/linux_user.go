package models

import "github.com/safaci2000/golug/dbmodels"

type LinuxUser struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Distro string `json:"distro,omitempty"`
}

func NewLinusUser(dbuser dbmodels.LinuxUser) *LinuxUser {
	u := &LinuxUser{}
	u.Id = dbuser.LinuxUserID
	u.Distro = dbuser.LinuxDistro.String
	u.Name = dbuser.LinuxUserName.String

	return u

}
