package handlers

import (
	"restaurant/services"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
)

func OrderCycleHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv}
}

func (h userHandler) PostOrderCycle(c *fiber.Ctx) error {
	response, err := h.userSrv.PostOrderCycle(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetOrdersCycle(c *fiber.Ctx) error {
	response, err := h.userSrv.GetOrdersCycle(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}


func (h userHandler) PutOrdersCycle(c *fiber.Ctx) error {
	response, err := h.userSrv.PutOrdersCycle(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) DeleteOrdersCycle(c *fiber.Ctx) error {
	response, err := h.userSrv.DeleteOrdersCycle(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}