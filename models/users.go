package models

import (
	"strconv"

	"github.com/emmohac/go-learnit/databases"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func GetPersons(count int) ([]User, error) {

	rows, err := databases.DB.Query("SELECT id, first_name, last_name, email, password from users LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		singleUser := User{}
		err = rows.Scan(&singleUser.Id, &singleUser.FirstName, &singleUser.LastName, &singleUser.Email, &singleUser.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, singleUser)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return users, err
}
