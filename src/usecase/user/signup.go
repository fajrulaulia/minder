package User

type UserRequest struct {
	Username string `json:"username"`
	Emal     string `json:"email"`
	Password string `json:"password"`
}
