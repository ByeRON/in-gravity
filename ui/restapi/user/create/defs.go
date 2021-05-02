package create

type RequestCreateUser struct {
	Name UserName
}

type ResponseCreateUser struct {
	UserId string
}
