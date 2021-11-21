package model

import (
	"time"
)

type Student struct {
	Id       int
	Name     string
	Gender   string
	Age      int
	JoinDate time.Time `db:"join_date"`
	IdCard   string    `db:"id_card"`
	Senior   bool
}
