package password

import (
	_ "embed"
	"strings"
)

//go:embed password.txt
var passwords []byte
var cachePassword []string

func GetPasswords() []string {
	if len(cachePassword) > 0 {
		return cachePassword
	} else {
		data := string(passwords)
		cachePassword = strings.Split(data, "\n")
		return cachePassword
	}
}
