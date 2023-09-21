package repositorys

import (
	"restaurant/errors"
	"restaurant/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductCategoryRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) PostCategory(c *fiber.Ctx) (*ResponseProductCategory, error) {

	request := models.ProductCategory{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.CategoryName == ""  {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	var count int64
	u.db.Model(&models.ProductCategory{}).Where("category_name = ?", request.CategoryName).Count(&count)

	if count >= 1 {
		return nil, errors.NewUnexpectedError("มีข้อมูลแล้ว")
	}

	// Insert 
	tx := u.db.Create(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	response := &ResponseProductCategory{
		Data:     nil,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) GetCategory(c *fiber.Ctx) (*ResponseProductCategory, error) {
	id := c.Params("id")
	product := []models.ProductCategory{}

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

	response := &ResponseProductCategory{
		Data: product,
	}

	return response, nil
}


func (u usersRepositoryDB) PutCategory(c *fiber.Ctx) (*ResponseProductCategory, error) {
	Id := c.Params("id")

	// claims := c.Locals("claims").(jwt.MapClaims)
	// userId := claims["iss"].(string)

	request := models.ProductCategory{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}

	if request.CategoryName == "" || Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err = u.db.Where("id = ?", Id).First(&models.ProductCategory{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Update
	tx := u.db.Model(&models.ProductCategory{}).Where("id = ?", Id).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	products := &ResponseProductCategory{
		Data: nil,
	}

	return products, nil
}

func (u usersRepositoryDB) DeleteCategory(c *fiber.Ctx) (*ResponseProductCategory, error) {
	Id := c.Params("id")

	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&models.ProductCategory{}).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}
	// Delete
	tx := u.db.Delete(&models.ProductCategory{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseProductCategory{
		Data: nil,
	}

	return user, nil
}


