package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"in-gravity/domains/entities"
	"in-gravity/domains/repositories"
)

type MySQLRepository struct {
}

func NewUserRepository() repositories.UserRepositoryInterface {
	return MySQLRepository{}
}

func (r MySQLRepository) Save(user entities.User) error {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/in_gravity",
	)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}
