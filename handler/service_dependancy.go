package handler

import (
	"tour_mgmt/repo"
	"tour_mgmt/service"
)

type depedecies struct {
	UserService    service.UserServiceInterface
	BookingService service.BookingServiceInterface
}

var depend depedecies

func InitDependancies() {
	userRepo := repo.InitUserRepo()
	userService := service.InitUserService(userRepo)

	bookingRepo := repo.InitBookingRepo()
	bookingService := service.InitBookingService(bookingRepo)

	depend.UserService = userService
	depend.BookingService = bookingService
}
