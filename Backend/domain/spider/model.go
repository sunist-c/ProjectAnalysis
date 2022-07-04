package spider

import "strconv"

type RawData struct {
	CountryName  string
	ProvinceName string
	Date         string
	Confirm      float64
	Death        float64
	Recovered    float64
}

func (r *RawData) Format(arr []string) {
	// todo: remove hard code
	r.CountryName = arr[0]
	r.ProvinceName = arr[1]
	r.Date = arr[2]
	r.Confirm, _ = strconv.ParseFloat(arr[3], 64)
	r.Death, _ = strconv.ParseFloat(arr[4], 64)
	r.Recovered, _ = strconv.ParseFloat(arr[5], 64)
}
