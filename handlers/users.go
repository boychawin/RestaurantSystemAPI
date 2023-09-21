package handlers

import (
	"restaurant/services"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv}
}

func (h userHandler) GetUsers(c *fiber.Ctx) error {
	response, err := h.userSrv.GetUsers(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}

	return utils.GenerateResponse(c, response.Data, response.Messages, nil)
}

func (h userHandler) PutUsers(c *fiber.Ctx) error {
	response, err := h.userSrv.PutUsers(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data, response.Messages, nil)
}

func (h userHandler) DeleteUsers(c *fiber.Ctx) error {
	response, err := h.userSrv.DeleteUsers(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data, response.Messages, nil)

}
