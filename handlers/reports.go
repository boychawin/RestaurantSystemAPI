package handlers

import (
	"restaurant/services"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
)

func ReportHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv}
}

func (h userHandler) GetTotalAmountIncome(c *fiber.Ctx) error {
	response, err := h.userSrv.GetTotalAmountIncome(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetProductCategory(c *fiber.Ctx) error {
	response, err := h.userSrv.GetProductCategory(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}


func (h userHandler) GetBillCategorySummary(c *fiber.Ctx) error {
	response, err := h.userSrv.GetBillCategorySummary(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetBillSummary(c *fiber.Ctx) error {
	response, err := h.userSrv.GetBillSummary(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetCustomerSummary(c *fiber.Ctx) error {
	response, err := h.userSrv.GetCustomerSummary(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetCustomerAgeGroupSummary(c *fiber.Ctx) error {
	response, err := h.userSrv.GetCustomerAgeGroupSummary(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetCustomerAenderSummary(c *fiber.Ctx) error {
	response, err := h.userSrv.GetCustomerAenderSummary(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}
func (h userHandler) GetRepeatCustomers(c *fiber.Ctx) error {
	response, err := h.userSrv.GetRepeatCustomers(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetTop10Food(c *fiber.Ctx) error {
	response, err := h.userSrv.GetTop10Food(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}