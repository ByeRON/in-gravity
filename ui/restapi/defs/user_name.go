package defs

type UserName string

const (
	MinUserNameLength int = 3
)

func NewUserName(name string) UserName {
	/*
		if utf8.RuneCountInString(name) < MinUserNameLength {
			//TODO: Call Error
		}
	*/
	return UserName(name)
}
