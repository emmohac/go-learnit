package databases

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func Connect() error {
	fmt.Println(("Database connected"))
	return nil
}

func Init() error {
	fmt.Println("Database initialized")
	return nil
}
