package repositorys

import (
	"restaurant/errors"
	"restaurant/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func OrderCycleRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) PostOrderCycle(c *fiber.Ctx) (*ResponseOrderCycle, error) {
	request := models.OrderCycle{}
	bill := models.Bill{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.BillID == 0 {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	var count int64
	u.db.Model(&models.Bill{}).Where("id = ?", request.BillID).Count(&count).First(&bill)
	if count == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูลบิล")
	}

	if count == 1 && bill.Status == models.StatusCloseBill {
		return nil, errors.NewUnexpectedError("บิลใช้งานไม่ได้")
	}

	// Insert
	tx := u.db.Create(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	response := &ResponseOrderCycle{
		Data:     nil,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) GetOrdersCycle(c *fiber.Ctx) (*ResponseOrderCycle, error) {
	id := c.Params("id")
	order := []models.OrderCycle{}

	result := u.db
	if id != "" {
		result = result.Where("id = ?", id).First(&order)
	} else {
		result = result.Find(&order)
	}

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
		return nil, result.Error
	}

	if len(order) == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
	}

	response := &ResponseOrderCycle{
		Data: order,
	}

	return response, nil
}

func (u usersRepositoryDB) PutOrdersCycle(c *fiber.Ctx) (*ResponseOrderCycle, error) {
	Id := c.Params("id")
	request := models.OrderCycle{}
	orders := models.OrderCycle{}
	bill := models.Bill{}

	err := c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}
	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err = u.db.Where("id = ?", Id).First(&orders).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	var count int64
	u.db.Model(&models.Bill{}).Where("id = ?", request.BillID).Count(&count).First(&bill)
	if count == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูลบิล")
	}

	if count == 1 && bill.Status == models.StatusCloseBill {
		return nil, errors.NewUnexpectedError("บิลใช้งานไม่ได้")
	}

	// Update
	tx := u.db.Model(&models.OrderCycle{}).Where("id = ?", Id).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	response := &ResponseOrderCycle{
		Data: nil,
	}

	return response, nil
}

func (u usersRepositoryDB) DeleteOrdersCycle(c *fiber.Ctx) (*ResponseOrderCycle, error) {
	Id := c.Params("id")

	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&models.OrderCycle{}).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	// Delete
	tx := u.db.Delete(&models.OrderCycle{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseOrderCycle{
		Data: nil,
	}

	return user, nil
}
