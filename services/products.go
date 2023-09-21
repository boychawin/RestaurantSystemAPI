package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func ProductService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) PostProducts(c *fiber.Ctx) (*repositorys.ResponseProduct,  error) {

	response, err := s.userRepo.PostProducts(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetProducts(c *fiber.Ctx) (*repositorys.ResponseProduct, error) {

	response, err := s.userRepo.GetProducts(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) PutProducts(c *fiber.Ctx) (*repositorys.ResponseProduct, error) {

	response, err := s.userRepo.PutProducts(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) DeleteProducts(c *fiber.Ctx) (*repositorys.ResponseProduct, error) {

	response, err := s.userRepo.DeleteProducts(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
