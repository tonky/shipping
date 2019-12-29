package domain

import "fmt"

type Port struct {
	Key          string
	Name         string
	City         string
	Country      string
	Aliases      []string
	Regions      []string
	Latitude     float64
	Longitude    float64
	Province     string
	LocationName string // time.LoadLocation
	Unlocs       []string
	Code         string
}

func (p Port) String() string {
	return fmt.Sprintf(`[%s](%s)-%s-%s-%s`, p.Key, p.Name, p.City, p.Country, p.Code)
}
