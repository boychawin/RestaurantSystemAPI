package configs

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"

	"restaurant/gorms"
	"restaurant/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func InitTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		viper.GetString("db_host"),
		viper.GetInt("db_port"),
		viper.GetString("db_username"),
		viper.GetString("db_password"),
		viper.GetString("db_database"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &gorms.SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection established successfully")

	db.AutoMigrate(models.Users{}, models.Table{}, models.Reservation{}, models.ProductCategory{}, models.Product{},
		models.Order{}, models.Membership{}, models.Bill{})
	return db
}

func InitCors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type,Authorization,Token",
		AllowCredentials: true,
	})
}

func FibersConfig() fiber.Config {
	return fiber.Config{
		BodyLimit: 200 * 1024 * 1024,
	}
}
