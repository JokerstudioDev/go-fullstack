package service

import (
	"dac"
	"fmt"
	"model"
)

type UserSvc struct{}

var userDac dac.IUserDac

func NewUserSvc(dac dac.IUserDac) *UserSvc {
	userDac = dac
	return &UserSvc{}
}

func (*UserSvc) Users() []model.User {
	users := userDac.Read()
	return users
}

func (*UserSvc) UserById(citizenId string) model.User {
	user := userDac.ReadById(citizenId)
	fmt.Printf(user.CitizenID)
	return user
}

func (*UserSvc) UpdateFirstnameFather(id, firstnameFather string) bool {
	user := userDac.ReadById(id)
	if user.CitizenID == "" {
		return false
	}
	status := userDac.Update(id, firstnameFather)
	return status
}

func (*UserSvc) AddUser() bool {
	status := userDac.Add()
	return status
}

func (*UserSvc) RemoveUser(id string) bool {
	status := userDac.Remove(id)
	return status
}
