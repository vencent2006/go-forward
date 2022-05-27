/**
 * @Author: vincent
 * @Description:
 * @File:  split_test
 * @Version: 1.0.0
 * @Date: 2022/3/11 18:57
 */

package demo

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplitStringSimple(t *testing.T) {
	str := "-100 123 200"

	//指定分隔符
	countSplit := strings.Split(str, " ")
	fmt.Println(countSplit, len(countSplit))

	//指定分割符号，指定分割次数
	countSplit = strings.SplitN(str, " ", 2)
	fmt.Println(countSplit, len(countSplit))
}

func TestSplitString(t *testing.T) {
	str := `INFO|2022-03-09 00:00:00|552000810282135|0|res|mysql|1646755200.0816|3.96|0|45880cbd976149208761e6f99a290bfb|{"host":"eos.grid.com.cn","port":"4004","user":"database_r","function":"connect","result":{},"code":0,"trace":[["\/library\/Dr\/Module.php",266],["\/library\/Model\/Profile\/Module.php",333],["\/controllers\/Api\/External\/UserProfile\/Module.php",136],["\/controllers\/Interface\/External\/Profile\/Module.php",89]]}`
	//指定分隔符
	countSplit := strings.SplitN(str, "|", 4)
	fmt.Println(len(countSplit))
	for _, item := range countSplit {
		fmt.Printf("%v\n", item)
	}
}

func TestDoSplit(t *testing.T) {
	idMaps := make(map[string]int64)
	fileName := "example.log"
	readFile(fileName, idMaps)
	fmt.Printf("len(idMaps) = %d\n", len(idMaps))
	//fmt.Printf("idMaps = %+v\n", idMaps)
}
