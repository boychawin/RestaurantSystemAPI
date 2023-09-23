package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func OrderCycleservice(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) PostOrderCycle(c *fiber.Ctx) (*repositorys.ResponseOrderCycle, error) {
	// Repository
	response, err := s.userRepo.PostOrderCycle(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetOrdersCycle(c *fiber.Ctx) (*repositorys.ResponseOrderCycle, error) {
	response, err := s.userRepo.GetOrdersCycle(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) PutOrdersCycle(c *fiber.Ctx) (*repositorys.ResponseOrderCycle, error) {
	response, err := s.userRepo.PutOrdersCycle(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) DeleteOrdersCycle(c *fiber.Ctx) (*repositorys.ResponseOrderCycle, error) {
	response, err := s.userRepo.DeleteOrdersCycle(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
