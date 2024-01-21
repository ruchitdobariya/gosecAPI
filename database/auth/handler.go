package auth

import (
	"github.com/ruchitdobariya/gosecAPI/database"
	"github.com/ruchitdobariya/gosecAPI/database/user"
)

func Check(username string, password string) (*user.User, *database.MyError) {
	
	users := user.GetAllUsers()

	for _, u := range users {
		if u.Username == username && u.Password == database.ConvertToMd5(password) {
			return &u, nil
		}
	}

	return nil, GetAuthError(1)
}
