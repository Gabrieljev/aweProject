package member

import (
	"golang.org/x/crypto/bcrypt"
)

func (s *service) encrypt(password string) (resultSha string) {

	byteArr, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	resultSha = string(byteArr)
	return
}

func (s *service) comparingHash(hashPwd string, password string) (res bool){
	res = true
	if err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password)); err != nil {
		res = false
	}
	return
}
