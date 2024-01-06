package routes

import (
	"encoding/json"
	"time"

	"golang.org/x/crypto/bcrypt"

	orm "backend/orm"
	"backend/platform/database"
)

func HandleLoginGet(resp string) ([]byte, error) {
	db := database.OpenDb()
	var user orm.User
	db.
		Where("full_name LIKE ?", "%"+resp+"%").
		First(&user)
	return json.Marshal(user)
}

func HandleLoginPost(userdata map[string]string) ([]byte, error) {
	db := database.OpenDb()
	var user orm.User
	db.
		Where("user_name LIKE ?", "%"+userdata["username"]+"%").
		First(&user)
	if user.ID != 0 && !user.Blocked {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userdata["password"])); err == nil {
			deltaChange := time.Since(user.CreatedAt)
			if user.UpdatedAt.IsZero() && deltaChange.Hours() < 365*24 {
				user.LastLogin = time.Now()
				user.Attempt = 0
				db.Save(&user)
				result := map[string]interface{}{
					"message":       "Authenticated",
					"access_token":  "createAccessToken(user.UserName)",
					"refresh_token": "createRefreshToken(user.UserName)",
				}
				return json.Marshal(result)
			}
			result := map[string]interface{}{
				"message": "Expired",
			}
			return json.Marshal(result)
		} else {
			if user.Attempt < 9 {
				user.Attempt++
			} else {
				user.Blocked = true
			}
			db.Save(&user)
		}
	}
	result := map[string]interface{}{
		"message": "Denied",
	}
	return json.Marshal(result)
}

func HandleLoginPatch(userdata map[string]string) ([]byte, error) {
	db := database.OpenDb()
	var user orm.User
	db.
		Where("user_name LIKE ?", "%"+userdata["username"]+"%").
		First(&user)
	if user.ID != 0 && !user.Blocked {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userdata["password"])); err == nil {
			user.Password, err = bcrypt.GenerateFromPassword([]byte(userdata["new_pswd"]), bcrypt.DefaultCost)
			if err != nil {
				result := map[string]interface{}{
					"message": "Denied",
				}
				return json.Marshal(result)
			}
			db.Save(&user)
			result := map[string]interface{}{
				"message": "Authenticated",
			}
			return json.Marshal(result)
		}
	}
	result := map[string]interface{}{
		"message": "Denied",
	}
	return json.Marshal(result)
}

func HandleLoginDelete(resp string) ([]byte, error) {
	db := database.OpenDb()
	var user []orm.User
	db.
		Where("full_name LIKE ?", "%"+resp+"%").
		First(&user)
	return json.Marshal(user)
}
