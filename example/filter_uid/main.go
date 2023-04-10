package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	chJob    = make(chan *job, 100)
	writeJob = make(chan string, 100)
)

type job struct {
	uid    string
	orders []string
}

const (
	inputFile = "order_export.txt"
)

func main() {
	doFilter()
}

func doFilter() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go produceJob(&wg)
	go consumeJob(&wg)
	go writeFile(&wg)
	wg.Wait()
}

func produceJob(wg *sync.WaitGroup) {
	defer wg.Done()
	var newJob *job
	// read file
	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
		return
	}
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	uid := "" // initial value
	i := 0
	for scanner.Scan() {
		i++
		fmt.Printf("%d | line: %s\n", i, scanner.Text())
		// uid order_id pid start_time end_time
		line := strings.SplitN(scanner.Text(), " ", 2)
		tmpUid := line[0]
		if uid == "" {
			uid = tmpUid
			newJob = &job{uid: uid, orders: []string{}}
			newJob.orders = append(newJob.orders, scanner.Text())
		} else if uid == tmpUid {
			// same uid
			newJob.orders = append(newJob.orders, scanner.Text())
		} else {
			// uid != tmpUid
			// 发送到队列
			chJob <- newJob // todo newJob是一个指针，会不会有问题
			uid = tmpUid
			newJob = &job{uid: uid, orders: []string{}}
			newJob.orders = append(newJob.orders, scanner.Text())
		}

	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}
	close(chJob)
}

func consumeJob(wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range chJob {
		go func(item *job) {
			if len(item.orders) > 3 {
				writeJob <- item.uid
			}
		}(j)
	}
	fmt.Println("consumeJob finished")
	close(writeJob)
}

func writeFile(wg *sync.WaitGroup) {
	defer wg.Done()
	fileName := "huomian_uid.txt"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
		return
	}
	defer f.Close()
	for uid := range writeJob {
		_, err = f.WriteString(uid + "\n")
	}
	fmt.Println("writeFile finished")
}

func generateFile() {
	f, err := os.OpenFile(inputFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
		return
	}
	defer f.Close()

	uid := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		for j := 0; j < rand.Intn(10); j++ {
			// uid order_id pid start_time end_time
			order_id := uid
			pid := 3
			startTime := uid
			endTime := uid + 36000
			_, err = f.WriteString(fmt.Sprintf("%d %d %d %d %d\n", uid, order_id, pid, startTime, endTime))
			if err != nil {
				panic(err)
			}
		}
		uid++
	}
}
