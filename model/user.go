package model

import "fmt"

type User struct {
	ID       int
	Username string
	Password string
	Nama     string
	Foto     string
}

var TableName = "user"

func ValidateUser(param User) error {
	if param.Username == "" || len(param.Username) < 4 {
		return fmt.Errorf("Username minimal 4 karakter / Tidak boleh kosong")
	}

	if param.Password == "" || len(param.Password) < 8 {
		return fmt.Errorf("Password minimal 8 karakter / Tidak boleh kosong")
	}

	if param.Nama == "" || len(param.Nama) < 6 {
		return fmt.Errorf("Nama minimal 6 karakter / Tidak boleh kosong")
	}

	return nil
}
