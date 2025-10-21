package db

import (
	"log"
	"testing"

	"github.com/ip2location/ip2location-go/v9"
)

// go test -v -timeout 30s -count=1 -run TestIp2Location AIL2-ProveNode/db
func TestIp2Location(t *testing.T) {
	db, err := ip2location.OpenDB("../IP2LOCATION-LITE-DB5.BIN/IP2LOCATION-LITE-DB5.BIN")
	if err != nil {
		log.Fatalf("failed to open ip2location lite db: %v", err)
	}
	defer db.Close()

	ip := "180.173.63.106" // "8.8.8.8"
	results, err := db.Get_all(ip)
	if err != nil {
		log.Fatalf("failed to get ip from db: %v", err)
	}

	log.Printf("country_short: %s\n", results.Country_short)
	log.Printf("country_long: %s\n", results.Country_long)
	log.Printf("region: %s\n", results.Region)
	log.Printf("city: %s\n", results.City)
	log.Printf("isp: %s\n", results.Isp)
	log.Printf("latitude: %f\n", results.Latitude)
	log.Printf("longitude: %f\n", results.Longitude)
	log.Printf("domain: %s\n", results.Domain)
	log.Printf("zipcode: %s\n", results.Zipcode)
	log.Printf("timezone: %s\n", results.Timezone)
	log.Printf("netspeed: %s\n", results.Netspeed)
	log.Printf("iddcode: %s\n", results.Iddcode)
	log.Printf("areacode: %s\n", results.Areacode)
	log.Printf("weatherstationcode: %s\n", results.Weatherstationcode)
	log.Printf("weatherstationname: %s\n", results.Weatherstationname)
	log.Printf("mcc: %s\n", results.Mcc)
	log.Printf("mnc: %s\n", results.Mnc)
	log.Printf("mobilebrand: %s\n", results.Mobilebrand)
	log.Printf("elevation: %f\n", results.Elevation)
	log.Printf("usagetype: %s\n", results.Usagetype)
	log.Printf("addresstype: %s\n", results.Addresstype)
	log.Printf("category: %s\n", results.Category)
	log.Printf("district: %s\n", results.District)
	log.Printf("asn: %s\n", results.Asn)
	log.Printf("as: %s\n", results.As)
	log.Printf("api version: %s\n", ip2location.Api_version())
}
