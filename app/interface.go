package app

import "github.com/angeliski/git-fork/domain/repository"

// BusinessService provides the core functions for the app
type BusinessService interface {
	Fork(model repository.Model) error
	Sync(model repository.Model) error
}
