package model

type User struct {
	ID        uint   `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`
	Email     string `json:"email" gorm:"column:email"`
	Salt      string `json:"salt" gorm:"column:salt"`
	Salted    string `json:"salted" gorm:"column:salted"`
	IconImage string `json:"icon_image"`
}
