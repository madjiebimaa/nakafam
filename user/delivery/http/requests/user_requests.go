package requests

type UserRegisterOrLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterAsLeader struct {
	Token string `json:"token"`
}
