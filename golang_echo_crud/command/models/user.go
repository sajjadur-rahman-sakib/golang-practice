package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}
