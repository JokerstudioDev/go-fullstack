package service

import (
	"model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserDacTestifyMock struct {
	mock.Mock
}

func (o *UserDacTestifyMock) Read() []model.User {
	args := o.Mock.Called()
	return args.Get(0).([]model.User)
}

func (o *UserDacTestifyMock) Add() bool {
	args := o.Mock.Called()
	return args.Bool(0)
}

func (o *UserDacTestifyMock) Remove(id string) bool {
	args := o.Mock.Called(id)
	return args.Bool(0)
}

func (o *UserDacTestifyMock) Update(id string, firstnameFather string) bool {
	args := o.Mock.Called(id, firstnameFather)
	return args.Bool(0)
}

func (o *UserDacTestifyMock) ReadById(citizenId string) model.User {
	args := o.Mock.Called(citizenId)
	return args.Get(0).(model.User)
}

func Test_UserById_Success_With_Mock(t *testing.T) {
	expected := "002"

	userDacMock := &UserDacTestifyMock{}
	userDacMock.On("ReadById", "002").Return(model.User{ID: 2, CitizenID: "002"})

	userSvc := NewUserSvc(userDacMock)
	actual := userSvc.UserById("002")

	userDacMock.AssertCalled(t, "ReadById", "002")
	userDacMock.AssertNumberOfCalls(t, "ReadById", 1)
	assert.Equal(t, expected, actual.CitizenID, "they should be equal")
}
