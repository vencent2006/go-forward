package filter_uid

import (
	"bufio"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	WORKER_NUM = 100
	inputFile  = "order.csv"
)

var (
	chJob       = make(chan *job, WORKER_NUM)
	writeJob    = make(chan string, WORKER_NUM)
	workerToken = make(chan struct{}, WORKER_NUM)
	eg          *errgroup.Group
	ctx         context.Context
)

type job struct {
	uid    string
	orders []string
}

func init() {
	for i := 0; i < WORKER_NUM; i++ {
		workerToken <- struct{}{}
	}
}

func doFilter() {
	eg, ctx = errgroup.WithContext(context.Background())
	eg.Go(produceJob)
	eg.Go(consumeJob)
	eg.Go(writeFile)

	if err := eg.Wait(); err != nil {
		panic(err)
	}
}

func produceJob() error {
	defer close(chJob)
	var newJob *job
	// read file
	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	uid := "" // initial value
	i := 0
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			fmt.Println("produceJob ctx.Done")
			return ctx.Err()
		default:

		}

		i++
		fmt.Printf("%d | line: %s\n", i, scanner.Text())
		// uid order_id pid start_time end_time
		line := strings.SplitN(scanner.Text(), ",", 2)
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
		return err
	} else {
		// 发最后一个
		chJob <- newJob
	}
	return nil
}

func consumeJob() error {
	defer close(writeJob)
	tmpEg, tmpCtx := errgroup.WithContext(ctx)
EXIT:
	for {
		select {
		case <-tmpCtx.Done():
			fmt.Println("consumeJob ctx.Done")
			return tmpCtx.Err()
		case j, ok := <-chJob:
			if !ok {
				break EXIT
			}
			<-workerToken
			tmpEg.Go(func() error {
				defer func() { workerToken <- struct{}{} }()
				time.Sleep(time.Duration(rand.Intn(20)+10) * time.Millisecond)
				writeJob <- j.uid
				return nil
			})
		}
	}
	if err := tmpEg.Wait(); err != nil {
		fmt.Println("consumeJob err:", err)
		return err
	}
	fmt.Println("consumeJob finished")

	return nil
}

func writeFile() error {
	fileName := "huomian_uid.txt"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

EXIT:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writeFile ctx.Done")
			return ctx.Err()
		case uid, ok := <-writeJob:
			if !ok {
				break EXIT
			}
			fmt.Printf("-------------- huomian|uid|%s\n", uid)
			_, err = f.WriteString(uid + "\n")
		}
	}

	fmt.Println("writeFile finished")
	return nil
}
