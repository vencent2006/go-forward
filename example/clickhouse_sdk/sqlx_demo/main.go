package main

import (
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

func main() {
	connect, err := sqlx.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		log.Fatal(err)
	}
	var items []struct {
		URL         string  `db:"URL"`
		AvgDuration float32 `db:"AvgDuration"`
	}

	if err := connect.Select(&items, `SELECT
    StartURL AS URL,
    AVG(Duration) AS AvgDuration
FROM tutorial.visits_v1
WHERE StartDate BETWEEN '2014-03-23' AND '2014-03-30'
GROUP BY URL
ORDER BY AvgDuration DESC
LIMIT 10`); err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		log.Printf("URL: %s, AvgDuration: %f\n", item.URL, item.AvgDuration)
	}
}
