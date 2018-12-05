package jwt

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-xorm/xorm"
	Users "github.com/learning/project/api/models/users"
)

var signKey string

func init() {
	signKey = os.Getenv("SECRET")
	if len(signKey) == 0 {
		log.Println("WARNING!! Secret not set for JWT.")
	}
}

// CreateToken - this creates a jwt token from a user id
func CreateToken(id int64) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = id

	token.Claims = claims

	tokenString, err = token.SignedString([]byte(signKey))

	return
}

// ParseToken - parsing token
func ParseToken(val string) (id int64, err error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})

	switch err.(type) {
	case nil:
		if !token.Valid {
			return 0, errors.New("Token is invalid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, errors.New("Token is invalid")
		}

		id = int64(claims["id"].(float64))

		return
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return 0, errors.New("token has expired")
		default:
			log.Println(vErr)
			return 0, errors.New("error parsing token or token does not exist in request")
		}
	default:
		return 0, errors.New("Unable to parse token")
	}
}

// GetUserFromToken - grab user data from the database via the token data
func GetUserFromToken(db *xorm.Engine, tokenVal string) (user Users.User, err error) {
	if len(tokenVal) == 0 {
		err = errors.New("No token present")
		return
	}

	userID, err := ParseToken(tokenVal)
	if err != nil {
		log.Println("ParseToken error")
		log.Println(err)
		err = errors.New("token is invalid")
		return
	}

	if userID < 1 {
		err = errors.New("token is missing required data")
		return
	}

	user.ID = userID
	users, err := Users.Index(db, &user)
	if err != nil || len(users) == 0 {
		log.Println(err)
		err = errors.New("Unable to get user from token")
		return
	}

	user = users[0]
	user.Password = ""

	return
}
