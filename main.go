package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
	coingecko "github.com/superoo7/go-gecko/v3"
)

var (
	data = [][]interface{}{}
)

func main() {
	os.Exit(renderResult(os.Stdout))
}

func renderResult(out io.Writer) int {
	var arg string
	days := "30"

	if len(os.Args) > 1 {
		arg = os.Args[1]

		if _, err := strconv.Atoi(arg); err != nil {
			fmt.Fprintf(out, "Argument `%s` is not a number", arg)
			return 1
		}

		days = arg
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	coingeckoClient := coingecko.NewClient(httpClient)

	chart, err := coingeckoClient.CoinsIDMarketChart("bitcoin", "usd", days)

	if err != nil {
		log.Fatal(out, err)
		return 1
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
	fmt.Fprint(out, table.String())

	return 0
}
