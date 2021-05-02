package create

<<<<<<< HEAD
type RequestCreateUser struct {
	Name UserName
}

type ResponseCreateUser struct {
	UserId string
=======
import "in-gravity/ui/restapi/defs"

type RequestCreateUser struct {
	Name defs.UserName `json:"Name" validate:"required" minLength:"3" example: "aiue ou"`
>>>>>>> Add UserName def
}
