package repositorys

import (
	"restaurant/errors"
	"restaurant/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func MembershipRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) PostMemberships(c *fiber.Ctx) (*ResponseMembership, error) {
	claims := c.Locals("claims").(jwt.MapClaims)
	userIdStr := claims["iss"].(string)
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}
	request := models.Membership{}
	err = c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.CardNumber == "" || request.ExpiryDate == ""  {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	var count int64
	u.db.Model(&models.Membership{}).Where("card_number = ?", request.CardNumber).Count(&count)

	if count >= 1 {
		return nil, errors.NewUnexpectedError("มีข้อมูลแล้ว")
	}
	expiryDate, err := time.Parse("2006-01-02", request.ExpiryDate)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	request.ExpiryDate = expiryDate.Format("2006-01-02")
	request.UserID = userIdInt
	// Insert
	tx := u.db.Create(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	response := &ResponseMembership{
		Data:     nil,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) GetMemberships(c *fiber.Ctx) (*ResponseMembership, error) {
	id := c.Params("id")
	membership := []models.Membership{}

	result := u.db
	if id != "" {
		result = result.Where("id = ?", id).First(&membership)
	} else {
		result = result.Find(&membership)
	}

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
		return nil, result.Error
	}

	if len(membership) == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
	}

	response := &ResponseMembership{
		Data: membership,
	}

	return response, nil
}

func (u usersRepositoryDB) PutMemberships(c *fiber.Ctx) (*ResponseMembership, error) {
	Id := c.Params("id")
	claims := c.Locals("claims").(jwt.MapClaims)
	userIdStr := claims["iss"].(string)
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	request := models.Membership{}
	err = c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}

	if request.ExpiryDate == "" || request.DiscountPercentage == 0 || Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err = u.db.Where("id = ?", Id).First(&models.Membership{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	expiryDate, err := time.Parse("2006-01-02", request.ExpiryDate)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	request.ExpiryDate = expiryDate.Format("2006-01-02")
	request.UserID = userIdInt


	// Update
	tx := u.db.Model(&models.Membership{}).Where("id = ?", Id).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	products := &ResponseMembership{
		Data: nil,
	}

	return products, nil
}

func (u usersRepositoryDB) DeleteMemberships(c *fiber.Ctx) (*ResponseMembership, error) {
	Id := c.Params("id")

	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&models.Membership{}).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Delete
	tx := u.db.Delete(&models.Membership{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseMembership{
		Data: nil,
	}

	return user, nil
}
