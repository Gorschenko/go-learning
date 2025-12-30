package database

import "pkg/configs"

type RepositoryDependencies struct {
	Database *Db
	Config   *configs.Config
}
