package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type Reset struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPassword struct {
	NewPassword     string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type MailData struct {
	To      string
	From    string
	Subject string
	Content string
}

type AppConfig struct {
	MailChan chan MailData
}
