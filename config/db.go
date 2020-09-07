package config

import (
	"log"
	"os"
	"github.com/go-pg/pg/v10"
)

func Connect() *pg.Db {
	opts := &pg.Options{
		User: "taurai",
		Password: "password",
		Addr: "localhost:5432",
		Database: "testdb",
	}
}
