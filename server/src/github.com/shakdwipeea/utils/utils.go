package utils
import (
	"crypto/sha1"
	"io"
)


type Response struct {
	Err bool
	Msg string
}

func HashPassword (pass string) string {
	h := sha1.New()
	io.WriteString(h, pass)
	return string(h.Sum(nil))
}