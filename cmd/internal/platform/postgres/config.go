package postgres

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConfigDb() (*sql.DB, error) {
	// local run: set DATABASE_URL=postgresql://postgres:admin@localhost:5432/DepartmentSearch
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
