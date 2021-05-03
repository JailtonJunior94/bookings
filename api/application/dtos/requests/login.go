package requests

import "errors"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *Login) IsValid() error {
	if l.Email == "" {
		return errors.New("O E-mail é obrigatório")
	}

	if l.Password == "" {
		return errors.New("A Senha é obrigatória")
	}

	return nil
}
