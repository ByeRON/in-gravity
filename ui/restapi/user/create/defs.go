package create

import "in-gravity/ui/restapi/defs"

type RequestCreateUser struct {
	Name defs.UserName `json:"Name" validate:"required" minLength:"3" example:"aiue ou"`
}

type ResponseCreateUser struct {
	UserId string
}
