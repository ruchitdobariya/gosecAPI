package user

import (
	"github.com/ruchitdobariya/gosecAPI/database"
)

type User struct {
	UserId   float64 `json:"userId"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}

// Returns All Users
func GetAllUsers() []User {
	

	users := GetUsersObject()
	return users
}

// Finds User By Id
func GetUsersById(userId string) (*User, *database.MyError) {
	users := GetUsersObject()
	theUser, err := FindUserById(userId, users)

	
	if theUser != nil {
		if theUser.UserId != -1 {
			return theUser, nil
		} else {
			
			return nil, GetUserError(1)
		}
	}

	return nil, err
}
