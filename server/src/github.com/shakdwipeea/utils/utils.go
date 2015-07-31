package utils
import (
	"crypto/sha1"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"errors"
)

const MySigningKey string = "madh165r35##@@#"

type Response struct {
	Err bool
	Msg string
}

func HashPassword (pass string) string {
	h := sha1.New()
	io.WriteString(h, pass)
	return string(h.Sum(nil))
}

func ErrorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"err": msg,
	})
}

func JwtAuthenticator(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		/*		if _,ok := token.Method.(*jwt.SigningMethodHS256); !ok {
					return nil, errors.New("Wrong token")
				}*/
		return []byte(MySigningKey), nil
	})

	if err != nil && !parsedToken.Valid {
		println("Tampereed Token")
		return nil, errors.New("Tampered Token")
	}

	return parsedToken, nil
}

func AuthenticateTokenGetSubject(tokenTest string) (string, error) {
	token, err := JwtAuthenticator(tokenTest)

	if err != nil {
		return "", err
	}

	subject, ok := token.Claims["subject"].(string)

	if !ok {
		return "", errors.New("Not Ok")
	}

	return subject, nil
}