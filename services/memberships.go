package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func MembershipService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) PostMemberships(c *fiber.Ctx) (*repositorys.ResponseMembership, error) {
	response, err := s.userRepo.PostMemberships(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetMemberships(c *fiber.Ctx) (*repositorys.ResponseMembership, error) {
	response, err := s.userRepo.GetMemberships(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) PutMemberships(c *fiber.Ctx) (*repositorys.ResponseMembership, error) {

	response, err := s.userRepo.PutMemberships(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) DeleteMemberships(c *fiber.Ctx) (*repositorys.ResponseMembership, error) {
	response, err := s.userRepo.DeleteMemberships(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
