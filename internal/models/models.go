package models

type Register struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Verify struct {
	Email      string `json:"email"`
	SecretCode string `json:"secretcode"`
}

type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OriginCreate struct {
	Origin string `json:"origin"`
}

type OriginGet struct {
	Id     string    `json:"id"`
	Origin string `json:"origin"`
}
