/**
 * @Author: vincent
 * @Description:
 * @File:  util
 * @Version: 1.0.0
 * @Date: 2022/3/11 19:14
 */

package demo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFile(fileName string, m map[string]int64) {
	// 第二种
	fd, err := os.Open(fileName)
	defer fd.Close()
	if err != nil {
		fmt.Println("read error:", err)
	}

	scanner := bufio.NewScanner(fd)
	buf := make([]byte, 0, 10*1024*1024)
	scanner.Buffer(buf, 10*1024*1024)
	i := 0
	for scanner.Scan() {
		data := scanner.Text() // or
		//line := scanner.Bytes()
		i++
		logId, err := doSplit(data, i)
		if err != nil {
			fmt.Printf("panic at %d\n", i)
			panic(err)
		}
		// add count
		m[logId]++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Cannot scanner text file: %s, line:%d, err: [%v]", fileName, i, err)
		return
	}

	/*
		buff := bufio.NewReader(fd)
		i := 0
		for {
			data, _, eof := buff.ReadLine()
			if eof == io.EOF {
				break
			}
			i++

			logId, err := doSplit(string(data), i)
			if err != nil {
				panic(err)
			}

			// add count
			m[logId]++
		}

	*/
}

func readFile2(fileName string, m map[string]int64) {
	// 第二种
	fd, err := os.Open(fileName)
	defer fd.Close()
	if err != nil {
		fmt.Println("read error:", err)
	}
	buff := bufio.NewReader(fd)
	i := 0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		i++

		logId, err := doSplit(string(data), i)
		if err != nil {
			panic(err)
		}

		// add count
		m[logId]++
	}
}

func doSplit(str string, line int) (string, error) {
	SPLIT_NUM := 4
	LOG_ID_IDX := 2

	//指定分隔符
	countSplit := strings.SplitN(str, "|", SPLIT_NUM)
	//fmt.Println(len(countSplit))
	if len(countSplit) < SPLIT_NUM {
		fmt.Printf("line(%d)---- str:\n%s\n", line, str)
		fmt.Printf("line(%d)---- splict:\n%+v\n", line, countSplit)
		return "", errors.New("not enough split data")
	}
	/*
		for _, item := range countSplit {
			log.Printf("%v\n", item)
		}
	*/
	return countSplit[LOG_ID_IDX], nil
}
