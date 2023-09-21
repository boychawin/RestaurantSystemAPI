package repositorys

import (
	"restaurant/errors"
	"restaurant/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ReservationRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) PostReservations(c *fiber.Ctx) (*ResponseReservation, error) {
	request := models.Reservation{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.CustomerName == "" || request.TableID == 0 {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	var count int64
	u.db.Model(&models.Table{}).Where("id = ? AND status LIKE 'ว่าง' ", request.TableID).Count(&count)

	if count != 1 {
		return nil, errors.NewUnexpectedError("โต๊ะไม่ว่าง")
	}

	// Insert
	tx := u.db.Create(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}
	request2 := models.Table{}
	request2.Status = "จอง" 
	// Update Tables
	tx = u.db.Where("id = ?", request.TableID).Updates(&request2)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	response := &ResponseReservation{
		Data:     nil,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) GetReservations(c *fiber.Ctx) (*ResponseReservation, error) {
	id := c.Params("id")
	rese := []models.Reservation{}

	result := u.db
	if id != "" {
		result = result.Where("id = ?", id).First(&rese)
	} else {
		result = result.Find(&rese)
	}

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
		return nil, result.Error
	}

	if len(rese) == 0 {
		return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
	}

	response := &ResponseReservation{
		Data: rese,
	}

	return response, nil
}

func (u usersRepositoryDB) PutReservations(c *fiber.Ctx) (*ResponseReservation, error) {
	Id := c.Params("id")
	request := models.Reservation{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}

	if request.CustomerName == "" ||Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err = u.db.Where("id = ? AND ", Id).First(&models.Reservation{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Update
	tx := u.db.Model(&models.Reservation{}).Where("id = ?", Id).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	products := &ResponseReservation{
		Data: nil,
	}

	return products, nil
}

func (u usersRepositoryDB) DeleteReservations(c *fiber.Ctx) (*ResponseReservation, error) {
	Id := c.Params("id")

	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&models.Reservation{}).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Delete
	tx := u.db.Delete(&models.Reservation{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseReservation{
		Data: nil,
	}

	return user, nil
}
