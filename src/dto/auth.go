package dto

type RegisterInput struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"required,email"`
	Password string ` json:"password" form:"password" binding:"required" validate:"required"`
	Username string `json:"username" form:"username" binding:"required"  validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" form:"password" binding:"required" validate:"required"`
}

type LoginOutput struct {
	AccessToken string `json:"access_token"`
}

type LoginGoogleInput struct {
	Code string `json:"code" form:"code" binding:"required"`
}
