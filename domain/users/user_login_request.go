package users

//LoginRequest truct
type LoginRequest struct {
	Email    string `json:"auth_id"`
	Password string `json:"auth_secret"`
}
