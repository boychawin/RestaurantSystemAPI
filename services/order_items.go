package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func Orderservice(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) PostOrders(c *fiber.Ctx) (*repositorys.ResponseOrder, error) {
	// Repository
	response, err := s.userRepo.PostOrders(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetOrders(c *fiber.Ctx) (*repositorys.ResponseOrder, error) {
	response, err := s.userRepo.GetOrders(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) PutOrders(c *fiber.Ctx) (*repositorys.ResponseOrder, error) {
	response, err := s.userRepo.PutOrders(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) DeleteOrders(c *fiber.Ctx) (*repositorys.ResponseOrder, error) {
	response, err := s.userRepo.DeleteOrders(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
