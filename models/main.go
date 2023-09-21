package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uint
	UserName  string `gorm:"type:varchar(100);default:null" json:"UserName"`
	PassWord  string `gorm:"type:varchar(255);default:null"  json:"PassWord"`
	FirstName string
	LastName  string
	Role      string
}

type Product struct {
	gorm.Model
	ID          uint
	Name        string
	Description string
	Image       string
	ImageHash   string
	Price       float64
	CategoryID  int
	UserID      int
}

type ProductCategory struct {
	gorm.Model
	ID           uint
	CategoryName string
}

type TableStatus string

const (
	StatusEmpty    TableStatus = "ว่าง"
	StatusReserved TableStatus = "จอง"
	StatusOpen     TableStatus = "เปิด"
	StatusNotInUse TableStatus = "ไม่ใช้งาน"
)

type Table struct {
	gorm.Model
	ID     uint
	Number string
	Status TableStatus
}

type Reservation struct {
	gorm.Model
	ID           uint
	TableID      int
	CustomerName string
}

type BillStatus string

const (
	StatusCloseBill BillStatus = "ปิด"
	StatusOpenBill  BillStatus = "เปิด"
)

type GenderStatus string

const (
	Male   GenderStatus = "ชาย"
	Female GenderStatus = "หญิง"
)

type Bill struct {
	gorm.Model
	ID            uint
	TableID       int
	Number        int
	Gender        GenderStatus
	AgeGroupStart int
	AgeGroupEnd   int
	AmountPaid    float64
	Change        float64
	Status        BillStatus
}

type OrderStatus string

const (
	StatusPending   OrderStatus = "รอยืนยัน"
	StatusPreparing OrderStatus = "กำลังเตรียมอาหาร"
	StatusServing   OrderStatus = "เสริฟอาหาร"
	StatusCancelled OrderStatus = "ยกเลิก"
)

type Order struct {
	gorm.Model
	ID        uint
	BillID    int
	ProductID int
	Status    OrderStatus
}

type Membership struct {
	gorm.Model
	ID                 uint
	UserID             int
	CardNumber         string
	ExpiryDate         string `gorm:"type:date" `
	DiscountPercentage float64
}
