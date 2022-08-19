package helper

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

func SecurePassword(password string)string{
	hash,er:=bcrypt.GenerateFromPassword([]byte(password),14)
	if er != nil {
		return ""
	}
	return string(hash)
}

func ComparePassword(hash string,password string)bool{
	er := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))

	return er != nil
}

func MarshalGiven(v interface{})[]byte {
	value,er := json.Marshal(v)

	if er != nil {
		return nil
	}
	return value
}