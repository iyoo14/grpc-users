package model

type User struct {
	Name  string
	Age   int
	Email string
	Id    int
}

type ListUserRequest struct {
	Order     string
	Limit     int32
	OrderType int32
}
type ListUserResponse struct {
	UserList []User
}

type DetailUserRequest struct {
	ID int
}
