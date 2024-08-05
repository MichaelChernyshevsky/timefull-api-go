package user

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	// cred
	Email    string `gorm:"size:250"`
	Password string `gorm:"size:250"`
	Phone    string `gorm:"size:250"`
	// PackagesId int
	Subscribed bool

	// info
	Sex     string `gorm:"size:250"`
	Name    string `gorm:"size:250"`
	Name2   string `gorm:"size:250"`
	Age     int
	Admin   bool
	Creator bool
}

func (u *UserModel) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":       u.ID,
		"email":    u.Email,
		"password": u.Password,
		"phone":    u.Phone,
		// "packagesId": u.Packages,
		"sex":        u.Sex,
		"name":       u.Name,
		"name2":      u.Name2,
		"age":        u.Age,
		"admin":      u.Admin,
		"creator":    u.Creator,
		"subscribed": u.Subscribed,
	}
}

// Методы поиска пользователя
func FindUserModelByEmail(db *gorm.DB, email string) (*UserModel, error) {
	var UserModel UserModel
	result := db.Where("email = ?", email).First(&UserModel)
	return &UserModel, result.Error
}

func FindUserModelByID(db *gorm.DB, id uint) (*UserModel, error) {
	var UserModel UserModel
	result := db.First(&UserModel, id)
	return &UserModel, result.Error
}
