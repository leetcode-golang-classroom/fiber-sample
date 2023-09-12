package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/leetcode-golang-classroom/fiber-sample/book"
	"github.com/leetcode-golang-classroom/fiber-sample/database"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully open")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/books/:id", book.GetSingleBook)
	app.Post("/api/v1/books", book.CreateBook)
	app.Delete("/api/v1/books/:id", book.DeleteBook)
}
func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)
	app.Listen(3000)
}
