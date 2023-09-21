package repositorys

import (
	"restaurant/models"

	"github.com/gofiber/fiber/v2"
)

type AuthResponse struct {
	Data         *models.Users `json:"data"`
	Messages     string        `json:"messages"`
	AccessToken  string        `json:"accessToken"`
	RefreshToken string        `json:"refreshToken"`
}

type Users struct {
	ID       uint
	Username string `gorm:"unique;size(50)"`
}
type SignupRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UsersRepository interface {
	Register(c *fiber.Ctx) (*AuthResponse, error)
	Login(c *fiber.Ctx) (*AuthResponse, error)
	Refresh(c *fiber.Ctx) (*AuthResponse, error)
	Session(c *fiber.Ctx) (*AuthResponse, error)
	/** User **/
	GetUsers(c *fiber.Ctx) (*ResponseUsers, error)
	PutUsers(c *fiber.Ctx) (*ResponseUsers, error)
	DeleteUsers(c *fiber.Ctx) (*ResponseUsers, error)
	/** Product **/
	PostProducts(c *fiber.Ctx) (*ResponseProduct, error)
	GetProducts(c *fiber.Ctx) (*ResponseProduct, error)
	PutProducts(c *fiber.Ctx) (*ResponseProduct, error)
	DeleteProducts(c *fiber.Ctx) (*ResponseProduct, error)

	/** Product Category **/
	PostCategory(c *fiber.Ctx) (*ResponseProductCategory, error)
	GetCategory(c *fiber.Ctx) (*ResponseProductCategory, error)
	PutCategory(c *fiber.Ctx) (*ResponseProductCategory, error)
	DeleteCategory(c *fiber.Ctx) (*ResponseProductCategory, error)

	/** Table **/
	PostTables(c *fiber.Ctx) (*ResponseTable, error)
	GetTables(c *fiber.Ctx) (*ResponseTable, error)
	PutTables(c *fiber.Ctx) (*ResponseTable, error)
	DeleteTables(c *fiber.Ctx) (*ResponseTable, error)

	/** Reservation **/
	PostReservations(c *fiber.Ctx) (*ResponseReservation, error)
	GetReservations(c *fiber.Ctx) (*ResponseReservation, error)
	PutReservations(c *fiber.Ctx) (*ResponseReservation, error)
	DeleteReservations(c *fiber.Ctx) (*ResponseReservation, error)

	/** Bill **/
	PostBills(c *fiber.Ctx) (*ResponseBill, error)
	GetBills(c *fiber.Ctx) (*ResponseBill, error)
	PutBills(c *fiber.Ctx) (*ResponseBill, error)
	DeleteBills(c *fiber.Ctx) (*ResponseBill, error)

	/** Payment system / bill check **/
	PostBillsCheck(c *fiber.Ctx) (*ResponseBillCheck, error)
	PostBillsClose(c *fiber.Ctx) (*ResponseBillCheck, error)
	GetBillsCheck(c *fiber.Ctx) (*ResponseBillCheck, error)

	/** Order **/
	PostOrders(c *fiber.Ctx) (*ResponseOrder, error)
	GetOrders(c *fiber.Ctx) (*ResponseOrder, error)
	PutOrders(c *fiber.Ctx) (*ResponseOrder, error)
	DeleteOrders(c *fiber.Ctx) (*ResponseOrder, error)

	/** Membership **/
	PostMemberships(c *fiber.Ctx) (*ResponseMembership, error)
	GetMemberships(c *fiber.Ctx) (*ResponseMembership, error)
	PutMemberships(c *fiber.Ctx) (*ResponseMembership, error)
	DeleteMemberships(c *fiber.Ctx) (*ResponseMembership, error)

	/** Report **/
	GetTotalAmountIncome(c *fiber.Ctx) (*ResponseReport, error)
	GetProductCategory(c *fiber.Ctx) (*ResponseReport, error)
	GetBillCategorySummary(c *fiber.Ctx) (*ResponseReport, error)
	GetBillSummary(c *fiber.Ctx) (*ResponseReport, error)
	GetCustomerSummary(c *fiber.Ctx) (*ResponseReport, error)
	GetCustomerAgeGroupSummary(c *fiber.Ctx) (*ResponseReport, error)
	GetCustomerAenderSummary(c *fiber.Ctx) (*ResponseReport, error)
	GetRepeatCustomers(c *fiber.Ctx) (*ResponseReport, error)
	GetTop10Food(c *fiber.Ctx) (*ResponseReport, error)
}

type ResponseUsers struct {
	Data     []models.Users `json:"data"`
	Messages string         `json:"messages"`
}

type ResponseProduct struct {
	Data     []models.Product `json:"data"`
	Messages string           `json:"messages"`
}

type ResponseProductCategory struct {
	Data     []models.ProductCategory `json:"data"`
	Messages string                   `json:"messages"`
}

type ResponseBill struct {
	Data     []models.Bill `json:"data"`
	Messages string        `json:"messages"`
}

type DataResponseBillCheck struct {
	ID                 uint
	TableID            int
	DiscountPercentage float64
	CardNumber         string
	// Gender        string
	// AgeGroup      string
	Amount      float64
	TotalAmount float64
	Change      float64
	// PaymentMethod string
	Status string
}

type ResponseBillCheck struct {
	Data     *DataResponseBillCheck `json:"data"`
	Messages string                 `json:"messages"`
}

type ResponseOrder struct {
	Data     []models.Order `json:"data"`
	Messages string         `json:"messages"`
}

type ResponseMembership struct {
	Data     []models.Membership `json:"data"`
	Messages string              `json:"messages"`
}

type ResponseReservation struct {
	Data     []models.Reservation `json:"data"`
	Messages string               `json:"messages"`
}

type ResponseTable struct {
	Data     []models.Table `json:"data"`
	Messages string         `json:"messages"`
}

type ResponseReport struct {
	Data     string `json:"data"`
	Messages string `json:"messages"`
}
