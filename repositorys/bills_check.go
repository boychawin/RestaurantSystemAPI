package repositorys

import (
	"restaurant/errors"
	"restaurant/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BillCheckRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) PostBillsCheck(c *fiber.Ctx) (*ResponseBillCheck, error) {
	billID := c.Params("id")
	bill := models.Bill{}
	membership := models.Membership{}
	request := DataResponseBillCheck{}
	err := c.BodyParser(&request)
	if err != nil {
		return nil, err
	}

	if billID == "" || request.Amount == 0 {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	result := u.db.Where("id = ? AND status != 'ปิด'", billID).First(&bill)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
		return nil, result.Error
	}
	var change float64
	var totaldiscount float64
	var totalAmount float64

	if request.CardNumber != "" {
		u.db.Where("card_number LIKE ?", request.CardNumber).First(&membership)
	}

	expiryDate, err := time.Parse("2006-01-02T00:00:00Z", membership.ExpiryDate)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	if membership.DiscountPercentage != 0 && expiryDate.After(time.Now()) {
		request.DiscountPercentage = membership.DiscountPercentage

		percentage := membership.DiscountPercentage / 100.0
		totalAmount = percentage * bill.AmountPaid
		totaldiscount = bill.AmountPaid - totalAmount

	} else {
		totaldiscount = bill.AmountPaid
	}
	// Calculate
	if request.Amount >= bill.AmountPaid {
		change = request.Amount - bill.AmountPaid + totalAmount
	}

	if request.Amount < totaldiscount {
		return nil, errors.NewUnexpectedError("ยอดเงินที่ได้รับไม่ถูกต้อง")
	}

	if change != 0 {
		// Update
		tx := u.db.Model(&models.Bill{}).Where("id = ?", billID).Updates(&models.Bill{Change: change})
		if tx.Error != nil {
			return nil, errors.NewUnexpectedError(tx.Error.Error())
		}
	}

	request.ID = bill.ID
	request.TableID = bill.TableID
	request.Status = string(bill.Status)
	request.Amount = bill.AmountPaid
	request.TotalAmount = totaldiscount
	request.Change = change

	response := &ResponseBillCheck{
		Data:     &request,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) PostBillsClose(c *fiber.Ctx) (*ResponseBillCheck, error) {
	billID := c.Params("id")
	bill := models.Bill{}

	if billID == "" {
		return nil, errors.NewUnexpectedError("ต้องระบุฟิลด์ให้ครบ")
	}

	result := u.db.Where("id = ? AND status != 'ปิด'", billID).First(&bill)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
		return nil, result.Error
	}

	// Update
	tx := u.db.Model(&models.Bill{}).Where("id = ?", billID).Updates(&models.Bill{Status: models.StatusCloseBill})
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	// Update
	tx = u.db.Model(&models.Table{}).Where("id = ?", bill.TableID).Updates(&models.Table{Status: models.StatusEmpty})
	if tx.Error != nil {
		return nil, errors.NewUnexpectedError(tx.Error.Error())
	}

	response := &ResponseBillCheck{
		Data:     nil,
		Messages: "",
	}
	return response, nil
}

func (u usersRepositoryDB) GetBillsCheck(c *fiber.Ctx) (*ResponseBillCheck, error) {
	billID := c.Params("id")
	billCheck := DataResponseBillCheck{}
	bill := models.Bill{}
	request := models.Membership{}
	membership := models.Membership{}
	c.BodyParser(&request)

	result := u.db.Where("id = ? AND status != 'ปิด'", billID).First(&bill)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewUnexpectedError("ไม่พบข้อมูล")
		}
		return nil, result.Error
	}

	var totalQuantity float64
	var totalAmount float64
	var totaldiscount float64

	u.db.Raw("SELECT SUM(products.price) as total_quantity FROM orders JOIN products ON orders.product_id = products.id WHERE orders.bill_id = ?", billID).Scan(&totalQuantity)
	if totalQuantity != 0 {
		// Update
		tx := u.db.Model(&models.Bill{}).Where("id = ?", billID).Updates(&models.Bill{AmountPaid: totalQuantity})
		if tx.Error != nil {
			return nil, errors.NewUnexpectedError(tx.Error.Error())
		}
	}

	if request.CardNumber != "" {
		u.db.Where("card_number LIKE ?", request.CardNumber).First(&membership)
	}

	expiryDate, err := time.Parse("2006-01-02T00:00:00Z", membership.ExpiryDate)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}

	if membership.DiscountPercentage != 0 && expiryDate.After(time.Now()) {
		billCheck.DiscountPercentage = membership.DiscountPercentage

		percentage := membership.DiscountPercentage / 100.0
		totalAmount = percentage * totalQuantity
		totaldiscount = totalQuantity - totalAmount

	} else {
		totaldiscount = totalQuantity
	}
	billCheck.CardNumber = request.CardNumber
	billCheck.ID = bill.ID
	billCheck.TableID = bill.TableID
	billCheck.Amount = totalQuantity
	billCheck.TotalAmount = totaldiscount
	billCheck.Status = string(bill.Status)

	response := &ResponseBillCheck{
		Data: &billCheck,
	}

	return response, nil
}
