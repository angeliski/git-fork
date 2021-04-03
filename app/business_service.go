package app

import (
	"errors"
	"fmt"
	"github.com/angeliski/git-fork/domain/repository"
)

type businessService struct {
	repository repository.Service
}

// NewBusinessService returns a service to run the app
func NewBusinessService(service repository.Service) BusinessService {
	return &businessService{
		repository: service,
	}
}

// Fork provides a way to fork and set upstream remote in your git repository
func (b *businessService) Fork(model repository.Model) error {
	return b.repository.AddRemote(model)
}

// Sync enables your repository to be up to date form the upstream
func (b *businessService) Sync(model repository.Model) error {
	if !b.repository.HasRemote("upstream") {
		return errors.New("remote upstream not configured. Use git-fork fork` to configure")
	}

	fmt.Println("Fetching")
	err := b.repository.Fetch(model)

	if err != nil {
		return err
	}

	fmt.Println("Pulling")
	err = b.repository.Pull(model)

	if err != nil {
		return err
	}

	fmt.Println("Pushing")

	return b.repository.Push(repository.Model{
		RemoteName: "origin",
		Branch:     model.Branch,
	})
}
