package dac

import (
	"database/sql"
	"fmt"
	"model"

	_ "github.com/go-sql-driver/mysql"
)

type IUserDac interface {
	Read() []model.User
	Add() bool
	Remove(id string) bool
	Update(id string, firstnameFather string) bool
	ReadById(citizenId string) model.User
}

type UserDac struct{}

var conString string

func NewUserDac(con string) *UserDac {
	conString = con
	return &UserDac{}
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", conString)

	if err != nil {
		fmt.Println("connect fail")
	}
	fmt.Println("connect success")
	return db
}

func (*UserDac) Read() []model.User {
	db := openDB()
	defer db.Close()
	results, _ := db.Query("SELECT * FROM user")
	var userList []model.User
	for results.Next() {
		var user model.User
		err := results.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.CitizenID, &user.BirthYear, &user.FirstnameFather, &user.LastnameFather, &user.FirstnameMother, &user.LastnameMother, &user.SoldierID, &user.AddressID)
		if err != nil {
			return []model.User{}
		}
		userList = append(userList, user)
	}
	return userList
}

func (*UserDac) Add() bool {
	db := openDB()
	defer db.Close()
	results, _ := db.Prepare(`Insert INTO user
	(citizen_id,firstname,lastname,birthyear
		,firstname_father,lastname_father,firstname_mother,lastname_mother,soldier_id
		,address_id)
		VALUES(?,?,?,?,?,?,?,?,?,?)`) //เป็นการเตรียมค่าเพื่อจะทำงาน

	_, err := results.Exec("1092018019065", "พอยพอย", "ซันชาย", 1995, "ตะวัน", "ซันชาย", "สมิตา", "ซันชาย", 988, 1)
	if err != nil {
		return false
	}
	return true
}

func (*UserDac) Remove(id string) bool {
	db := openDB()
	defer db.Close()
	statement, _ := db.Prepare("DELETE FROM testsck.user WHERE user_id=?")
	defer statement.Close()
	_, err := statement.Exec(id)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func (*UserDac) Update(id string, firstnameFather string) bool {
	db := openDB()
	defer db.Close()
	statement, _ := db.Prepare("UPDATE `user` SET  firstname_father= ? WHERE user_id=?")
	defer statement.Close()
	_, err := statement.Exec(firstnameFather, id)
	if err != nil {
		return false
	}
	return true
}

func (*UserDac) ReadById(citizenId string) model.User {
	db, err := sql.Open("mysql", "root:ploy@tcp(127.0.0.1:3306)/testsck")

	if err != nil {
		fmt.Println("connect fail")
	}
	fmt.Println("connect success")
	defer db.Close()
	results, _ := db.Query("SELECT * FROM user WHERE citizen_id=?", citizenId)
	var user model.User
	for results.Next() {
		err := results.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.CitizenID, &user.BirthYear, &user.FirstnameFather, &user.LastnameFather, &user.FirstnameMother, &user.LastnameMother, &user.SoldierID, &user.AddressID)
		if err != nil {
			return model.User{}
		}
	}
	return user
}
