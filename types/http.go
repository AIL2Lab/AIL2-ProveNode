package types

type BaseHttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LocationResponse struct {
	BaseHttpResponse
	ClientIP        string `json:"client_ip"`
	BandwidthRegion string `json:"bandwidth_region"`
	// ip2location.IP2Locationrecord
	CountryShort       string  `json:"country_short"`
	CountryLong        string  `json:"country_long"`
	Region             string  `json:"region"`
	City               string  `json:"city"`
	Isp                string  `json:"isp"`
	Latitude           float32 `json:"latitude"`
	Longitude          float32 `json:"longitude"`
	Domain             string  `json:"domain"`
	Zipcode            string  `json:"zipcode"`
	Timezone           string  `json:"timezone"`
	Netspeed           string  `json:"netspeed"`
	Iddcode            string  `json:"iddcode"`
	Areacode           string  `json:"areacode"`
	Weatherstationcode string  `json:"weather_station_code"`
	Weatherstationname string  `json:"weather_station_name"`
	Mcc                string  `json:"mcc"`
	Mnc                string  `json:"mnc"`
	Mobilebrand        string  `json:"mobilebrand"`
	Elevation          float32 `json:"elevation"`
	Usagetype          string  `json:"usage_type"`
	Addresstype        string  `json:"address_type"`
	Category           string  `json:"category"`
	District           string  `json:"district"`
	Asn                string  `json:"asn"`
	As                 string  `json:"as"`
}

type CalculatePointResponse struct {
	BaseHttpResponse
	CalcPoint float64 `json:"calc_point"`
}
