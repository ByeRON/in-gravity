package vo

import (
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type UserName string

const (
	MinUserNameLength int = 3
)

func NewUserName(name string) (UserName, error) {
	if utf8.RuneCountInString(name) < MinUserNameLength {
		baseError := xerrors.New("Domain Error")
		return UserName(""),
			xerrors.Errorf("User name length have to be more than %d, but %d : %v",
				MinUserNameLength,
				utf8.RuneCountInString(name),
				baseError,
			)
	}

	return UserName(name), nil
}

func (n UserName) String() string {
	return string(n)
}
