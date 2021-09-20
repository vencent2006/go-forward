/**
 * @Author: vincent
 * @Description:
 * @File:  core
 * @Version: 1.0.0
 * @Date: 2021/9/19 23:05
 */

package framework

import "net/http"

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
