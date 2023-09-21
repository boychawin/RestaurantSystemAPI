package handlers

import (
	"restaurant/services"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
)

func MembershipHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv}
}

func (h userHandler) PostMemberships(c *fiber.Ctx) error {
	response, err := h.userSrv.PostMemberships(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) GetMemberships(c *fiber.Ctx) error {
	response, err := h.userSrv.GetMemberships(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}


func (h userHandler) PutMemberships(c *fiber.Ctx) error {
	response, err := h.userSrv.PutMemberships(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}

func (h userHandler) DeleteMemberships(c *fiber.Ctx) error {
	response, err := h.userSrv.DeleteMemberships(c)
	if err != nil {
		return utils.GenerateResponse(c, nil, "", err)
	}
	return utils.GenerateResponse(c, response.Data,response.Messages, nil)
}