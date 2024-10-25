package user

type User struct {
	Login    string `json:"login"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdated struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
