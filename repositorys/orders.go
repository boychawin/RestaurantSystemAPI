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
	err := c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.BillID == 0 || request.ProductID == 0  {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}
	request.Status = models.StatusPending
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

	if request.Status == models.StatusCancelled && orders.Status == models.StatusServing  {
		return nil, errors.NewUnexpectedError("รายการนี้ยกเลิกไม่ได้")
	}

	if(orders.Status == models.StatusCancelled){
		return nil, errors.NewUnexpectedError("รายการนี้ยกเลิกไปแล้ว")
	}

	if(orders.Status ==request.Status){
		return nil, errors.NewUnexpectedError("สถานะเดิม")
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
	// Delete
	tx := u.db.Delete(&models.Order{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseOrder{
		Data: nil,
	}

	return user, nil
}


