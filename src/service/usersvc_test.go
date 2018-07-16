package service

import (
	"model"
	"testing"
)

func Test_UserById_Success(t *testing.T) {
	expected := "002"
	userDac := &UserDacMock{}
	userSvc := NewUserSvc(userDac)

	user := userSvc.UserById("002")

	if user.CitizenID != expected {
		t.Fatalf("%s != %s", user.CitizenID, expected)
	}
}

type UserDacMock struct{}

var users = []model.User{
	model.User{ID: 1, CitizenID: "001"},
	model.User{ID: 2, CitizenID: "002"},
}

func (*UserDacMock) Read() []model.User {
	return users
}

func (*UserDacMock) Add() bool {
	return true
}

func (*UserDacMock) Remove(id string) bool {
	return true
}

func (*UserDacMock) Update(id string, firstnameFather string) bool {
	return true
}

func (*UserDacMock) ReadById(citizenId string) model.User {
	user := model.User{}
	for _, u := range users {
		if u.CitizenID == citizenId {
			user = u
		}
	}
	return user
}
