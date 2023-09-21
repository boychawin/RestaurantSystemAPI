package handlers

import (
	"restaurant/services"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv services.UserService
}

func NewMainHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv}
}

func (h userHandler) Session(c *fiber.Ctx) error {

	response, err := h.userSrv.Session(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data, response.Messages, nil)

}

func (h userHandler) Refresh(c *fiber.Ctx) error {

	users, err := h.userSrv.Refresh(c)

	if err != nil {

		return c.JSON(fiber.Map{
			"Status":   "error",
			"Messages": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Status":       "ok",
		"Messages":     users.Messages,
		"AccessToken":  users.AccessToken,
		"RefreshToken": users.RefreshToken,
	})

}

func (h userHandler) Login(c *fiber.Ctx) error {

	users, err := h.userSrv.Login(c)

	if err != nil {

		return c.JSON(fiber.Map{
			"Status":   "error",
			"Messages": err.Error(),
		})
	}

	response := fiber.Map{
		"Status":       "ok",
		"Messages":     users.Messages,
		"AccessToken":  users.AccessToken,
		"RefreshToken": users.RefreshToken,
	}

	return c.JSON(response)

}

func (h userHandler) Register(c *fiber.Ctx) error {

	response, err := h.userSrv.Register(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data, response.Messages, nil)

}
