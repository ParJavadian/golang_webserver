package main

//import (
//	"database/sql"
//	"fmt"
//)

//import "os"
//
//func main() {
//	a := App{}
//	a.Initialize(
//		os.Getenv("APP_DB_USERNAME"),
//		os.Getenv("APP_DB_PASSWORD"),
//		os.Getenv("APP_DB_NAME"),
//	)
//	a.Run(":8001")
//}

//package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "howyoudoin"
	dbname   = "biz_database"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	insertStmt := `insert into "biz_table" ("name", "family","age", "sex","createdAt", "id") values ("parmida","javadian", 1, "female", 5000, 5)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// check db
	//err = db.Ping()
	//CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//// // import (
//// //    "database/sql" // add this
//// //    "fmt"
//// //    "log"
//// //    "os"
//// //    "github.com/lib/pq" // add this   "- "
//// //
//// //    "github.com/gofiber/fiber/v2"
//// // )
//// // http://localhost/
//// //  connStr := "postgresql://<postgres>:<how you doin>@<biz_server_web>/todos?sslmode=disable
//// // "
//// //    // Connect to database
//// //    db, err := sql.Open("postgres", connStr)
//// //    if err != nil {
//// //        log.Fatal(err)
//// //    }
//// //
//// //
//// //  func main() {
//// //    app := fiber.New()
//// //
//// //    app.Get("/", indexHandler) // Add this
//// //
//// //    app.Post("/", postHandler) // Add this
//// //
//// //    app.Put("/update", putHandler) // Add this
//// //
//// //    app.Delete("/delete", deleteHandler) // Add this
//// //
//// //    port := os.Getenv("PORT")
//// //    if port == "" {
//// //        port = "3000"
//// //    }
//// //    log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
//// // }
//// //
//// //
//// // func indexHandler(c *fiber.Ctx) error {
//// //    return c.SendString("Hello")
//// // }
//// // func postHandler(c *fiber.Ctx) error {
//// //    return c.SendString("Hello")
//// // }
//// // func putHandler(c *fiber.Ctx) error {
//// //    return c.SendString("Hello")
//// // }
//// // func deleteHandler(c *fiber.Ctx) error {
//// //    return c.SendString("Hello")
//// // }
////
////
//// package main
////
//// import (
////    "fmt"
////    "log"
////    "os"
////
////    "github.com/gofiber/fiber/v2"
//// )
////
//// func main() {
////    app := fiber.New()
////    port := os.Getenv("PORT")
////    if port == "" {
////        port = "3000"
////    }
////    log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
//// }
////
////  func main() {
////    app := fiber.New()
////
////    app.Get("/", indexHandler) // Add this
////
////    app.Post("/", postHandler) // Add this
////
////    app.Put("/update", putHandler) // Add this
////
////    app.Delete("/delete", deleteHandler) // Add this
////
////    port := os.Getenv("PORT")
////    if port == "" {
////        port = "3000"
////    }
////    log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
//// }
////
////
//// func indexHandler(c *fiber.Ctx) error {
////    return c.SendString("Hello")
//// }
//// func postHandler(c *fiber.Ctx) error {
////    return c.SendString("Hello")
//// }
//// func putHandler(c *fiber.Ctx) error {
////    return c.SendString("Hello")
//// }
//// func deleteHandler(c *fiber.Ctx) error {
////    return c.SendString("Hello")
//// }
////
//
//package service
//
//import (
//	_ "database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//)
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "how you doin"
//	dbname   = "biz_server_web"
//)
//
//func main() {
//	psqlInfo := fmt.Sprintf("host=localhost port=55432 user=postgres "+
//		"password=how you doin dbname=biz_server_web sslmode=disable",
//		host, port, user, password, dbname)
//}
