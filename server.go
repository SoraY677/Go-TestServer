package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func read_csv() []map[string]string{}{

	fp, err := os.Open("crazy.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(fp)

	i := 0
	crazy_list := []map[string]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if i != 0 {

			item :=
				map[string]string{
					"date":   record[0],
					"action": record[1],
				}

			crazy_list = append(crazy_list, item)
		}
		i++
	}

	return crazy_list
}

func main() {

	



}
