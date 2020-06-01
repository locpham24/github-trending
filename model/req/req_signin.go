package req

type ReqSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ReqUpdateProfile struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
