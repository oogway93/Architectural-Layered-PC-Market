package user

type UserCreate struct {
	ID       string `'json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}