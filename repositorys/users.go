package repositorys

import (
	"restaurant/errors"
	"restaurant/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func UsersRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) GetUsers(c *fiber.Ctx) (*ResponseUsers, error) {
	id := c.Params("id")
	users := []models.Users{}

	result := u.db
	if id != "" {
		result = result.Where("id = ?", id).First(&users)
	} else {
		result = result.Find(&users)
	}

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
		return nil, result.Error
	}

	if len(users) == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
	}

	response := &ResponseUsers{
		Data: users,
	}

	return response, nil
}

func (u usersRepositoryDB) PutUsers(c *fiber.Ctx) (*ResponseUsers, error) {
	Id := c.Params("id")

	claims := c.Locals("claims").(jwt.MapClaims)
	userId := claims["iss"].(string)
	var IdUser string
	if Id == "" {
		IdUser = userId
	} else {
		IdUser = Id
	}

	request := models.Users{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}

	if request.FirstName == "" || request.LastName == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err = u.db.Where("id = ?", IdUser).First(&models.Users{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Update
	tx := u.db.Model(&models.Users{}).Where("id = ?", IdUser).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseUsers{
		Data: nil,
	}

	return user, nil
}

func (u usersRepositoryDB) DeleteUsers(c *fiber.Ctx) (*ResponseUsers, error) {
	Id := c.Params("id")

	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&models.Users{}).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Delete
	tx := u.db.Delete(&models.Users{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseUsers{
		Data: nil,
	}

	return user, nil
}
