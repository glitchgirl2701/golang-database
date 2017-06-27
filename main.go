package main

import (
  "fmt"

  //"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
  host = "localhost"
  port = 5432
  user = "kaylathomsen"
  password = "broadway"
  dbname = "usegolang_dev"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  us, err := models.NewUserService(psqlInfo)
    if err != nil {
      panic(err)
    }
  defer us.Close()
  us.DestructiveReset()
}

//   user, err := us.ByID(2)
//   if err != nil {
//     panic(err)
//   }
//   fmt.Println(user)
// }
