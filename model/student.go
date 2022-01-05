package model

import (
	"time"
)

type Student struct {
	Id        int
	Name      string
	Gender    string
	Age       int
	JoinDate  *time.Time `db:"join_date"`
	IdCard    string     `db:"id_card"`
	Senior    bool
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
