package requests

import "errors"

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Register) IsValid() error {
	if r.Name == "" {
		return errors.New("O Nome é obrigatório")
	}

	if r.Email == "" {
		return errors.New("O E-mail é obrigatório")
	}

	if r.Password == "" {
		return errors.New("A Senha é obrigatória")
	}

	return nil
}
