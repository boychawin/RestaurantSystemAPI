package repositorys

import (
	"restaurant/errors"
	"restaurant/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func OrderRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) PostOrders(c *fiber.Ctx) (*ResponseOrder, error) {
	request := models.Order{}
	bill := models.Bill{}
	OrderCycle := models.OrderCycle{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.OrderCycleID == 0 || request.BillID == 0 || request.ProductID == 0 {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}
	request.Status = models.StatusPending
	var countOrderCycle int64

	u.db.Model(&models.OrderCycle{}).Where("id = ?", request.OrderCycleID).Count(&countOrderCycle).First(&OrderCycle)
	if countOrderCycle == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
	}

	var count int64
	u.db.Model(&models.Bill{}).Where("id = ?", request.BillID).Count(&count).First(&bill)
	if count == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูลบิล")
	}

	if count == 1 && bill.Status == models.StatusCloseBill {
		return nil, errors.NewUnexpectedError("บิลใช้งานไม่ได้")
	}

	var countProduct int64
	u.db.Model(&models.Product{}).Where("id = ?", request.ProductID).Count(&countProduct)
	if countProduct == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูลสินค้า")
	}

	// Insert
	tx := u.db.Create(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	response := &ResponseOrder{
		Data:     nil,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) GetOrders(c *fiber.Ctx) (*ResponseOrder, error) {
	id := c.Params("id")
	bill := []models.Order{}

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

	response := &ResponseOrder{
		Data: bill,
	}

	return response, nil
}

func (u usersRepositoryDB) PutOrders(c *fiber.Ctx) (*ResponseOrder, error) {
	Id := c.Params("id")
	request := models.Order{}
	orders := models.Order{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}
	if request.Status == "" || Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err = u.db.Where("id = ?", Id).First(&orders).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	if request.Status == models.StatusCancelled && orders.Status == models.StatusServing {
		return nil, errors.NewUnexpectedError("รายการนี้ยกเลิกไม่ได้")
	}

	if orders.Status == models.StatusCancelled {
		return nil, errors.NewUnexpectedError("รายการนี้ยกเลิกไปแล้ว")
	}

	if orders.Status == request.Status {
		return nil, errors.NewUnexpectedError("สถานะเดิม")
	}

	var countProduct int64
	u.db.Model(&models.Product{}).Where("id = ?", request.ProductID).Count(&countProduct)
	if countProduct == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูลสินค้า")
	}

	// Update
	tx := u.db.Model(&models.Order{}).Where("id = ?", Id).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	response := &ResponseOrder{
		Data: nil,
	}

	return response, nil
}

func (u usersRepositoryDB) DeleteOrders(c *fiber.Ctx) (*ResponseOrder, error) {
	Id := c.Params("id")

	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&models.Order{}).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	// Update Order Status
	tx := u.db.Model(&models.Order{}).Where("id = ?", Id).Updates(&models.Order{Status: models.StatusCancelled})
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	// Delete
	tx = u.db.Delete(&models.Order{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseOrder{
		Data: nil,
	}

	return user, nil
}
