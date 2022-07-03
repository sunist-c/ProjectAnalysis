package _map

import "encoding/json"

var (
	Info Map
)

func init() {
	Info = Map{Counties: make(map[string]Record)}
}

type Record struct {
	Country   string   `json:"country"`
	Provinces []string `json:"provinces"`
}

type Map struct {
	Counties map[string]Record `json:"counties"`
}

func (m *Map) AddRecord(country, province string) {
	// todo: use sync.Map to avoid threading errors
	if v, ok := m.Counties[country]; ok {
		for _, i := range v.Provinces {
			if province == i {
				break
			}
		}

		v.Provinces = append(v.Provinces, province)
	} else {
		m.Counties[country] = Record{
			Country:   country,
			Provinces: []string{province},
		}
	}
}

func (m Map) ToString() string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}
