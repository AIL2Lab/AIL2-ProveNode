package db

import (
	"fmt"

	"github.com/ip2location/ip2location-go/v9"
)

var IP2LDB *ip2location.DB = nil

func InitIP2LDB(dbpath string) error {
	var err error
	IP2LDB, err = ip2location.OpenDB(dbpath)
	if err != nil {
		return fmt.Errorf("failed to open ip2location lite db %v at %v", err, dbpath)
	}
	return nil
}

func GetPositionOfIP(ip string) (ip2location.IP2Locationrecord, error) {
	// results, err := IP2LDB.Get_all(ip)
	// if err != nil {
	// 	return 0, 0, err
	// }
	// return results.Longitude, results.Latitude, nil
	return IP2LDB.Get_all(ip)
}

/*
“华南”和“华北”是中国地理区域划分的常用术语，主要是根据地理位置、气候条件、历史文化等因素进行区分。以下是具体的划分和包含的省份：

1. 华北地区
华北地区位于中国的北部，通常包括以下省份和直辖市：
北京市
天津市
河北省
山西省
内蒙古自治区（东部地区）

华北地区的特点是气候较为干旱，冬季寒冷，夏季炎热，是中国重要的农业区和工业区之一。

2. 华南地区
华南地区位于中国的南部，通常包括以下省份和自治区：
广东省
广西壮族自治区
海南省
香港特别行政区
澳门特别行政区
福建省（部分地区）
台湾省（中国的省份）

华南地区的特点是气候温暖湿润，四季如春，是中国经济最发达的地区之一，尤其是广东省。

3. 其他相关地理区域
为了更全面地理解中国的区域划分，以下是与“华南”和“华北”相关的其他地理区域：
华东地区：包括上海、江苏、浙江、安徽、江西、山东、福建等。
华中地区：包括河南、湖北、湖南等。
西北地区：包括陕西、甘肃、青海、宁夏、新疆等。
西南地区：包括四川、重庆、贵州、云南、西藏等。
东北地区：包括辽宁、吉林、黑龙江等。

这些区域划分有助于更好地理解中国的地理、经济和文化差异。
*/

func GetBandwidthRegion(loc *ip2location.IP2Locationrecord) string {
	switch loc.Country_long {
	case "China":
		switch loc.Region {
		case "Beijing", "Tianjin", "Hebei", "Shanxi", "Nei Mongol":
			return "North China"
		case "Guangdong", "Guangxi", "Guangxi Zhuangzu", "Hainan":
			return "South China"
		case "Henan", "Hubei", "Hunan":
			return "Central China"
		case "Shanghai", "Jiangsu", "Zhejiang", "Anhui", "Jiangxi", "Shandong", "Fujian":
			return "East China"
		case "Shaanxi", "Gansu", "Qinghai", "Ningxia", "Ningxia Huizu", "Xinjiang":
			return "Northwest China"
		case "Sichuan", "Chongqing", "Guizhou", "Yunnan", "Xizang":
			return "Southwest China"
		case "Liaoning", "Jilin", "Heilongjiang":
			return "Northeast China"
		case "Taiwan":
			return "Taiwan, China"
		case "Hong Kong":
			return "Hong Kong, China"
		default:
			return fmt.Sprintf("%v, %v", loc.Region, loc.Country_long)
		}
	case "Taiwan (Province of China)":
		return "Taiwan, China"
	case "Hong Kong":
		return "Hong Kong, China"
	case "India":
		switch loc.Region {
		case "Uttar Pradesh", "Maharashtra", "Bihar":
			return loc.Region
		default:
			return fmt.Sprintf("%v, %v", loc.Region, loc.Country_long)
		}
	case "United States of America":
		switch loc.Region {
		case "California", "Texas", "Florida", "New York", "Pennsylvania", "Illinois", "Ohio", "Georgia", "Michigan", "North Carolina":
			return loc.Region
		default:
			return "Other Regions of the USA"
		}
	case "Russian Federation":
		// switch loc.Region {
		// case "Moskva":
		// 	return "Moscow"
		// case "Sankt-Peterburg":
		// 	return "Saint Petersburg"
		// default:
		// 	return "Other parts of Russia"
		// }
		switch loc.City {
		case "Moscow", "Saint Petersburg":
			return loc.City
		default:
			return "Other parts of Russia"
		}
	case "Korea (the Republic of)":
		return "South Korea"
	case "Viet Nam":
		return "Vietnam"
	case "Turkiye":
		return "Turkey"
	case "United Kingdom of Great Britain and Northern Ireland":
		return "United Kingdom"
	case "Netherlands (Kingdom of the)":
		return "Netherlands"
	default:
		return loc.Country_long
	}
}
