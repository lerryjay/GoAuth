package models

type Login struct {
	LoginId  string
	Password string
}

type VerifyToken struct {
	LoginId string
	Token   int
}

type ResetPassword struct {
	LoginId  string
	Password string
	Token    int
}

type ChangePassword struct {
	OldPassword string
	NewPassword string
}
