package login

// Output ...
type Output struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}

// Input ...
type Input struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
