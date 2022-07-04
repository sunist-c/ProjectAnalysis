package visualization

type ResponseCalculateCountry struct {
	Confirm     int    `json:"confirm"`
	Death       int    `json:"death"`
	Recovered   int    `json:"recovered"`
	RefreshTime int    `json:"refreshTime"`
	CountryName string `json:"countryName"`
}

type ResponseCalculateProvince struct {
	Confirm      int    `json:"confirm"`
	Death        int    `json:"death"`
	Recovered    int    `json:"recovered"`
	RefreshTime  int    `json:"refreshTime"`
	CountryName  string `json:"countryName"`
	ProvinceName string `json:"provinceName"`
}

type ResponseImport string

type BaseHiveProcessResponse struct {
	Date string `json:"date"`
	Uuid string `json:"uuid"`
}

type CalculateCountryHiveProcessResponse struct {
	BaseHiveProcessResponse
	Data ResponseCalculateCountry `json:"data"`
}

type CalculateProvinceHiveProcessResponse struct {
	BaseHiveProcessResponse
	Data ResponseCalculateProvince `json:"data"`
}

type ImportHiveProcessResponse struct {
	BaseHiveProcessResponse
	Data ResponseImport `json:"data"`
}
