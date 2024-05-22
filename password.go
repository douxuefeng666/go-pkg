/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 17:30:01
 * @LastEditTime: 2024-05-21 17:30:12
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package pkg

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func ValidatePassword(encodePassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(inputPassword))
	return err == nil
}
