package repositorys

import (
	"restaurant/errors"
	"restaurant/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TableRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) PostTables(c *fiber.Ctx) (*ResponseTable, error) {
	request := models.Table{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.Number == "" || request.Status == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	// var count int64
	// u.db.Model(&models.Table{}).Where("category_name = ?", request.Number).Count(&count)

	// if count >= 1 {
	// 	return nil, errors.NewUnexpectedError("มีข้อมูลแล้ว")
	// }

	// Insert
	tx := u.db.Create(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	response := &ResponseTable{
		Data:     nil,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) GetTables(c *fiber.Ctx) (*ResponseTable, error) {
	id := c.Params("id")
	product := []models.Table{}

	result := u.db
	if id != "" {
		result = result.Where("id = ?", id).First(&product)
	} else {
		result = result.Find(&product)
	}

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
		return nil, result.Error
	}

	if len(product) == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
	}

	response := &ResponseTable{
		Data: product,
	}

	return response, nil
}

func (u usersRepositoryDB) PutTables(c *fiber.Ctx) (*ResponseTable, error) {
	Id := c.Params("id")
	request := models.Table{}
	table := models.Table{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}

	if request.Number == "" || request.Status == "" || Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err = u.db.Where("id = ?", Id).First(&table).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	if request.Status == table.Status {
		return nil, errors.NewUnexpectedError("สถานะเดิม")
	}

	// Update
	tx := u.db.Model(&models.Table{}).Where("id = ?", Id).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	products := &ResponseTable{
		Data: nil,
	}

	return products, nil
}

func (u usersRepositoryDB) DeleteTables(c *fiber.Ctx) (*ResponseTable, error) {
	Id := c.Params("id")

	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&models.Table{}).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Delete
	tx := u.db.Delete(&models.Table{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseTable{
		Data: nil,
	}

	return user, nil
}
