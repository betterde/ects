package utils

import (
	"database/sql"
	"fmt"
	"github.com/betterde/ects/config"
	"log"
)

var (
	DB  *sql.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=%s",
		config.Conf.Database.User,
		config.Conf.Database.Pass,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Char,
	)

	DB, err = sql.Open("mysql", dsn)
}

func IsDatabaseExist() bool {
	Init()
	defer func() {
		if err := DB.Close(); err != nil {
			// TODO
		}
	}()

	statement := fmt.Sprintf("SHOW DATABASES LIKE '%s'", config.Conf.Database.Name)

	var (
		rows     *sql.Rows
		Database string
	)

	rows, err = DB.Query(statement)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&Database); err != nil {
			log.Println(err)
		}
		if Database == config.Conf.Database.Name {
			return true
		}
	}

	return false
}

func CreateDatabase() error {
	Init()
	statement := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET %s DEFAULT COLLATE %s", config.Conf.Database.Name, "utf8mb4", "utf8mb4_unicode_ci")
	_, err := DB.Query(statement)
	return err
}
