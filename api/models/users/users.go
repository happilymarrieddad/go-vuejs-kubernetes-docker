package users

// Users - slice of User
type Users []User

// User - a model of user from the DB
type User struct {
	ID       int64  `xorm:"id"`
	Email    string `xorm:"email"`
	Password string `xorm:"password"`
}

// TableName - sets the correct table name in the DB
func (u *User) TableName() string {
	return "users"
}
