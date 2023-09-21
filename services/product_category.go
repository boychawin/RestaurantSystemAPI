package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func ProductCategoryService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) PostCategory(c *fiber.Ctx) (*repositorys.ResponseProductCategory, error) {

	response, err := s.userRepo.PostCategory(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetCategory(c *fiber.Ctx) (*repositorys.ResponseProductCategory, error) {

	response, err := s.userRepo.GetCategory(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) PutCategory(c *fiber.Ctx) (*repositorys.ResponseProductCategory, error) {

	response, err := s.userRepo.PutCategory(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) DeleteCategory(c *fiber.Ctx) (*repositorys.ResponseProductCategory, error) {

	response, err := s.userRepo.DeleteCategory(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
