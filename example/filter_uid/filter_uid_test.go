package filter_uid

import (
	"runtime"
	"testing"
)

// awk -F, '{print $1}' order.csv | sort -n | uniq | wc -l
// wc -l huomian_uid.txt
func Test_doFilter(t *testing.T) {
	runtime.GOMAXPROCS(2)
	doFilter()
}

/*
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

func Test_generate_file(t *testing.T) {
	generateFile()
}

*/
