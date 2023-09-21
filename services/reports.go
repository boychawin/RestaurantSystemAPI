package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func ReportService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) GetTotalAmountIncome(c *fiber.Ctx) (*repositorys.ResponseReport, error) {

	response, err := s.userRepo.GetTotalAmountIncome(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetProductCategory(c *fiber.Ctx) (*repositorys.ResponseReport, error) {

	response, err := s.userRepo.GetProductCategory(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetBillCategorySummary(c *fiber.Ctx) (*repositorys.ResponseReport, error) {

	response, err := s.userRepo.GetBillCategorySummary(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetBillSummary(c *fiber.Ctx) (*repositorys.ResponseReport, error) {

	response, err := s.userRepo.GetBillSummary(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetCustomerSummary(c *fiber.Ctx) (*repositorys.ResponseReport, error) {

	response, err := s.userRepo.GetCustomerSummary(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetCustomerAgeGroupSummary(c *fiber.Ctx) (*repositorys.ResponseReport, error) {

	response, err := s.userRepo.GetCustomerAgeGroupSummary(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetCustomerAenderSummary(c *fiber.Ctx) (*repositorys.ResponseReport, error) {

	response, err := s.userRepo.GetCustomerAenderSummary(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetRepeatCustomers(c *fiber.Ctx) (*repositorys.ResponseReport, error) {

	response, err := s.userRepo.GetRepeatCustomers(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}

func (s userService) GetTop10Food(c *fiber.Ctx) (*repositorys.ResponseReport, error) {

	response, err := s.userRepo.GetTop10Food(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return response, nil
}
