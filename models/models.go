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

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
}
