package api

// This structs parameters must start with a capital letter and match
// the JSON keys. These comments are used because the json package uses reflection to
// match the keys and values. In other words, the json package can read the
// metadata when is marshaling and unmarshaling.
type CEXResponse struct {
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"`
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}
