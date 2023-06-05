package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"
	_ "github.com/lib/pq"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	host2     = "localhost"
	port2     = 5432
	user2     = "postgres"
	password2 = "howyoudoin"
	dbname2   = "biz_database"
)

func CreateServer() *fiber.App {
	app := fiber.New()

	return app
}

func main() {

	app := CreateServer()

	app.Use(cors.New())

	app.Get("/hello", func(c *fiber.Ctx) {
		return c.SendString("Hello, World!")
	}

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":5432"))

	create_database()
}

func create_database() {

	conninfo := "user=postgres password=howyoudoin host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		log.Fatal(err)
	}
	dbName := "biz_database"
	_, err = db.Exec("create database " + dbName)
	if err != nil {
		//handle the error
		log.Fatal(err)

		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host2, port2, user2, password2, dbname2)

		// open database
		db, err := sql.Open("postgres", psqlconn)
		CheckError2(err)

		//todo:code below creates a table
		_, err = db.Query(`CREATE TABLE "biz_server" (name TEXT, family TEXT,sex TEXT,age bigint,createdat bigint,id bigint);`)
		CheckError2(err)

		defer db.Close()
	}
}

func CheckError2(err error) {
	if err != nil {
		panic(err)
	}
}
