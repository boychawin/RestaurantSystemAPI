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
	table := models.Table{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.TableID == 0 || request.Number == 0 || request.AgeGroupStart == 0 || request.AgeGroupEnd == 0 || request.Gender == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	var count int64
	u.db.Model(&models.Table{}).Where("id = ?", request.TableID).Count(&count).First(&table)
	if count == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูลโต๊ะ")
	}

	if count == 1 && table.Status == models.StatusOpen {
		return nil, errors.NewUnexpectedError("มีข้อมูลแล้ว")
	}

	request.Status = models.StatusOpenBill
	// Insert
	tx := u.db.Create(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	// Update Table Status
	tx = u.db.Model(&models.Table{}).Where("id = ?", request.TableID).Updates(&models.Table{Status: models.StatusOpen})
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

	request := models.Bill{}
	bill := models.Bill{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}

	if request.TableID == 0 || request.Number == 0 || request.AgeGroupStart == 0 || request.AgeGroupEnd == 0 || request.Gender == "" || Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err = u.db.Where("id = ? AND status = 'เปิด'", Id).First(&bill).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	err = u.db.Where("id = ? AND status = 'ว่าง' ", request.TableID).First(&models.Table{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูลโต๊ะที่ว่าง")
		}
	}

	// Update
	tx := u.db.Model(&models.Bill{}).Where("id = ?", Id).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	// Update Table Status Old
	tx = u.db.Model(&models.Table{}).Where("id = ?", bill.TableID).Updates(&models.Table{Status: models.StatusEmpty})
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	// Update Table Status New
	tx = u.db.Model(&models.Table{}).Where("id = ?", request.TableID).Updates(&models.Table{Status: models.StatusOpen})
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	products := &ResponseBill{
		Data: nil,
	}

	return products, nil
}

func (u usersRepositoryDB) DeleteBills(c *fiber.Ctx) (*ResponseBill, error) {
	Id := c.Params("id")
	bill := models.Bill{}
	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&bill).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	// Update Bill Status
	tx := u.db.Model(&models.Bill{}).Where("id = ?", Id).Updates(&models.Bill{Status: models.StatusCloseBill})
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	// Update Table Status
	tx = u.db.Model(&models.Table{}).Where("id = ?", bill.TableID).Updates(&models.Table{Status: models.StatusEmpty})
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	// Delete
	tx = u.db.Delete(&models.Bill{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseBill{
		Data: nil,
	}

	return user, nil
}
