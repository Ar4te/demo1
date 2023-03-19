package dto

import (
	"ginDemo/model"
	"strconv"
)

type UserDto struct {
	Id        string `json:"Id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id:        strconv.FormatInt(int64(user.ID), 10),
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
