package requests

type UserRegisterOrLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToLeaderRole struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}
