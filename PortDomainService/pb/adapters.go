package ports

import "github.com/tonky/shipping/PortDomainService/domain"

func FromDomain(p domain.Port) Port {
	return Port{
		Key:          p.Key,
		Name:         p.Name,
		City:         p.City,
		Country:      p.Country,
		Alias:        p.Aliases,
		Regions:      p.Regions,
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		Province:     p.Province,
		LocationName: p.LocationName,
		Unlocs:       p.Unlocs,
		Code:         p.Code,
	}
}

func ToDomain(p Port) domain.Port {
	return domain.Port{
		Key:          p.Key,
		Name:         p.Name,
		City:         p.City,
		Country:      p.Country,
		Aliases:      p.Alias,
		Regions:      p.Regions,
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		Province:     p.Province,
		LocationName: p.LocationName,
		Unlocs:       p.Unlocs,
		Code:         p.Code,
	}
}
