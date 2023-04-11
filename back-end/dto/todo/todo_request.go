package tododto

type CreateTodoRequest struct {
	Title     string `json:"title" gorm:"type:varchar(225)"`
	Body      string `json:"body" gorm:"type:varchar(225)"`
	StartDate string `json:"startDate" gorm:"type:varchar(225)"`
	EndDate   string `json:"endDate" gorm:"type:varchar(225)"`
}

type UpdateTodoRequest struct {
	Title     string `json:"title" gorm:"type:varchar(225)"`
	Body      string `json:"body" gorm:"type:varchar(225)"`
	StartDate string `json:"startDate" gorm:"type:varchar(225)"`
	EndDate   string `json:"endDate" gorm:"type:varchar(225)"`
}
