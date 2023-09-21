package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func TableService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) PostTables(c *fiber.Ctx) (*repositorys.ResponseTable, error) {

	response, err := s.userRepo.PostTables(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetTables(c *fiber.Ctx) (*repositorys.ResponseTable, error) {

	response, err := s.userRepo.GetTables(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) PutTables(c *fiber.Ctx) (*repositorys.ResponseTable, error) {

	response, err := s.userRepo.PutTables(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) DeleteTables(c *fiber.Ctx) (*repositorys.ResponseTable, error) {

	response, err := s.userRepo.DeleteTables(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
