/**
 * @Author: vincent
 * @Description:
 * @File:  admin
 * @Version: 1.0.0
 * @Date: 2022/5/31 21:51
 */

package dto

import "time"

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}
