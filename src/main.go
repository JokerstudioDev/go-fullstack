package main

import (
	"dac"
	"fmt"
	"service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	userDac := dac.NewUserDac("root:ploy@tcp(127.0.0.1:3306)/testsck")
	userSvc := service.NewUserSvc(userDac)
	user := userSvc.Users()
	fmt.Println(user)
}
