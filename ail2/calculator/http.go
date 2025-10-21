package calculator

import (
	"AIL2-ProveNode/log"
	"AIL2-ProveNode/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalculatePointFromHttp(ctx *gin.Context) {
	miReq := types.DeepLinkMachineInfoST{}
	if err := ctx.ShouldBindJSON((&miReq)); err != nil {
		ctx.JSON(http.StatusBadRequest, types.BaseHttpResponse{
			Code:    int(types.ErrCodeParam),
			Message: "invalid machine info",
		})
		return
	}

	if err := miReq.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, types.BaseHttpResponse{
			Code:    int(types.ErrCodeParam),
			Message: err.Error(),
		})
		return
	}
	miReq.ClientIP = ctx.ClientIP()

	calcPoint, err := CalculatePointExactFromReport(
		miReq.GPUNames,
		miReq.GPUMemoryTotal,
		miReq.MemoryTotal,
	)
	if calcPoint == 0 {
		log.Log.Errorf("calculate gpu point from http %v failed %v", miReq, err)
		calcPoint, err = CalculatePointFuzzyFromReport(
			miReq.GPUNames,
			miReq.MemoryTotal,
		)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.BaseHttpResponse{
			Code:    int(types.ErrCodeMachineInfo),
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, types.CalculatePointResponse{
			BaseHttpResponse: types.BaseHttpResponse{
				Code:    0,
				Message: "ok",
			},
			CalcPoint: calcPoint,
		})
	}
}
