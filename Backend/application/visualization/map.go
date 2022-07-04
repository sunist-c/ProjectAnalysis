package visualization

import _map "ProjectAnalysis/domain/map"

func (a Application) QueryCountryMap(country string) []string {
	if provinces, ok := _map.Info.Counties[country]; ok {
		return provinces.Provinces
	} else {
		return nil
	}
}

func (a Application) QueryWholeMap() map[string][]string {
	result := make(map[string][]string)
	for key, record := range _map.Info.Counties {
		result[key] = record.Provinces
	}

	return result
}
