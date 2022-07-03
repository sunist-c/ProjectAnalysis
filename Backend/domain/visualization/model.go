package visualization

import (
	"encoding/json"
	"time"
)

// LocationType the enums of types of location
type LocationType string

const (
	TypeCountry  LocationType = "Country"  // Country type, example: China
	TypeProvince LocationType = "Province" // Province type, example: Guangdong
	TypeCity     LocationType = "City"     // City type, example: Guangzhou
)

// toString convert LocationType to string case
func (l LocationType) toString() string {
	return string(l)
}

type TypeRequest string

const (
	RequestImport   TypeRequest = "import"
	RequestCountry  TypeRequest = "country"
	RequestProvince TypeRequest = "province"
)

func (t TypeRequest) ToString() string {
	return string(t)
}

// Record the structure of processing record stored in mysql database
type Record struct {
	Uuid                 string    `xorm:"pk varchar(16) index notnull unique"`
	RefreshDate          time.Time `xorm:"notnull"`
	LocationName         string    `xorm:"varchar(255) notnull"`
	LocationType         int       `xorm:"notnull"`
	DailyConfirmCase     int       `xorm:"notnull"`
	DailyDeathCase       int       `xorm:"notnull"`
	DailyRecoveredCase   int       `xorm:"notnull"`
	WeeklyConfirmCase    int       `xorm:"notnull"`
	WeeklyDeathCase      int       `xorm:"notnull"`
	WeeklyRecoveredCase  int       `xorm:"notnull"`
	MonthlyConfirmCase   int       `xorm:"notnull"`
	MonthlyDeathCase     int       `xorm:"notnull"`
	MonthlyRecoveredCase int       `xorm:"notnull"`
	TotalConfirmCase     int       `xorm:"notnull"`
	TotalDeathCase       int       `xorm:"notnull"`
	TotalRecoveredCase   int       `xorm:"notnull"`
}

type RequestCache struct {
	Uuid      string `json:"uuid"`
	Date      string `json:"date"`
	Location  string `json:"location"`
	Days1     bool   `json:"days1"`
	Days7     bool   `json:"days7"`
	Days28    bool   `json:"days28"`
	DaysTotal bool   `json:"daysTotal"`
	Imported  bool   `json:"imported"`
	Record    Record `json:"record"`
}

func (r RequestCache) toString() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}

func (r *RequestCache) Format(str string) {
	json.Unmarshal([]byte(str), r)
}

func (r RequestCache) Ready() bool {
	return r.Days1 && r.Days7 && r.Days28 && r.DaysTotal && r.Imported
}
