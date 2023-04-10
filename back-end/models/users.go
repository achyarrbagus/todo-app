package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" gorm:"type: varchar(255);unique;not null"`
	Password string `json:"password"`
	Todo     []Todo `json:"todos"`
}

type UsersTodoResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (UsersTodoResponse) TableName() string {
	return "users"
}
