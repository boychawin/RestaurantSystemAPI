package repositorys

import (
	"restaurant/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ReportRepositoryDB(db *gorm.DB) UsersRepository {
	return usersRepositoryDB{db}
}

func (u usersRepositoryDB) GetTotalAmountIncome(c *fiber.Ctx) (*ResponseReport, error) {

	var revenueResponse RevenueResponse

	timePeriods := []struct {
		Field     *float64
		StartDate time.Time
	}{
		{&revenueResponse.Last7DaysRevenue, time.Now().AddDate(0, 0, -7)},
		{&revenueResponse.Last15DaysRevenue, time.Now().AddDate(0, 0, -15)},
		{&revenueResponse.Last1MonthRevenue, time.Now().AddDate(0, -1, 0)},
		{&revenueResponse.Last3MonthsRevenue, time.Now().AddDate(0, -3, 0)},
	}

	for _, tp := range timePeriods {
		err := u.db.
			Table("orders").
			Select("SUM(products.price) as total_revenue").
			Joins("JOIN products ON orders.product_id = products.id").
			Where("orders.created_at >= ?", tp.StartDate).
			Row().
			Scan(tp.Field)

		if err != nil {
			return nil, err
		}
	}

	response := &ResponseReport{
		Data: &revenueResponse,
	}

	return response, nil
}

func (u usersRepositoryDB) GetProductCategory(c *fiber.Ctx) (*ResponseReportCategory, error) {

	var categoryRevenues []CategoryRevenueResponse


	categories := []models.ProductCategory{}
	if err := u.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	for _, category := range categories {
		var categoryRevenue CategoryRevenueResponse
		categoryRevenue.CategoryName = category.CategoryName

		timePeriods := []struct {
			Field     *float64
			StartDate time.Time
		}{
			{&categoryRevenue.Last7DaysRevenue, time.Now().AddDate(0, 0, -7)},
			{&categoryRevenue.Last15DaysRevenue, time.Now().AddDate(0, 0, -15)},
			{&categoryRevenue.Last1MonthRevenue, time.Now().AddDate(0, -1, 0)},
			{&categoryRevenue.Last3MonthsRevenue, time.Now().AddDate(0, -3, 0)},
		}

		for _, tp := range timePeriods {
			err := u.db.
				Table("orders").
				Select("SUM(products.price) as total_revenue").
				Joins("JOIN products ON orders.product_id = products.id").
				Where("products.category_id = ?", category.ID).
				Where("orders.created_at >= ?", tp.StartDate).
				Row().
				Scan(tp.Field)

			if err != nil {
				return nil, err
			}
		}

		categoryRevenues = append(categoryRevenues, categoryRevenue)
	}

	response := &ResponseReportCategory{
		Data: categoryRevenues,
	}

	return response, nil
}

func (u usersRepositoryDB) GetBillCategorySummary(c *fiber.Ctx) (*ResponseReportCategoryBillCount, error) {
	var categoryBillCounts []CategoryBillCount
	var categoryBillCount CategoryBillCount

	var categories []models.ProductCategory
	if err := u.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	timePeriods := []struct {
		Field     *int64
		StartDate time.Time
	}{
		{&categoryBillCount.Last7DaysCount, time.Now().AddDate(0, 0, -7)},
		{&categoryBillCount.Last15DaysCount, time.Now().AddDate(0, 0, -15)},
		{&categoryBillCount.Last1MonthCount, time.Now().AddDate(0, -1, 0)},
		{&categoryBillCount.Last3MonthsCount, time.Now().AddDate(0, -3, 0)},
	}

	for _, category := range categories {

		categoryBillCount.CategoryName = category.CategoryName

		for _, tp := range timePeriods {
			err := u.db.
				Table("bills").
				Joins("JOIN orders ON bills.id = orders.bill_id").
				Joins("JOIN products ON orders.product_id = products.id").
				Where("products.category_id = ?", category.ID).
				Where("bills.created_at >= ?", tp.StartDate).
				Count(tp.Field).Error

			if err != nil {
				return nil, err
			}
		}

		categoryBillCounts = append(categoryBillCounts, categoryBillCount)
	}

	response := &ResponseReportCategoryBillCount{
		Data: categoryBillCounts,
	}

	return response, nil
}

func (u usersRepositoryDB) GetBillSummary(c *fiber.Ctx) (*ResponseReportBillCount, error) {

	var billCountResponse BillCount

	timePeriods := []struct {
		Field     *int64
		StartDate time.Time
	}{
		{&billCountResponse.Last7DaysCount, time.Now().AddDate(0, 0, -7)},
		{&billCountResponse.Last15DaysCount, time.Now().AddDate(0, 0, -15)},
		{&billCountResponse.Last1MonthCount, time.Now().AddDate(0, -1, 0)},
		{&billCountResponse.Last3MonthsCount, time.Now().AddDate(0, -3, 0)},
	}

	for _, tp := range timePeriods {
		err := u.db.
			Table("bills").
			Joins("JOIN orders ON bills.id = orders.bill_id").
			Joins("JOIN products ON orders.product_id = products.id").
			Where("bills.created_at >= ?", tp.StartDate).
			Count(tp.Field).Error

		if err != nil {
			return nil, err
		}
	}

	response := &ResponseReportBillCount{
		Data: billCountResponse,
	}

	return response, nil

}

func (u usersRepositoryDB) GetCustomerSummary(c *fiber.Ctx) (*ResponseReportBillCount, error) {
	var customerCountResponse BillCount

	timePeriods := []struct {
		Field     *int64
		StartDate time.Time
	}{
		{&customerCountResponse.Last7DaysCount, time.Now().AddDate(0, 0, -7)},
		{&customerCountResponse.Last15DaysCount, time.Now().AddDate(0, 0, -15)},
		{&customerCountResponse.Last1MonthCount, time.Now().AddDate(0, -1, 0)},
		{&customerCountResponse.Last3MonthsCount, time.Now().AddDate(0, -3, 0)},
	}

	for _, tp := range timePeriods {
		err := u.db.
			Table("bills").
			Joins("JOIN orders ON bills.id = orders.bill_id").
			Joins("JOIN products ON orders.product_id = products.id").
			Where("bills.created_at >= ?", tp.StartDate).
			Pluck("SUM(bills.number)", tp.Field).Error
		// Count(tp.Field).Error

		if err != nil {
			return nil, err
		}
	}

	response := &ResponseReportBillCount{
		Data: customerCountResponse,
	}

	return response, nil
}

func (u usersRepositoryDB) GetCustomerAgeGroupSummary(c *fiber.Ctx) (*ResponseReportAgeGroupSummary, error) {

	var customerCountResponses []AgeGroupSummaryCount
	var customerCountResponse AgeGroupSummaryCount
	timePeriods := []struct {
		Field     *int64
		StartDate time.Time
	}{
		{&customerCountResponse.Last7DaysCount, time.Now().AddDate(0, 0, -7)},
		{&customerCountResponse.Last15DaysCount, time.Now().AddDate(0, 0, -15)},
		{&customerCountResponse.Last1MonthCount, time.Now().AddDate(0, -1, 0)},
		{&customerCountResponse.Last3MonthsCount, time.Now().AddDate(0, -3, 0)},
	}

	ageGroups := []struct {
		MinAge int
		MaxAge int
	}{
		{0, 18},
		{19, 30},
		{31, 45},
		{46, 60},
		{61, 100},
	}

	for _, ageGroup := range ageGroups {

		customerCountResponse.AgeGroupStart = ageGroup.MinAge
		customerCountResponse.AgeGroupEnd = ageGroup.MaxAge

		for _, tp := range timePeriods {
			err := u.db.
				Table("bills").
				Where("bills.created_at >= ?", tp.StartDate).
				Where("bills.age_group_start >= ? AND bills.age_group_end <= ?", ageGroup.MinAge, ageGroup.MaxAge).
				Count(tp.Field).Error

			if err != nil {
				return nil, err
			}
		}

		customerCountResponses = append(customerCountResponses, customerCountResponse)
	}

	response := &ResponseReportAgeGroupSummary{
		Data: customerCountResponses,
	}

	return response, nil
}

func (u usersRepositoryDB) GetCustomerGenderSummary(c *fiber.Ctx) (*ResponseReportGenderSummary, error) {
	var customerCountResponses []GenderSummaryCount
	var customerCountResponse GenderSummaryCount
	timePeriods := []struct {
		Field     *int64
		StartDate time.Time
	}{
		{&customerCountResponse.Last7DaysCount, time.Now().AddDate(0, 0, -7)},
		{&customerCountResponse.Last15DaysCount, time.Now().AddDate(0, 0, -15)},
		{&customerCountResponse.Last1MonthCount, time.Now().AddDate(0, -1, 0)},
		{&customerCountResponse.Last3MonthsCount, time.Now().AddDate(0, -3, 0)},
	}

	genders := []string{string(models.Male), string(models.Female)}
	for _, gender := range genders {

		customerCountResponse.Gender = gender

		for _, tp := range timePeriods {
			err := u.db.
				Table("bills").
				Joins("JOIN orders ON bills.id = orders.bill_id").
				Joins("JOIN products ON orders.product_id = products.id").
				Where("bills.created_at >= ?", tp.StartDate).
				Where("bills.gender LIKE ?", gender).
				Count(tp.Field).Error

			if err != nil {
				return nil, err
			}
		}

		customerCountResponses = append(customerCountResponses, customerCountResponse)

	}

	response := &ResponseReportGenderSummary{
		Data: customerCountResponses,
	}

	return response, nil
}

func (u usersRepositoryDB) GetRepeatCustomers(c *fiber.Ctx) (*ResponseReportCustomerRepeatVisits, error) {
	var repeatVisitsList []CustomerRepeatVisits

	now := time.Now()
	startDate15Days := now.AddDate(0, 0, -15)
	startDate1Month := now.AddDate(0, -1, 0)
	startDate3Months := now.AddDate(0, -3, 0)
	startDate6Months := now.AddDate(0, -6, 0)
	startDate12Months := now.AddDate(0, -12, 0)

	err := u.db.Table("reservations").
		Select("customer_name, COUNT(*) as \"total_visits\"").
		Where("created_at >= ?", startDate12Months).
		Group("customer_name").
		Having("COUNT(*) >= ?", 2).
		Scan(&repeatVisitsList).
		Error

	if err != nil {
		return nil, err
	}

	for i, customer := range repeatVisitsList {
		var totalVisits15Days, totalVisits1Month, totalVisits3Months, totalVisits6Months, totalVisit12Months int64

		u.db.Table("reservations").
			Where("customer_name = ? AND created_at >= ?", customer.CustomerName, startDate15Days).
			Count(&totalVisits15Days)

		u.db.Table("reservations").
			Where("customer_name = ? AND created_at >= ?", customer.CustomerName, startDate1Month).
			Count(&totalVisits1Month)

		u.db.Table("reservations").
			Where("customer_name = ? AND created_at >= ?", customer.CustomerName, startDate3Months).
			Count(&totalVisits3Months)

		u.db.Table("reservations").
			Where("customer_name = ? AND created_at >= ?", customer.CustomerName, startDate6Months).
			Count(&totalVisits6Months)

		u.db.Table("reservations").
			Where("customer_name = ? AND created_at >= ?", customer.CustomerName, startDate12Months).
			Count(&totalVisit12Months)

		repeatVisitsList[i].TotalVisits15Days = totalVisits15Days
		repeatVisitsList[i].TotalVisits1Month = totalVisits1Month
		repeatVisitsList[i].TotalVisits3Months = totalVisits3Months
		repeatVisitsList[i].TotalVisits6Months = totalVisits6Months
		repeatVisitsList[i].TotalVisits12Months = totalVisit12Months
	}

	response := &ResponseReportCustomerRepeatVisits{
		Data: repeatVisitsList,
	}

	return response, nil

}

func (u usersRepositoryDB) GetTop10Food(c *fiber.Ctx) (*ResponseReportMonthlyTopFood, error) {
	var monthlyTopFoods []MonthlyTopFood

	now := time.Now()
	currentMonth := now.Month()
	currentYear := now.Year()

	startDate := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.UTC)

	err := u.db.Table("orders").
		Select("EXTRACT(MONTH FROM orders.created_at) as month, EXTRACT(YEAR FROM orders.created_at) as year, products.name as food_name, SUM(products.price) as total_sales").
		Joins("JOIN products ON orders.product_id = products.id").
		Where("orders.created_at >= ?", startDate).
		Group("month, year, food_name").
		Order("total_sales DESC").
		Limit(10).
		Scan(&monthlyTopFoods).Error

	if err != nil {
		return nil, err
	}

	response := &ResponseReportMonthlyTopFood{
		Data: monthlyTopFoods,
	}

	return response, nil

}
