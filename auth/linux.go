//go:build linux

package auth

import (
	"github.com/msteinert/pam"
)

func Authenticate(user, pass string) (bool, error) {
	t, err := pam.StartFunc("login", user, func(s pam.Style, msg string) (string, error) {
		switch s {
		case pam.PromptEchoOff:
			return pass, nil
		}
		return "", nil
	})
	if err != nil {
		return false, err
	}
	err = t.Authenticate(0)
	if err != nil {
		return false, err
	}
	return true, nil
}
