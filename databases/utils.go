package databases

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() (*pgx.Conn, error) {
	connString := "postgresql://postgres:password@postgres/?sslmode=disable"
	DB, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer DB.Close(context.Background())

	fmt.Println(("Database connected"))
	return DB, nil
}

func Init(db *pgx.Conn) error {
	var userTable = `
		CREATE TABLE IF NOT EXISTS USERS (
			Id SERIAL PRIMARY KEY,
			FirstName VARCHAR(255),
			LastName VARCHAR(255),
			Email VARCHAR(255) NOT NULL UNIQUE,
			Password VARCHAR(255),
			Permission VARCHAR(5)
		);
	`
	returnVal, err := db.Exec(context.Background(), userTable)
	fmt.Println(returnVal)

	var insertAdmin = `
		INSERT INTO users (Id, FirstName, LastName, Email, Password, Permission)
		VALUES (1, 'Huy Minh', 'Tran', 'huy@admin.com', 'password', '1')
		ON CONFLICT (Id) DO NOTHING;
	`

	db.Exec(context.Background(), insertAdmin)
	return err
}
