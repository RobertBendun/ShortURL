package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func updateURLDatabase(csvFilename string, ticker *time.Ticker) {
	for range ticker.C {
		file, err := os.Open(csvFilename)
		defer file.Close()
		if err != nil {
			continue
		}
		w := csv.NewWriter(file)
		for url, entry := range urlEntries {
			record := make([]string, 3)
			record[0] = url
			record[1] = entry.url
			record[2] = fmt.Sprintf("%d", entry.visited)
			w.Write(record)
		}
		w.Flush()
	}
}
