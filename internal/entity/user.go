package entity

type User struct {
	ID        int64  `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Password  string `json:"password" db:"password"`
}
