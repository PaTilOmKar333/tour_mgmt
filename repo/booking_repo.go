package repo

import (
	"tour_mgmt/app"

	"github.com/jmoiron/sqlx"
)

type BookingRepoInterface interface {
}

type bookingRepo struct {
	db *sqlx.DB
}

func InitBookingRepo() BookingRepoInterface {
	var br bookingRepo

	br.db = app.GetDB()
	return &br

}
