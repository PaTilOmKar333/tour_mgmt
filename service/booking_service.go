package service

import "tour_mgmt/repo"

type BookingServiceInterface interface {
}

type bookingService struct {
	repo repo.BookingRepoInterface
}

func InitBookingService(r repo.BookingRepoInterface) BookingServiceInterface {
	return bookingService{
		repo: r,
	}
}
