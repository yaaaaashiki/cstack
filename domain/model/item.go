package model

type Item struct {
	ID                  uint   `json:"id" gorm:"column:id"`
	UserID              uint   `json:"user_id" gorm:"column:user_id"`
	Name                string `json:"name" gorm:"column:name"`
	Price               int `json:"price" gorm:"column:price"`
	CurrentPaymentPrice int `json:"current_payment_price" gorm:"column:current_payment_price"`
	IconImage           string `json:"icon_image" gorm:"column:icon_image"`
	Description         string `json:"description" gorm:"column:description"`
}
