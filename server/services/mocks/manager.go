// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	crypto "crypto"

	config "go.woodpecker-ci.org/woodpecker/v2/server/services/config"

	environment "go.woodpecker-ci.org/woodpecker/v2/server/services/environment"

	forge "go.woodpecker-ci.org/woodpecker/v2/server/forge"

	mock "github.com/stretchr/testify/mock"

	model "go.woodpecker-ci.org/woodpecker/v2/server/model"

	registry "go.woodpecker-ci.org/woodpecker/v2/server/services/registry"

	secret "go.woodpecker-ci.org/woodpecker/v2/server/services/secret"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// ConfigServiceFromRepo provides a mock function with given fields: repo
func (_m *Manager) ConfigServiceFromRepo(repo *model.Repo) config.Service {
	ret := _m.Called(repo)

	if len(ret) == 0 {
		panic("no return value specified for ConfigServiceFromRepo")
	}

	var r0 config.Service
	if rf, ok := ret.Get(0).(func(*model.Repo) config.Service); ok {
		r0 = rf(repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Service)
		}
	}

	return r0
}

// EnvironmentService provides a mock function with given fields:
func (_m *Manager) EnvironmentService() environment.Service {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EnvironmentService")
	}

	var r0 environment.Service
	if rf, ok := ret.Get(0).(func() environment.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(environment.Service)
		}
	}

	return r0
}

// ForgeFromRepo provides a mock function with given fields: repo
func (_m *Manager) ForgeFromRepo(repo *model.Repo) (forge.Forge, error) {
	ret := _m.Called(repo)

	if len(ret) == 0 {
		panic("no return value specified for ForgeFromRepo")
	}

	var r0 forge.Forge
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Repo) (forge.Forge, error)); ok {
		return rf(repo)
	}
	if rf, ok := ret.Get(0).(func(*model.Repo) forge.Forge); ok {
		r0 = rf(repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(forge.Forge)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Repo) error); ok {
		r1 = rf(repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForgeFromUser provides a mock function with given fields: user
func (_m *Manager) ForgeFromUser(user *model.User) (forge.Forge, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for ForgeFromUser")
	}

	var r0 forge.Forge
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.User) (forge.Forge, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*model.User) forge.Forge); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(forge.Forge)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForgeMain provides a mock function with given fields:
func (_m *Manager) ForgeMain() (forge.Forge, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ForgeMain")
	}

	var r0 forge.Forge
	var r1 error
	if rf, ok := ret.Get(0).(func() (forge.Forge, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() forge.Forge); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(forge.Forge)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryService provides a mock function with given fields:
func (_m *Manager) RegistryService() registry.Service {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RegistryService")
	}

	var r0 registry.Service
	if rf, ok := ret.Get(0).(func() registry.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(registry.Service)
		}
	}

	return r0
}

// RegistryServiceFromRepo provides a mock function with given fields: repo
func (_m *Manager) RegistryServiceFromRepo(repo *model.Repo) registry.Service {
	ret := _m.Called(repo)

	if len(ret) == 0 {
		panic("no return value specified for RegistryServiceFromRepo")
	}

	var r0 registry.Service
	if rf, ok := ret.Get(0).(func(*model.Repo) registry.Service); ok {
		r0 = rf(repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(registry.Service)
		}
	}

	return r0
}

// SecretService provides a mock function with given fields:
func (_m *Manager) SecretService() secret.Service {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SecretService")
	}

	var r0 secret.Service
	if rf, ok := ret.Get(0).(func() secret.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(secret.Service)
		}
	}

	return r0
}

// SecretServiceFromRepo provides a mock function with given fields: repo
func (_m *Manager) SecretServiceFromRepo(repo *model.Repo) secret.Service {
	ret := _m.Called(repo)

	if len(ret) == 0 {
		panic("no return value specified for SecretServiceFromRepo")
	}

	var r0 secret.Service
	if rf, ok := ret.Get(0).(func(*model.Repo) secret.Service); ok {
		r0 = rf(repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(secret.Service)
		}
	}

	return r0
}

// SignaturePublicKey provides a mock function with given fields:
func (_m *Manager) SignaturePublicKey() crypto.PublicKey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SignaturePublicKey")
	}

	var r0 crypto.PublicKey
	if rf, ok := ret.Get(0).(func() crypto.PublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PublicKey)
		}
	}

	return r0
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
