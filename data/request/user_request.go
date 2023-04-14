package request

type CreateUserRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Password string `validate:"required,min=3,max=50" json:"password"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
}

type UpdateUserRequest struct {
	Id       int    `validate:"required" json:"id"`
	Username string `validate:"required,min=2,max=100" json:"username"`
	Password string `validate:"required,min=3,max=50" json:"password"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
}

type LoginRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Password string `validate:"required,min=3,max=50" json:"password"`
}
