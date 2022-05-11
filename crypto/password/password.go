package password

import (
	"github.com/lfin-open/lfin.utils.golang/strings"
	"golang.org/x/crypto/bcrypt"
)

// HashBcryptPassword bcrypt 로 패스워드 해싱, empty("") 이면 empty("") 리턴
func HashBcryptPassword(password string, cost int) (string, error) {
	if strings.IsEmptyString(password) {
		return "", nil
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// CompareHashBcryptAndPassword bcrypt 해싱 패스워드와 비교
func CompareHashBcryptAndPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
