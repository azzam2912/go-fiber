package main

import (
	"fmt"
	"github.com/azzam2912/go-fiber/database"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App){
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}