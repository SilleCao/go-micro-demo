package model

type Credentials struct {
	Username string `form:"username" json:"username" example:"admin"`
	Password string `form:"password" json:"password" example:"111000"`
}
