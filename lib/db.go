package lib

import (
	"log"

	db "upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var settings = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "",
	Database: "members",
}

// Sessão do banco de dados
var Sesss db.Database

func init() {
	var err error
	Sesss, err = mysql.Open(settings)

	if err != nil {
		log.Fatal(err.Error())
	}
}
