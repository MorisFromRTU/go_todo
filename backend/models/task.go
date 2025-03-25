package models

type Task struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" validate:"required"`
	Done   bool   `json:"done"`
	UserID uint   `json:"user_id"`
}

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"username" validate:"required" gorm:"unique;not null"`
	Password string `json:"-" validate:"required" gorm:"not null"`
	Tasks    []Task `json:"tasks" form:"foreignKey:UserID"`
}
