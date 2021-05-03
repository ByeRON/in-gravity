package entities

import "in-gravity/domains/vo"

type User struct {
	Id   vo.UserId
	Name vo.UserName
}

func NewUser(id vo.UserId, name vo.UserName) User {
	return User{
		Id:   id,
		Name: name,
	}
}
