package services

import (
	"context"
)

// ListDistributions return list of linux distributions supported
func (s *MagicService) ListDistributions() ([]string, error) {
	return s.query.GetLinuxDistros(context.Background())
}
