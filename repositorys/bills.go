package repositorys

import (
	"restaurant/errors"
	"restaurant/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BillRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) PostBills(c *fiber.Ctx) (*ResponseBill, error) {
	request := models.Bill{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.TableID == 0 || request.Number == 0 || request.AgeGroupStart == 0 || request.AgeGroupEnd == 0 || request.Gender == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	// Insert
	tx := u.db.Create(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	response := &ResponseBill{
		Data:     nil,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) GetBills(c *fiber.Ctx) (*ResponseBill, error) {
	id := c.Params("id")
	bill := []models.Bill{}

	result := u.db
	if id != "" {
		result = result.Where("id = ?", id).First(&bill)
	} else {
		result = result.Find(&bill)
	}

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
		return nil, result.Error
	}

	if len(bill) == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
	}

	response := &ResponseBill{
		Data: bill,
	}

	return response, nil
}

func (u usersRepositoryDB) PutBills(c *fiber.Ctx) (*ResponseBill, error) {
	Id := c.Params("id")

	// claims := c.Locals("claims").(jwt.MapClaims)
	// userId := claims["iss"].(string)

	request := models.Bill{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}

	if request.TableID == 0 || request.Number == 0 || request.AgeGroupStart == 0 || request.AgeGroupEnd == 0 || request.Gender == "" || Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err = u.db.Where("id = ?", Id).First(&models.Bill{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Update
	tx := u.db.Model(&models.Bill{}).Where("id = ?", Id).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	products := &ResponseBill{
		Data: nil,
	}

	return products, nil
}

func (u usersRepositoryDB) DeleteBills(c *fiber.Ctx) (*ResponseBill, error) {
	Id := c.Params("id")

	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&models.Bill{}).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Delete
	tx := u.db.Delete(&models.Bill{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseBill{
		Data: nil,
	}

	return user, nil
}
