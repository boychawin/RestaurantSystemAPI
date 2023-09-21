package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func BillCheckService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) PostBillsCheck(c *fiber.Ctx) (*repositorys.ResponseBillCheck, error) {
	// Repository
	response, err := s.userRepo.PostBillsCheck(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) PostBillsClose(c *fiber.Ctx) (*repositorys.ResponseBillCheck, error) {
	// Repository
	response, err := s.userRepo.PostBillsClose(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}


func (s userService) GetBillsCheck(c *fiber.Ctx) (*repositorys.ResponseBillCheck, error) {
	response, err := s.userRepo.GetBillsCheck(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
