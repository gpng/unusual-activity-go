package download

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/montanaflynn/stats"
)

const months = 5
const cutoff = 8
const days = 3

// Ticker struct
type Ticker struct {
	Name   string
	Code   string
	Symbol string
}

const baseURL = "https://query1.finance.yahoo.com/v8/finance/chart/"

func getTickers() ([]Ticker, error) {
	csvFile, err := os.Open("data/tickers.csv")
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csvFile)

	tickers := []Ticker{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		tickers = append(tickers, Ticker{
			Name:   record[0],
			Code:   record[1],
			Symbol: record[2],
		})
	}

	return tickers, nil
}

type dataResponseQuote struct {
	Volume []float64 `json:"volume"`
}

type dataResponseIndicators struct {
	Quote []dataResponseQuote `json:"quote"`
}

type dataResponseResult struct {
	Indicators dataResponseIndicators `json:"indicators"`
}

type dataResponseChart struct {
	Result []dataResponseResult `json:"result"`
}

type dataResponse struct {
	Chart dataResponseChart `json:"chart"`
}

func (t *Ticker) getData() error {
	endDate := time.Now().UTC()
	endDate = endDate.Truncate(24 * time.Hour)
	startDate := endDate.AddDate(0, -months, 0)

	client := &http.Client{}

	req, err := http.NewRequest("GET", baseURL+t.Code+".SI", nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("period1", strconv.FormatInt(startDate.Unix(), 10))
	q.Add("period2", strconv.FormatInt(endDate.Unix(), 10))
	q.Add("interval", "1d")

	req.URL.RawQuery = q.Encode()
	// log.Println(req.URL.String())

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body := dataResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil || len(body.Chart.Result) < 1 || len(body.Chart.Result[0].Indicators.Quote) < 1 {
		return err
	}

	volume := body.Chart.Result[0].Indicators.Quote[0].Volume
	if len(volume) < days {
		return fmt.Errorf("Not enough trading days")
	}
	mean, err := stats.Mean(volume)
	if err != nil {
		return err
	}
	std, err := stats.StandardDeviation(volume)
	if err != nil {
		return err
	}
	anomalyCutoff := std * cutoff
	limit := mean + anomalyCutoff

	lastVolumes := volume[len(volume)-days:]
	for _, volume := range lastVolumes {
		if volume > limit {
			log.Printf("%s.SI\thttps://sg.finance.yahoo.com/quote/%s.SI\t%s - Unusually high volume", t.Code, t.Code, t.Name)
		}
	}

	return nil
}

// Download ticker
func Download() {
	defer timeTrack(time.Now(), "Finding anomalies")
	tickers, err := getTickers()
	if err != nil {
		log.Println(err)
	}
	for _, ticker := range tickers {
		ticker.getData()
		// if err != nil {
		// 	log.Println(err)
		// }
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
