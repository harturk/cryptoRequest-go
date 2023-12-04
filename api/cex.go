package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/harturk/go-first/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}

	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		// Check if the flow of bytes was interrupted
		if err != nil {
			return nil, err
		}
		// Go always makes a clone of the value when passing it to a function
		// So we need to pass a pointer to the value
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
