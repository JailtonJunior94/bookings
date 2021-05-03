package adapters

import "golang.org/x/crypto/bcrypt"

type IHashAdapter interface {
	GenerateHash(str string) (string, error)
	CheckHash(hash, str string) bool
}

type HashAdapter struct {
}

func NewHashAdapter() IHashAdapter {
	return &HashAdapter{}
}

func (h *HashAdapter) GenerateHash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 5)
	return string(bytes), err
}

func (h *HashAdapter) CheckHash(hash, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
