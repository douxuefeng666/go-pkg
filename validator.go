/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-22 10:49:58
 * @LastEditTime: 2024-05-22 10:53:15
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package pkg

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 表单验证
// https://github.com/go-playground/validator
var (
	uni           *ut.UniversalTranslator
	Validate      *validator.Validate
	ValidateTrans ut.Translator
)

func init() {
	zh2 := zh.New()
	uni = ut.New(zh2, zh2)
	ValidateTrans, _ = uni.GetTranslator("zh")
	Validate = validator.New()
	// 收集结构体中的comment标签，用于替换英文字段名称，这样返回错误就能展示中文字段名称了
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})
	if err := zhtranslations.RegisterDefaultTranslations(Validate, ValidateTrans); err != nil {
		fmt.Printf("RegisterDefaultTranslations %v", err)
	}
}

func ValidateStruct(s interface{}) error {
	err := Validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(ValidateTrans) {
			if len(e) > 0 {
				return errors.New(e)
			}
		}
	}
	return nil
}
