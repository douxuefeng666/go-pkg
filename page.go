/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 17:32:08
 * @LastEditTime: 2024-05-21 17:32:13
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package pkg

import "math"

type Page struct {
	Total       int64       `json:"total"`
	CurrentPage int         `json:"current_page"`
	PerPage     int         `json:"per_page"`
	LastPage    int         `json:"last_page"`
	Data        interface{} `json:"data"`
}

func NewPage(total int64, currentPage, perPage int, data interface{}) *Page {
	return &Page{
		Total:       total,
		CurrentPage: currentPage,
		PerPage:     perPage,
		LastPage:    lastPage(total, perPage),
		Data:        data,
	}
}

func lastPage(total int64, perPage int) int {
	return int(math.Ceil(float64(total) / float64(perPage)))
}
