package user

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ruchitdobariya/gosecAPI/database"
)

func GetUsersObject() []User {

	var users []User
	path := database.GetPath() + "users.json"

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("The database could not be found. If you have executed the 'gosec' program and the database did not get created, please provide feedback on github.com/ruchitdobariya/GoSec.")
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &users)

	return users
}

func FindUserById(id string, users []User) (*User, *database.MyError) {
	userId, err := database.ConverToFloat64(id)

	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if u.UserId == userId {
			return &u, nil
		}
	}

	return &User{UserId: -1}, nil
}

// Checks if the entered id is valid or not.
func CheckValidUserId(id float64) *database.MyError {
	users := GetAllUsers()

	check := false
	for _, u := range users {
		if u.UserId == id {
			check = true
			break
		} else {
			check = false
		}
	}

	if check {
		return nil
	}

	return GetUserError(1)
}
