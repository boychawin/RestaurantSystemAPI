package repositorys

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ReportRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) GetTotalAmountIncome(c *fiber.Ctx) (*ResponseReport, error) {
	
	return nil, nil
}


func (u usersRepositoryDB) GetProductCategory(c *fiber.Ctx) (*ResponseReport, error) {
	
	return nil, nil
}


func (u usersRepositoryDB) GetBillCategorySummary(c *fiber.Ctx) (*ResponseReport, error) {
	
	return nil, nil
}

func (u usersRepositoryDB) GetBillSummary(c *fiber.Ctx) (*ResponseReport, error) {
	
	return nil, nil
}


func (u usersRepositoryDB) GetCustomerSummary(c *fiber.Ctx) (*ResponseReport, error) {
	
	return nil, nil
}

func (u usersRepositoryDB) GetCustomerAgeGroupSummary(c *fiber.Ctx) (*ResponseReport, error) {
	
	return nil, nil
}

func (u usersRepositoryDB) GetCustomerAenderSummary(c *fiber.Ctx) (*ResponseReport, error) {
	
	return nil, nil
}

func (u usersRepositoryDB) GetRepeatCustomers(c *fiber.Ctx) (*ResponseReport, error) {
	
	return nil, nil
}

func (u usersRepositoryDB) GetTop10Food(c *fiber.Ctx) (*ResponseReport, error) {
	
	return nil, nil
}
