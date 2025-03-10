package repository

type User struct {
	ID    interface{} `db:"id"`
	Name  interface{} `db:"name"`
	Email interface{} `db:"email"`
	Age   interface{} `db:"age"`
}

type UserListEntity struct {
	UserList []User
}
