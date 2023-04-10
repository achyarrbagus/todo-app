package models

import "time"

type Todo struct {
	ID        int               `json:"id"`
	Title     string            `json:"title"`
	UserID    int               `json:"-"`
	User      UsersTodoResponse `json:"users"`
	Body      string            `json:"body"`
	StartDate string            `json:"starDate"`
	EndDate   string            `json:"endDate"`
	CreatedAt time.Time         `json:"-"`
	UpdateAt  time.Time         `json:"-"`
}
