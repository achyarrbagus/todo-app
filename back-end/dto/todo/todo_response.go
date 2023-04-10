package tododto

import "time"

type TodoResponse struct {
	Title     string    `json:"title" gorm:"type:varchar(225)"`
	Body      string    `json:"body" gorm:"type:varchar(225)"`
	StartDate string    `json:"startDate" gorm:"type:varchar(225)"`
	EndDate   string    `json:"endDate" gorm:"type:varchar(225)"`
	CreatedAt time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}
