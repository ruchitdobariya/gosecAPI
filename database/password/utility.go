package password

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ruchitdobariya/gosecAPI/database"
	"github.com/ruchitdobariya/gosecAPI/database/config"
	"github.com/ruchitdobariya/gosecAPI/database/user"
	"github.com/ruchitdobariya/gosecAPI/myencode"
	"github.com/ruchitdobariya/gosecAPI/settings"
)

// Returns all passwords
func GetPasswordObjects() []Password {
	var passwords []Password
	path := database.GetPath() + "password.json"

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("The database could not be found. If you have executed the 'gosec' program and the database did not get created, please provide feedback on github.com/ruchitdobariya/GoSec.")
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &passwords)

	return passwords
}

func FindPasswordByUserId(id string, passwords []Password) ([]Password, *database.MyError) {
	userId, convertionErr := database.ConverToFloat64(id)

	if convertionErr != nil {
		return nil, convertionErr
	}

	
	validIdErr := user.CheckValidUserId(userId)
	if validIdErr != nil {
		return nil, validIdErr
	}

	var allPasswords []Password

	for _, u := range passwords {
		if u.UserId == userId {
			allPasswords = append(allPasswords, u)
		}
	}

	if len(allPasswords) == 0 {
		return nil, GetPasswordError(1)
	}

	return DecodePasswords(id, allPasswords), nil
}

func DecodePasswords(userId string, passwords []Password) []Password {
	
	config, _ := config.GetConfigByUserId(userId)

	userDecryptedSecret, _ := myencode.Decrypt(settings.GetSecretForSecrets(), config.Secret)

	var newPasswords []Password

	
	for _, p := range passwords {
		decrpytedPassword, _ := myencode.Decrypt([]byte(userDecryptedSecret), p.Password)
		newPasswords = append(newPasswords, Password{p.PasswordId, p.UserId, p.Title, p.Url, decrpytedPassword})
	}

	return newPasswords
}
