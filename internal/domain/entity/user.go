package entity

type User struct {
	Id       int64
	Username string
	Name     string
}

type UserOpt struct {
	Username string
	Password string
}
