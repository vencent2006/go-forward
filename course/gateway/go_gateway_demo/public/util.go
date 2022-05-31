/**
 * @Author: vincent
 * @Description:
 * @File:  util
 * @Version: 1.0.0
 * @Date: 2022/5/31 15:50
 */

package public

import (
	"crypto/sha256"
	"fmt"
)

func GenSaltPassword(salt, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	return fmt.Sprintf("%x", s2.Sum(nil))
}
