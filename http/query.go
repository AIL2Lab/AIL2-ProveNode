package http

import (
	"AIL2-ProveNode/db"
	"AIL2-ProveNode/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Location(ctx *gin.Context) {
	var response types.LocationResponse
	// response := types.LocationResponse{
	// 	BaseHttpResponse: types.BaseHttpResponse{
	// 		Code:    0,
	// 		Message: "ok",
	// 	},
	// 	ClientIP: ctx.ClientIP(),
	// }
	ip := ctx.ClientIP()
	type UserIp struct {
		Ip string `json:"ip" form:"ip"`
	}
	user_ip := UserIp{}
	if err := ctx.ShouldBindQuery(&user_ip); err != nil {
		response.Code = int(types.ErrCodeParse)
		response.Message = types.ErrCodeParse.String()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if user_ip.Ip != "" {
		ip = user_ip.Ip
	}
	loc, err := db.GetPositionOfIP(ip)
	if err != nil {
		// response.Code = int(types.ErrCodeIp2Location)
		// response.Message = err.Error()
		response = types.LocationResponse{
			BaseHttpResponse: types.BaseHttpResponse{
				Code:    int(types.ErrCodeIp2Location),
				Message: err.Error(),
			},
			ClientIP: ip,
		}
	} else {
		response = types.LocationResponse{
			BaseHttpResponse: types.BaseHttpResponse{
				Code:    0,
				Message: "ok",
			},
			ClientIP:           ip,
			BandwidthRegion:    db.GetBandwidthRegion(&loc),
			CountryShort:       loc.Country_short,
			CountryLong:        loc.Country_long,
			Region:             loc.Region,
			City:               loc.City,
			Isp:                loc.Isp,
			Latitude:           loc.Latitude,
			Longitude:          loc.Longitude,
			Domain:             loc.Domain,
			Zipcode:            loc.Zipcode,
			Timezone:           loc.Timezone,
			Netspeed:           loc.Netspeed,
			Iddcode:            loc.Iddcode,
			Areacode:           loc.Areacode,
			Weatherstationcode: loc.Weatherstationcode,
			Weatherstationname: loc.Weatherstationname,
			Mcc:                loc.Mcc,
			Mnc:                loc.Mnc,
			Mobilebrand:        loc.Mobilebrand,
			Elevation:          loc.Elevation,
			Usagetype:          loc.Usagetype,
			Addresstype:        loc.Addresstype,
			Category:           loc.Category,
			District:           loc.District,
			Asn:                loc.Asn,
			As:                 loc.As,
		}
	}
	// response.IP2Locationrecord = loc
	ctx.JSON(http.StatusOK, response)
}
