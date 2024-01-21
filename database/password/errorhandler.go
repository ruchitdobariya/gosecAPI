package password

import "github.com/ruchitdobariya/gosecAPI/database"

func GetPasswordError(errorId float64) *database.MyError {
	switch errorId {
	case 1:
		return database.MyErrorInit(errorId, "Password Error", "This user does not have a registered password.")
	}

	return nil
}
