package entity

type Users struct {
	// gorm.Model
	ID       uint
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
