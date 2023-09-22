package services

import (
	"restaurant/models"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	Data         *models.Users `json:"data"`
	Messages     string        `json:"messages"`
	AccessToken  string        `json:"accessToken"`
	RefreshToken string        `json:"refreshToken"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	Login(c *fiber.Ctx) (*UserResponse, error)
	Register(c *fiber.Ctx) (*UserResponse, error)
	Refresh(c *fiber.Ctx) (*UserResponse, error)
	Session(c *fiber.Ctx) (*UserResponse, error)
	/** User **/
	GetUsers(c *fiber.Ctx) (*ResponseUsers, error)
	PutUsers(c *fiber.Ctx) (*repositorys.ResponseUsers, error)
	DeleteUsers(c *fiber.Ctx) (*repositorys.ResponseUsers, error)
	/** Product **/
	PostProducts(c *fiber.Ctx) (*repositorys.ResponseProduct, error)
	GetProducts(c *fiber.Ctx) (*repositorys.ResponseProduct, error)
	PutProducts(c *fiber.Ctx) (*repositorys.ResponseProduct, error)
	DeleteProducts(c *fiber.Ctx) (*repositorys.ResponseProduct, error)

	/** Product Category **/
	PostCategory(c *fiber.Ctx) (*repositorys.ResponseProductCategory, error)
	GetCategory(c *fiber.Ctx) (*repositorys.ResponseProductCategory, error)
	PutCategory(c *fiber.Ctx) (*repositorys.ResponseProductCategory, error)
	DeleteCategory(c *fiber.Ctx) (*repositorys.ResponseProductCategory, error)

	/** Table **/
	PostTables(c *fiber.Ctx) (*repositorys.ResponseTable, error)
	GetTables(c *fiber.Ctx) (*repositorys.ResponseTable, error)
	PutTables(c *fiber.Ctx) (*repositorys.ResponseTable, error)
	DeleteTables(c *fiber.Ctx) (*repositorys.ResponseTable, error)

	/** Reservation **/
	PostReservations(c *fiber.Ctx) (*repositorys.ResponseReservation, error)
	GetReservations(c *fiber.Ctx) (*repositorys.ResponseReservation, error)
	PutReservations(c *fiber.Ctx) (*repositorys.ResponseReservation, error)
	DeleteReservations(c *fiber.Ctx) (*repositorys.ResponseReservation, error)

	/** Bill **/
	PostBills(c *fiber.Ctx) (*repositorys.ResponseBill, error)
	GetBills(c *fiber.Ctx) (*repositorys.ResponseBill, error)
	PutBills(c *fiber.Ctx) (*repositorys.ResponseBill, error)
	DeleteBills(c *fiber.Ctx) (*repositorys.ResponseBill, error)

	/** Bill Check **/
	PostBillsCheck(c *fiber.Ctx) (*repositorys.ResponseBillCheck, error)
	PostBillsClose(c *fiber.Ctx) (*repositorys.ResponseBillCheck, error)
	GetBillsCheck(c *fiber.Ctx) (*repositorys.ResponseBillCheck, error)

	/** Order **/
	PostOrders(c *fiber.Ctx) (*repositorys.ResponseOrder, error)
	GetOrders(c *fiber.Ctx) (*repositorys.ResponseOrder, error)
	PutOrders(c *fiber.Ctx) (*repositorys.ResponseOrder, error)
	DeleteOrders(c *fiber.Ctx) (*repositorys.ResponseOrder, error)

	/** Membership **/
	PostMemberships(c *fiber.Ctx) (*repositorys.ResponseMembership, error)
	GetMemberships(c *fiber.Ctx) (*repositorys.ResponseMembership, error)
	PutMemberships(c *fiber.Ctx) (*repositorys.ResponseMembership, error)
	DeleteMemberships(c *fiber.Ctx) (*repositorys.ResponseMembership, error)

	/** Report **/
	GetTotalAmountIncome(c *fiber.Ctx) (*repositorys.ResponseReport, error)
	GetProductCategory(c *fiber.Ctx) (*repositorys.ResponseReportCategory, error)
	GetBillCategorySummary(c *fiber.Ctx) (*repositorys.ResponseReportCategoryBillCount, error)
	GetBillSummary(c *fiber.Ctx) (*repositorys.ResponseReportBillCount, error)
	GetCustomerSummary(c *fiber.Ctx) (*repositorys.ResponseReportBillCount, error)
	GetCustomerAgeGroupSummary(c *fiber.Ctx) (*repositorys.ResponseReportAgeGroupSummary, error)
	GetCustomerGenderSummary(c *fiber.Ctx) (*repositorys.ResponseReportGenderSummary, error)
	GetRepeatCustomers(c *fiber.Ctx) (*repositorys.ResponseReportCustomerRepeatVisits, error)
	GetTop10Food(c *fiber.Ctx) (*repositorys.ResponseReportMonthlyTopFood, error)
}

type DataResponseUsers struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ResponseUsers struct {
	Data     []DataResponseUsers `json:"data"`
	Messages string              `json:"messages"`
}
