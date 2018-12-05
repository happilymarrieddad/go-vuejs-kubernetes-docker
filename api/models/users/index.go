package users

import "github.com/go-xorm/xorm"

// Index - gets a slice of users based on the passed parametes
func Index(db *xorm.Engine, findBy *User) (users Users, err error) {

	if findBy == nil {
		findBy = &User{}
	}

	sess := db.Where("1 = 1")

	if findBy.ID > 0 {
		sess = sess.Where("id = ?", findBy.ID)
	}

	if len(findBy.Email) > 0 {
		sess = sess.Where("email = ?", findBy.Email)
	}

	if err = sess.Find(&users); err != nil {
		return
	}

	return
}
