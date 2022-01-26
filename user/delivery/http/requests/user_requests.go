package requests

type UserRegisterOrLogin struct {
	Email    string `json:"email" faker:"email"`
	Password string `json:"password" faker:"password"`
}

type UserToLeaderRole struct {
	Token    string `json:"token" faker:"uuid_digit"`
	Password string `json:"password" faker:"password"`
}
