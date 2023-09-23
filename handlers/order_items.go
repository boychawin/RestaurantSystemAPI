package handlers

import (
	"restaurant/services"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
)

func OrderHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv}
}

func (h userHandler) PostOrders(c *fiber.Ctx) error {
	response, err := h.userSrv.PostOrders(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetOrders(c *fiber.Ctx) error {
	response, err := h.userSrv.GetOrders(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}


func (h userHandler) PutOrders(c *fiber.Ctx) error {
	response, err := h.userSrv.PutOrders(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) DeleteOrders(c *fiber.Ctx) error {
	response, err := h.userSrv.DeleteOrders(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}