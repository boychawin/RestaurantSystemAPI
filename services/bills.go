package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func BillService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) PostBills(c *fiber.Ctx) (*repositorys.ResponseBill, error) {
	// Repository
	response, err := s.userRepo.PostBills(c)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetBills(c *fiber.Ctx) (*repositorys.ResponseBill, error) {
	response, err := s.userRepo.GetBills(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) PutBills(c *fiber.Ctx) (*repositorys.ResponseBill, error) {
	response, err := s.userRepo.PutBills(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) DeleteBills(c *fiber.Ctx) (*repositorys.ResponseBill, error) {
	response, err := s.userRepo.DeleteBills(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
