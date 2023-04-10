package userdto

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Token    string `json:"token" validate:"required"`
}
