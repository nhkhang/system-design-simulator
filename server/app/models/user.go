package models

type UserModel struct {
	BaseModel
	ID string `json:"id" gorm:"column:id; primaryKey; not null"`
}
