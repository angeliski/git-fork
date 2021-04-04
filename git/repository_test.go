package git_test

import (
	"github.com/angeliski/git-fork/domain/repository"
	"github.com/angeliski/git-fork/git"
	mockexecutor "github.com/angeliski/git-fork/mocks/executor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperations(t *testing.T) {
	var tests = []struct {
		name         string
		model        repository.Model
		expectedArgs []string
		path         string
		call         func(*testing.T, repository.Model, repository.Service)
	}{
		{
			name:         "Fetch",
			model:        repository.Model{RemoteName: "origin", Branch: "master"},
			expectedArgs: []string{"fetch", "origin", "master"},
			path:         "my/path",
			call: func(t *testing.T, model repository.Model, service repository.Service) {
				err := service.Fetch(model)

				assert.Nil(t, err)
			},
		},
		{
			name:         "Pull",
			model:        repository.Model{RemoteName: "origin", Branch: "master"},
			expectedArgs: []string{"pull", "origin", "master"},
			path:         "my/path",
			call: func(t *testing.T, model repository.Model, service repository.Service) {
				err := service.Pull(model)

				assert.Nil(t, err)
			},
		},
		{
			name:         "Push",
			model:        repository.Model{RemoteName: "origin", Branch: "master"},
			expectedArgs: []string{"push", "origin", "master"},
			path:         "my/path",
			call: func(t *testing.T, model repository.Model, service repository.Service) {
				err := service.Push(model)

				assert.Nil(t, err)
			},
		},
		{
			name:         "AddRemote",
			model:        repository.Model{RemoteName: "origin", URLRemote: "url.git"},
			expectedArgs: []string{"remote", "add", "origin", "url.git"},
			path:         "my/path",
			call: func(t *testing.T, model repository.Model, service repository.Service) {
				err := service.AddRemote(model)

				assert.Nil(t, err)
			},
		},
		{
			name:         "HasRemote",
			model:        repository.Model{RemoteName: "origin", Branch: "master"},
			expectedArgs: []string{"remote", "get-url", "origin"},
			path:         "my/path",
			call: func(t *testing.T, model repository.Model, service repository.Service) {
				hasRemote := service.HasRemote(model.RemoteName)

				assert.True(t, hasRemote)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock := &mockexecutor.CommandExecutor{}

			service, err := git.NewRepositoryService(test.path, mock)

			assert.Nil(t, err)

			mock.On("RunGitOperation", test.expectedArgs, test.path).Return(nil)

			test.call(t, test.model, service)

			mock.AssertCalled(t, "RunGitOperation", test.expectedArgs, test.path)
		})
	}
}
