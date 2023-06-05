// // import (
// //    "database/sql" // add this
// //    "fmt"
// //    "log"
// //    "os"
// //    "github.com/lib/pq" // add this   "- "
// //
// //    "github.com/gofiber/fiber/v2"
// // )
// // http://localhost/
// //  connStr := "postgresql://<postgres>:<how you doin>@<biz_server_web>/todos?sslmode=disable
// // "
// //    // Connect to database
// //    db, err := sql.Open("postgres", connStr)
// //    if err != nil {
// //        log.Fatal(err)
// //    }
// //
// //
// //  func main() {
// //    app := fiber.New()
// //
// //    app.Get("/", indexHandler) // Add this
// //
// //    app.Post("/", postHandler) // Add this
// //
// //    app.Put("/update", putHandler) // Add this
// //
// //    app.Delete("/delete", deleteHandler) // Add this
// //
// //    port := os.Getenv("PORT")
// //    if port == "" {
// //        port = "3000"
// //    }
// //    log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
// // }
// //
// //
// // func indexHandler(c *fiber.Ctx) error {
// //    return c.SendString("Hello")
// // }
// // func postHandler(c *fiber.Ctx) error {
// //    return c.SendString("Hello")
// // }
// // func putHandler(c *fiber.Ctx) error {
// //    return c.SendString("Hello")
// // }
// // func deleteHandler(c *fiber.Ctx) error {
// //    return c.SendString("Hello")
// // }
//
//
// package main
//
// import (
//    "fmt"
//    "log"
//    "os"
//
//    "github.com/gofiber/fiber/v2"
// )
//
// func main() {
//    app := fiber.New()
//    port := os.Getenv("PORT")
//    if port == "" {
//        port = "3000"
//    }
//    log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
// }
//
//  func main() {
//    app := fiber.New()
//
//    app.Get("/", indexHandler) // Add this
//
//    app.Post("/", postHandler) // Add this
//
//    app.Put("/update", putHandler) // Add this
//
//    app.Delete("/delete", deleteHandler) // Add this
//
//    port := os.Getenv("PORT")
//    if port == "" {
//        port = "3000"
//    }
//    log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
// }
//
//
// func indexHandler(c *fiber.Ctx) error {
//    return c.SendString("Hello")
// }
// func postHandler(c *fiber.Ctx) error {
//    return c.SendString("Hello")
// }
// func putHandler(c *fiber.Ctx) error {
//    return c.SendString("Hello")
// }
// func deleteHandler(c *fiber.Ctx) error {
//    return c.SendString("Hello")
// }
//


package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "how you doin"
  dbname   = "biz_server_web"
)


func main() {
  psqlInfo := fmt.Sprintf("host=localhost port=55432 user=postgres "+
    "password=how you doin dbname=biz_server_web sslmode=disable",
    host, port, user, password, dbname)
}