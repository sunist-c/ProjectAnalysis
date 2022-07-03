package visualization

import (
	"ProjectAnalysis/infrastructure/common"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

// HiveProcessRequest the structure of message send to hive via kafka
type HiveProcessRequest struct {
	Date   string `json:"date"`
	Type   int    `json:"type"`
	Uuid   string `json:"uuid"`
	Data   string `json:"data"`
	length int
}

// FormatDate format date field in HiveProcessRequest from time
func (h *HiveProcessRequest) FormatDate(t time.Time) {
	h.Date = t.Format("2006-01-02")
}

// FormatData format data field in HiveProcessRequest from struct
func (h *HiveProcessRequest) FormatData(object interface{}) (err error) {
	if bytes, err := json.Marshal(object); err != nil {
		return err
	} else {
		h.Data = string(bytes)
		return nil
	}
}

// FormatType format type field in HiveProcessRequest from enum
func (h *HiveProcessRequest) FormatType(t string) (err error) {
	switch t {
	case RequestImport.ToString():
		h.Type = 1
		return nil
	case RequestCountry.ToString():
		h.Type = 2
		return nil
	case RequestProvince.ToString():
		h.Type = 3
		return nil
	default:
		return errors.New("unknown request type")
	}
}

// FormatUuid generate uuid in HiveProcessRequest from date and location
func (h *HiveProcessRequest) FormatUuid(location string) {
	h.Uuid = common.GenerateMd5Len16(location, h.Date)
	h.Uuid = strings.Join([]string{h.Uuid, strconv.Itoa(h.Type), strconv.Itoa(h.length)}, "-")
}

func (h *HiveProcessRequest) FormatLength(data RequestCalculateData) {
	h.length = int(data)
}

// ToString format HiveProcessRequest to string with json
func (h *HiveProcessRequest) ToString() string {
	bytes, _ := json.Marshal(h)
	return string(bytes)
}

// RequestImportData the structure of type-import request data field
type RequestImportData struct {
	Confirm          int    `json:"confirm"`
	Death            int    `json:"death"`
	Recovered        int    `json:"recovered"`
	RefreshTime      string `json:"refreshTime"`
	LocationCountry  string `json:"locationCountry"`
	LocationProvince string `json:"locationProvince"`
}

// FormatRefreshTime format date field in RequestImportData from time
func (r *RequestImportData) FormatRefreshTime(t time.Time) {
	r.RefreshTime = t.Format("2006-01-02")
}

// RequestCalculateData the structure of type-country/type-province request data field
type RequestCalculateData int

// CalculateToday set RequestCalculateData to today
func (r *RequestCalculateData) CalculateToday() {
	*r = RequestCalculateData(1)
}

// Calculate7Days set RequestCalculateData to 7-day
func (r *RequestCalculateData) Calculate7Days() {
	*r = RequestCalculateData(7)
}

// Calculate28Days set RequestCalculateData to 28-day
func (r *RequestCalculateData) Calculate28Days() {
	*r = RequestCalculateData(28)
}

// CalculateAllDays set RequestCalculateData to all-day
func (r *RequestCalculateData) CalculateAllDays() {
	*r = RequestCalculateData(-1)
}
