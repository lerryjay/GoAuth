package models

type ResponseError struct {
	Description string
	Errors      any
}

type ResponseObject struct {
	Message string
	Status  int
	Data    any
	Error   any
}
