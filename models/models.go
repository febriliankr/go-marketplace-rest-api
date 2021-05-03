package models

type Seller struct {
	Id    string `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	DisplayName string `json:"display_name" db:"display_name"`
	Email string `json:"email" db:"email"`
	DisplayPictureUrl string `json:"display_picture_url" db:"display_picture_url"`
	Address string `json:"address" db:"address"`
}

type User struct {
	Id    string `json:"id"`
	Username  string `json:"username"`
	Email string `json:"email"`
	DisplayPictureUrl string `json:"display_picture_url"`
}

type RegisterUser struct {
	Username  string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UsernameAndPassword struct {
	Username  string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type EmailAndPassword struct {
	Email string `json:"email" db:"email" form:"email"`
	Password string `json:"password" db:"password" form:"password"`
}

type Email struct {
	Username string `json:"email" db:"email"`
}

type Username struct {
	Username string `json:"username" db:"username"`
}