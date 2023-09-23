package main

import (
	"fmt"

	configs "restaurant/config"
	"restaurant/handlers"
	"restaurant/repositorys"
	"restaurant/services"
	"restaurant/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	configs.InitTimeZone()
	configs.InitConfig()
	db := configs.InitDatabase()

	app := fiber.New(configs.FibersConfig())

	app.Use(configs.InitCors())
	/** Auth **/
	mainRepositoryDB := repositorys.NewMainRepositoryDB(db)
	mainService := services.NewMainService(mainRepositoryDB)
	mainHandler := handlers.NewMainHandler(mainService)
	/** Rounts **/
	app.Static("uploads", "./uploads")
	/** Auth **/
	app.Post("api/login", mainHandler.Login)
	app.Get("api/refresh", utils.ValidateJWT(mainHandler.Refresh))
	app.Get("api/session", utils.ValidateJWT(mainHandler.Session))

	/** User **/
	app.Post("api/users", mainHandler.Register) // สร้างผู้ใช้ใหม่
	app.Get("api/users", utils.ValidateJWT(mainHandler.GetUsers))
	app.Get("api/users/:id", utils.ValidateJWT(mainHandler.GetUsers)) //  ดึงข้อมูลผู้ใช้ด้วย ID
	app.Put("api/users", utils.ValidateJWT(mainHandler.PutUsers))
	app.Put("api/users/:id", utils.ValidateJWT(mainHandler.PutUsers))       // แก้ไขข้อมูลผู้ใช้ด้วย ID
	app.Delete("api/users/:id", utils.ValidateJWT(mainHandler.DeleteUsers)) // ลบผู้ใช้ด้วย ID

	/** Product **/
	app.Post("api/products", utils.ValidateJWT(mainHandler.PostProducts))         // เพิ่มสินค้าใหม่
	app.Get("api/products", utils.ValidateJWT(mainHandler.GetProducts))           //  ดึงข้อมูลสินค้าด้วย ID
	app.Put("api/products/:id", utils.ValidateJWT(mainHandler.PutProducts))       // แก้ไขข้อมูลสินค้าด้วย ID
	app.Delete("api/products/:id", utils.ValidateJWT(mainHandler.DeleteProducts)) // ลบสินค้าด้วย ID

	/** Product Category **/
	app.Post("api/product/category", utils.ValidateJWT(mainHandler.PostCategory)) // เพิ่มหมวดหมู่สินค้าใหม่
	app.Get("api/product/category", utils.ValidateJWT(mainHandler.GetCategory))
	app.Get("api/product/category/:id", utils.ValidateJWT(mainHandler.GetCategory))       //  ดึงข้อมูลหมวดหมู่สินค้าด้วย ID
	app.Put("api/product/category/:id", utils.ValidateJWT(mainHandler.PutCategory))       // แก้ไขข้อมูลหมวดหมู่สินค้าด้วย ID
	app.Delete("api/product/category/:id", utils.ValidateJWT(mainHandler.DeleteCategory)) // ลบหมวดหมู่สินค้าด้วย ID

	/** Table **/
	app.Post("api/tables", utils.ValidateJWT(mainHandler.PostTables)) // เพิ่มโต๊ะใหม่
	app.Get("api/tables", utils.ValidateJWT(mainHandler.GetTables))
	app.Get("api/tables/:id", utils.ValidateJWT(mainHandler.GetTables))       //  ดึงข้อมูลโต๊ะด้วย ID
	app.Put("api/tables/:id", utils.ValidateJWT(mainHandler.PutTables))       // แก้ไขข้อมูลโต๊ะด้วย ID
	app.Delete("api/tables/:id", utils.ValidateJWT(mainHandler.DeleteTables)) // ลบโต๊ะด้วย ID

	/** Reservation **/
	app.Post("api/reservations", utils.ValidateJWT(mainHandler.PostReservations)) // จองโต๊ะ
	app.Get("api/reservations", utils.ValidateJWT(mainHandler.GetReservations))
	app.Get("api/reservations/:id", utils.ValidateJWT(mainHandler.GetReservations))       //  ดึงข้อมูลการจองด้วย ID
	app.Put("api/reservations/:id", utils.ValidateJWT(mainHandler.PutReservations))       // แก้ไขข้อมูลการจองด้วย ID
	app.Delete("api/reservations/:id", utils.ValidateJWT(mainHandler.DeleteReservations)) // ยกเลิกการจองด้วย ID

	/** Bill **/
	app.Post("api/bills", utils.ValidateJWT(mainHandler.PostBills)) // เปิดบิลใหม่
	app.Get("api/bills", utils.ValidateJWT(mainHandler.GetBills))
	app.Get("api/bills/:id", utils.ValidateJWT(mainHandler.GetBills))       //  ดึงข้อมูลบิลด้วย ID
	app.Put("api/bills/:id", utils.ValidateJWT(mainHandler.PutBills))       // แก้ไขข้อมูลบิลด้วย ID (เช่น เพิ่มรายการสั่งอาหาร)
	app.Delete("api/bills/:id", utils.ValidateJWT(mainHandler.DeleteBills)) //  ปิดบิลด้วย ID (ชำระเงิน)

	/** Order **/
	app.Post("api/orders", utils.ValidateJWT(mainHandler.PostOrderCycle))
	app.Get("api/orders", utils.ValidateJWT(mainHandler.GetOrdersCycle))
	app.Get("api/orders/:id", utils.ValidateJWT(mainHandler.GetOrdersCycle))
	app.Put("api/orders/:id", utils.ValidateJWT(mainHandler.PutOrdersCycle))
	app.Delete("api/orders/:id", utils.ValidateJWT(mainHandler.DeleteOrdersCycle))

	/** Order Iitems**/
	app.Post("api/order-items", utils.ValidateJWT(mainHandler.PostOrders))
	app.Get("api/order-items", utils.ValidateJWT(mainHandler.GetOrders))
	app.Get("api/order-items/:id", utils.ValidateJWT(mainHandler.GetOrders))
	app.Put("api/order-items/:id", utils.ValidateJWT(mainHandler.PutOrders))
	app.Delete("api/order-items/:id", utils.ValidateJWT(mainHandler.DeleteOrders))

	/** Membership **/
	app.Post("api/memberships", utils.ValidateJWT(mainHandler.PostMemberships)) // สมัครสมาชิกใหม่
	app.Get("api/memberships", utils.ValidateJWT(mainHandler.GetMemberships))
	app.Get("api/memberships/:id", utils.ValidateJWT(mainHandler.GetMemberships))       //  ดึงข้อมูลสมาชิกด้วย ID
	app.Put("api/memberships/:id", utils.ValidateJWT(mainHandler.PutMemberships))       // แก้ไขข้อมูลสมาชิกด้วย ID
	app.Delete("api/memberships/:id", utils.ValidateJWT(mainHandler.DeleteMemberships)) //  ยกเลิกสมาชิกด้วย ID

	/** Payment system / bill check **/
	app.Get("api/bills/check/:id", utils.ValidateJWT(mainHandler.GetBillsCheck))
	app.Post("api/bills/check/:id", utils.ValidateJWT(mainHandler.PostBillsCheck))
	app.Post("api/bills/close/:id", utils.ValidateJWT(mainHandler.PostBillsClose))

	/** Report **/
	app.Get("api/report/total-amount-income", utils.ValidateJWT(mainHandler.GetTotalAmountIncome))              // จำนวนรายรับทั้งหมด
	app.Get("api/report/product-category", utils.ValidateJWT(mainHandler.GetProductCategory))                   // จำนวนรายรับ แยกตาม Product Category
	app.Get("api/report/bill-category-summary", utils.ValidateJWT(mainHandler.GetBillCategorySummary))          // สรุปข้อมูลบิล แยกตาม Product Category
	app.Get("api/report/bill-summary", utils.ValidateJWT(mainHandler.GetBillSummary))                           // สรุปข้อมูลบิลทั้งหมด
	app.Get("api/report/customer-summary", utils.ValidateJWT(mainHandler.GetCustomerSummary))                   // สรุปข้อมูลลูกค้า
	app.Get("api/report/customer-age-group-summary", utils.ValidateJWT(mainHandler.GetCustomerAgeGroupSummary)) // สรุปข้อมูลลูกค้า แยกตาม ช่วงอายุ
	app.Get("api/report/customer-gender-summary", utils.ValidateJWT(mainHandler.GetCustomerGenderSummary))      // สรุปข้อมูลลูกค้า แยกตาม เพศ
	app.Get("api/report/repeat-customers", utils.ValidateJWT(mainHandler.GetRepeatCustomers))                   // รายชื่อลูกค้าที่เข้าซ้ำและจำนวนครั้งที่เข้าร้านทั้งหมดในช่วง 15 วัน, 1 เดือน, 3 เดือน, 6 เดือน, 12 เดือน
	app.Get("api/report/top-10-food", utils.ValidateJWT(mainHandler.GetTop10Food))                              //รายการอาหาร top 10 ของร้านในแต่ละเดือน

	// Start the server and listen on port 8000
	err := app.Listen(fmt.Sprintf(":%v", viper.GetInt("app_port")))
	if err != nil {
		panic(err)
	}
}
