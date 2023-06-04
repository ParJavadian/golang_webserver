// package main
// import "fmt"
// func main() {
//     fmt.Println("hello world")
// }

package main;

import "time"

type User struct {
  name string
  family string
  id int
  age int
  sex string
  createdAt time.Time
}
