package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "sync"
    "time"
)

const (
    LOOP_SLEEP_MS  = time.Millisecond * 100
    MAX_WORKER_NUM = 10
)

var (
    workerChan = make(chan interface{}, MAX_WORKER_NUM)
)

func main() {

    defer close(workerChan)

    blockNumber := int64(1) // block number ready to process
    chain := "eth"
    typ := "coin"

    // 1. load data: block_on_process, wallet_addr_list, contract_addr_list, contract_config
    // 2. main loop
    mainLoop(chain, typ, blockNumber)

}

func mainLoop(chain string, typ string, blockNumber int64) {

    for i := 0; i < MAX_WORKER_NUM; i++ {
        workerChan <- 1
    }

    for {
        time.Sleep(LOOP_SLEEP_MS)
        txHashes := MustGetTxHashByBlockNumber(chain, typ, blockNumber)
        txHashesLen := len(*txHashes)
        if txHashesLen <= 0 {
            blockNumber++ // 开始下一个blockNumber
            continue
        }
        wg := sync.WaitGroup{}
        wg.Add(txHashesLen)
        for i, txHash := range *txHashes {
            _, ok := <-workerChan
            if !ok {
                return // because chan closed
            }

            go processOneTxHash(i, txHash, &wg)
        }
        wg.Wait()
        fmt.Printf("********************* fininsh block %d\n", blockNumber)
        blockNumber++ // 开始下一个blockNumber

        if blockNumber >= 3 {
            break
        }
    }
}

func MustGetTxHashByBlockNumber(chain string, typ string, blockNumber int64) *[]string {
    var txHashes *[]string
    var err error
    for {
        txHashes, err = getTxHashByBlockNumber(chain, typ, blockNumber)
        if err == nil {
            return txHashes
        }
        time.Sleep(LOOP_SLEEP_MS)
    }
}

func loadData() error {
    return nil
}

func getTxHashByBlockNumber(chain string, typ string, blockNumber int64) (*[]string, error) {
    cnt := 20
    txHashes := make([]string, 0, cnt)
    for i := 0; i < cnt; i++ {
        txHashes = append(txHashes, strconv.Itoa(i))
    }
    fmt.Printf("----- %+v\n", txHashes)
    return &txHashes, nil
}

func processOneTxHash(i int, txHash string, wg *sync.WaitGroup) {
    defer wg.Done()
    defer func() { workerChan <- 1 }()
    fmt.Printf("%d ready to processOneTxHash\n", i)
    n := rand.Intn(20)
    time.Sleep(time.Second * time.Duration(n))
    fmt.Printf("%d: sleep %d seconds, finished %s\n", i, n, txHash)
}
