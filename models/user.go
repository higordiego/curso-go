package models

import "curso/lib"

// Tabelas usuários representação
type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

// User Model
var UserModel = lib.Sesss.Collection("user")
