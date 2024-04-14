package auth

type authResponse struct {
	ID    int64    `json:"id"`
	Email string `json:"email"`
}