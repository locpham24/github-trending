package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Sql struct {
	DB *sqlx.DB
}

func (s *Sql) Connect() {
	s.DB = sqlx.MustConnect("postgres", os.Getenv("DB_URI"))

	if err := s.DB.Ping(); err != nil {
		log.Error(err.Error())
		return
	}

	fmt.Println("Connected to database....")
}

func (s *Sql) Close() {
	s.DB.Close()
}
