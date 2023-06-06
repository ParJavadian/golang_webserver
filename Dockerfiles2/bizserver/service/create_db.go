package service

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
)

const (
	host2     = "localhost"
	port2     = 5432
	user2     = "postgres"
	password2 = "12345"
	dbname2   = "biz_database4"
)

func CreateServer() *fiber.App {
	app := fiber.New()

	return app
}

func main() {

	//app := CreateServer()
	//
	//app.Use(cors.New())
	//
	//app.Get("/hello", func(c *fiber.Ctx) error {
	//	create_database()
	//	return c.SendString("Hello, World!")
	//})
	//
	//// 404 Handler
	//app.Use(func(c *fiber.Ctx) error {
	//	return c.SendStatus(404) // => 404 "Not Found"
	//})
	//
	//log.Fatal(app.Listen(":3000"))

	create_database()
}

func create_database() {

	conninfo := "user=postgres password=12345 host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		log.Fatal(err)
	}
	//dbName := "biz_database3"
	_, err = db.Exec("create database " + dbname2)
	if err != nil {
		//handle the error
		log.Fatal(err)
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host2, port2, user2, password2, dbname2)

	// open database
	db, err = sql.Open("postgres", psqlconn)
	CheckError2(err)

	//todo:code below creates a table
	_, err = db.Query(`CREATE TABLE "biz_server3" (name TEXT, family TEXT,sex TEXT,age bigint,createdat bigint,id bigint);`)
	CheckError2(err)

	defer db.Close()
}

func CheckError2(err error) {
	if err != nil {
		panic(err)
	}
}
