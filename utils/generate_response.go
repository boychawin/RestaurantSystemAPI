package utils

import "github.com/gofiber/fiber/v2"


func GenerateResponse(c *fiber.Ctx, data interface{},messages string, err error) error {
	if err != nil {
        return c.JSON(fiber.Map{
            "Status":   "error",
            "Messages": err.Error(),
        })
    }

    response := fiber.Map{
        "Status": "ok",
    }

	if messages != "" {
		response["Messages"] = messages
	}

    if data != nil {
        response["Data"] = data
    }

    return c.JSON(response)
}
