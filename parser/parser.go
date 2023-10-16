package parser

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gocolly/colly"
)

type RequestData struct {
	Webpage  string `json:"webpage"`
	Selector string `json:"selector"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Result  []string    `json:"result"`
	Params  RequestData `json:"params"`
}

func HandleApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var (
		data     RequestData
		response ResponseData
	)

	bodyBytes, _ := io.ReadAll(r.Body)
	bodyString := string(bodyBytes)

	json.Unmarshal([]byte(bodyString), &data)

	response = ResponseData{"", nil, data}

	if data.Webpage == "" || data.Selector == "" {
		response.Message = "webpage or selector fields are empty"
		json.NewEncoder(w).Encode(response)

		return
	}

	parsedData := parseData(data.Webpage, data.Selector)

	if len(parsedData) == 0 {
		response.Message = "no data"
		json.NewEncoder(w).Encode(response)

		return
	}

	response.Message = "done"
	response.Result = parsedData

	json.NewEncoder(w).Encode(response)
}

func parseData(url string, selector string) []string {
	var data []string

	c := colly.NewCollector()

	c.OnHTML(selector, func(e *colly.HTMLElement) {
		data = append(data, e.Text)
	})

	c.Visit(url)

	return data
}
