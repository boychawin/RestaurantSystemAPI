package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func ReservationService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) PostReservations(c *fiber.Ctx) (*repositorys.ResponseReservation, error) {

	response, err := s.userRepo.PostReservations(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetReservations(c *fiber.Ctx) (*repositorys.ResponseReservation, error) {

	response, err := s.userRepo.GetReservations(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) PutReservations(c *fiber.Ctx) (*repositorys.ResponseReservation, error) {

	response, err := s.userRepo.PutReservations(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) DeleteReservations(c *fiber.Ctx) (*repositorys.ResponseReservation, error) {

	response, err := s.userRepo.DeleteReservations(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
