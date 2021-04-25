package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func read_csv() []map[string]string {

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

	rand.Seed(time.Now().UnixNano())

	word_list := read_csv()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		index := rand.Intn(len(word_list))
		bytes, err := json.Marshal(word_list[index])
		if err != nil {
			fmt.Println("JSON marshal error: ", err)
		}
		return c.String(http.StatusOK, string(bytes))
	})
	e.Logger.Fatal(e.Start(":1323"))

}
