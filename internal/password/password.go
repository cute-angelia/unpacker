package password

import (
	_ "embed"
	"strings"
)

//go:embed password.txt
var passwords []byte

func GetPasswords() []string {
	data := string(passwords)
	return strings.Split(data, "\n")
}
