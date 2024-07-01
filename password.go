/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 17:30:01
 * @LastEditTime: 2024-06-30 13:41:53
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
		fmt.Println("Ecode Password Err==:", err)
		return ""
	}
	return string(hash)
}

func ValidatePassword(encodePassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(inputPassword))
	return err == nil
}
