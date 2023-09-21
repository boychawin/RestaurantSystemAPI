package handlers

import (
	"restaurant/services"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
)

func ReservationHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv}
}

func (h userHandler) PostReservations(c *fiber.Ctx) error {
	response, err := h.userSrv.PostReservations(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetReservations(c *fiber.Ctx) error {
	response, err := h.userSrv.GetReservations(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}


func (h userHandler) PutReservations(c *fiber.Ctx) error {
	response, err := h.userSrv.PutReservations(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) DeleteReservations(c *fiber.Ctx) error {
	response, err := h.userSrv.DeleteReservations(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}