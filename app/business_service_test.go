package app_test

import (
	"github.com/angeliski/git-fork/app"
	"github.com/angeliski/git-fork/domain/repository"
	mockrepository "github.com/angeliski/git-fork/mocks/domain/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBusinessServiceFork(t *testing.T) {
	t.Run("When repository not has upstream should fail", func(t *testing.T) {
		mock := mockrepository.Service{}

		mock.On("HasRemote", "upstream").Return(false)

		businessService := app.NewBusinessService(&mock)

		model := repository.Model{
			RemoteName: "upstream",
			Branch:     "master",
		}
		err := businessService.Sync(model)
		assert.NotNil(t, err)
		mock.AssertCalled(t, "HasRemote", "upstream")

	})
	t.Run("When repository has upstream should sync the repository", func(t *testing.T) {
		mock := mockrepository.Service{}

		mock.On("HasRemote", "upstream").Return(true)

		businessService := app.NewBusinessService(&mock)

		model := repository.Model{
			RemoteName: "upstream",
			Branch:     "master",
		}

		pushModel := repository.Model{
			RemoteName: "origin",
			Branch:     model.Branch,
		}

		mock.On("Fetch", model).Return(nil)
		mock.On("Pull", model).Return(nil)
		mock.On("Push", pushModel).Return(nil)

		err := businessService.Sync(model)
		assert.Nil(t, err)
		mock.AssertCalled(t, "Fetch", model)
		mock.AssertCalled(t, "Pull", model)
		mock.AssertCalled(t, "Push", pushModel)

	})
}
