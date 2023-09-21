package handlers

import (
	"restaurant/services"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
)

func ProductHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv}
}

func (h userHandler) PostProducts(c *fiber.Ctx) error {
	response, err := h.userSrv.PostProducts(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetProducts(c *fiber.Ctx) error {
	response, err := h.userSrv.GetProducts(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}


func (h userHandler) PutProducts(c *fiber.Ctx) error {
	response, err := h.userSrv.PutProducts(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) DeleteProducts(c *fiber.Ctx) error {
	response, err := h.userSrv.DeleteProducts(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}