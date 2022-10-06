/**
 * @Author: vincent
 * @Description:
 * @File:  conn
 * @Version: 1.0.0
 * @Date: 2022/10/1 20:57
 */

package resp

type Connection interface {
	Write([]byte) error
	GetDBIndex() int
	SelectDB(int)
}
