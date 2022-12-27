package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexeyco/simpletable"
	coingecko "github.com/superoo7/go-gecko/v3"
)

var (
	data = [][]interface{}{}
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	coingeckoClient := coingecko.NewClient(httpClient)

	chart, err := coingeckoClient.CoinsIDMarketChart("bitcoin", "usd", "30")

	if err != nil {
		panic(err)
	}

	for i, v := range *chart.Prices {
		if i%12 != 0 {
			continue
		}

		data = append(data, []interface{}{
			time.Unix(int64(v[0])/1000, int64(v[0])%1000).Local().Format(time.RFC1123), float64(v[1]),
		})
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Date"},
			{Align: simpletable.AlignCenter, Text: "Price"},
		},
	}

	for _, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: row[0].(string)},
			{Text: fmt.Sprintf("$%.2f", row[1].(float64))},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleRounded)
	fmt.Println(table.String())
}
