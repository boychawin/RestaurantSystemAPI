package repositorys

import (
	"restaurant/errors"
	"restaurant/models"
	"restaurant/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func ProductRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) PostProducts(c *fiber.Ctx) (*ResponseProduct, error) {
	claims := c.Locals("claims").(jwt.MapClaims)
	userIdStr := claims["iss"].(string)
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}
	request := models.Product{}
	err = c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if request.Name == "" || request.Description == "" || request.Price == 0 || request.CategoryID == 0 {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	var count int64
	u.db.Model(&models.Product{}).Where("name = ?", request.Name).Count(&count)

	if count >= 1 {
		return nil, errors.NewUnexpectedError("มีข้อมูลแล้ว")
	}

	image, err := utils.UploadAndResizeImageHandler(c, "uploads/product/")
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	request.UserID = userIdInt
	request.Image = image.Filename
	request.ImageHash = image.MD5Hash

	// Insert
	tx := u.db.Create(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	response := &ResponseProduct{
		Data:     nil,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) GetProducts(c *fiber.Ctx) (*ResponseProduct, error) {
	id := c.Params("id")
	product := []models.Product{}

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

	response := &ResponseProduct{
		Data: product,
	}

	return response, nil
}

func (u usersRepositoryDB) PutProducts(c *fiber.Ctx) (*ResponseProduct, error) {
	Id := c.Params("id")

	claims := c.Locals("claims").(jwt.MapClaims)
	userIdStr := claims["iss"].(string)
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	request := models.Product{}
	err = c.BodyParser(&request)
	if err != nil {
		return nil, errors.NewUnexpectedError("รูปแบบไม่ถูกต้อง")
	}

	if Id == "" || request.Name == "" || request.Description == "" || request.Price == 0 || request.CategoryID == 0 {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	// request.ID = uint(Id)
	// request.UserID = userIdInt

	err = u.db.Where("id = ?", Id).First(&models.Product{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	image, err := utils.UploadAndResizeImageHandler(c, "uploads/product/")

	if image != nil {
		request.UserID = userIdInt
		request.Image = image.Filename
		request.ImageHash = image.MD5Hash
	}
	// Update

	tx := u.db.Model(&models.Product{}).Where("id = ?", Id).Updates(&request)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}



	products := &ResponseProduct{
		Data: nil,
	}

	return products, nil
}

func (u usersRepositoryDB) DeleteProducts(c *fiber.Ctx) (*ResponseProduct, error) {
	Id := c.Params("id")
	product := models.Product{}
	if Id == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	err := u.db.Where("id = ?", Id).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
	}

	err = utils.DeleteImage(product.Image)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	// Delete
	tx := u.db.Delete(&models.Product{}, Id)
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &ResponseProduct{
		Data: nil,
	}

	return user, nil
}
