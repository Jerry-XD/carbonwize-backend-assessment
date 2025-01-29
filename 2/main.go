package main

import (
	"time"

	"carbonwize-backend-assessment/domain"
	"carbonwize-backend-assessment/handlers"
	"carbonwize-backend-assessment/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	var dsn = "host= user= password= dbname= port= sslmode="
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic(err)
	}

	var autoMigrate = []interface{}{
		&domain.EmissionFactors{},
	}
	err = db.AutoMigrate(autoMigrate...)
	if err != nil {
		panic(err)
	}

	var calculateRepo = repository.NewCalculateRepo(db)
	var calculateHandler = handlers.NewCalculateHandler(calculateRepo)

	var app = fiber.New()

	app.Post("/api/carbon/footprint/calculate", calculateHandler.Calculate)

	err = app.Listen("localhost:8080")
	if err != nil {
		panic(err)
	}
}
