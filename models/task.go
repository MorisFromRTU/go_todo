package models

type Task struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" validate:"required"`
	Done  bool   `json:"done"`
}
