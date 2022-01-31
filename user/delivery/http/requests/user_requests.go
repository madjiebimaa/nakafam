package requests

type UserRegisterOrLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserToLeaderRole struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}
