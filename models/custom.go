package models

type DistroCount struct {
	DistroName  string `bun:"linux_distro"`
	DistroCount int32  `bun:"count"`
}
