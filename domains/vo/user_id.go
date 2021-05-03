package vo

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UserId string

func NewUserId() (UserId, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		baseError := xerrors.New("Domain Error")
		return UserId(""),
			xerrors.Errorf("Faild to create new ID in internal error : %v", baseError)
	}

	return UserId(id.String()), nil
}

func (i UserId) String() string {
	return string(i)
}
