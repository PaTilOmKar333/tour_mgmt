package repo

import (
	"tour_mgmt/app"

	"github.com/jmoiron/sqlx"
)

type UserRepoInterface interface {
}

type userRepo struct {
	db *sqlx.DB
}

func InitUserRepo() UserRepoInterface {
	var ur userRepo

	ur.db = app.GetDB()
	return &ur
}
