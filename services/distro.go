package services

import (
	"context"

	"github.com/csg33k/golug/dbmodels"
	log "github.com/sirupsen/logrus"
)

// ListDistributions return list of linux distributions supported
func (s *MagicService) ListDistributions() ([]string, error) {
	return s.query.GetLinuxDistros(context.Background())
}

// LinuxDistroCount demonstrates a few more raw queries patterns
func (s *MagicService) LinuxDistroCount() ([]dbmodels.GetLinuxDistroCountRow, error) {
	// select linux_distro, count(*) from linux_user group by linux_distro;
	r, err := s.query.GetLinuxDistroCount(context.Background())
	log.Info(r)

	res := []dbmodels.GetLinuxDistroCountRow{}
	err = s.queryBuilder.Select(&res, "select linux_distro, count(*) from linux_user group by linux_distro order by count(*)")
	user := dbmodels.LinuxUser{}
	err = s.queryBuilder.Get(&user, "SELECT * FROM linux_user WHERE linux_user_id=$1", 1)
	log.Info(user)

	return res, err
}
