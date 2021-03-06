// Code generated by mockery v2.7.4. DO NOT EDIT.

package executor

import mock "github.com/stretchr/testify/mock"

// CommandExecutor is an autogenerated mock type for the CommandExecutor type
type CommandExecutor struct {
	mock.Mock
}

// RunGitOperation provides a mock function with given fields: commands, path
func (_m *CommandExecutor) RunGitOperation(commands []string, path string) error {
	ret := _m.Called(commands, path)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, string) error); ok {
		r0 = rf(commands, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
