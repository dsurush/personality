package helperfunc

import "golang.org/x/crypto/bcrypt"

type TimeInterval struct {
	From string
	To   string
}

type PasswordChange struct {
	Pass    string `json:"pass"`
	NewPass string `json:"new_pass"`
}

func MaxOftoInt(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MinOftoInt(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
