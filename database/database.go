package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBCon *sql.DB
)
