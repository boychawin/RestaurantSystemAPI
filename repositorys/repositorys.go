package repositorys

import (
	"restaurant/models"
	"time"

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
	GetProductCategory(c *fiber.Ctx) (*ResponseReportCategory, error)
	GetBillCategorySummary(c *fiber.Ctx) (*ResponseReportCategoryBillCount, error)
	GetBillSummary(c *fiber.Ctx) (*ResponseReportBillCount, error)
	GetCustomerSummary(c *fiber.Ctx) (*ResponseReportBillCount, error)
	GetCustomerAgeGroupSummary(c *fiber.Ctx) (*ResponseReportAgeGroupSummary, error)
	GetCustomerGenderSummary(c *fiber.Ctx) (*ResponseReportGenderSummary, error)
	GetRepeatCustomers(c *fiber.Ctx) (*ResponseReportCustomerRepeatVisits, error)
	GetTop10Food(c *fiber.Ctx) (*ResponseReportMonthlyTopFood, error)
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

type RevenueResponse struct {
	Last7DaysRevenue   float64
	Last15DaysRevenue  float64
	Last1MonthRevenue  float64
	Last3MonthsRevenue float64
}

type ResponseReport struct {
	Data     *RevenueResponse `json:"data"`
	Messages string           `json:"messages"`
}

type CategoryRevenueResponse struct {
	CategoryName       string
	Last7DaysRevenue   float64
	Last15DaysRevenue  float64
	Last1MonthRevenue  float64
	Last3MonthsRevenue float64
}

type ResponseReportCategory struct {
	Data     []CategoryRevenueResponse `json:"data"`
	Messages string                    `json:"messages"`
}

type CategoryBillCount struct {
	CategoryName     string
	Last7DaysCount   int64
	Last15DaysCount  int64
	Last1MonthCount  int64
	Last3MonthsCount int64
}

type ResponseReportCategoryBillCount struct {
	Data     []CategoryBillCount `json:"data"`
	Messages string              `json:"messages"`
}

type BillCount struct {
	Last7DaysCount   int64
	Last15DaysCount  int64
	Last1MonthCount  int64
	Last3MonthsCount int64
}

type ResponseReportBillCount struct {
	Data     BillCount `json:"data"`
	Messages string    `json:"messages"`
}

type GenderSummaryCount struct {
	Gender           string
	Last7DaysCount   int64
	Last15DaysCount  int64
	Last1MonthCount  int64
	Last3MonthsCount int64
}

type ResponseReportGenderSummary struct {
	Data     []GenderSummaryCount `json:"data"`
	Messages string               `json:"messages"`
}

type AgeGroupSummaryCount struct {
	AgeGroupStart    int
	AgeGroupEnd      int
	Last7DaysCount   int64
	Last15DaysCount  int64
	Last1MonthCount  int64
	Last3MonthsCount int64
}

type ResponseReportAgeGroupSummary struct {
	Data     []AgeGroupSummaryCount `json:"data"`
	Messages string                 `json:"messages"`
}

type MonthlyTopFood struct {
	Month      time.Month
	Year       int
	FoodName   string
	TotalSales int
}

type ResponseReportMonthlyTopFood struct {
	Data     []MonthlyTopFood `json:"data"`
	Messages string           `json:"messages"`
}

type CustomerRepeatVisits struct {
	CustomerName        string
	TotalVisits15Days   int64
	TotalVisits1Month   int64
	TotalVisits3Months  int64
	TotalVisits6Months  int64
	TotalVisits12Months int64
}

type ResponseReportCustomerRepeatVisits struct {
	Data     []CustomerRepeatVisits `json:"data"`
	Messages string                 `json:"messages"`
}
