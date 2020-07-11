package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readURLDatabase(filename string) (map[string]entry, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0700)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.ReuseRecord = true
	reader.FieldsPerRecord = 3

	entries := make(map[string]entry)
	for {
		record, err := reader.Read()
		switch err {
		case nil:
			break
		case io.EOF:
			return entries, nil
		default:
			return nil, err
		}
		e := entry{url: record[1]}
		public := 0
		fmt.Sscanf(record[2], "%d", &public)
		e.public = public != 0
		entries[record[0]] = e
	}
}
