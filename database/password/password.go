package password

import "github.com/ruchitdobariya/gosecAPI/database"

type Password struct {
	PasswordId float64 `json:"passwordId"`
	UserId     float64 `json:"userId"`
	Title      string  `json:"title"`
	Url        string  `json:"url"`
	Password   string  `json:"password"`
}


func GetPasswordsByUserId(userId string) ([]Password, *database.MyError) {
	passwords := GetPasswordObjects()
	thePasswords, err := FindPasswordByUserId(userId, passwords)

	if thePasswords != nil {
		return thePasswords, nil
	}

	return nil, err
}
