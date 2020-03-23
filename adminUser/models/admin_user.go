package models

type AdminUserReqBody struct {
	Nickname string
	Email    string
	Password string
	IsAdmin  bool
	Platform int
}
