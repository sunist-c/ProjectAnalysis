package visualization

// BaseResponse the common structure of response
type BaseResponse struct {
	ErrorCode int    `json:"err_code"`
	Message   string `json:"msg"`
}

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
