package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDataVerifier interface {
	Verify() User
}

type LoginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (loginData LoginData) Verify() bool {
	var user User

	if err := DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		return false
	}

	return true
}
